package hdf5

// #include "hdf5.h"
// #include "hdf5_hl.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"unsafe"
)

// getName returns the name of the object identified by id.
// For objects that have multiple links it attempts to return the name with
// which the object was opened.
//
// See the documentation of H5Iget_name for more details.
func getName(id C.hid_t) string {
	sz := int(C.H5Iget_name(id, nil, 0)) + 1
	if sz < 0 {
		return ""
	}
	buf := string(make([]byte, sz))
	c_buf := C.CString(buf)
	defer C.free(unsafe.Pointer(c_buf))
	sz = int(C.H5Iget_name(id, c_buf, C.size_t(sz)))
	if sz < 0 {
		return ""
	}
	return C.GoString(c_buf)
}

// getFile returns an open File with which the object identified by id is associated.
// Returns nil if the file could not be opened.
func getFile(id C.hid_t) *File {
	fid := C.H5Iget_file_id(id)
	if fid < 0 {
		return nil
	}
	return &File{fid}
}
