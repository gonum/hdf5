// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"gonum.org/v1/hdf5"
)

const (
	fname  string = "SDScompound.h5"
	dsname string = "ArrayOfStructures"
	mbr1   string = "A_name"
	mbr2   string = "B_name"
	mbr3   string = "C_name"
	length uint   = 10
	rank   int    = 1
)

type s1Type struct {
	a int
	b float32
	c float64
	d [3]int
	e string
}

type s2Type struct {
	c float64
	a int
}

func main() {

	// initialize data
	// s1 := make([]s1_t, LENGTH)
	// for i:=0; i<LENGTH; i++ {
	// 	s1[i] = s1_t{a:i, b:float32(i*i), c:1./(float64(i)+1)}
	// }
	// fmt.Printf(":: data: %v\n", s1)
	s1 := [length]s1Type{}
	for i := 0; i < int(length); i++ {
		s1[i] = s1Type{
			a: i,
			b: float32(i * i),
			c: 1. / (float64(i) + 1),
			d: [...]int{i, i * 2, i * 3},
			e: fmt.Sprintf("--%d--", i),
		}
		//s1[i].d = []float64{float64(i), float64(2*i), 3.*i}}
	}
	fmt.Printf(":: data: %v\n", s1)

	// create data space
	dims := []uint{length}
	space, err := hdf5.CreateSimpleDataspace(dims, nil)
	if err != nil {
		panic(err)
	}

	// create the file
	f, err := hdf5.CreateFile(fname, hdf5.F_ACC_TRUNC)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Printf(":: file [%s] created (id=%d)\n", fname, f.Id())

	// create the memory data type
	dtype, err := hdf5.NewDatatypeFromValue(s1[0])
	if err != nil {
		panic("could not create a dtype")
	}

	// create the dataset
	dset, err := f.CreateDataset(dsname, dtype, space)
	if err != nil {
		panic(err)
	}
	fmt.Printf(":: dset (id=%d)\n", dset.Id())

	// write data to the dataset
	fmt.Printf(":: dset.Write...\n")
	err = dset.Write(&s1)
	if err != nil {
		panic(err)
	}
	fmt.Printf(":: dset.Write... [ok]\n")

	// release resources
	dset.Close()
	f.Close()

	// open the file and the dataset
	f, err = hdf5.OpenFile(fname, hdf5.F_ACC_RDONLY)
	if err != nil {
		panic(err)
	}
	dset, err = f.OpenDataset(dsname)
	if err != nil {
		panic(err)
	}

	// read it back into a new slice
	s2 := make([]s1Type, length)
	err = dset.Read(&s2)
	if err != nil {
		panic(err)
	}

	// display the fields
	fmt.Printf(":: data: %v\n", s2)

	// release resources
	dset.Close()
	f.Close()
}
