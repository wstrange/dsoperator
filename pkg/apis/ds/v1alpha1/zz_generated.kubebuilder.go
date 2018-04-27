package v1alpha1

import (
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: "ds.forgerock.com", Version: "v1alpha1"}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&DirectoryService{},
		&DirectoryServiceList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DirectoryServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DirectoryService `json:"items"`
}

// CRD Generation
func getFloat(f float64) *float64 {
	return &f
}

var (
	// Define CRDs for resources
	DirectoryServiceCRD = v1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: "directoryservices.ds.forgerock.com",
		},
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group:   "ds.forgerock.com",
			Version: "v1alpha1",
			Names: v1beta1.CustomResourceDefinitionNames{
				Kind:   "DirectoryService",
				Plural: "directoryservices",
			},
			Scope: "Namespaced",
			Validation: &v1beta1.CustomResourceValidation{
				OpenAPIV3Schema: &v1beta1.JSONSchemaProps{
					Type: "object",
					Properties: map[string]v1beta1.JSONSchemaProps{
						"apiVersion": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"kind": v1beta1.JSONSchemaProps{
							Type: "string",
						},
						"metadata": v1beta1.JSONSchemaProps{
							Type: "object",
						},
						"spec": v1beta1.JSONSchemaProps{
							Type: "object",
							Properties: map[string]v1beta1.JSONSchemaProps{
								"basedn": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"dirManager": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"image": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"initLoadURL": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"password": v1beta1.JSONSchemaProps{
									Type: "string",
								},
								"replicas": v1beta1.JSONSchemaProps{
									Type:   "integer",
									Format: "int32",
								},
							},
						},
						"status": v1beta1.JSONSchemaProps{
							Type: "object",
							Properties: map[string]v1beta1.JSONSchemaProps{
								"availableReplicas": v1beta1.JSONSchemaProps{
									Type:   "integer",
									Format: "int32",
								},
							},
						},
					},
				},
			},
		},
	}
)
