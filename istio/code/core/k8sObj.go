package core

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func GetRestConfigOutCluster() *rest.Config {
	var kubeConfig *string
	if home := homeDir(); home != "" {
		kubeConfig = flag.String("kube_config", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kube config file")
	} else {
		kubeConfig = flag.String("kube_config", "", "absolute path to the kube config file")
	}
	flag.Parse()

	// use the current context in kubeConfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err.Error())
	}
	return config
}

func GetRestConfigInCluster() *rest.Config {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	return config
}

func getClientSet() *kubernetes.Clientset {
	// OutCluster 集群外部访问，必须要有证书才行
	config := GetRestConfigOutCluster()
	// 集群内部访问pod列表不需要任何东西
	//config := GetRestConfigInCluster()
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientSet
}
