package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
// inline static
// hid_t _go_hdf5_H5P_DEFAULT() { return H5P_DEFAULT; }
import "C"

import (
	//"unsafe"

	"fmt"
	"runtime"
)

type PropType C.hid_t

// --- H5P: Property List Interface ---

type PropList struct {
	id C.hid_t
}

var (
	P_DEFAULT *PropList = new_proplist(C._go_hdf5_H5P_DEFAULT())
)

func new_proplist(id C.hid_t) *PropList {
	p := &PropList{id: id}
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

// Creates a new property as an instance of a property list class.
// hid_t H5Pcreate(hid_t cls_id )
func NewPropList(cls_id PropType) (*PropList, error) {
	hid := C.H5Pcreate(C.hid_t(cls_id))
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	p := new_proplist(hid)
	return p, err
}

// Terminates access to a property list.
// herr_t H5Pclose(hid_t plist )
func (p *PropList) Close() error {
	err := C.H5Pclose(p.id)
	return togo_err(err)
}

// Copies an existing property list to create a new property list.
// hid_t H5Pcopy(hid_t plist )
func (p *PropList) Copy() (*PropList, error) {

	hid := C.H5Pcopy(p.id)
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	o := new_proplist(hid)
	return o, err

}

// EOF
