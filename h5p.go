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
)

type PropType C.hid_t

type PropList struct {
	Location
}

var (
	P_DEFAULT *PropList = newPropList(C._go_hdf5_H5P_DEFAULT())
)

func newPropList(id C.hid_t) *PropList {
	p := &PropList{Location{id}}
	runtime.SetFinalizer(p, (*PropList).finalizer)
	return p
}

func (p *PropList) finalizer() {
	err := p.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing PropList: %s", err))
	}
	return
}

// NewPropList creates a new PropList as an instance of a property list class.
func NewPropList(cls_id PropType) (*PropList, error) {
	hid := C.H5Pcreate(C.hid_t(cls_id))
	err := h5err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	p := newPropList(hid)
	return p, err
}

// Close terminates access to a PropList.
func (p *PropList) Close() error {
	err := C.H5Pclose(p.id)
	return h5err(err)
}

// Copy copies an existing PropList to create a new PropList.
func (p *PropList) Copy() (*PropList, error) {
	hid := C.H5Pcopy(p.id)
	err := h5err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	o := newPropList(hid)
	return o, err
}
