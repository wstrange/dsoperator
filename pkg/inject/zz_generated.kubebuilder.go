package inject

import (
	dsv1alpha1 "github.com/ForgeRock/dsoperator/pkg/apis/ds/v1alpha1"
	rscheme "github.com/ForgeRock/dsoperator/pkg/client/clientset/versioned/scheme"
	"github.com/ForgeRock/dsoperator/pkg/controller/directoryservice"
	"github.com/ForgeRock/dsoperator/pkg/inject/args"
	"github.com/kubernetes-sigs/kubebuilder/pkg/inject/run"
	appsv1 "k8s.io/api/apps/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"
)

func init() {
	rscheme.AddToScheme(scheme.Scheme)

	// Inject Informers
	Inject = append(Inject, func(arguments args.InjectArgs) error {
		Injector.ControllerManager = arguments.ControllerManager

		if err := arguments.ControllerManager.AddInformerProvider(&dsv1alpha1.DirectoryService{}, arguments.Informers.Ds().V1alpha1().DirectoryServices()); err != nil {
			return err
		}

		// Add Kubernetes informers
		if err := arguments.ControllerManager.AddInformerProvider(&appsv1.StatefulSet{}, arguments.KubernetesInformers.Apps().V1().StatefulSets()); err != nil {
			return err
		}

		if c, err := directoryservice.ProvideController(arguments); err != nil {
			return err
		} else {
			arguments.ControllerManager.AddController(c)
		}
		return nil
	})

	// Inject CRDs
	Injector.CRDs = append(Injector.CRDs, &dsv1alpha1.DirectoryServiceCRD)
	// Inject PolicyRules
	Injector.PolicyRules = append(Injector.PolicyRules, rbacv1.PolicyRule{
		APIGroups: []string{"ds.forgerock.com"},
		Resources: []string{"*"},
		Verbs:     []string{"*"},
	})
	// Inject GroupVersions
	Injector.GroupVersions = append(Injector.GroupVersions, schema.GroupVersion{
		Group:   "ds.forgerock.com",
		Version: "v1alpha1",
	})
	Injector.RunFns = append(Injector.RunFns, func(arguments run.RunArguments) error {
		Injector.ControllerManager.RunInformersAndControllers(arguments)
		return nil
	})
}
