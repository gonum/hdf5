// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
// static inline hid_t _go_hdf5_H5P_DEFAULT() { return H5P_DEFAULT; }
// static inline hid_t _go_hdf5_H5P_DATASET_CREATE() { return H5P_DATASET_CREATE; }
import "C"

import (
	"compress/zlib"
	"fmt"
	"unsafe"
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
func (p *PropList) SetChunk(dim []uint) error {
	ndims := len(dim)
	if ndims <= 0 {
		return fmt.Errorf("number of dimensions must be same size as the rank of the dataset, but zero received")
	}
	c_dim := (*C.hsize_t)(unsafe.Pointer(&dim[0]))
	if err := h5err(C.H5Pset_chunk(C.hid_t(p.id), C.int(ndims), c_dim)); err != nil {
		return err
	}
	return nil
}

// SetDeflate sets deflate (GNU gzip) compression method and compression level.
// If level is set as DefaultCompression, 6 will be used.
// https://support.hdfgroup.org/HDF5/doc/RM/RM_H5P.html#Property-SetDeflate
func (p *PropList) SetDeflate(level int) error {
	if level == DefaultCompression {
		level = 6
	}
	if level < 0 {
		return fmt.Errorf("unsupported compression level: %d", level)
	}
	if err := h5err(C.H5Pset_deflate(C.hid_t(p.id), C.uint(level))); err != nil {
		return err
	}
	return nil
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
