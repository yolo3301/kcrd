package reconciler

import (
	"context"
	"time"

	log "github.com/Sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	logger     *log.Entry
	reconciler *Reconciler
	queue      workqueue.RateLimitingInterface
}

func (c *Controller) Run(stopCh <-chan struct{}) {
	defer runtime.HandleCrash()
	defer c.queue.ShutDown()

	c.logger.Info("Starting controller...")
	wait.Until(func() {
		for c.processNextWorkItem() {
		}
	}, time.Second, stopCh)
}

func (c *Controller) processNextWorkItem() bool {
	c.logger.Info("Processing next item...")

	obj, shutdown := c.queue.Get()
	if shutdown {
		return false
	}

	defer c.queue.Done(obj)

	key := obj.(string)

	if err := c.reconciler.Reconcile(context.TODO(), key); err != nil {
		c.logger.Errorf("error syncing %q: %v", key, err)
	} else {
		c.logger.Infof("Successfully synced %q", key)
	}

	c.queue.Forget(obj)
	return true
}
