// Package helm provides functionality to manage helm charts on a remote cluster
package helm

// HelmHelperContract declares the contract that can manage helm charts on remote cluster
type HelmHelperContract interface {
	// AddRepository adds the new repository to the local helm repo list
	// name: Mandaory. the helm repo name to add
	// url: Mandaory. the helm repo url to add
	// Returns error if something goes wrong
	AddRepository(name, url string) error

	// UpdateCharts updates the charts list for the local helm repo
	// Returns error if something goes wrong
	UpdateCharts() error

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
	InstallChart(kubeconfig, namespace, name, repo, chart string, args map[string]string) error
}
