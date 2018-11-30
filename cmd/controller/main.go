package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/Sirupsen/logrus"
	clientset "github.com/yolo3301/kcrd/pkg/client/clientset/versioned"
	informers "github.com/yolo3301/kcrd/pkg/client/informers/externalversions"
	"github.com/yolo3301/kcrd/pkg/reconciler"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// kubeConfigPath := os.Getenv("HOME") + "/.kube/config"
	cfg, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	// _, err = kubernetes.NewForConfig(cfg)
	// if err != nil {
	// 	log.Fatalf("Error building kubernetes clientset: %v", err)
	// }

	queueClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building queue clientset: %v", err)
	}

	informerFactory := informers.NewSharedInformerFactory(queueClient, 10*time.Hour)
	informer := informerFactory.Queue().V1alpha1().Queues()

	controller := reconciler.NewController(informer)

	stopCh := make(chan struct{})
	defer close(stopCh)

	informerFactory.Start(stopCh)
	if ok := cache.WaitForCacheSync(stopCh, informer.Informer().HasSynced); !ok {
		log.Fatal("failed to wait for cache")
	}

	go controller.Run(stopCh)

	// use a channel to handle OS signals to terminate and gracefully shut
	// down processing
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, syscall.SIGTERM)
	signal.Notify(sigTerm, syscall.SIGINT)
	<-sigTerm
}
