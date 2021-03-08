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

// GType describes the type of an object inside a Group or File.
type GType int

const (
	H5G_UNKNOWN GType = C.H5G_UNKNOWN // Unknown object type
	H5G_GROUP   GType = C.H5G_GROUP   // Object is a group
	H5G_DATASET GType = C.H5G_DATASET // Object is a dataset
	H5G_TYPE    GType = C.H5G_TYPE    // Object is a named data type
	H5G_LINK    GType = C.H5G_LINK    // Object is a symbolic link
	H5G_UDLINK  GType = C.H5G_UDLINK  // Object is a user-defined link
)

func (typ GType) String() string {
	switch typ {
	case H5G_UNKNOWN:
		return "unknown"
	case H5G_GROUP:
		return "group"
	case H5G_DATASET:
		return "dataset"
	case H5G_TYPE:
		return "type"
	case H5G_LINK:
		return "link"
	case H5G_UDLINK:
		return "udlink"
	default:
		return fmt.Sprintf("GType(%d)", int(typ))
	}
}

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
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	link_flags := C.hid_t(C.H5P_DEFAULT)
	grp_c_flags := C.hid_t(C.H5P_DEFAULT)
	hid := C.H5Gcreate2(g.id, c_name, link_flags, grp_c_flags, P_DEFAULT.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	group := &Group{CommonFG{Identifier{hid}}}
	return group, nil
}

// CreateDataset creates a new Dataset. The returned dataset must be
// closed by the user when it is no longer needed.
func (g *CommonFG) CreateDataset(name string, dtype *Datatype, dspace *Dataspace) (*Dataset, error) {
	return createDataset(g.id, name, dtype, dspace, P_DEFAULT)
}

// CreateDatasetWith creates a new Dataset with a user-defined PropList.
// The returned dataset must be closed by the user when it is no longer needed.
func (g *CommonFG) CreateDatasetWith(name string, dtype *Datatype, dspace *Dataspace, dcpl *PropList) (*Dataset, error) {
	return createDataset(g.id, name, dtype, dspace, dcpl)
}

// CreateAttribute creates a new attribute at this location. The returned
// attribute must be closed by the user when it is no longer needed.
func (g *Group) CreateAttribute(name string, dtype *Datatype, dspace *Dataspace) (*Attribute, error) {
	return createAttribute(g.id, name, dtype, dspace, P_DEFAULT)
}

// CreateAttributeWith creates a new attribute at this location with a user-defined
// PropList. The returned dataset must be closed by the user when it is no longer needed.
func (g *Group) CreateAttributeWith(name string, dtype *Datatype, dspace *Dataspace, acpl *PropList) (*Attribute, error) {
	return createAttribute(g.id, name, dtype, dspace, acpl)
}

// Opens an existing attribute. The returned attribute must be closed
// by the user when it is no longer needed.
func (g *Group) OpenAttribute(name string) (*Attribute, error) {
	return openAttribute(g.id, name)
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
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Gopen2(g.id, c_name, P_DEFAULT.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	group := &Group{CommonFG{Identifier{hid}}}
	return group, nil
}

// OpenDataset opens and returns a named Dataset. The returned
// dataset must be closed by the user when it is no longer needed.
func (g *CommonFG) OpenDataset(name string) (*Dataset, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Dopen2(g.id, c_name, P_DEFAULT.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newDataset(hid, nil), nil
}

// OpenDatasetWith opens and returns a named Dataset with a user-defined PropList.
// The returned dataset must be closed by the user when it is no longer needed.
func (g *CommonFG) OpenDatasetWith(name string, dapl *PropList) (*Dataset, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Dopen2(g.id, c_name, dapl.id)
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

// ObjectTypeByIndex returns the type of the object at idx.
func (g *CommonFG) ObjectTypeByIndex(idx uint) (GType, error) {
	cidx := C.hsize_t(idx)
	gtyp := GType(C.H5Gget_objtype_by_idx(g.id, cidx))
	if gtyp < H5G_GROUP {
		return H5G_UNKNOWN, fmt.Errorf("could not get object type")
	}

	return gtyp, nil
}

// CreateTable creates a packet table to store fixed-length packets.
// The returned table must be closed by the user when it is no longer needed.
func (g *Group) CreateTable(name string, dtype *Datatype, chunkSize, compression int) (*Table, error) {
	return createTable(g.id, name, dtype, chunkSize, compression)
}

// CreateTableFrom creates a packet table to store fixed-length packets.
// The returned table must be closed by the user when it is no longer needed.
func (g *Group) CreateTableFrom(name string, dtype interface{}, chunkSize, compression int) (*Table, error) {
	return createTableFrom(g.id, name, dtype, chunkSize, compression)
}

// OpenTable opens an existing packet table. The returned table must be
// closed by the user when it is no longer needed.
func (g *Group) OpenTable(name string) (*Table, error) {
	return openTable(g.id, name)
}

// LinkExists returns whether a link with the specified name exists in the group.
func (g *CommonFG) LinkExists(name string) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	return C.H5Lexists(g.id, c_name, 0) > 0
}
