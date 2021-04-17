package helm

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/decentralized-cloud/edge-cluster/services/edgecluster/types"
	"github.com/gofrs/flock"
	commonErrors "github.com/micro-business/go-core/system/errors"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/cli/values"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
	"helm.sh/helm/v3/pkg/strvals"
)

type helmHelper struct {
	logger   *zap.Logger
	settings *cli.EnvSettings
}

// NewHelmHelperService creates new instance of the helmHelper, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// Returns the new service or error if something goes wrong
func NewHelmHelperService(
	logger *zap.Logger) (HelmHelperContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	service := helmHelper{
		logger:   logger,
		settings: cli.New(),
	}

	if err := service.AddRepository("decentralized-cloud", "https://decentralized-cloud.github.io/helm"); err != nil {
		return nil, types.NewUnknownErrorWithError("failed add helm repo", err)
	}

	if err := service.AddRepository("portainer", "https://portainer.github.io/k8s"); err != nil {
		return nil, types.NewUnknownErrorWithError("failed add helm repo", err)
	}

	return &service, nil
}

// AddRepository adds the new repository to the local helm repo list
// name: Mandaory. the helm repo name to add
// url: Mandaory. the helm repo url to add
// Returns error if something goes wrong
func (service *helmHelper) AddRepository(name, url string) error {
	if strings.TrimSpace(name) == "" {
		return commonErrors.NewArgumentError("name", "name is required")
	}

	if strings.TrimSpace(url) == "" {
		return commonErrors.NewArgumentError("url", "url is required")
	}

	repoFile := service.settings.RepositoryConfig

	// Ensure the file directory exists as it is required for file locking
	err := os.MkdirAll(filepath.Dir(repoFile), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return err
	}

	// Acquire a file lock for process synchronization
	fileLock := flock.New(strings.Replace(repoFile, filepath.Ext(repoFile), ".lock", 1))
	lockCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	locked, err := fileLock.TryLockContext(lockCtx, time.Second)
	if err == nil && locked {
		defer func() {
			_ = fileLock.Unlock()
		}()
	}

	if err != nil {
		return err
	}

	b, err := ioutil.ReadFile(repoFile)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	var repositoryFile repo.File
	if err := yaml.Unmarshal(b, &repositoryFile); err != nil {
		return err
	}

	if repositoryFile.Has(name) {
		service.logger.Info("repository already exists", zap.String("name", name))

		return nil
	}

	config := repo.Entry{
		Name: name,
		URL:  url,
	}

	chartRepository, err := repo.NewChartRepository(&config, getter.All(service.settings))
	if err != nil {
		return err
	}

	if _, err := chartRepository.DownloadIndexFile(); err != nil {
		return errors.Wrapf(err, "looks like %q is not a valid chart repository or cannot be reached", url)
	}

	repositoryFile.Update(&config)

	if err := repositoryFile.WriteFile(repoFile, 0644); err != nil {
		return err
	}

	service.logger.Info("repository has been added to local helm repositories", zap.String("name", name))

	return nil
}

// UpdateCharts updates the charts list for the local helm repo
// Returns error if something goes wrong
func (service *helmHelper) UpdateCharts() error {
	repoFile, err := repo.LoadFile(service.settings.RepositoryConfig)
	if os.IsNotExist(errors.Cause(err)) || len(repoFile.Repositories) == 0 {
		return errors.New("no repositories found. You must add one before updating")
	}

	var repositories []*repo.ChartRepository
	for _, cfg := range repoFile.Repositories {
		chartRepository, err := repo.NewChartRepository(cfg, getter.All(service.settings))
		if err != nil {
			return err
		}

		repositories = append(repositories, chartRepository)
	}

	if len(repositories) > 0 {
		service.logger.Info("Hang tight while we grab the latest from your chart repositories...")

		errorsChan := make(chan error)
		waitGroupDoneChan := make(chan bool)

		var waitGroup sync.WaitGroup
		for _, repository := range repositories {
			waitGroup.Add(1)

			go func(re *repo.ChartRepository) {
				defer waitGroup.Done()

				if _, err := re.DownloadIndexFile(); err != nil {
					errorsChan <- errors.Wrapf(err, "...Unable to get an update from the %q chart repository (%s):\n\t%s\n", re.Config.Name, re.Config.URL, err)
				} else {
					service.logger.Info("...Successfully got an update from the chart repository", zap.String("name", re.Config.Name))
				}
			}(repository)
		}

		go func() {
			waitGroup.Wait()
			close(waitGroupDoneChan)
		}()

		select {
		case <-waitGroupDoneChan:
			close(errorsChan)

			break

		case err := <-errorsChan:
			close(errorsChan)

			return err
		}

		service.logger.Info("Update Complete. ⎈ Happy Helming!⎈")
	}

	return nil
}

