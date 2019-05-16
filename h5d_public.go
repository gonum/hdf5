// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

// #include "hdf5.h"
import "C"

// Used to unset chunk cache configuration parameter.
// https://support.hdfgroup.org/HDF5/doc/RM/RM_H5P.html#Property-SetChunkCache
const (
	D_CHUNK_CACHE_NSLOTS_DEFAULT int     = -1 // The number of chunk slots in the raw data chunk cache for this dataset
	D_CHUNK_CACHE_NBYTES_DEFAULT int     = -1 // The total size of the raw data chunk cache for this dataset
	D_CHUNK_CACHE_W0_DEFAULT     float64 = -1 // The chunk preemption policy for this dataset
)
