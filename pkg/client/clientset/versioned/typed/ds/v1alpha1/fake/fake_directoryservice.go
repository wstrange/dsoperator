package fake

import (
	v1alpha1 "github.com/ForgeRock/dsoperator/pkg/apis/ds/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDirectoryServices implements DirectoryServiceInterface
type FakeDirectoryServices struct {
	Fake *FakeDsV1alpha1
	ns   string
}

var directoryservicesResource = schema.GroupVersionResource{Group: "ds.forgerock.com", Version: "v1alpha1", Resource: "directoryservices"}

var directoryservicesKind = schema.GroupVersionKind{Group: "ds.forgerock.com", Version: "v1alpha1", Kind: "DirectoryService"}

// Get takes name of the directoryService, and returns the corresponding directoryService object, and an error if there is any.
func (c *FakeDirectoryServices) Get(name string, options v1.GetOptions) (result *v1alpha1.DirectoryService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(directoryservicesResource, c.ns, name), &v1alpha1.DirectoryService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryService), err
}

// List takes label and field selectors, and returns the list of DirectoryServices that match those selectors.
func (c *FakeDirectoryServices) List(opts v1.ListOptions) (result *v1alpha1.DirectoryServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(directoryservicesResource, directoryservicesKind, c.ns, opts), &v1alpha1.DirectoryServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DirectoryServiceList{}
	for _, item := range obj.(*v1alpha1.DirectoryServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested directoryServices.
func (c *FakeDirectoryServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(directoryservicesResource, c.ns, opts))

}

// Create takes the representation of a directoryService and creates it.  Returns the server's representation of the directoryService, and an error, if there is any.
func (c *FakeDirectoryServices) Create(directoryService *v1alpha1.DirectoryService) (result *v1alpha1.DirectoryService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(directoryservicesResource, c.ns, directoryService), &v1alpha1.DirectoryService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryService), err
}

// Update takes the representation of a directoryService and updates it. Returns the server's representation of the directoryService, and an error, if there is any.
func (c *FakeDirectoryServices) Update(directoryService *v1alpha1.DirectoryService) (result *v1alpha1.DirectoryService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(directoryservicesResource, c.ns, directoryService), &v1alpha1.DirectoryService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDirectoryServices) UpdateStatus(directoryService *v1alpha1.DirectoryService) (*v1alpha1.DirectoryService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(directoryservicesResource, "status", c.ns, directoryService), &v1alpha1.DirectoryService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryService), err
}

// Delete takes name of the directoryService and deletes it. Returns an error if one occurs.
func (c *FakeDirectoryServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(directoryservicesResource, c.ns, name), &v1alpha1.DirectoryService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDirectoryServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(directoryservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DirectoryServiceList{})
	return err
}

// Patch applies the patch and returns the patched directoryService.
func (c *FakeDirectoryServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DirectoryService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(directoryservicesResource, c.ns, name, data, subresources...), &v1alpha1.DirectoryService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DirectoryService), err
}
