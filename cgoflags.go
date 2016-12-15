package hdf5

// #cgo LDFLAGS: -lhdf5 -lhdf5_hl
// #cgo darwin CFLAGS: -I/usr/local/include
// #cgo darwin LDFLAGS: -L/usr/local/lib
// #cgo linux CFLAGS: -I/usr/local/include, -I/usr/lib/x86_64-linux-gnu/hdf5/serial/include
// #cgo linux LDFLAGS: -L/usr/local/lib, -L/usr/lib/x86_64-linux-gnu/hdf5/serial/
// #include "hdf5.h"
import "C"
