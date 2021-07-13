// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Attribute struct {
	Identifier
}

func newAttribute(id C.hid_t) *Attribute {
	return &Attribute{Identifier{id}}
}

func createAttribute(id C.hid_t, name string, dtype *Datatype, dspace *Dataspace, acpl *PropList) (*Attribute, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	hid := C.H5Acreate2(id, c_name, dtype.id, dspace.id, acpl.id, P_DEFAULT.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newAttribute(hid), nil
}

func openAttribute(id C.hid_t, name string) (*Attribute, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Aopen(id, c_name, P_DEFAULT.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newAttribute(hid), nil
}

// Access the type of an attribute
func (s *Attribute) GetType() Identifier {
	ftype := C.H5Aget_type(s.id)
	return Identifier{ftype}
}

// Close releases and terminates access to an attribute.
func (s *Attribute) Close() error {
	return s.closeWith(h5aclose)
}

func h5aclose(id C.hid_t) C.herr_t {
	return C.H5Aclose(id)
}

// Space returns an identifier for a copy of the dataspace for a attribute.
func (s *Attribute) Space() *Dataspace {
	hid := C.H5Aget_space(s.id)
	if int(hid) > 0 {
		return newDataspace(hid)
	}
	return nil
}

// Read reads raw data from a attribute into a buffer.
func (s *Attribute) Read(data interface{}, dtype *Datatype) error {
	var (
		addr unsafe.Pointer
		rv   = reflect.ValueOf(data)
	)
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("hdf5: read expects a pointer value")
	}
	v := reflect.Indirect(rv)

	switch v.Kind() {

	case reflect.Array:
		addr = unsafe.Pointer(v.UnsafeAddr())

	case reflect.String:
		dtAttr, err := copyDatatype(s.GetType().id)
		if err != nil {
			return fmt.Errorf("hdf5: could not access attribute datatype: %v", err)
		}
		defer dtAttr.Close()

		dtype = dtAttr
		dlen := dtype.Size()
		cstr := (*C.char)(unsafe.Pointer(C.malloc(C.size_t(uint(unsafe.Sizeof(byte(0))) * (dlen + 1)))))
		defer C.free(unsafe.Pointer(cstr))
		switch {
		case C.H5Tis_variable_str(dtAttr.Identifier.id) != 0:
			addr = unsafe.Pointer(&cstr)
		default:
			addr = unsafe.Pointer(cstr)
		}
		defer func() {
			v.SetString(C.GoString(cstr))
		}()

	default:
		addr = unsafe.Pointer(v.UnsafeAddr())
	}

	rc := C.H5Aread(s.id, dtype.id, addr)
	err := h5err(rc)
	return err
}

// Write writes raw data from a buffer to an attribute.
func (s *Attribute) Write(data interface{}, dtype *Datatype) error {
	var addr unsafe.Pointer
	v := reflect.Indirect(reflect.ValueOf(data))
	switch v.Kind() {

	case reflect.Array:
		addr = unsafe.Pointer(v.UnsafeAddr())

	case reflect.String:
		str := C.CString(v.String())
		defer C.free(unsafe.Pointer(str))
		addr = unsafe.Pointer(&str)

	case reflect.Ptr:
		addr = unsafe.Pointer(v.Pointer())

	default:
		addr = unsafe.Pointer(v.UnsafeAddr())
	}

	rc := C.H5Awrite(s.id, dtype.id, addr)
	err := h5err(rc)
	return err
}
