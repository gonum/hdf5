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
	Location
}

type SpaceClass C.H5S_class_t

const (
	S_NO_CLASS SpaceClass = -1 // error
	S_SCALAR   SpaceClass = 0  // scalar variable
	S_SIMPLE   SpaceClass = 1  // simple data space
	S_NULL     SpaceClass = 2  // null data space
)

func newDataspace(id C.hid_t) *Dataspace {
	ds := &Dataspace{Location{id}}
	runtime.SetFinalizer(ds, (*Dataspace).finalizer)
	return ds
}

// CreateDataspace creates a new dataspace of a specified type.
func CreateDataspace(class SpaceClass) (*Dataspace, error) {
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

// Copy creates an exact copy of a dataspace.
func (s *Dataspace) Copy() (*Dataspace, error) {
	hid := C.H5Scopy(s.id)
	err := h5err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	o := newDataspace(hid)
	return o, err
}

// Close releases and terminates access to a dataspace.
func (s *Dataspace) Close() error {
	err := C.H5Sclose(s.id)
	return h5err(err)
}

func (s *Dataspace) Id() int {
	return int(s.id)
}

// CreateSimpleDataspace creates a new simple dataspace and opens it for access.
func CreateSimpleDataspace(dims, maxDims []uint) (*Dataspace, error) {
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

// IsSimple returns whether a dataspace is a simple dataspace.
func (s *Dataspace) IsSimple() bool {
	return int(C.H5Sis_simple(s.id)) > 0
}

// SetOffset sets the offset of a simple dataspace.
func (s *Dataspace) SetOffset(offset []uint) error {
	rank := len(offset)
	if rank == 0 {
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

// SelectHyperslab creates a subset of the data space.
func (s *Dataspace) SelectHyperslab(offset, stride, count, block []uint) error {
	rank := len(offset)
	if rank == 0 {
		err := C.H5Soffset_simple(s.id, nil)
		return h5err(err)
	}
	if rank != s.SimpleExtentNDims() {
		err := errors.New("size of offset does not match extent")
		return err
	}

	var c_offset, c_stride, c_count, c_block *C.hsize_t
	if offset != nil {
		c_offset = (*C.hsize_t)(unsafe.Pointer(&offset[0]))
	}
	if stride != nil {
		c_stride = (*C.hsize_t)(unsafe.Pointer(&stride[0]))
	}
	if count != nil {
		c_count = (*C.hsize_t)(unsafe.Pointer(&count[0]))
	}
	if block != nil {
		c_block = (*C.hsize_t)(unsafe.Pointer(&block[0]))
	}
	err := C.H5Sselect_hyperslab(s.id, C.H5S_SELECT_SET, c_offset, c_stride, c_count, c_block)
	return h5err(err)
}

// SimpleExtentDims returns dataspace dimension size and maximum size.
func (s *Dataspace) SimpleExtentDims() (dims, maxdims []uint, err error) {
	rank := s.SimpleExtentNDims()
	dims = make([]uint, rank)
	maxdims = make([]uint, rank)

	c_dims := (*C.hsize_t)(unsafe.Pointer(&dims[0]))
	c_maxdims := (*C.hsize_t)(unsafe.Pointer(&maxdims[0]))
	rc := C.H5Sget_simple_extent_dims(s.id, c_dims, c_maxdims)
	err = h5err(C.herr_t(rc))
	return
}

// SimpleExtentNDims returns the dimensionality of a dataspace.
func (s *Dataspace) SimpleExtentNDims() int {
	return int(C.H5Sget_simple_extent_ndims(s.id))
}

// SimpleExtentNPoints returns the number of elements in a dataspace.
func (s *Dataspace) SimpleExtentNPoints() int {
	return int(C.H5Sget_simple_extent_npoints(s.id))
}

// SimpleExtentType returns the current class of a dataspace.
func (s *Dataspace) SimpleExtentType() SpaceClass {
	return SpaceClass(C.H5Sget_simple_extent_type(s.id))
}
