package edgecluster

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

//GetKubeConfigInsideCluster get Kubernates Config from cluster
//rest.Config return config to connect to kubernetes
func GetKubeConfigInsideCluster() (*rest.Config, error) {
	config, err := rest.InClusterConfig()

	return config, err
}

//GetKubeConfigOutOfCluster get Kubernates Config from os
//rest.Config return config to connect to kubernetes
func GetKubeConfigOutOfCluster(configName string) (*rest.Config, error) {
	var kubeconfig string

	//get hoem directory path
	homeDir, _ := GetHomeDirectoryPath()

	if homeDir != "" {
		flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(homeDir, ".kube", configName), "(optional) path to config file")
	} else {
		flag.StringVar(&kubeconfig, "kubeconfig", "", "path to kube config file")
	}

	configContext, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	return configContext, err
}

//ConnectToCluster connect to kubernates cluster
//rest.Config: config to initialize kubernates client
//ClientSet: return kubernetes client object
func ConnectToCluster(configContext *rest.Config) (kubernetes.Clientset, error) {

	client, err := kubernetes.NewForConfig(configContext)

	return *client, err
}

func GetHomeDirectoryPath() (string, error) {

	homePath, err := os.UserHomeDir()

	if homePath != "" {
		log.Print("linux mode")
		return homePath, err
	}

	return os.Getenv("USERPROFILE"), err
}

func Int32Ptr(i int32) *int32 {
	return &i
}
