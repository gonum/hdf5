package hdf5

// #include "hdf5.h"
// #include "hdf5_hl.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"unsafe"
)

type IDComponent struct {
	id C.hid_t
}

func (c IDComponent) Id() int {
	return int(c.id)
}

// Name returns the full name of the IDComponent
func (c *IDComponent) Name() string {
	sz := int(C.H5Iget_name(c.id, nil, 0)) + 1
	if sz < 0 {
		return ""
	}
	buf := string(make([]byte, sz))
	c_buf := C.CString(buf)
	defer C.free(unsafe.Pointer(c_buf))
	sz = int(C.H5Iget_name(c.id, c_buf, C.size_t(sz)))
	if sz < 0 {
		return ""
	}
	return C.GoString(c_buf)
}

// File returns the file associated with this IDComponent.
func (c *IDComponent) File() *File {
	fid := C.H5Iget_file_id(c.id)
	if fid < 0 {
		return nil
	}
	return &File{CommonFG{Location{IDComponent{fid}}}}
}
