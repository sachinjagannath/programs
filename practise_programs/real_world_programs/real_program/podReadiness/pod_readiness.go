package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type PodReadiness struct {
	PodName   string
	Ready     bool
	Condition string
}

func GetPodReadinessInfo(clientSet *kubernetes.Clientset, namespace string) ([]PodReadiness, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	svcPods, err := clientSet.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list pods: %w", err)
	}
	totalPodCount := len(svcPods.Items)
	if totalPodCount == 0 {
		fmt.Println("No pods found in namespace:", namespace)
		return nil, nil
	}

	smRespChan := make(chan PodReadiness, totalPodCount)
	var wg sync.WaitGroup

	for _, pod := range svcPods.Items {
		wg.Add(1)
		go func(p v1.Pod) {
			defer wg.Done()
			ready := false
			status := "Not Ready"

			for _, cond := range p.Status.Conditions {
				if cond.Type == v1.PodReady && cond.Status == v1.ConditionTrue {
					ready = true
					status = "Ready"
					break
				}
			}
			smRespChan <- PodReadiness{
				PodName:   p.Name,
				Ready:     ready,
				Condition: status,
			}
		}(pod)
	}
	wg.Wait()
	close(smRespChan)

	var readinessList []PodReadiness

	for r := range smRespChan {
		readinessList = append(readinessList, r)
	}

	return readinessList, nil
}

func main() {
	// Try in-cluster config first
	config, err := rest.InClusterConfig()
	if err != nil {
		// Fallback to local kubeconfig for Minikube
		kubeconfig := filepath.Join(
			os.Getenv("USERPROFILE"), ".kube", "config",
		)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatalf("Failed to load kubeconfig: %v", err)
		}
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create Kubernetes client: %v", err)
	}

	namespace := "default"
	fmt.Println("Checking pod readiness in namespace:", namespace)

	pods, err := GetPodReadinessInfo(clientSet, namespace)
	if err != nil {
		log.Fatalf("Error checking pod readiness: %v", err)
	}

	fmt.Println("\n--- Pod Readiness Report ---")
	for _, pod := range pods {
		fmt.Printf("Pod: %-30s | Status: %s\n", pod.PodName, pod.Condition)
	}
}
