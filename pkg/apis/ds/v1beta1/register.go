// Experimental dsoperator

// NOTE: Boilerplate only.  Ignore this file.

// Package v1beta1 contains API Schema definitions for the ds v1beta1 API group
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen=package,register
// +k8s:conversion-gen=github.com/ForgeRock/dsoperator/pkg/apis/ds
// +k8s:defaulter-gen=TypeMeta
// +groupName=ds.forgeock.com
package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/runtime/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "ds.forgeock.com", Version: "v1beta1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}
)
