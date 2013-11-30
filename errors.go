package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
// #include "cgo_hdf5.h"
import "C"

import (
	"fmt"
)

// DisplayErrors enables/disables HDF5's automatic error printing
func DisplayErrors(b bool) error {
	switch b {
	case true:
		if err := togo_err(C._go_hdf5_unsilence_errors()); err != nil {
			return fmt.Errorf("hdf5: could not call H5E_set_auto(): %v", err)
		}
	default:
		if err := togo_err(C._go_hdf5_silence_errors()); err != nil {
			return fmt.Errorf("hdf5: could not call H5E_set_auto(): %v", err)
		}
	}
	return nil
}

func init() {
	err := DisplayErrors(false)
	if err != nil {
		panic(err.Error())
	}
}

// EOF
