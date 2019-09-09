// license that can be found in the LICENSE file.

package hdf5

// #include "hdf5.h"
// #include "hdf5_hl.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"errors"
	"image"
	"unsafe"
)

// newImage takes a image object and convert it to a hdf5 format and write to the id node
func newImage(id C.hid_t, name string, img image.Image) error {
	if img == nil {
		return errors.New("nil image!")
	}
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	var c_width, c_height C.hsize_t
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	var gbuf []uint8
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gbuf = append(gbuf, uint8(r>>8))
			gbuf = append(gbuf, uint8(g>>8))
			gbuf = append(gbuf, uint8(b>>8))
		}
	}
	c_width = C.hsize_t(width)
	c_height = C.hsize_t(height)
	c_image := (*C.uchar)(unsafe.Pointer(&gbuf[0]))

	status := C.H5IMmake_image_24bit(id, c_name, c_width, c_height, C.CString("INTERLACE_PIXEL"), c_image)
	if status < 0 {
		return errors.New("Failed to create HDF5 true color image")
	}
	return nil
}
