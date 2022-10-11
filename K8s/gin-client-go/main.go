package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"k8s.io/client-go/tools/homedir"
	"k8s.io/client-go/tools/kubernetes"
	"k8s.io/client-go/util/clientcmd"
	"k8s.io/klog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	var kubeConfig *string
	ctx := context.Background()
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "absolute path to")
	} else {
		kubeConfig = flag.String("kubeConfig", "", "absolute path to kubeConfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		klog.Fatal(err)
		return
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
		return
	}
	namespaceList, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Fatal(err)
		return
	}
	namespaces := namespaceList.Items
	for _, ns := range namespaces {
		fmt.Println(ns.Name)
	}

}
