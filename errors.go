package hdf5

// #include "hdf5.h"
//
// herr_t _go_hdf5_unsilence_errors(void) {
//   return H5Eset_auto2(H5E_DEFAULT, (H5E_auto2_t)(H5Eprint), stderr);
// }
//
// herr_t _go_hdf5_silence_errors(void) {
//   return H5Eset_auto2(H5E_DEFAULT, NULL, NULL);
// }
import "C"

import (
	"fmt"
)

// DisplayErrors enables/disables HDF5's automatic error printing
func DisplayErrors(b bool) error {
	switch b {
	case true:
		if err := h5err(C._go_hdf5_unsilence_errors()); err != nil {
			return fmt.Errorf("hdf5: could not call H5E_set_auto(): %v", err)
		}
	default:
		if err := h5err(C._go_hdf5_silence_errors()); err != nil {
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
