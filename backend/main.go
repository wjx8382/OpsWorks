package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const (
	Namespace = "default"
	Kind      = "Pod"
)

func main() {
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		flag.StringVar(&kubeconfig, "kubeconfig", fmt.Sprintf("%s/.kube/config", home), "(optional) absolute path to the kubeconfig file")
	} else {
		flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// Build config from kubeconfig file or in-cluster config
	config, err := buildConfig(kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// Create clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// List pods
	pods, err := clientset.CoreV1().Pods(Namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	// Print pod names
	for _, pod := range pods.Items {
		fmt.Printf("Pod: %s\n", pod.Name)
	}
}

func buildConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig == "" {
		return rest.InClusterConfig()
	}
	kubeConfig, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		return nil, err
	}
	restConfig, err := clientcmd.NewDefaultClientConfig(*kubeConfig, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return nil, err
	}
	return restConfig, nil
}
