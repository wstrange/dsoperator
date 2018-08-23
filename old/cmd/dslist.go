package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	ds "github.com/ForgeRock/dsoperator/pkg/client/clientset/versioned"
	dsop "github.com/ForgeRock/dsoperator/pkg/operator"
)

var (
	kuberconfig = flag.String("kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	master      = flag.String("master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
)

func main() {
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(*master, *kuberconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %v", err)
	}

	exampleClient, err := ds.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %v", err)
	}

	list, err := exampleClient.ForgerockV1().DirectoryServices("default").List(metav1.ListOptions{})
	//forgerockV1().DirectoryServers("default").List(metav1.ListOptions{})

	if err != nil {
		glog.Fatalf("Error listing all databases: %v", err)
	}

	for _, db := range list.Items {
		fmt.Printf("Directory Service found name %s with user %q basedn %q\n", db.Name, db.Spec.DirManager,
			db.Spec.BaseDN)
	}

	dsop.Start(exampleClient)
}
