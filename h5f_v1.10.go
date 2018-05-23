// Copyright ©2018 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//+build hdf5_v1.10

package hdf5

// #include "hdf5.h"
import "C"

// StartSWMRWrite enables SWMR writing mode for this file.
func (f *File) StartSWMRWrite() error {
	return h5err(C.H5Fstart_swmr_write(f.id))
}