// InstallChart installs chart on a remote cluster using the provided kubeconfig.
// If the helm chart was already registered, the method will try to upgrade the chart
// using the new provided value
// kubeconfig: Mandatory. string represents the kubeconfig of the remote cluster
// namespace: Mandatory. the namespace the helm chart should be installed to
// name: Mandaory. the name of the helm chart release
// chart: Mandaory. the name of the chart to install
// repo: Mandaory. the name of the repo to install
// args: Mandaory. extra arguments to install the helm chart with
// Returns error if something goes wrong
func (service *helmHelper) InstallChart(kubeconfig, namespace, name, repo, chart string, args map[string]string) error {
	if strings.TrimSpace(kubeconfig) == "" {
		return commonErrors.NewArgumentError("kubeconfig", "kubeconfig is required")
	}

	if strings.TrimSpace(namespace) == "" {
		return commonErrors.NewArgumentError("namespace", "namespace is required")
	}

	if strings.TrimSpace(name) == "" {
		return commonErrors.NewArgumentError("name", "name is required")
	}

	if strings.TrimSpace(repo) == "" {
		return commonErrors.NewArgumentError("repo", "repo is required")
	}

	if strings.TrimSpace(chart) == "" {
		return commonErrors.NewArgumentError("chart", "chart is required")
	}

	isDeployed, err := service.isHelmChartDeployed(kubeconfig, namespace, name)
	if err != nil {
		return err
	}

	if isDeployed {
		return service.upgrade(kubeconfig, namespace, name, repo, chart, args)
	}

	return service.install(kubeconfig, namespace, name, repo, chart, args)
}

func (service *helmHelper) install(kubeconfig, namespace, name, repo, chart string, args map[string]string) error {
	restClient := newKubeconfigClientGetter(namespace, kubeconfig)

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(restClient, namespace, os.Getenv("HELM_DRIVER"), service.debug); err != nil {
		return err
	}

	client := action.NewInstall(actionConfig)
	client.Namespace = namespace
	client.CreateNamespace = true

	if client.Version == "" && client.Devel {
		client.Version = ">0.0.0-0"
	}

	client.ReleaseName = name
	cp, err := client.ChartPathOptions.LocateChart(fmt.Sprintf("%s/%s", repo, chart), service.settings)
	if err != nil {
		return err
	}

	p := getter.All(service.settings)
	valueOpts := &values.Options{}
	vals, err := valueOpts.MergeValues(p)
	if err != nil {
		return err
	}

	// Add args
	if err := strvals.ParseInto(args["set"], vals); err != nil {
		return errors.Wrap(err, "failed parsing --set data")
	}

	// Check chart dependencies to make sure all are present in /charts
	chartRequested, err := loader.Load(cp)
	if err != nil {
		return err
	}

	validInstallableChart, err := isChartInstallable(chartRequested)
	if !validInstallableChart {
		return err
	}

	if req := chartRequested.Metadata.Dependencies; req != nil {
		if err := action.CheckDependencies(chartRequested, req); err != nil {
			if client.DependencyUpdate {
				man := &downloader.Manager{
					Out:              os.Stdout,
					ChartPath:        cp,
					Keyring:          client.ChartPathOptions.Keyring,
					SkipUpdate:       false,
					Getters:          p,
					RepositoryConfig: service.settings.RepositoryConfig,
					RepositoryCache:  service.settings.RepositoryCache,
				}
				if err := man.Update(); err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}

	if _, err = client.Run(chartRequested, vals); err != nil {
		return err
	}

	return nil
}

func (service *helmHelper) upgrade(kubeconfig, namespace, name, repo, chart string, args map[string]string) error {
	restClient := newKubeconfigClientGetter(namespace, kubeconfig)
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(restClient, namespace, os.Getenv("HELM_DRIVER"), service.debug); err != nil {
		return err
	}

	client := action.NewUpgrade(actionConfig)
	client.Namespace = namespace

	if client.Version == "" && client.Devel {
		client.Version = ">0.0.0-0"
	}

	cp, err := client.ChartPathOptions.LocateChart(fmt.Sprintf("%s/%s", repo, chart), service.settings)
	if err != nil {
		return err
	}

	p := getter.All(service.settings)
	valueOpts := &values.Options{}
	vals, err := valueOpts.MergeValues(p)
	if err != nil {
		return err
	}

	// Add args
	if err := strvals.ParseInto(args["set"], vals); err != nil {
		return errors.Wrap(err, "failed parsing --set data")
	}

	// Check chart dependencies to make sure all are present in /charts
	chartRequested, err := loader.Load(cp)
	if err != nil {
		return err
	}

	validInstallableChart, err := isChartInstallable(chartRequested)
	if !validInstallableChart {
		return err
	}

	if _, err = client.Run(name, chartRequested, vals); err != nil {
		return err
	}

	return nil
}

func (service *helmHelper) debug(format string, v ...interface{}) {
	service.logger.Info(fmt.Sprintf(format, v...))
}

func isChartInstallable(ch *chart.Chart) (bool, error) {
	switch ch.Metadata.Type {
	case "", "application":
		return true, nil
	}

	return false, errors.Errorf("%s charts are not installable", ch.Metadata.Type)
}

func (service *helmHelper) isHelmChartDeployed(kubeconfig, namespace, name string) (bool, error) {
	restClient := newKubeconfigClientGetter(namespace, kubeconfig)
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(restClient, namespace, os.Getenv("HELM_DRIVER"), service.debug); err != nil {
		return false, err
	}

	client := action.NewList(actionConfig)
	releases, err := client.Run()
	if err != nil {
		return false, err
	}

	for _, release := range releases {
		if release.Name == name && release.Namespace == namespace {
			return true, nil
		}
	}

	return false, nil
}
