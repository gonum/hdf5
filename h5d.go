package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"fmt"

	"reflect"
	"runtime"
	"unsafe"
)

type Dataset struct {
	id C.hid_t
}

func newDataset(id C.hid_t) *Dataset {
	d := &Dataset{id: id}
	runtime.SetFinalizer(d, (*Dataset).finalizer)
	return d
}

func createDataset(id C.hid_t, name string, dtype *Datatype, dspace *Dataspace, dcpl *PropList) (*Dataset, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	hid := C.H5Dcreate2(id, c_name, dtype.id, dspace.id, P_DEFAULT.id, dcpl.id, P_DEFAULT.id)
	if err := h5err(C.herr_t(int(hid))); err != nil {
		return nil, err
	}
	return newDataset(hid), nil
}

func openDataset(id C.hid_t, name string) (*Dataset, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Dopen2(id, c_name, P_DEFAULT.id)
	if err := h5err(C.herr_t(int(hid))); err != nil {
		return nil, err
	}
	return newDataset(hid), nil
}

func (s *Dataset) finalizer() {
	err := s.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing dset: %s", err))
	}
}

func (s *Dataset) Name() string {
	return getName(s.id)
}

func (s *Dataset) Id() int {
	return int(s.id)
}

func (s *Dataset) File() *File {
	return getFile(s.id)
}

// Releases and terminates access to a dataset.
// herr_t H5Dclose( hid_t space_id )
func (s *Dataset) Close() error {
	if s.id > 0 {
		err := C.H5Dclose(s.id)
		s.id = 0
		return h5err(err)
	}
	return nil
}

// Returns an identifier for a copy of the dataspace for a dataset.
// hid_t H5Dget_space(hid_t dataset_id )
func (s *Dataset) Space() *Dataspace {
	hid := C.H5Dget_space(s.id)
	if int(hid) > 0 {
		return new_dataspace(hid)
	}
	return nil
}

// Reads raw data from a dataset into a buffer.
// herr_t H5Dread(hid_t dataset_id, hid_t mem_type_id, hid_t mem_space_id, hid_t file_space_id, hid_t xfer_plist_id, void * buf )
func (s *Dataset) Read(data interface{}, dtype *Datatype) error {
	var addr uintptr
	v := reflect.ValueOf(data)

	//fmt.Printf(":: read[%s]...\n", v.Kind())
	switch v.Kind() {

	case reflect.Array:
		addr = v.UnsafeAddr()

	case reflect.Slice:
		addr = v.Pointer()

	case reflect.String:
		str := (*reflect.StringHeader)(unsafe.Pointer(v.UnsafeAddr()))
		addr = str.Data

	case reflect.Ptr:
		addr = v.Pointer()

	default:
		addr = v.UnsafeAddr()
	}

	rc := C.H5Dread(s.id, dtype.id, 0, 0, 0, unsafe.Pointer(addr))
	err := h5err(rc)
	return err
}

// Writes raw data from a buffer to a dataset.
// herr_t H5Dwrite(hid_t dataset_id, hid_t mem_type_id, hid_t mem_space_id, hid_t file_space_id, hid_t xfer_plist_id, const void * buf )
func (s *Dataset) Write(data interface{}, dtype *Datatype) error {
	var addr uintptr
	v := reflect.ValueOf(data)

	//fmt.Printf(":: write[%s]...\n", v.Kind())
	switch v.Kind() {

	case reflect.Array:
		addr = v.UnsafeAddr()

	case reflect.Slice:
		addr = v.Pointer()

	case reflect.String:
		str := (*reflect.StringHeader)(unsafe.Pointer(v.UnsafeAddr()))
		addr = str.Data

	case reflect.Ptr:
		addr = v.Pointer()

	default:
		addr = v.Pointer()
	}

	rc := C.H5Dwrite(s.id, dtype.id, 0, 0, 0, unsafe.Pointer(addr))
	err := h5err(rc)
	return err
}
