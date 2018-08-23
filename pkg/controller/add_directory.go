// Experimental dsoperator

package controller

import (
	"github.com/ForgeRock/dsoperator/pkg/controller/directory"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, directory.Add)
}
