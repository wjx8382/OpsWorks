package api

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	Namespace = "default"
	Kind      = "Pod"
)

func GetPodList(config *rest.Config) ([]string, error) {
	// Create clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	// List pods
	pods, err := clientset.CoreV1().Pods(Namespace).List(context.Background(), v1.ListOptions{})
	if err != nil {
		return nil, err
	}

	// Get pod names
	names := make([]string, 0, len(pods.Items))
	for _, pod := range pods.Items {
		names = append(names, pod.Name)
	}

	return names, nil
}
