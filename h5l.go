package hdf5

// #include "hdf5.h"
// #include "hdf5_hl.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"unsafe"
)

type Container interface {
	Close() error
	Name() string
	FileName() string
	File() *File
	CreateGroup(name string) (*Group, error)
	OpenGroup(name string) (*Group, error)
	OpenDatatype(name string, tapl_id int) (*Datatype, error)
	NumObjects() (uint, error)
	ObjectNameByIndex(idx uint) (string, error)
	CreateDataset(name string, dtype *Datatype, dspace *Dataspace) (*Dataset, error)
	CreateDatasetWith(name string, dtype *Datatype, dspace *Dataspace) (*Dataset, error)
	OpenDataset(name string) (*Dataset, error)
	CreateTable(name string, dtype *Datatype, chunkSize, compression int) (*Table, error)
	CreateTableFrom(name string, dtype interface{}, chunkSize, compression int) (*Table, error)
	OpenTable(name string) (*Table, error)
}

type Location struct {
	id C.hid_t
}

func (l Location) Id() int {
	return int(l.id)
}

// Name returns the full name (path) of the Location.
// For Locations that have multiple links it attempts to return the name with
// which the Location was opened.
func (l *Location) Name() string {
	sz := int(C.H5Iget_name(l.id, nil, 0)) + 1
	if sz < 0 {
		return ""
	}
	buf := string(make([]byte, sz))
	c_buf := C.CString(buf)
	defer C.free(unsafe.Pointer(c_buf))
	sz = int(C.H5Iget_name(l.id, c_buf, C.size_t(sz)))
	if sz < 0 {
		return ""
	}
	return C.GoString(c_buf)
}

// File returns the file associated with this Location.
func (l *Location) File() *File {
	fid := C.H5Iget_file_id(l.id)
	if fid < 0 {
		return nil
	}
	return &File{Location{fid}}
}
