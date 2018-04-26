// Api versions allow the api contract for a resource to be changed while keeping
// backward compatibility by support multiple concurrent versions
// of the same resource

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/ForgeRock/dsoperator/pkg/apis/ds
// +k8s:defaulter-gen=TypeMeta
// +groupName=ds.forgerock.com
package v1alpha1 // import "github.com/ForgeRock/dsoperator/pkg/apis/ds/v1alpha1"
