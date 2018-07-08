// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"gonum.org/v1/hdf5"
)

const (
	fname    string = "ex_table_01.h5"
	tname    string = "table"
	nfields  int    = 5
	nrecords int    = 8
)

type particle struct {
	// name        string  `hdf5:"Name"`      // FIXME(sbinet)
	Lati        int32   `hdf5:"Latitude"`
	Longi       int64   `hdf5:"Longitude"`
	Pressure    float32 `hdf5:"Pressure"`
	Temperature float64 `hdf5:"Temperature"`
	// isthep      []int                     // FIXME(sbinet)
	// jmohep [2][2]int64                    // FIXME(sbinet)
}

func (p *particle) Equal(o *particle) bool {
	return p.Lati == o.Lati && p.Longi == o.Longi && p.Pressure == o.Pressure && p.Temperature == o.Temperature
}

func main() {

	// define an array of particles
	particles := []particle{
		{0, 0, 0.0, 0.},
		{10, 10, 1.0, 10.},
		{20, 20, 2.0, 20.},
		{30, 30, 3.0, 30.},
		{40, 40, 4.0, 40.},
		{50, 50, 5.0, 50.},
		{60, 60, 6.0, 60.},
		{70, 70, 7.0, 70.},
	}

	// define an array of particles
	// p_data := []particle_t{
	// 	{"zero", 0, 0, 0.0, 0.},
	// 	{"one", 10, 10, 1.0, 10.},
	// 	{"two", 20, 20, 2.0, 20.},
	// 	{"three", 30, 30, 3.0, 30.},
	// 	{"four", 40, 40, 4.0, 40.},
	// 	{"five", 50, 50, 5.0, 50.},
	// 	{"six", 60, 60, 6.0, 60.},
	// 	{"seven", 70, 70, 7.0, 70.},
	// }

	fmt.Printf(":: reference data: %v\n", particles)

	// open a file
	f, err := hdf5.OpenFile(fname, hdf5.F_ACC_RDONLY)
	if err != nil {
		panic(err)
	}
	fmt.Printf(":: file [%s] opened (id=%d)\n", f.Name(), f.Id())

	// create a fixed-length packet table within the file
	table, err := f.OpenTable(tname)
	if err != nil {
		panic(err)
	}
	fmt.Printf(":: table [%s] opened (id=%d)\n", tname, 3)

	// iterate through packets
	for i := 0; i != nrecords; i++ {
		//p := []particle_t{{}}
		p := make([]particle, 1)
		err := table.Next(&p)
		if err != nil {
			panic(err)
		}
		fmt.Printf(":: data[%d]: %v -> [%v]\n", i, p, p[0].Equal(&particles[i]))
	}

	// reset index
	table.CreateIndex()
	parts := make([]particle, nrecords)
	err = table.ReadPackets(0, nrecords, &parts)
	if err != nil {
		panic(err)
	}
	fmt.Printf(":: whole data: %v\n", parts)

	fmt.Printf(":: bye.\n")
}
