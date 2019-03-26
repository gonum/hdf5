// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
// inline static
// hid_t _go_hdf5_H5P_DEFAULT() { return H5P_DEFAULT; }
// hid_t _go_hdf5_H5P_DATASET_CREATE() { return H5P_DATASET_CREATE; }
// hid_t _go_hdf5_H5_SZIP_EC_OPTION_MASK() { return H5_SZIP_EC_OPTION_MASK; }
// hid_t _go_hdf5_H5_SZIP_NN_OPTION_MASK() { return H5_SZIP_NN_OPTION_MASK; }
import "C"

import "unsafe"

type PropType C.hid_t

type PropList struct {
	Identifier
}

var (
	P_DEFAULT              *PropList = newPropList(C._go_hdf5_H5P_DEFAULT())
	H5P_DATASET_CREATE     PropType  = PropType(C._go_hdf5_H5P_DATASET_CREATE()) // Properties for dataset creation
	H5_SZIP_EC_OPTION_MASK           = uint(C._go_hdf5_H5_SZIP_EC_OPTION_MASK()) // Selects entropy coding method.
	H5_SZIP_NN_OPTION_MASK           = uint(C._go_hdf5_H5_SZIP_NN_OPTION_MASK()) // Selects nearest neighbor coding method.
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

// Sets the size of the chunks used to store a chunked layout dataset.
// https://support.hdfgroup.org/HDF5/doc/RM/RM_H5P.html#Property-SetChunk
func (p *PropList) SetChunk(ndims uint, dim []uint) error {
	var c_dim *C.hsize_t
	c_dim = (*C.hsize_t)(unsafe.Pointer(&dim[0]))
	if err := h5err(C.H5Pset_chunk(C.hid_t(p.id), C.int(ndims), c_dim)); err != nil {
		return err
	}
	return nil
}

// Sets deflate (GNU gzip) compression method and compression level.
// https://support.hdfgroup.org/HDF5/doc/RM/RM_H5P.html#Property-SetDeflate
func (p *PropList) SetDeflate(level uint) error {
	if err := h5err(C.H5Pset_deflate(C.hid_t(p.id), C.uint(level))); err != nil {
		return err
	}
	return nil
}

// Sets up use of the SZIP compression filter.
// https://support.hdfgroup.org/HDF5/doc/RM/RM_H5P.html#Property-SetSzip
func (p *PropList) SetSzip(optionsMask, pixelsPerBlock uint) error {
	if err := h5err(C.H5Pset_szip(C.hid_t(p.id), C.uint(optionsMask), C.uint(pixelsPerBlock))); err != nil {
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
