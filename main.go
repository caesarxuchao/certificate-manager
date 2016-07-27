package main

import (
	"fmt"

	"k8s.io/kubernetes/pkg/api"
	client "k8s.io/kubernetes/pkg/client/clientset_generated/release_1_4"
	"k8s.io/kubernetes/pkg/client/restclient"
)

func NewClient() client.Interface {
	config, err := restclient.NewInCluster()
	if err != nil {
		panic(err)
	}
	return client.NewForConfigOrDie(config)
}

func PrintPods() {
	client := NewClient()
	pods, err := client.Core().Pods(api.NamespaceAll).List()
	if err != nil {
		panic(err)
	}
	fmt.Println(pods)
}

func main() {
	PrintPods()
}
