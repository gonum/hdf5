// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
// static inline hid_t _go_hdf5_H5P_DEFAULT() { return H5P_DEFAULT; }
// static inline hid_t _go_hdf5_H5P_DATASET_CREATE() { return H5P_DATASET_CREATE; }
// static inline hid_t _go_hdf5_H5P_DATASET_ACCESS() { return H5P_DATASET_ACCESS; }
import "C"

import (
	"compress/zlib"
	"fmt"
)

const (
	NoCompression      = zlib.NoCompression
	BestSpeed          = zlib.BestSpeed
	BestCompression    = zlib.BestCompression
	DefaultCompression = zlib.DefaultCompression
)

type PropType C.hid_t

type PropList struct {
	Identifier
}

var (
	P_DEFAULT        *PropList = newPropList(C._go_hdf5_H5P_DEFAULT())
	P_DATASET_CREATE PropType  = PropType(C._go_hdf5_H5P_DATASET_CREATE()) // Properties for dataset creation
	P_DATASET_ACCESS PropType  = PropType(C._go_hdf5_H5P_DATASET_ACCESS()) // Properties for dataset access
)

func newPropList(id C.hid_t) *PropList {
	return &PropList{Identifier{id}}
}

// NewPropList creates a new PropList as an instance of a property list class.
// The returned proplist must be closed by the user when it is no longer needed.
func NewPropList(cls_id PropType) (*PropList, error) {
	hid := C.H5Pcreate(C.hid_t(cls_id))
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newPropList(hid), nil
}

// Close terminates access to a PropList.
func (p *PropList) Close() error {
	return p.closeWith(h5pclose)
}

// SetChunk sets the size of the chunks used to store a chunked layout dataset.
// https://support.hdfgroup.org/HDF5/doc/RM/RM_H5P.html#Property-SetChunk
func (p *PropList) SetChunk(dims []uint) error {
	ndims := len(dims)
	if ndims <= 0 {
		return fmt.Errorf("number of dimensions must be same size as the rank of the dataset, but zero received")
	}
	c_dim := make([]C.hsize_t, ndims)
	for i := range dims {
		c_dim[i] = C.hsize_t(dims[i])
	}
	return h5err(C.H5Pset_chunk(C.hid_t(p.id), C.int(ndims), &c_dim[0]))
}

// GetChunk retrieves the size of chunks for the raw data of a chunked layout dataset.
// https://support.hdfgroup.org/HDF5/doc/RM/RM_H5P.html#Property-GetChunk
func (p *PropList) GetChunk(ndims int) (dims []uint, err error) {
	if ndims <= 0 {
		err = fmt.Errorf("number of dimensions must be same size as the rank of the dataset, but nonpositive value received")
		return
	}
	c_dims := make([]C.hsize_t, ndims)
	if err = h5err(C.H5Pget_chunk(C.hid_t(p.id), C.int(ndims), &c_dims[0])); err != nil {
		return
	}
	dims = make([]uint, ndims)
	for i := range dims {
		dims[i] = uint(c_dims[i])
	}
	return
}

// SetDeflate sets deflate (GNU gzip) compression method and compression level.
// If level is set as DefaultCompression, 6 will be used.
// https://support.hdfgroup.org/HDF5/doc/RM/RM_H5P.html#Property-SetDeflate
func (p *PropList) SetDeflate(level int) error {
	if level == DefaultCompression {
		level = 6
	}
	return h5err(C.H5Pset_deflate(C.hid_t(p.id), C.uint(level)))
}

// SetChunkCache sets the raw data chunk cache parameters.
// To reset them as default, use `D_CHUNK_CACHE_NSLOTS_DEFAULT`, `D_CHUNK_CACHE_NBYTES_DEFAULT` and `D_CHUNK_CACHE_W0_DEFAULT`.
// https://support.hdfgroup.org/HDF5/doc/RM/RM_H5P.html#Property-SetChunkCache
func (p *PropList) SetChunkCache(nslots, nbytes int, w0 float64) error {
	return h5err(C.H5Pset_chunk_cache(C.hid_t(p.id), C.size_t(nslots), C.size_t(nbytes), C.double(w0)))
}

// GetChunkCache retrieves the number of chunk slots in the raw data chunk cache hash table.
// https://support.hdfgroup.org/HDF5/doc/RM/RM_H5P.html#Property-GetChunkCache
func (p *PropList) GetChunkCache() (nslots, nbytes int, w0 float64, err error) {
	var (
		c_nslots C.size_t
		c_nbytes C.size_t
		c_w0     C.double
	)
	err = h5err(C.H5Pget_chunk_cache(C.hid_t(p.id), &c_nslots, &c_nbytes, &c_w0))
	return int(c_nslots), int(c_nbytes), float64(c_w0), err
}

func h5pclose(id C.hid_t) C.herr_t {
	return C.H5Pclose(id)
}

// Copy copies an existing PropList to create a new PropList.
func (p *PropList) Copy() (*PropList, error) {
	hid := C.H5Pcopy(p.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newPropList(hid), nil
}
