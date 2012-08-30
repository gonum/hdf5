package main

import (
	"github.com/sbinet/go-hdf5/pkg/hdf5"

	"fmt"
)

func main() {

	fname := "SDSextendible.h5"
	//dsname:= "ExtendibleArray"
	//NX := 10
	//NY :=  5
	//RANK:= 2

	//dims := []int{3, 3} // dset dimensions at creation
	//maxdims:= []int{hdf5.S_UNLIMITED, hdf5.S_UNLIMITED}

	//mspace := hdf5.CreateDataSpace(dims, maxdims)

	// create a new file
	f, err := hdf5.CreateFile(fname, hdf5.F_ACC_TRUNC)
	if err != nil {
		panic(err)
	}

	fmt.Printf(":: file [%s] created\n", f.Name())

}

// EOF
