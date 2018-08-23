package fake

import (
	v1alpha1 "github.com/ForgeRock/dsoperator/pkg/client/clientset/versioned/typed/ds/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeDsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeDsV1alpha1) DirectoryServices(namespace string) v1alpha1.DirectoryServiceInterface {
	return &FakeDirectoryServices{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeDsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
