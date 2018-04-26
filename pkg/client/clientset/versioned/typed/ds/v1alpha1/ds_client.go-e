package v1alpha1

import (
	v1alpha1 "github.com/ForgeRock/dsoperator/pkg/apis/ds/v1alpha1"
	"github.com/ForgeRock/dsoperator/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type DsV1alpha1Interface interface {
	RESTClient() rest.Interface
	DirectoryServicesGetter
}

// DsV1alpha1Client is used to interact with features provided by the ds.forgerock.com group.
type DsV1alpha1Client struct {
	restClient rest.Interface
}

func (c *DsV1alpha1Client) DirectoryServices(namespace string) DirectoryServiceInterface {
	return newDirectoryServices(c, namespace)
}

// NewForConfig creates a new DsV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*DsV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &DsV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new DsV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *DsV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new DsV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *DsV1alpha1Client {
	return &DsV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *DsV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
