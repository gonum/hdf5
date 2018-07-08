// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"reflect"

	"gonum.org/v1/hdf5"
)

const (
	fname    string = "ex_table_01.h5"
	tname    string = "table"
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
	// particles := []particle{
	// 	{"zero", 0, 0, 0.0, 0., []int{0, 0}, [2][2]int{{0, 0}, {0, 0}}},
	// 	{"one", 10, 10, 1.0, 10., []int{0, 0}, [2][2]int{{1, 0}, {0, 1}}},
	// 	{"two", 20, 20, 2.0, 20., []int{0, 0}, [2][2]int{{2, 0}, {0, 2}}},
	// 	{"three", 30, 30, 3.0, 30., []int{0, 0}, [2][2]int{{3, 0}, {0, 3}}},
	// 	{"four", 40, 40, 4.0, 40., []int{0, 0}, [2][2]int{{4, 0}, {0, 4}}},
	// 	{"five", 50, 50, 5.0, 50., []int{0, 0}, [2][2]int{{5, 0}, {0, 5}}},
	// 	{"six", 60, 60, 6.0, 60., []int{0, 0}, [2][2]int{{6, 0}, {0, 6}}},
	// 	{"seven", 70, 70, 7.0, 70., []int{0, 0}, [2][2]int{{7, 0}, {0, 7}}},
	// }

	chunkSize := 10
	compress := 0

	// create a new file using default properties
	f, err := hdf5.CreateFile(fname, hdf5.F_ACC_TRUNC)
	if err != nil {
		panic(fmt.Errorf("CreateFile failed: %s", err))
	}
	defer f.Close()
	fmt.Printf(":: file [%s] created (id=%d)\n", fname, f.Id())

	// create a fixed-length packet table within the file
	table, err := f.CreateTableFrom(tname, particle{}, chunkSize, compress)
	if err != nil {
		panic(fmt.Errorf("CreateTableFrom failed: %s", err))
	}
	defer table.Close()
	fmt.Printf(":: table [%s] created (id=%d)\n", tname, table.Id())

	if !table.IsValid() {
		panic("table is invalid")
	}

	// write one packet to the packet table
	if err = table.Append(particles[0]); err != nil {
		panic(fmt.Errorf("Append failed with single packet: %s", err))
	}

	// write several packets
	if err = table.Append(particles[1], particles[2], particles[3],
		particles[4], particles[5], particles[6], particles[7],
	); err != nil {
		panic(fmt.Errorf("Append failed with multiple packets: %s", err))
	}

	// get the number of packets
	n, err := table.NumPackets()
	if err != nil {
		panic(fmt.Errorf("NumPackets failed: %s", err))
	}
	fmt.Printf(":: nbr entries: %d\n", n)
	if n != nrecords {
		panic(fmt.Errorf(
			"Wrong number of packets reported, expected %d but got %d",
			nrecords, n,
		))
	}

	// iterate through packets
	for i := 0; i != n; i++ {
		p := make([]particle, 1)
		if err := table.Next(&p); err != nil {
			panic(fmt.Errorf("Next failed: %s", err))
		}
		fmt.Printf(":: data[%d]: %v\n", i, p)
	}

	// reset index
	table.CreateIndex()
	parts := make([]particle, nrecords)
	if err = table.ReadPackets(0, nrecords, &parts); err != nil {
		panic(fmt.Errorf("ReadPackets failed: %s", err))
	}

	if !reflect.DeepEqual(parts, particles) {
		panic(fmt.Errorf("particles differ.\ngot= %#v\nwant=%#v\n", parts, particles))
	}

	fmt.Printf(":: whole data: %v\n", parts)
	fmt.Printf(":: bye.\n")
}
