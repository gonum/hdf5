package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"fmt"
)

// initialize the hdf5 library
func init() {
	err := h5err(C.H5open())
	if err != nil {
		err_str := fmt.Sprintf("pb calling H5open(): %s", err)
		panic(err_str)
	}
}

// utils
type hdferror struct {
	code int
}

func (h *hdferror) Error() string {
	return fmt.Sprintf("**hdf5 error** code=%d", h.code)
}

func h5err(herr C.herr_t) error {
	if herr >= C.herr_t(0) {
		return nil
	}
	return &hdferror{code: int(herr)}
}

// Close flushes all data to disk, closes all open identifiers, and cleans up memory.
// It should generally be called before your application exits.
func Close() error {
	return h5err(C.H5close())
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
	err = h5err(herr)
	if err == nil {
		majnum = uint(c_majnum)
		minnum = uint(c_minnum)
		relnum = uint(c_relnum)
	}
	return
}

// Garbage collects on all free-lists of all types.
func GarbageCollect() error {
	return h5err(C.H5garbage_collect())
}

type Object interface {
	Name() string
	Id() int
	File() *File
}
