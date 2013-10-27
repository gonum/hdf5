package hdf5

// #include "hdf5.h"
import "C"

func disableErrorPrinting() error {
	return h5err(C.H5Eset_auto(C.H5E_DEFAULT, nil, nil))
}
