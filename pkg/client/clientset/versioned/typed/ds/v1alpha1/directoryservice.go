package v1alpha1

import (
	v1alpha1 "github.com/ForgeRock/dsoperator/pkg/apis/ds/v1alpha1"
	scheme "github.com/ForgeRock/dsoperator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DirectoryServicesGetter has a method to return a DirectoryServiceInterface.
// A group's client should implement this interface.
type DirectoryServicesGetter interface {
	DirectoryServices(namespace string) DirectoryServiceInterface
}

// DirectoryServiceInterface has methods to work with DirectoryService resources.
type DirectoryServiceInterface interface {
	Create(*v1alpha1.DirectoryService) (*v1alpha1.DirectoryService, error)
	Update(*v1alpha1.DirectoryService) (*v1alpha1.DirectoryService, error)
	UpdateStatus(*v1alpha1.DirectoryService) (*v1alpha1.DirectoryService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DirectoryService, error)
	List(opts v1.ListOptions) (*v1alpha1.DirectoryServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DirectoryService, err error)
	DirectoryServiceExpansion
}

// directoryServices implements DirectoryServiceInterface
type directoryServices struct {
	client rest.Interface
	ns     string
}

// newDirectoryServices returns a DirectoryServices
func newDirectoryServices(c *DsV1alpha1Client, namespace string) *directoryServices {
	return &directoryServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the directoryService, and returns the corresponding directoryService object, and an error if there is any.
func (c *directoryServices) Get(name string, options v1.GetOptions) (result *v1alpha1.DirectoryService, err error) {
	result = &v1alpha1.DirectoryService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("directoryservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DirectoryServices that match those selectors.
func (c *directoryServices) List(opts v1.ListOptions) (result *v1alpha1.DirectoryServiceList, err error) {
	result = &v1alpha1.DirectoryServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("directoryservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested directoryServices.
func (c *directoryServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("directoryservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a directoryService and creates it.  Returns the server's representation of the directoryService, and an error, if there is any.
func (c *directoryServices) Create(directoryService *v1alpha1.DirectoryService) (result *v1alpha1.DirectoryService, err error) {
	result = &v1alpha1.DirectoryService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("directoryservices").
		Body(directoryService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a directoryService and updates it. Returns the server's representation of the directoryService, and an error, if there is any.
func (c *directoryServices) Update(directoryService *v1alpha1.DirectoryService) (result *v1alpha1.DirectoryService, err error) {
	result = &v1alpha1.DirectoryService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("directoryservices").
		Name(directoryService.Name).
		Body(directoryService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *directoryServices) UpdateStatus(directoryService *v1alpha1.DirectoryService) (result *v1alpha1.DirectoryService, err error) {
	result = &v1alpha1.DirectoryService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("directoryservices").
		Name(directoryService.Name).
		SubResource("status").
		Body(directoryService).
		Do().
		Into(result)
	return
}

// Delete takes name of the directoryService and deletes it. Returns an error if one occurs.
func (c *directoryServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("directoryservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *directoryServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("directoryservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched directoryService.
func (c *directoryServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DirectoryService, err error) {
	result = &v1alpha1.DirectoryService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("directoryservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
