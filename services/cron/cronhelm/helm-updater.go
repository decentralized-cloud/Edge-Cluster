// Package cronhelm provides a cron job to keep the local helm repository updated
package cronhelm

import (
	cronContract "github.com/decentralized-cloud/edge-cluster/services/cron"
	"github.com/decentralized-cloud/edge-cluster/services/edgecluster/helm"
	commonErrors "github.com/micro-business/go-core/system/errors"
	cron "github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type helmCronService struct {
	logger      *zap.Logger
	cronSpec    string
	cron        *cron.Cron
	helmService helm.HelmHelperContract
}

// NewhelmCronService creates new instance of the helmCronService, setting up all dependencies and returns the instance
// logger: Mandatory. Reference to the logger service
// configurationService: Mandatory. Reference to the service that provides required configurations
// Returns the new service or error if something goes wrong
func NewhelmCronService(
	logger *zap.Logger,
	helmService helm.HelmHelperContract) (cronContract.CronContract, error) {
	if logger == nil {
		return nil, commonErrors.NewArgumentNilError("logger", "logger is required")
	}

	if helmService == nil {
		return nil, commonErrors.NewArgumentNilError("helmService", "helmService is required")
	}

	return &helmCronService{
		logger:      logger,
		cronSpec:    "@every 30m",
		cron:        cron.New(),
		helmService: helmService,
	}, nil
}

// Stop starts the helm crom service
// Returns error if something goes wrong
func (service *helmCronService) Start() error {
	service.logger.Info("cron Helm service started")

	_, err := service.cron.AddFunc(service.cronSpec, service.updateHelmCharts)
	if err != nil {
		return err
	}

	service.cron.Start()

	go service.updateHelmCharts()

	return nil
}

// Stop stops the helm crom service
// Returns error if something goes wrong
func (service *helmCronService) Stop() error {
	service.cron.Stop()

	return nil
}

// updateHelmCharts invoke helm service to update the list of the helm charts
func (service *helmCronService) updateHelmCharts() {
	if err := service.helmService.UpdateCharts(); err != nil {
		service.logger.Error("failed to update helm chart repositories", zap.Error(err))
	}

	service.logger.Info("finished updating helm repositories.")
}
