// Experimental dsoperator

package apis

import (
	"github.com/ForgeRock/dsoperator/pkg/apis/ds/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes, v1beta1.SchemeBuilder.AddToScheme)
}
