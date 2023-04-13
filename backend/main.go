package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"OpsWorks/api"
)

func main() {
	var kubeconfig string
	flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()

	// Build config from kubeconfig file or in-cluster config
	config, err := buildConfig(kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	// Create router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/pods", func(w http.ResponseWriter, r *http.Request) {
		names, err := api.GetPodList(config)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, name := range names {
			fmt.Fprintf(w, "Pod: %s\n", name)
		}
	})

	// Start server
	log.Fatal(http.ListenAndServe(":8080", router))
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
