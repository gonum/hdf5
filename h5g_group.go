// Copyright ©2017 The Gonum Authors. All rights reserved.
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

// CreateGroup creates and returns a new empty group and links it to a location
// in the file. The returned group must be closed by the user when it is no
// longer needed.
func (g *CommonFG) CreateGroup(name string) (*Group, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	linkFlags := C.hid_t(C.H5P_DEFAULT)
	grpCFlags := C.hid_t(C.H5P_DEFAULT)
	hid := C.H5Gcreate2(g.id, cName, linkFlags, grpCFlags, P_DEFAULT.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	group := &Group{CommonFG{Identifier{hid}}}
	return group, nil
}

// CreateDataset creates a new Dataset. The returned dataset must be
// closed by the user when it is no longer needed.
func (g *CommonFG) CreateDataset(name string, typ *Datatype, spc *Dataspace) (*Dataset, error) {
	return createDataset(g.id, name, typ, spc, P_DEFAULT)
}

// CreateDatasetWith creates a new Dataset with a user-defined PropList.
// The returned dataset must be closed by the user when it is no longer needed.
func (g *CommonFG) CreateDatasetWith(name string, typ *Datatype, spc *Dataspace, dcpl *PropList) (*Dataset, error) {
	return createDataset(g.id, name, typ, spc, dcpl)
}

// CreateAttribute creates a new attribute at this location. The returned
// attribute must be closed by the user when it is no longer needed.
func (g *Group) CreateAttribute(name string, typ *Datatype, spc *Dataspace) (*Attribute, error) {
	return createAttribute(g.id, name, typ, spc, P_DEFAULT)
}

// CreateAttributeWith creates a new attribute at this location with a user-defined
// PropList. The returned dataset must be closed by the user when it is no longer needed.
func (g *Group) CreateAttributeWith(name string, typ *Datatype, spc *Dataspace, acpl *PropList) (*Attribute, error) {
	return createAttribute(g.id, name, typ, spc, acpl)
}

// Close closes the Group.
func (g *Group) Close() error {
	return g.closeWith(h5gclose)
}

func h5gclose(id C.hid_t) C.herr_t {
	return C.H5Gclose(id)
}

// OpenGroup opens and returns an existing child group from this Group.
// The returned group must be closed by the user when it is no longer needed.
func (g *CommonFG) OpenGroup(name string) (*Group, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	hid := C.H5Gopen2(g.id, cName, P_DEFAULT.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	group := &Group{CommonFG{Identifier{hid}}}
	return group, nil
}

// OpenDataset opens and returns a named Dataset. The returned
// dataset must be closed by the user when it is no longer needed.
func (g *CommonFG) OpenDataset(name string) (*Dataset, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	hid := C.H5Dopen2(g.id, cName, P_DEFAULT.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newDataset(hid, nil), nil
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
// The returned table must be closed by the user when it is no longer needed.
func (g *Group) CreateTable(name string, typ *Datatype, chunkSize, compression int) (*Table, error) {
	return createTable(g.id, name, typ, chunkSize, compression)
}

// CreateTableFrom creates a packet table to store fixed-length packets.
// The returned table must be closed by the user when it is no longer needed.
func (g *Group) CreateTableFrom(name string, typ interface{}, chunkSize, compression int) (*Table, error) {
	return createTableFrom(g.id, name, typ, chunkSize, compression)
}

// OpenTable opens an existing packet table. The returned table must be
// closed by the user when it is no longer needed.
func (g *Group) OpenTable(name string) (*Table, error) {
	return openTable(g.id, name)
}
