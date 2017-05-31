// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
// inline static
// hid_t _go_hdf5_H5P_DEFAULT() { return H5P_DEFAULT; }
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

type PropType C.hid_t

type PropList struct {
	Identifier
}

var (
	P_DEFAULT *PropList = newPropList(C._go_hdf5_H5P_DEFAULT())
)

func newPropList(id C.hid_t) *PropList {
	p := &PropList{Identifier{id}}
	runtime.SetFinalizer(p, (*PropList).finalizer)
	return p
}

func (p *PropList) finalizer() {
	if err := p.Close(); err != nil {
		panic(fmt.Errorf("error closing PropList: %s", err))
	}
}

func AttachPropList(id C.hid_t) *PropList {
	p := &PropList{Identifier{id}}
	runtime.SetFinalizer(p, (*PropList).finalizer)
	return p
}

// NewPropList creates a new PropList as an instance of a property list class.
func NewPropList(cls_id PropType) (*PropList, error) {
	hid := C.H5Pcreate(C.hid_t(cls_id))
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newPropList(hid), nil
}

// NewPropList creates a new PropList as an instance of a property list class.
func NewDefaultDatasetPropList() (*PropList, error) {
	hid := C.H5Pcreate(C.H5P_CLS_DATASET_CREATE_ID_g)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newPropList(hid), nil
}

// Close terminates access to a PropList.
func (p *PropList) Close() error {
	if p.id == 0 {
		return nil
	}
	err := h5err(C.H5Pclose(p.id))
	p.id = 0
	return err
}

// Copy copies an existing PropList to create a new PropList.
func (p *PropList) Copy() (*PropList, error) {
	hid := C.H5Pcopy(p.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newPropList(hid), nil
}

/* Object creation property list (OCPL) routines */

//H5_DLL int H5Pget_nfilters(hid_t plist_id);
func (p *PropList) H5Pget_nfilters() int {
	return int(C.H5Pget_nfilters(p.id))
}

//H5_DLL herr_t H5Pset_deflate(hid_t plist_id, unsigned aggression);
func (p *PropList) H5Pset_deflate(aggression uint) error {
	return h5err(C.H5Pset_deflate(p.id, C.uint(aggression)))
}

/* Dataset creation property list (DCPL) routines */

//H5_DLL herr_t H5Pset_layout(hid_t plist_id, H5D_layout_t layout);
func (p *PropList) H5Pset_layout(layout_code uint) error {
	return h5err(C.H5Pset_layout(p.id, C.H5D_layout_t(layout_code)))
}

//H5_DLL H5D_layout_t H5Pget_layout(hid_t plist_id);
func (p *PropList) H5Pget_layout() uint {
	return uint(C.H5Pget_layout(p.id))
}

//H5_DLL herr_t H5Pset_chunk(hid_t plist_id, int ndims, const hsize_t dim[/*ndims*/]);
func (p *PropList) H5Pset_chunk(ndims uint, dims []uint) error {
	return h5err(C.H5Pset_chunk(p.id, C.int(ndims), (*C.hsize_t)(unsafe.Pointer(&dims[0]))))
}

//H5_DLL int H5Pget_chunk(hid_t plist_id, int max_ndims, hsize_t dim[]/*out*/);
func (p *PropList) H5Pget_chunk() uint {
	ndims := uint(0xffff)
	dims := []uint{ndims}
	c_dims := (*C.hsize_t)(unsafe.Pointer(&dims[0]))
	res := C.H5Pget_chunk(p.id, C.int(ndims), c_dims)

	return uint(res)
}
