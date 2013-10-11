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

type Group struct {
	id C.hid_t
}

func numObjects(id C.hid_t) (uint, error) {
	var info C.H5G_info_t
	err := h5err(C.H5Gget_info(id, &info))
	return uint(info.nlinks), err
}

func objectNameByIndex(id C.hid_t, idx uint) (string, error) {
	cidx := C.hsize_t(idx)
	size := C.H5Lget_name_by_idx(id, cdot, C.H5_INDEX_NAME, C.H5_ITER_INC, cidx, nil, 0, C.H5P_DEFAULT)
	if size < 0 {
		return "", fmt.Errorf("could not get name")
	}

	name := make([]C.char, size+1)
	size = C.H5Lget_name_by_idx(id, cdot, C.H5_INDEX_NAME, C.H5_ITER_INC, cidx, &name[0], C.size_t(size)+1, C.H5P_DEFAULT)

	if size < 0 {
		return "", fmt.Errorf("could not get name")
	}
	return C.GoString(&name[0]), nil
}

func createGroup(id C.hid_t, name string, link_flags, grp_c_flags, grp_a_flags int) (*Group, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Gcreate2(id, c_name, C.hid_t(link_flags), C.hid_t(grp_c_flags), P_DEFAULT.id)
	if err := h5err(C.herr_t(int(hid))); err != nil {
		return nil, err
	}
	g := &Group{id: hid}
	runtime.SetFinalizer(g, (*Group).finalizer)
	return g, nil
}

func openGroup(id C.hid_t, name string, gapl_flag C.hid_t) (*Group, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Gopen2(id, c_name, gapl_flag)
	if err := h5err(C.herr_t(int(hid))); err != nil {
		return nil, err
	}
	g := &Group{id: hid}
	runtime.SetFinalizer(g, (*Group).finalizer)
	return g, nil
}

// FIXME
// Creates a new empty group and links it to a location in the file.
func (g *Group) CreateGroup(name string, link_flags, grp_c_flags, grp_a_flags int) (*Group, error) {
	return createGroup(g.id, name, C.H5P_DEFAULT, C.H5P_DEFAULT, C.H5P_DEFAULT)
}

func (g *Group) CreateDataset(name string, dtype *Datatype, dspace *Dataspace, dcpl *PropList) (*Dataset, error) {
	return createDataset(g.id, name, dtype, dspace, dcpl)
}

func (g *Group) finalizer() {
	err := g.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing group: %s", err))
	}
}

// Closes the specified group.
// herr_t H5Gclose(hid_t group_id)
func (g *Group) Close() error {
	return h5err(C.H5Gclose(g.id))
}

func (g *Group) Name() string {
	return getName(g.id)
}

func (g *Group) Id() int {
	return int(g.id)
}

func (g *Group) File() *File {
	return getFile(g.id)
}

// Opens an existing group in a file.
func (g *Group) OpenGroup(name string) (*Group, error) {
	return openGroup(g.id, name, P_DEFAULT.id)
}

func (g *Group) OpenDataset(name string) (*Dataset, error) {
	return openDataset(g.id, name)
}

// Opens a named datatype.
// hid_t H5Topen2( hid_t loc_id, const char * name, hid_t tapl_id )
func (g *Group) OpenDatatype(name string, tapl_id int) (*Datatype, error) {
	return openDatatype(g.id, name, tapl_id)
}

func (g *Group) NumObjects() (uint, error) {
	return numObjects(g.id)
}

func (g *Group) ObjectNameByIndex(idx uint) (string, error) {
	return objectNameByIndex(g.id, idx)
}

// Creates a packet table to store fixed-length packets.
func (g *Group) CreateTable(name string, dtype *Datatype, chunkSize, compression int) (*Table, error) {
	return createTable(g.id, name, dtype, chunkSize, compression)
}

// Creates a packet table to store fixed-length packets.
func (g *Group) CreateTableFrom(name string, dtype interface{}, chunkSize, compression int) (*Table, error) {
	return createTableFrom(g.id, name, dtype, chunkSize, compression)
}

// Opens an existing packet table.
func (g *Group) OpenTable(name string) (*Table, error) {
	return openTable(g.id, name)
}
