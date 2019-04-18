// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

// #include "hdf5.h"
import "C"

const (
	// Used to unset chunk cache configuration parameter.
	// https://support.hdfgroup.org/HDF5/doc/RM/RM_H5P.html#Property-SetChunkCache
	D_CHUNK_CACHE_NSLOTS_DEFAULT int     = -1 // rdcc_nslots
	D_CHUNK_CACHE_NBYTES_DEFAULT int     = -1 // rdcc_nbytes
	D_CHUNK_CACHE_W0_DEFAULT     float64 = -1 // rdcc_w0
)
