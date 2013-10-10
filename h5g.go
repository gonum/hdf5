package hdf5

// #include "hdf5.h"
// #include "hdf5_hl.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

type Group struct {
	id C.hid_t
}

func createGroup(id C.hid_t, name string, link_flags, grp_c_flags, grp_a_flags int) (*Group, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Gcreate2(id, c_name, C.hid_t(link_flags), C.hid_t(grp_c_flags), P_DEFAULT.id)
	if err := togo_err(C.herr_t(int(hid))); err != nil {
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
	if err := togo_err(C.herr_t(int(hid))); err != nil {
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
	return togo_err(C.H5Gclose(g.id))
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
func (g *Group) OpenDataType(name string, tapl_id int) (*Datatype, error) {
	return openDatatype(g.id, name, tapl_id)
}

/* Packet table methods */

// Creates a packet table to store fixed-length packets.
// hid_t H5PTcreate_fl( hid_t loc_id, const char * dset_name, hid_t dtype_id, hsize_t chunk_size, int compression )
func (g *Group) CreateTable(name string, dtype *Datatype, chunk_size, compression int) (*Table, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_chunk := C.hsize_t(chunk_size)
	c_compr := C.int(compression)
	hid := C.H5PTcreate_fl(g.id, c_name, dtype.id, c_chunk, c_compr)
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	table := new_packet_table(hid)
	return table, err
}

// Creates a packet table to store fixed-length packets.
// hid_t H5PTcreate_fl( hid_t loc_id, const char * dset_name, hid_t dtype_id, hsize_t chunk_size, int compression )
func (g *Group) CreateTableFrom(name string, dtype interface{}, chunk_size, compression int) (*Table, error) {
	switch dt := dtype.(type) {
	case reflect.Type:
		hdf_dtype := new_dataTypeFromType(dt)
		return g.CreateTable(name, hdf_dtype, chunk_size, compression)
	case *Datatype:
		return g.CreateTable(name, dt, chunk_size, compression)
	default:
		hdf_dtype := new_dataTypeFromType(reflect.TypeOf(dtype))
		return g.CreateTable(name, hdf_dtype, chunk_size, compression)
	}
}

// Opens an existing packet table.
// hid_t H5PTopen( hid_t loc_id, const char *dset_name )
func (g *Group) OpenTable(name string) (*Table, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5PTopen(g.id, c_name)
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	table := new_packet_table(hid)
	return table, err
}

// EOF
