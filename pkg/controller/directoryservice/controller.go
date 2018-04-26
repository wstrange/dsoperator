package directoryservice

import (
	"fmt"
	"log"

	"github.com/kubernetes-sigs/kubebuilder/pkg/controller"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/eventhandlers"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/predicates"
	"github.com/kubernetes-sigs/kubebuilder/pkg/controller/types"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/record"

	dsv1alpha1 "github.com/ForgeRock/dsoperator/pkg/apis/ds/v1alpha1"
	dsv1alpha1client "github.com/ForgeRock/dsoperator/pkg/client/clientset/versioned/typed/ds/v1alpha1"
	dsv1alpha1informer "github.com/ForgeRock/dsoperator/pkg/client/informers/externalversions/ds/v1alpha1"
	dsv1alpha1lister "github.com/ForgeRock/dsoperator/pkg/client/listers/ds/v1alpha1"
	"github.com/ForgeRock/dsoperator/pkg/inject/args"
)

const (
	// SuccessSynced is used as part of the Event 'reason' when a Foo is synced
	SuccessSynced = "Synced"
	// ErrResourceExists is used as part of the Event 'reason' when a Foo fails
	// to sync due to a Deployment of the same name already existing.
	ErrResourceExists = "ErrResourceExists"

	// MessageResourceExists is the message used for Events when a resource
	// fails to sync due to a Deployment already existing
	MessageResourceExists = "Resource %q already exists and is not managed by DirectoryService"
	// MessageResourceSynced is the message used for an Event fired when a Foo
	// is synced successfully
	MessageResourceSynced = "DirectoryService synced successfully"
)

// EDIT THIS FILE
// This files was created by "kubebuilder create resource" for you to edit.
// Controller implementation logic for DirectoryService resources goes here.

// Reconcile the current state vs. desired state
func (bc *DirectoryServiceController) Reconcile(k types.ReconcileKey) error {
	// INSERT YOUR CODE HERE
	namespace, name := k.Namespace, k.Name
	log.Printf("Implement the Reconcile function on directoryservice.DirectoryServiceController to reconcile %s\n", name)
	//ds, err := bc.Informers.DirectoryServiceController().DsV1alpha1().DirectoryServices().Lister().DirectoryService(namespace).Get(name)

	ds, err := bc.directoryserviceLister.DirectoryServices(namespace).Get(name)

	if err != nil {
		log.Printf("Got an error %v", err)
		// The  resource may no longer exist, in which case we stop
		// processing.
		if errors.IsNotFound(err) {
			runtime.HandleError(fmt.Errorf("Resource '%s' in work queue no longer exists", k.Name))
			return nil
		}

		return err
	}
	log.Printf("got directory service object %s", ds.Name)

	setName := ds.Spec.StatefulSetName

	log.Printf("got set name %s", setName)

	if setName == "" {
		runtime.HandleError(fmt.Errorf("%s: statefulSet name must be specified", k))
		return nil
	}

	statefulSet, err := bc.KubernetesInformers.Apps().V1().StatefulSets().Lister().StatefulSets(ds.Namespace).Get(setName)

	if errors.IsNotFound(err) {
		log.Printf("Creating new statefulset %s", setName)
		statefulSet, err = bc.KubernetesClientSet.AppsV1().StatefulSets(ds.Namespace).Create(NewDSSet(ds))

		log.Printf("Created new set %v", statefulSet)
	}

	// If an error occurs during Get/Create, we'll requeue the item so we can
	// attempt processing again later. This could have been caused by a
	// temporary network failure, or any other transient reason.
	if err != nil {
		return err
	}

	if !metav1.IsControlledBy(statefulSet, ds) {
		msg := fmt.Sprintf(MessageResourceExists, statefulSet.Name)
		bc.directoryservicerecorder.Event(ds, corev1.EventTypeWarning, ErrResourceExists, msg)
		return fmt.Errorf(msg)
	}

	// todo: check for number of replicas...

	log.Printf("I got a statefulset set %s", statefulSet.Name)

	err = bc.updateDStatus(ds, statefulSet)

	if err != nil {
		return err
	}

	bc.directoryservicerecorder.Event(ds, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)

	return nil
}

func (bc *DirectoryServiceController) updateDStatus(ds *dsv1alpha1.DirectoryService, statefulSet *appsv1.StatefulSet) error {
	dsCopy := ds.DeepCopy()
	dsCopy.Status.AvailableReplicas = statefulSet.Status.Replicas

	_, err := bc.Clientset.DsV1alpha1().DirectoryServices(ds.Namespace).Update(dsCopy)
	return err
}

// LookupDirectoryService looks up a named directory service
func (bc *DirectoryServiceController) LookupDirectoryService(r types.ReconcileKey) (interface{}, error) {
	return bc.directoryserviceLister.DirectoryServices(r.Namespace).Get(r.Name)
}

// +kubebuilder:controller:group=ds,version=v1alpha1,kind=DirectoryService,resource=directoryservices
// +informers:group=apps,version=v1,kind=StatefulSet
// +rbac:rbac:groups=apps,resources=StatefulSet,verbs=get;list;watch;create;update;patch;delete
type DirectoryServiceController struct {
	// INSERT ADDITIONAL FIELDS HERE
	directoryserviceLister dsv1alpha1lister.DirectoryServiceLister
	directoryserviceclient dsv1alpha1client.DsV1alpha1Interface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	directoryservicerecorder record.EventRecorder

	// Added by Warren. Do we need this??
	args.InjectArgs
}

// ProvideController provides a controller that will be run at startup.  Kubebuilder will use codegeneration
// to automatically register this controller in the inject package
func ProvideController(arguments args.InjectArgs) (*controller.GenericController, error) {
	// INSERT INITIALIZATIONS FOR ADDITIONAL FIELDS HERE
	bc := &DirectoryServiceController{
		directoryserviceLister:   arguments.ControllerManager.GetInformerProvider(&dsv1alpha1.DirectoryService{}).(dsv1alpha1informer.DirectoryServiceInformer).Lister(),
		directoryserviceclient:   arguments.Clientset.DsV1alpha1(),
		directoryservicerecorder: arguments.CreateRecorder("DirectoryServiceController"),
		InjectArgs:               arguments,
	}

	// Create a new controller that will call DirectoryServiceController.Reconcile on changes to DirectoryServices
	gc := &controller.GenericController{
		Name:             "DirectoryServiceController",
		Reconcile:        bc.Reconcile,
		InformerRegistry: arguments.ControllerManager,
	}
	if err := gc.Watch(&dsv1alpha1.DirectoryService{}); err != nil {
		return gc, err
	}

	// INSERT ADDITIONAL WATCHES HERE BY CALLING gc.Watch.*() FUNCTIONS
	// NOTE: Informers for Kubernetes resources *MUST* be registered in the pkg/inject package so that they are started.
	// Added by warren.
	// todo: We may not want to directly watch these resources.
	// https://github.com/kubernetes/community/blob/8cafef897a22026d42f5e5bb3f104febe7e29830/contributors/devel/controllers.md
	// if err := gc.Watch(&appsv1.StatefulSet{}); err != nil {
	// 	return gc, err
	// }

	if err := gc.WatchControllerOf(&appsv1.StatefulSet{}, eventhandlers.Path{bc.LookupDirectoryService},
		predicates.ResourceVersionChanged); err != nil {
		return gc, err
	}

	return gc, nil
}
