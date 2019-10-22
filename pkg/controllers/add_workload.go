package controller

import (
	"github.com/xkcp0324/workload-controller/pkg/controllers/workload"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, workload.Add)
}
