package hdf5

import (
	"os"
	"reflect"
)
import "testing"

const (
	fname    string = "ex_table_01.h5"
	tname    string = "table"
	nrecords int    = 8
)

type particle struct {
	// name        string  `hdf5:"Name"`      // FIXME(sbinet)
	lati        int32   `hdf5:"Latitude"`
	longi       int64   `hdf5:"Longitude"`
	pressure    float32 `hdf5:"Pressure"`
	temperature float64 `hdf5:"Temperature"`
	// isthep      []int                     // FIXME(sbinet)
	// jmohep [2][2]int64                    // FIXME(sbinet)
}

func TestTable(t *testing.T) {

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
	f, err := CreateFile(fname, F_ACC_TRUNC)
	if err != nil {
		t.Fatalf("CreateFile failed: %s", err)
	}
	defer os.Remove(fname)
	defer f.Close()

	// create a fixed-length packet table within the file
	table, err := f.CreateTableFrom(tname, particle{}, chunkSize, compress)
	if err != nil {
		t.Fatalf("CreateTableFrom failed: %s", err)
	}
	defer table.Close()

	if !table.IsValid() {
		t.Fatal("table is invalid")
	}

	// write one packet to the packet table
	if err = table.Append(&particles[0]); err != nil {
		t.Fatalf("Append failed with single packet: %s", err)
	}

	// write several packets
	parts := particles[1:]
	if err = table.Append(&parts); err != nil {
		t.Fatalf("Append failed with multiple packets: %s", err)
	}

	// get the number of packets
	n, err := table.NumPackets()
	if err != nil {
		t.Fatalf("NumPackets failed: %s", err)
	}
	if n != nrecords {
		t.Fatalf("Wrong number of packets reported, expected %d but got %d", nrecords, n)
	}

	// iterate through packets
	for i := 0; i != n; i++ {
		p := make([]particle, 1)
		if err := table.Next(p); err != nil {
			t.Fatalf("Next failed: %s", err)
		}
	}

	// reset index
	table.CreateIndex()
	parts = make([]particle, nrecords)
	if err = table.ReadPackets(0, nrecords, &parts); err != nil {
		t.Fatalf("ReadPackets failed: %s", err)
	}

	if !reflect.DeepEqual(parts, particles) {
		t.Fatalf("particles differ.\ngot= %#v\nwant=%#v\n", parts, particles)
	}
}
