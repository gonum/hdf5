package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

type CString struct {
	cstr *C.char
}

func NewCString(s string) *CString {
	c_s := &CString{cstr: C.CString(s)}
	runtime.SetFinalizer(c_s, (*CString).finalizer)
	return c_s
}

func (s *CString) finalizer() {
	C.free(unsafe.Pointer(s.cstr))
}

// initialize the hdf5 library
func init() {
	err := togo_err(C.H5open())
	if err != nil {
		err_str := fmt.Sprintf("pb calling H5open(): %s", err)
		panic(err_str)
	}
}

// Close flushes all data to disk, closes all open identifiers, and cleans up memory.
// It should generally be called before your application exits.
func Close() error {
	return togo_err(C.H5close())
}

// Returns the HDF library release number.
func GetLibVersion() (majnum, minnum, relnum uint, err error) {
	err = nil
	majnum = 0
	minnum = 0
	relnum = 0

	c_majnum := C.uint(majnum)
	c_minnum := C.uint(minnum)
	c_relnum := C.uint(relnum)

	herr := C.H5get_libversion(&c_majnum, &c_minnum, &c_relnum)
	err = togo_err(herr)
	if err == nil {
		majnum = uint(c_majnum)
		minnum = uint(c_minnum)
		relnum = uint(c_relnum)
	}
	return
}

// Garbage collects on all free-lists of all types.
func GarbageCollect() error {
	return togo_err(C.H5garbage_collect())
}

type Object interface {
	Name() string
	Id() int
	File() *File
}
