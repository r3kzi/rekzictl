package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	var namespace string

	flag.StringVar(&namespace, "namespace", "default", "namespace")
	flag.Parse()

	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")

	log.Println("Using kubeconfig: ", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	coreV1Api := clientset.CoreV1()

	namespaces, err := coreV1Api.Namespaces().List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("failed to get namespaces: %v", err)
	}

	fmt.Println(strings.Repeat("-", 20))
	for i, namespace := range namespaces.Items {
		fmt.Printf("[%d] %s\n", i, namespace.Name)
	}

	pods, err := coreV1Api.Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalf("failed to get pods: %v", err)
	}

	fmt.Println(strings.Repeat("-", 20))
	for i, pod := range pods.Items {
		fmt.Printf("[%d] %s\n", i, pod.GetName())
	}
}
