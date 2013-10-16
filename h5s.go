package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"errors"
	"fmt"

	"runtime"
	"unsafe"
)

type Dataspace struct {
	id C.hid_t
}

type SpaceClass C.H5S_class_t

const (
	S_NO_CLASS SpaceClass = -1 // error
	S_SCALAR   SpaceClass = 0  // scalar variable
	S_SIMPLE   SpaceClass = 1  // simple data space
	S_NULL     SpaceClass = 2  // null data space
)

func newDataspace(id C.hid_t) *Dataspace {
	ds := &Dataspace{id: id}
	runtime.SetFinalizer(ds, (*Dataspace).finalizer)
	return ds
}

// Creates a new dataspace of a specified type.
// hid_t H5Screate( H5S_class_t type )
func CreateDataSpace(class SpaceClass) (*Dataspace, error) {
	hid := C.H5Screate(C.H5S_class_t(class))
	err := h5err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	ds := newDataspace(hid)
	return ds, nil
}

func (s *Dataspace) finalizer() {
	err := s.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing dspace: %s", err))
	}
}

// Creates an exact copy of a dataspace.
// hid_t H5Scopy( hid_t space_id )
func (s *Dataspace) Copy() (*Dataspace, error) {
	hid := C.H5Scopy(s.id)
	err := h5err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	o := newDataspace(hid)
	return o, err
}

// Releases and terminates access to a dataspace.
// herr_t H5Sclose( hid_t space_id )
func (s *Dataspace) Close() error {
	err := C.H5Sclose(s.id)
	return h5err(err)
}

func (s *Dataspace) Id() int {
	return int(s.id)
}

func (s *Dataspace) Name() string {
	return getName(s.id)
}

// Creates a new simple dataspace and opens it for access.
func CreateSimpleDataSpace(dims, maxDims []uint) (*Dataspace, error) {
	var c_dims, c_maxdims *C.hsize_t

	rank := C.int(0)
	if dims != nil {
		rank = C.int(len(dims))
		c_dims = (*C.hsize_t)(unsafe.Pointer(&dims[0]))

	}
	if maxDims != nil {
		rank = C.int(len(maxDims))
		c_maxdims = (*C.hsize_t)(unsafe.Pointer(&maxDims[0]))

	}
	if len(dims) != len(maxDims) && (dims != nil && maxDims != nil) {
		return nil, errors.New("lengths of dims and maxDims do not match")
	}

	hid := C.H5Screate_simple(rank, c_dims, c_maxdims)
	if hid < 0 {
		return nil, fmt.Errorf("failed to create dataspace")
	}
	return newDataspace(hid), nil
}

// IsSimple determines whether a dataspace is a simple dataspace.
func (s *Dataspace) IsSimple() bool {
	return int(C.H5Sis_simple(s.id)) > 0
}

// SetOffset sets the offset of a simple dataspace.
func (s *Dataspace) SetOffset(offset []int) error {
	rank := len(offset)
	if rank == 0 || offset == nil {
		err := C.H5Soffset_simple(s.id, nil)
		return h5err(err)
	}
	if rank != s.SimpleExtentNDims() {
		err := errors.New("size of offset does not match extent")
		return err
	}

	c_offset := (*C.hssize_t)(unsafe.Pointer(&offset[0]))
	err := C.H5Soffset_simple(s.id, c_offset)
	return h5err(err)
}

// Retrieves dataspace dimension size and maximum size.
func (s *Dataspace) SimpleExtentDims() (dims, maxdims []int, err error) {
	rank := s.SimpleExtentNDims()
	dims = make([]int, rank)
	maxdims = make([]int, rank)

	c_dims := (*C.hsize_t)(unsafe.Pointer(&dims[0]))
	c_maxdims := (*C.hsize_t)(unsafe.Pointer(&maxdims[0]))
	rc := C.H5Sget_simple_extent_dims(s.id, c_dims, c_maxdims)
	err = h5err(C.herr_t(rc))
	return
}

// Determines the dimensionality of a dataspace.
// int H5Sget_simple_extent_ndims( hid_t space_id )
func (s *Dataspace) SimpleExtentNDims() int {
	return int(C.H5Sget_simple_extent_ndims(s.id))
}

// Determines the number of elements in a dataspace.
// hssize_t H5Sget_simple_extent_npoints( hid_t space_id )
func (s *Dataspace) SimpleExtentNPoints() int {
	return int(C.H5Sget_simple_extent_npoints(s.id))
}

// Determines the current class of a dataspace.
// H5S_class_t H5Sget_simple_extent_type( hid_t space_id )
func (s *Dataspace) SimpleEventType() SpaceClass {
	return SpaceClass(C.H5Sget_simple_extent_type(s.id))
}
