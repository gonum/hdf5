// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"gonum.org/v1/hdf5"
)

func main() {

	fname := "SDSextendible.h5"
	//dsname:= "ExtendibleArray"
	//NX := 10
	//NY :=  5
	//RANK:= 2

	//dims := []int{3, 3} // dset dimensions at creation
	//maxDims:= []int{hdf5.S_UNLIMITED, hdf5.S_UNLIMITED}

	//mSpace := hdf5.CreateDataspace(dims, maxDims)

	// create a new file
	f, err := hdf5.CreateFile(fname, hdf5.F_ACC_TRUNC)
	if err != nil {
		panic(err)
	}

	fmt.Printf(":: file [%s] created\n", f.Name())

}
