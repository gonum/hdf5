package hdf5

/*
 #cgo LDFLAGS: -lhdf5
 #include "hdf5.h"
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