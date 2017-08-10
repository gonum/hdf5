// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

// #include "hdf5.h"
// #include "hdf5_hl.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

// CommonFG is for methods common to both File and Group
type CommonFG struct {
	Identifier
}

// Group is an HDF5 container object. It can contain any Location.
type Group struct {
	CommonFG
}

// CreateGroup creates a new empty group and links it to a location in the file.
func (g *CommonFG) CreateGroup(name string) (*Group, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	link_flags := C.hid_t(C.H5P_DEFAULT)
	grp_c_flags := C.hid_t(C.H5P_DEFAULT)
	hid := C.H5Gcreate2(g.id, c_name, link_flags, grp_c_flags, P_DEFAULT.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	group := &Group{CommonFG{Identifier{hid}}}
	runtime.SetFinalizer(group, (*Group).finalizer)
	return group, nil
}

// CreateDataset creates a new Dataset.
func (g *CommonFG) CreateDataset(name string, dtype *Datatype, dspace *Dataspace) (*Dataset, error) {
	return createDataset(g.id, name, dtype, dspace, P_DEFAULT)
}

// CreateDatasetWith creates a new Dataset with a user-defined PropList.
func (g *CommonFG) CreateDatasetWith(name string, dtype *Datatype, dspace *Dataspace, dcpl *PropList) (*Dataset, error) {
	return createDataset(g.id, name, dtype, dspace, dcpl)
}

// CreateAttribute creates a new attribute at this location.
func (g *Group) CreateAttribute(name string, dtype *Datatype, dspace *Dataspace) (*Attribute, error) {
	return createAttribute(g.id, name, dtype, dspace, P_DEFAULT)
}

// CreateAttributeWith creates a new attribute at this location with a user-defined PropList.
func (g *Group) CreateAttributeWith(name string, dtype *Datatype, dspace *Dataspace, acpl *PropList) (*Attribute, error) {
	return createAttribute(g.id, name, dtype, dspace, acpl)
}

func (g *Group) finalizer() {
	if err := g.Close(); err != nil {
		panic(fmt.Errorf("error closing group: %s", err))
	}
}

// Close closes the Group.
func (g *Group) Close() error {
	if g.id == 0 {
		return nil
	}
	err := h5err(C.H5Gclose(g.id))
	g.id = 0
	return err
}

// OpenGroup opens an existing child group from this Group.
func (g *CommonFG) OpenGroup(name string) (*Group, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Gopen2(g.id, c_name, P_DEFAULT.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	group := &Group{CommonFG{Identifier{hid}}}
	runtime.SetFinalizer(group, (*Group).finalizer)
	return group, nil
}

// OpenDataset opens a named Dataset.
func (g *CommonFG) OpenDataset(name string) (*Dataset, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Dopen2(g.id, c_name, P_DEFAULT.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newDataset(hid), nil
}

// NumObjects returns the number of objects in the Group.
func (g *CommonFG) NumObjects() (uint, error) {
	var info C.H5G_info_t
	err := h5err(C.H5Gget_info(g.id, &info))
	return uint(info.nlinks), err
}

// ObjectNameByIndex returns the name of the object at idx.
func (g *CommonFG) ObjectNameByIndex(idx uint) (string, error) {
	cidx := C.hsize_t(idx)
	size := C.H5Lget_name_by_idx(g.id, cdot, C.H5_INDEX_NAME, C.H5_ITER_INC, cidx, nil, 0, C.H5P_DEFAULT)
	if size < 0 {
		return "", fmt.Errorf("could not get name")
	}

	name := make([]C.char, size+1)
	size = C.H5Lget_name_by_idx(g.id, cdot, C.H5_INDEX_NAME, C.H5_ITER_INC, cidx, &name[0], C.size_t(size)+1, C.H5P_DEFAULT)

	if size < 0 {
		return "", fmt.Errorf("could not get name")
	}
	return C.GoString(&name[0]), nil
}

// CreateTable creates a packet table to store fixed-length packets.
func (g *Group) CreateTable(name string, dtype *Datatype, chunkSize, compression int) (*Table, error) {
	return createTable(g.id, name, dtype, chunkSize, compression)
}

// CreateTableFrom creates a packet table to store fixed-length packets.
func (g *Group) CreateTableFrom(name string, dtype interface{}, chunkSize, compression int) (*Table, error) {
	return createTableFrom(g.id, name, dtype, chunkSize, compression)
}

// OpenTable opens an existing packet table.
func (g *Group) OpenTable(name string) (*Table, error) {
	return openTable(g.id, name)
}
