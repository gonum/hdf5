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

type IType C.H5I_type_t

const (
	FILE      IType = C.H5I_FILE
	GROUP     IType = C.H5I_GROUP
	DATATYPE  IType = C.H5I_DATATYPE
	DATASPACE IType = C.H5I_DATASPACE
	DATASET   IType = C.H5I_DATASET
	ATTRIBUTE IType = C.H5I_ATTR
	BAD_ID    IType = C.H5I_BADID
)

func (typ IType) String() string {
	switch typ {
	case FILE:
		return "file"
	case GROUP:
		return "group"
	case DATATYPE:
		return "datatype"
	case DATASPACE:
		return "dataspace"
	case DATASET:
		return "dataset"
	case ATTRIBUTE:
		return "attribute"
	case BAD_ID:
		return "bad_id"
	default:
		return fmt.Sprintf("IType=%d", int(typ))
	}
}

// Identifier is a simple wrapper around a C hid_t. It has basic methods
// which apply to every type in the go-hdf5 API.
type Identifier struct {
	id C.hid_t
}

// ID returns the integer value of an identifier.
func (i Identifier) ID() int64 {
	return int64(i.id)
}

// Name returns the full name of the Identifier
func (i Identifier) Name() string {
	sz := int(C.H5Iget_name(i.id, nil, 0)) + 1
	if sz < 0 {
		return ""
	}
	buf := string(make([]byte, sz))
	c_buf := C.CString(buf)
	defer C.free(unsafe.Pointer(c_buf))
	sz = int(C.H5Iget_name(i.id, c_buf, C.size_t(sz)))
	if sz < 0 {
		return ""
	}
	return C.GoString(c_buf)
}

// File returns the file associated with this Identifier.
func (i Identifier) File() *File {
	fid := C.H5Iget_file_id(i.id)
	if fid < 0 {
		return nil
	}
	return &File{CommonFG{Identifier{fid}}}
}

// Type returns the type of the identifier.
func (i Identifier) Type() IType {
	return IType(C.H5Iget_type(i.id))
}

func (i *Identifier) closeWith(fn func(C.hid_t) C.herr_t) error {
	if i.id == 0 {
		return nil
	}
	err := h5err(fn(i.id))
	i.id = 0
	return err
}
