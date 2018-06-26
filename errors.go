// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
func DisplayErrors(on bool) error {
	var err error
	if on {
		err = h5err(C._go_hdf5_unsilence_errors())
	} else {
		err = h5err(C._go_hdf5_silence_errors())
	}
	if err != nil {
		return fmt.Errorf("hdf5: could not call H5E_set_auto(): %s", err)
	}
	return nil
}

func init() {
	if err := DisplayErrors(false); err != nil {
		panic(err)
	}
}

// EOF
