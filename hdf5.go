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

type Version struct {
	Major   uint
	Minor   uint
	Release uint
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Release)
}

// LibVersion returns version information for the HDF5 library.
func LibVersion() (Version, error) {
	var maj, min, rel C.uint
	var v Version
	err := h5err(C.H5get_libversion(&maj, &min, &rel))
	if err == nil {
		v.Major = uint(maj)
		v.Minor = uint(min)
		v.Release = uint(rel)
	}
	return v, err
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
