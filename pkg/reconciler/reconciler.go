package reconciler

import (
	"context"

	log "github.com/Sirupsen/logrus"
	informers "github.com/yolo3301/kcrd/pkg/client/informers/externalversions/queue/v1alpha1"
	listers "github.com/yolo3301/kcrd/pkg/client/listers/queue/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Reconciler struct {
	logger      *log.Entry
	queueLister listers.QueueLister
}

func NewController(queueInformer informers.QueueInformer) *Controller {
	c := &Reconciler{
		queueLister: queueInformer.Lister(),
		logger:      log.NewEntry(log.New()),
	}
	impl := &Controller{
		reconciler: c,
		queue: workqueue.NewNamedRateLimitingQueue(
			workqueue.DefaultControllerRateLimiter(),
			"Queues",
		),
		logger: c.logger,
	}

	c.logger.Info("Setting up event handlers")
	queueInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			// convert the resource object into a key (in this case
			// we are just doing it in the format of 'namespace/name')
			key, err := cache.MetaNamespaceKeyFunc(obj)
			c.logger.Infof("Add crd: %s", key)
			if err == nil {
				// add the key to the queue for the handler to get
				impl.queue.Add(key)
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(newObj)
			log.Infof("Update crd: %s", key)
			if err == nil {
				impl.queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			// DeletionHandlingMetaNamsespaceKeyFunc is a helper function that allows
			// us to check the DeletedFinalStateUnknown existence in the event that
			// a resource was deleted but it is still contained in the index
			//
			// this then in turn calls MetaNamespaceKeyFunc
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			log.Infof("Delete crd: %s", key)
			if err == nil {
				impl.queue.Add(key)
			}
		},
	})
	return impl
}

func (c *Reconciler) Reconcile(ctx context.Context, key string) error {
	// Convert the namespace/name string into a distinct namespace and name
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		c.logger.Errorf("invalid resource key: %s", key)
		return nil
	}

	// Get the Configuration resource with this namespace/name
	original, err := c.queueLister.Queues(namespace).Get(name)
	if errors.IsNotFound(err) {
		// The resource no longer exists, in which case we stop processing.
		c.logger.Errorf("crd %q in work queue no longer exists", key)
		return nil
	} else if err != nil {
		return err
	}

	original.DeepCopy()

	c.logger.Infof("crd %q in work queue exists", key)
	return nil
}
