package operator

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	forgerock_v1 "github.com/ForgeRock/dsoperator/pkg/apis/forgerock.com/v1"
	ds "github.com/ForgeRock/dsoperator/pkg/client/clientset/versioned"
	"github.com/golang/glog"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

const maxRetries = 5

// Controller object
type Controller struct {
	//logger       *logrus.Entry
	//clientset kubernetes.Interface
	clientset *ds.Clientset
	queue     workqueue.RateLimitingInterface
	informer  cache.SharedIndexInformer
	//eventHandler handlers.Handler
}

// Start an operator
func Start(kubeClient *ds.Clientset) {

	c := newResourceController(kubeClient)

	stopCh := make(chan struct{})
	defer close(stopCh)

	go c.Run(stopCh)

	glog.Info("Starting controller")
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM)
	signal.Notify(sigterm, syscall.SIGINT)
	<-sigterm
}

func newResourceController(client *ds.Clientset) *Controller {
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	resourceType := "directoryservice"

	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.ForgerockV1().DirectoryServices(meta_v1.NamespaceAll).List(meta_v1.ListOptions{})
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.ForgerockV1().DirectoryServices(meta_v1.NamespaceAll).Watch(options)
			},
		},
		&forgerock_v1.DirectoryService{},
		0, // skip resync
		cache.Indexers{},
	)

	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			glog.Infof("Processing add to %v: %s", resourceType, key)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: func(old, new interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(new)
			glog.Infof("Processing update to %v: %s", resourceType, key)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			glog.Infof("Processing delete to %v: %s", resourceType, key)
			if err == nil {
				queue.Add(key)
			}
		},
	})

	return &Controller{
		clientset: client,
		informer:  informer,
		queue:     queue,
		//eventHandler: eventHandler,
	}
}

// Run the go routine
func (c *Controller) Run(stopCh <-chan struct{}) {
	//defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()

	//c.logger.Info("Starting kubewatch controller")

	go c.informer.Run(stopCh)

	if !cache.WaitForCacheSync(stopCh, c.HasSynced) {
		// utilruntime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}

	//c.logger.Info("Kubewatch controller synced and ready")
	glog.Info("Kubewatch controller synced and ready")
	wait.Until(c.runWorker, time.Second, stopCh)
}

// HasSynced is required for the cache.Controller interface.
func (c *Controller) HasSynced() bool {
	return c.informer.HasSynced()
}

func (c *Controller) runWorker() {
	for c.processNextItem() {
		// continue looping
	}
}

func (c *Controller) processNextItem() bool {
	key, quit := c.queue.Get()

	if quit {
		return false
	}
	defer c.queue.Done(key)

	err := c.processItem(key.(string))
	if err == nil {
		// No error, reset the ratelimit counters
		c.queue.Forget(key)
	} else if c.queue.NumRequeues(key) < maxRetries {
		glog.Errorf("Error processing %s (will retry): %v", key, err)
		c.queue.AddRateLimited(key)
	} else {
		// err != nil and too many retries
		glog.Errorf("Error processing %s (giving up): %v", key, err)
		c.queue.Forget(key)
		//utilruntime.HandleError(err)
	}

	return true
}

// This is where the handlers will get called from when something intersting happens
// to a DS crd. Example: If a new DS CRD is created, we should create a DS cluster.
// If a DS CRD is updated (e.g. replicas moves from 2 to 3) we should add a new replica
func (c *Controller) processItem(key string) error {
	obj, exists, err := c.informer.GetIndexer().GetByKey(key)
	if err != nil {
		return fmt.Errorf("Error fetching object with key %s from store: %v", key, err)
	}

	if !exists {
		//c.eventHandler.ObjectDeleted(obj)
		glog.Infof("Thing deleted key %v obj is %v", key, obj)
		return nil
	}

	glog.Infof("Thing created is %v", obj)
	//c.eventHandler.ObjectCreated(obj)
	return nil
}
