package hdf5

// #include "hdf5.h"
// #include "hdf5_hl.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"errors"
	"fmt"

	"reflect"
	"runtime"
	"unsafe"
)

// -------- The H5F API for accessing HDF5 files. ---------

// File constants
const (

	// absence of rdwr => rd-only
	F_ACC_RDONLY int = 0x0000

	// open for read and write
	F_ACC_RDWR int = 0x0001

	// Truncate file, if it already exists, erasing all data previously stored in the file.
	F_ACC_TRUNC int = 0x0002

	// Fail if file already exists.
	F_ACC_EXCL int = 0x0004

	// print debug info
	F_ACC_DEBUG int = 0x0008

	// create non-existing files
	F_ACC_CREAT int = 0x0010

	// value passed to set_elink_acc_flags to cause flags to be taken from the parent file
	F_ACC_DEFAULT int = 0xffff
)

// The difference between a single file and a set of mounted files
type Scope C.H5F_scope_t

const (

	// specified file handle only
	F_SCOPE_LOCAL Scope = 0

	// entire virtual file
	F_SCOPE_GLOBAL Scope = 1
)

// a HDF5 file
type File struct {
	id C.hid_t
}

func (f *File) h5f_finalizer() {
	err := f.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing file: %s", err))
	}
}

func new_file(id C.hid_t) *File {
	f := &File{id: id}
	runtime.SetFinalizer(f, (*File).h5f_finalizer)
	return f
}

// Creates an HDF5 file.
// hid_t H5Fcreate( const char *name, unsigned flags, hid_t fcpl_id, hid_t fapl_id )
func CreateFile(name string, flags int) (f *File, err error) {
	f = nil
	err = nil

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	// FIXME: file props
	hid := C.H5Fcreate(c_name, C.uint(flags), P_DEFAULT.id, P_DEFAULT.id)
	err = togo_err(C.herr_t(int(hid)))
	if err != nil {
		return
	}
	f = new_file(hid)
	return
}

// Opens an existing HDF5 file.
// hid_t H5Fopen( const char *name, unsigned flags, hid_t fapl_id )
func OpenFile(name string, flags int) (f *File, err error) {
	f = nil
	err = nil

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	// FIXME: file props
	hid := C.H5Fopen(c_name, C.uint(flags), P_DEFAULT.id)
	err = togo_err(C.herr_t(int(hid)))
	if err != nil {
		return
	}
	f = new_file(hid)
	return
}

// Returns a new identifier for a previously-opened HDF5 file.
func (self *File) ReOpen() (f *File, err error) {
	f = nil
	err = nil

	hid := C.H5Freopen(self.id)
	err = togo_err(C.herr_t(int(hid)))
	if err != nil {
		return
	}
	f = new_file(hid)
	return
}

// Determines whether a file is in the HDF5 format.
func IsHdf5(name string) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	o := int(C.H5Fis_hdf5(c_name))
	if o > 0 {
		return true
	}
	return false
}

// Terminates access to an HDF5 file.
func (f *File) Close() error {
	var err error = nil
	if f.id > 0 {
		err = togo_err(C.H5Fclose(f.id))
		f.id = 0
	}
	return err
}

// Flushes all buffers associated with a file to disk.
// herr_t H5Fflush(hid_t object_id, H5F_scope_t scope )
func (f *File) Flush(scope Scope) error {
	return togo_err(C.H5Fflush(f.id, C.H5F_scope_t(scope)))
}

func (f *File) Name() string {
	return getName(f.id)
}

// FIXME
// Retrieves name of file to which object belongs.
// ssize_t H5Fget_name(hid_t obj_id, char *name, size_t size )
func (f *File) FileName() string {
	sz := int(C.H5Fget_name(f.id, nil, 0)) + 1
	if sz < 0 {
		return ""
	}
	buf := string(make([]byte, sz))
	c_buf := C.CString(buf)
	defer C.free(unsafe.Pointer(c_buf))
	sz = int(C.H5Fget_name(f.id, c_buf, C.size_t(sz)))
	if sz < 0 {
		return ""
	}
	return C.GoString(c_buf)

}

