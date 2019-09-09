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

// File constants
const (
	F_ACC_RDONLY  int = 0x0000 // absence of rdwr => rd-only
	F_ACC_RDWR    int = 0x0001 // open for read and write
	F_ACC_TRUNC   int = 0x0002 // Truncate file, if it already exists, erasing all data previously stored in the file.
	F_ACC_EXCL    int = 0x0004 // Fail if file already exists.
	F_ACC_DEBUG   int = 0x0008 // print debug info
	F_ACC_CREAT   int = 0x0010 // create non-existing files
	F_ACC_DEFAULT int = 0xffff // value passed to set_elink_acc_flags to cause flags to be taken from the parent file
)

// The difference between a single file and a set of mounted files.
type Scope C.H5F_scope_t

const (
	F_SCOPE_LOCAL  Scope = 0 // specified file handle only.
	F_SCOPE_GLOBAL Scope = 1 // entire virtual file.
)

// a HDF5 file
type File struct {
	CommonFG
}

func newFile(id C.hid_t) *File {
	return &File{CommonFG{Identifier{id}}}
}

// Creates an HDF5 file.
func CreateFile(name string, flags int) (*File, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	// FIXME: file props
	hid := C.H5Fcreate(c_name, C.uint(flags), P_DEFAULT.id, P_DEFAULT.id)
	if err := checkID(hid); err != nil {
		return nil, fmt.Errorf("error creating hdf5 file: %s", err)
	}
	return newFile(hid), nil
}

// Open opens and returns an an existing HDF5 file. The returned
// file must be closed by the user when it is no longer needed.
func OpenFile(name string, flags int) (*File, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	// FIXME: file props
	hid := C.H5Fopen(c_name, C.uint(flags), P_DEFAULT.id)
	if err := checkID(hid); err != nil {
		return nil, fmt.Errorf("error opening hdf5 file: %s", err)
	}
	return newFile(hid), nil
}

// ReOpen returns a new identifier for a previously-opened HDF5 file.
// The returned file must be closed by the user when it is no longer needed.
func (f *File) ReOpen() (*File, error) {
	hid := C.H5Freopen(f.id)
	if err := checkID(hid); err != nil {
		return nil, fmt.Errorf("error reopening hdf5 file: %s", err)
	}
	return newFile(hid), nil
}

// IsHDF5 Determines whether a file is in the HDF5 format.
func IsHDF5(name string) bool {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	return C.H5Fis_hdf5(c_name) > 0
}

// Close closes the file.
func (f *File) Close() error {
	return f.closeWith(h5fclose)
}

func h5fclose(id C.hid_t) C.herr_t {
	return C.H5Fclose(id)
}

// Flushes all buffers associated with a file to disk.
func (f *File) Flush(scope Scope) error {
	// herr_t H5Fflush(hid_t object_id, H5F_scope_t scope )
	return h5err(C.H5Fflush(f.id, C.H5F_scope_t(scope)))
}

// FIXME
// Retrieves name of file to which object belongs.
func (f *File) FileName() string {
	// ssize_t H5Fget_name(hid_t obj_id, char *name, size_t size )
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

var cdot = C.CString(".")

// Creates a packet table to store fixed-length packets. The returned
// table must be closed by the user when it is no longer needed.
func (f *File) CreateTable(name string, dtype *Datatype, chunkSize, compression int) (*Table, error) {
	// hid_t H5PTcreate_fl( hid_t loc_id, const char * dset_name, hid_t dtype_id, hsize_t chunk_size, int compression )
	return createTable(f.id, name, dtype, chunkSize, compression)
}

// Creates a packet table to store fixed-length packets. The returned
// table must be closed by the user when it is no longer needed.
func (f *File) CreateTableFrom(name string, dtype interface{}, chunkSize, compression int) (*Table, error) {
	// hid_t H5PTcreate_fl( hid_t loc_id, const char * dset_name, hid_t dtype_id, hsize_t chunk_size, int compression )
	return createTableFrom(f.id, name, dtype, chunkSize, compression)
}

// Opens an existing packet table. The returned table must be closed
// by the user when it is no longer needed.
func (f *File) OpenTable(name string) (*Table, error) {
	// hid_t H5PTopen( hid_t loc_id, const char *dset_name )
	return openTable(f.id, name)
}
