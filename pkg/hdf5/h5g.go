package hdf5

/*
 #cgo LDFLAGS: -lhdf5 -lhdf5_hl
 #include "hdf5.h"
 #include "hdf5_hl.h"

 #include <stdlib.h>
 #include <string.h>
*/
import "C"

import (
	"errors"
	"fmt"

	"reflect"
	"runtime"
	"unsafe"
)

type Group struct {
	id C.hid_t
}

// FIXME
// Creates a new empty group and links it to a location in the file. 
// hid_t H5Gcreate2( hid_t loc_id, const char *name, hid_t lcpl_id, hid_t gcpl_id, hid_t gapl_id ) 
func (self *Group) CreateGroup(name string, link_flags, grp_c_flags, grp_a_flags int) (g *Group, err error) {
	g = nil
	err = nil

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Gcreate2(self.id, c_name, C.hid_t(link_flags), C.hid_t(grp_c_flags), P_DEFAULT.id)
	err = togo_err(C.herr_t(int(hid)))
	if err != nil {
		return
	}
	g = &Group{id: hid}
	runtime.SetFinalizer(g, (*Group).h5g_finalizer)
	return
}

func (g *Group) h5g_finalizer() {
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

func (g *Group) Id() int {
	return int(g.id)
}

// Opens an existing group in a file.
// hid_t H5Gopen( hid_t loc_id, const char * name, hid_t gapl_id ) 
func (self *Group) OpenGroup(name string, gapl_flag int) (g *Group, err error) {
	g = nil
	err = nil

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Gopen(self.id, c_name, C.hid_t(gapl_flag))
	err = togo_err(C.herr_t(int(hid)))
	if err != nil {
		return
	}
	g = &Group{id: hid}
	runtime.SetFinalizer(g, (*Group).h5g_finalizer)
	return
}

// Opens a named datatype.
// hid_t H5Topen2( hid_t loc_id, const char * name, hid_t tapl_id ) 
func (g *Group) OpenDataType(name string, tapl_id int) (*DataType, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Topen2(g.id, c_name, C.hid_t(tapl_id))
	err := togo_err(C.herr_t(hid))
	if err != nil {
		return nil, err
	}
	dt := &DataType{id: hid}
	runtime.SetFinalizer(dt, (*DataType).h5t_finalizer)
	return dt, err
}

// Creates a packet table to store fixed-length packets.
// hid_t H5PTcreate_fl( hid_t loc_id, const char * dset_name, hid_t dtype_id, hsize_t chunk_size, int compression )
func (g *Group) CreateTable(name string, dtype *DataType, chunk_size, compression int) (*Table, error) {
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

	case *DataType:
		return g.CreateTable(name, dt, chunk_size, compression)

	default:
		hdf_dtype := new_dataTypeFromType(reflect.TypeOf(dtype))
		return g.CreateTable(name, hdf_dtype, chunk_size, compression)
	}
	panic("unreachable")
	return nil, errors.New("unreachable")
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