// Creates a new empty group and links it to a location in the file.
func (f *File) CreateGroup(name string) (*Group, error) {
	return createGroup(f.id, name, C.H5P_DEFAULT, C.H5P_DEFAULT, C.H5P_DEFAULT)
}

func (f *File) Id() int {
	return int(f.id)
}

// Opens an existing group in a file.
// hid_t H5Gopen2( hid_t loc_id, const char * name, hid_t gapl_id )
func (f *File) OpenGroup(name string, gapl_flag int) (g *Group, err error) {
	g = nil
	err = nil

	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Gopen2(f.id, c_name, C.hid_t(gapl_flag))
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
func (f *File) OpenDataType(name string, tapl_id int) (*DataType, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Topen2(f.id, c_name, C.hid_t(tapl_id))
	err := togo_err(C.herr_t(hid))
	if err != nil {
		return nil, err
	}
	dt := &DataType{id: hid}
	runtime.SetFinalizer(dt, (*DataType).h5t_finalizer)
	return dt, err
}

// Creates a new dataset at this location.
// hid_t H5Dcreate2( hid_t loc_id, const char *name, hid_t dtype_id, hid_t space_id, hid_t lcpl_id, hid_t dcpl_id, hid_t dapl_id )
func (f *File) CreateDataSet(name string, dtype *DataType, dspace *Dataspace, dcpl *PropList) (*Dataset, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	hid := C.H5Dcreate2(f.id, c_name, dtype.id, dspace.id, P_DEFAULT.id, dcpl.id, P_DEFAULT.id)
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	dset := new_dataset(hid)
	return dset, err
}

// Opens an existing dataset.
// hid_t H5Dopen( hid_t loc_id, const char *name, hid_t dapl_id )
func (f *File) OpenDataSet(name string) (*Dataset, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Dopen2(f.id, c_name, P_DEFAULT.id)
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	dset := new_dataset(hid)
	return dset, err
}

// Creates a packet table to store fixed-length packets.
// hid_t H5PTcreate_fl( hid_t loc_id, const char * dset_name, hid_t dtype_id, hsize_t chunk_size, int compression )
func (f *File) CreateTable(name string, dtype *DataType, chunk_size, compression int) (*Table, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_chunk := C.hsize_t(chunk_size)
	c_compr := C.int(compression)
	hid := C.H5PTcreate_fl(f.id, c_name, dtype.id, c_chunk, c_compr)
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	table := new_packet_table(hid)
	return table, err
}

// Creates a packet table to store fixed-length packets.
// hid_t H5PTcreate_fl( hid_t loc_id, const char * dset_name, hid_t dtype_id, hsize_t chunk_size, int compression )
func (f *File) CreateTableFrom(name string, dtype interface{}, chunk_size, compression int) (table *Table, err error) {
	switch dt := dtype.(type) {
	case reflect.Type:
		hdf_dtype := new_dataTypeFromType(dt)
		table, err = f.CreateTable(name, hdf_dtype, chunk_size, compression)
		return

	case *DataType:
		table, err = f.CreateTable(name, dt, chunk_size, compression)
		return

	default:
		hdf_dtype := new_dataTypeFromType(reflect.TypeOf(dtype))
		table, err = f.CreateTable(name, hdf_dtype, chunk_size, compression)
		return
	}
	panic("unreachable")
	return nil, errors.New("unreachable")
}

// Opens an existing packet table.
// hid_t H5PTopen( hid_t loc_id, const char *dset_name )
func (f *File) OpenTable(name string) (*Table, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5PTopen(f.id, c_name)
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		println("===")
		return nil, err
	}
	table := new_packet_table(hid)
	return table, err
}

// EOF
