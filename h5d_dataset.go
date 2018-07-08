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

type Dataset struct {
	Identifier

	typ *Datatype
}

func newDataset(id C.hid_t, typ *Datatype) *Dataset {
	return &Dataset{Identifier: Identifier{id}, typ: typ}
}

func createDataset(id C.hid_t, name string, dType *Datatype, dSpace *Dataspace, dcpl *PropList) (*Dataset, error) {
	dType, err := dType.Copy() // For safety
	if err != nil {
		return nil, err
	}
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	hid := C.H5Dcreate2(id, cName, dType.id, dSpace.id, P_DEFAULT.id, dcpl.id, P_DEFAULT.id)
	if err := checkId(hid); err != nil {
		return nil, err
	}
	return newDataset(hid, dType), nil
}

// Close releases and terminates access to a dataset.
func (s *Dataset) Close() error {
	return s.closeWith(h5dclose)
}

func h5dclose(id C.hid_t) C.herr_t {
	return C.H5Dclose(id)
}

// Space returns an identifier for a copy of the dataspace for a dataset.
func (s *Dataset) Space() *Dataspace {
	hid := C.H5Dget_space(s.id)
	if int(hid) > 0 {
		return newDataspace(hid)
	}
	return nil
}

// ReadSubset reads a subset of raw data from a dataset into a buffer.
func (s *Dataset) ReadSubset(data interface{}, memSpace, fileSpace *Dataspace) error {
	dType, err := s.Datatype()
	defer dType.Close()
	if err != nil {
		return err
	}

	var addr unsafe.Pointer
	v := reflect.Indirect(reflect.ValueOf(data))

	switch v.Kind() {

	case reflect.Array:
		addr = unsafe.Pointer(v.UnsafeAddr())

	case reflect.Slice:
		slice := (*reflect.SliceHeader)(unsafe.Pointer(v.UnsafeAddr()))
		addr = unsafe.Pointer(slice.Data)

	case reflect.String:
		str := (*reflect.StringHeader)(unsafe.Pointer(v.UnsafeAddr()))
		addr = unsafe.Pointer(str.Data)

	case reflect.Ptr:
		addr = unsafe.Pointer(v.Pointer())

	default:
		addr = unsafe.Pointer(v.UnsafeAddr())
	}

	var fileSpaceId, memSpaceId C.hid_t = 0, 0
	if memSpace != nil {
		memSpaceId = memSpace.id
	}
	if fileSpace != nil {
		fileSpaceId = fileSpace.id
	}
	rc := C.H5Dread(s.id, dType.id, memSpaceId, fileSpaceId, 0, addr)
	err = h5err(rc)
	return err
}

// Read reads raw data from a dataset into a buffer.
func (s *Dataset) Read(data interface{}) error {
	return s.ReadSubset(data, nil, nil)
}

// WriteSubset writes a subset of raw data from a buffer to a dataset.
func (s *Dataset) WriteSubset(data interface{}, memSpace, fileSpace *Dataspace) error {
	dType, err := s.Datatype()
	defer dType.Close()
	if err != nil {
		return err
	}

	addr := unsafe.Pointer(nil)
	v := reflect.Indirect(reflect.ValueOf(data))

	switch v.Kind() {

	case reflect.Array:
		addr = unsafe.Pointer(v.UnsafeAddr())

	case reflect.Slice:
		slice := (*reflect.SliceHeader)(unsafe.Pointer(v.UnsafeAddr()))
		addr = unsafe.Pointer(slice.Data)

	case reflect.String:
		str := (*reflect.StringHeader)(unsafe.Pointer(v.UnsafeAddr()))
		addr = unsafe.Pointer(str.Data)

	case reflect.Ptr:
		addr = unsafe.Pointer(v.Pointer())

	default:
		addr = unsafe.Pointer(v.UnsafeAddr())
	}

	var fileSpaceId, memSpaceId C.hid_t = 0, 0
	if memSpace != nil {
		memSpaceId = memSpace.id
	}
	if fileSpace != nil {
		fileSpaceId = fileSpace.id
	}
	rc := C.H5Dwrite(s.id, dType.id, memSpaceId, fileSpaceId, 0, addr)
	err = h5err(rc)
	return err
}

// Write writes raw data from a buffer to a dataset.
func (s *Dataset) Write(data interface{}) error {
	return s.WriteSubset(data, nil, nil)
}

// Creates a new attribute at this location. The returned attribute
// must be closed by the user when it is no longer needed.
func (s *Dataset) CreateAttribute(name string, dType *Datatype, dSpace *Dataspace) (*Attribute, error) {
	return createAttribute(s.id, name, dType, dSpace, P_DEFAULT)
}

// Creates a new attribute at this location. The returned
// attribute must be closed by the user when it is no longer needed.
func (s *Dataset) CreateAttributeWith(name string, dType *Datatype, dSpace *Dataspace, acpl *PropList) (*Attribute, error) {
	return createAttribute(s.id, name, dType, dSpace, acpl)
}

// Opens an existing attribute. The returned attribute must be closed
// by the user when it is no longer needed. The returned attribute
// must be closed by the user when it is no longer needed.
func (s *Dataset) OpenAttribute(name string) (*Attribute, error) {
	return openAttribute(s.id, name)
}

// Datatype returns the HDF5 Datatype of the Dataset. The returned
// datatype must be closed by the user when it is no longer needed.
func (s *Dataset) Datatype() (*Datatype, error) {
	dTypeId := C.H5Dget_type(s.id)
	if dTypeId < 0 {
		return nil, fmt.Errorf("couldn't open Datatype from Dataset %q", s.Name())
	}
	return NewDatatype(dTypeId), nil
}

// hasIllegalGoPointer returns whether the Dataset is known to have
// a Go pointer to Go pointer chain. If the Dataset was created by
// a call to OpenDataset without a read operation, it will be false,
// but will not be a valid reflection of the real situation.
func (s *Dataset) hasIllegalGoPointer() bool {
	return s.typ.hasIllegalGoPointer()
}
