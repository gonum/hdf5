package hdf5

/*
 #cgo LDFLAGS: -lhdf5
 #include "hdf5.h"

 #include <stdlib.h>
 #include <string.h>
*/
import "C"

import (
	"os"
	"fmt"
)

// utils
type hdferror struct {
	code int
}
func (h *hdferror) String() string {
	return fmt.Sprintf("**hdf5 error** code=%d", h.code)
}

func togo_err(herr C.herr_t) os.Error {
	if herr >= C.herr_t(0) {
		return nil
	}
	return &hdferror{code:int(herr)}
}

// initialize the hdf5 library
func init() {
	err := init_hdf5()
	if err != nil {
		err_str := fmt.Sprintf("pb calling H5open(): %s", err)
		panic(err_str)
	}
}

// Initializes the HDF5 library. 
func init_hdf5() os.Error {
	return togo_err(C.H5open())
}

// Flushes all data to disk, closes all open identifiers, and cleans up memory. 
func close_hdf5() os.Error {
	return togo_err(C.H5close())
}

// Returns the HDF library release number. 
func GetLibVersion() (majnum, minnum, relnum uint, err os.Error) {
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
func GarbageCollect() os.Error {
	return togo_err(C.H5garbage_collect())
}

// constants


// EOF
