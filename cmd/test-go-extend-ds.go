package main

import (
	"hdf5"

	"fmt"
)

func main() {
	
	fname := "SDSextendible.h5"
	dsname:= "ExtendibleArray"
	NX := 10
	NY :=  5
	RANK:= 2

	dims := []int{3, 3} // dset dimensions at creation
	maxdims:= []int{hdf5.S_UNLIMITED, hdf5.S_UNLIMITED}

	//mspace := hdf5.CreateDataSpace(dims, maxdims)

	// create a new file
	f := hdf5.CreateFile(fname)

}