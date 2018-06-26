// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
// inline static
// hid_t _go_hdf5_H5P_DEFAULT() { return H5P_DEFAULT; }
import "C"

type PropType C.hid_t

type PropList struct {
	Identifier
}

var (
	P_DEFAULT *PropList = newPropList(C._go_hdf5_H5P_DEFAULT())
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
