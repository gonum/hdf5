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
	"image/color"
	"unsafe"
)

// newImage takes a image object and convert it to a hdf5 format and write to the id node
func newImage(id C.hid_t, name string, img image.Image) error {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	var c_width, c_height C.hsize_t
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	var buf []byte
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			buf = append(buf, byte(r>>8), byte(g>>8), byte(b>>8))
		}
	}
	c_width = C.hsize_t(width)
	c_height = C.hsize_t(height)
	c_image := (*C.uchar)(unsafe.Pointer(&buf[0]))

	status := C.H5IMmake_image_24bit(id, c_name, c_width, c_height, C.CString("INTERLACE_PIXEL"), c_image)
	if status < 0 {
		return errors.New("hdf5: failed to create true color image")
	}
	return nil
}

//
func getImage(id C.hid_t, name string) (image.Image, error) {
	//TODO(zyc-sudo) Should handle interlace and npal better, yet these two are not needed for simple image read and write.
	var width, height, planes C.hsize_t
	var npals C.hssize_t
	var interlace C.char
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	rc := C.H5IMget_image_info(id, c_name, &width, &height, &planes, &interlace, &npals)
	err := h5err(rc)
	if err != nil {
		return nil, err
	}
	gbuf := make([]uint8, width*height*planes)
	c_image := (*C.uchar)(unsafe.Pointer(&gbuf[0]))
	rc = C.H5IMread_image(id, c_name, c_image)

	err = h5err(rc)
	if err != nil {
		return nil, err
	}

	g_width := int(width)
	g_height := int(height)
	img := image.NewRGBA(image.Rect(0, 0, g_width, g_height))
	for y := 0; y < g_height; y++ {
		for x := 0; x < g_width; x++ {
			img.Set(x, y, color.RGBA{gbuf[y*g_width*3+x*3], gbuf[y*g_width*3+x*3+1], gbuf[y*g_width*3+x*3+2], 255})
		}
	}
	return img, err
}
