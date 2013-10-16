package hdf5

// #cgo LDFLAGS: -lhdf5 -lhdf5_hl
// #cgo darwin CFLAGS: -I/opt/local/include
// #cgo darwin LDFLAGS: -L/opt/local/lib
// #include "hdf5.h"
import "C"
