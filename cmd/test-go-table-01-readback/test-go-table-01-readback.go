package main

import (
	"fmt"

	"github.com/sbinet/go-hdf5"
)

const (
	FNAME      string = "ex_table_01.h5"
	TABLE_NAME string = "table"
	NFIELDS    int    = 5
	NRECORDS   int    = 8
)

type particle_t struct {
	name        string  //"Name"
	lati        int     "Latitude"
	longi       int     "Longitude"
	pressure    float32 "Pressure"
	temperature float64 "Temperature"
}

func (p *particle_t) Equal(o *particle_t) bool {
	return p.name == o.name && p.lati == o.lati && p.longi == o.longi && p.pressure == o.pressure && p.temperature == o.temperature
}

func main() {

	// define an array of particles
	p_data := []particle_t{
		{"zero", 0, 0, 0.0, 0.},
		{"one", 10, 10, 1.0, 10.},
		{"two", 20, 20, 2.0, 20.},
		{"three", 30, 30, 3.0, 30.},
		{"four", 40, 40, 4.0, 40.},
		{"five", 50, 50, 5.0, 50.},
		{"six", 60, 60, 6.0, 60.},
		{"seven", 70, 70, 7.0, 70.},
	}
	fmt.Printf(":: reference data: %v\n", p_data)

	// open a file
	f, err := hdf5.OpenFile(FNAME, hdf5.F_ACC_RDONLY)
	if err != nil {
		panic(err)
	}
	fmt.Printf(":: file [%s] opened (id=%d)\n", f.Name(), f.Id())

	// create a fixed-length packet table within the file
	table, err := f.OpenTable(TABLE_NAME)
	if err != nil {
		panic(err)
	}
	fmt.Printf(":: table [%s] opened (id=%d)\n", TABLE_NAME, 3)

	// iterate through packets
	for i := 0; i != NRECORDS; i++ {
		//p := []particle_t{{}}
		p := make([]particle_t, 1)
		p[0].name = "+++++++    +"
		err := table.Next(p)
		if err != nil {
			panic(err)
		}
		fmt.Printf(":: data[%d]: %v -> [%v]\n", i, p, p[0].Equal(&p_data[i]))
	}

	// reset index
	table.CreateIndex()
	dst_buf := make([]particle_t, NRECORDS)
	err = table.ReadPackets(0, NRECORDS, dst_buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf(":: whole data: %v\n", dst_buf)

	fmt.Printf(":: bye.\n")
}
