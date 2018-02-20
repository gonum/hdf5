// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import (
	"os"
	"reflect"
)
import "testing"

const (
	fname     string = "ex_table_01.h5"
	tname     string = "table"
	chunkSize int    = 10
	compress  int    = 0
)

type particle struct {
	//name        string  `hdf5:"Name"`		 // FIXME(TacoVox)
	vehicle_no  uint8   `hdf5:"Vehicle Number"`
	sattelites  int8    `hdf5:"Sattelites"`
	cars_no     int16   `hdf5:"Number of Cars"`
	trucks_no   int16   `hdf5:"Number of Trucks"`
	min_speed   uint32  `hdf5:"Minimum Speed"`
	lati        int32   `hdf5:"Latitude"`
	max_speed   uint64  `hdf5:"Maximum Speed"`
	longi       int64   `hdf5:"Longitude"`
	pressure    float32 `hdf5:"Pressure"`
	temperature float64 `hdf5:"Temperature"`
	accurate    bool    `hdf5:"Accurate"`
	// isthep      []int                     // FIXME(sbinet)
	// jmohep [2][2]int64                    // FIXME(sbinet)
}

func testTable(t *testing.T, dType interface{}, data interface{}, nrecords int) {
	var table *Table

	typeString := reflect.TypeOf(data).String()

	// create a new file using default properties
	f, err := CreateFile(fname, F_ACC_TRUNC)
	if err != nil {
		t.Fatalf("CreateFile failed for %s: %s", typeString, err)
	}
	defer os.Remove(fname)
	defer f.Close()

	table, err = f.CreateTableFrom(tname, dType, chunkSize, compress)
	if err != nil {
		t.Fatalf("CreateTableFrom struct failed for %s: %s", typeString, err)
	}
	defer table.Close()

	if !table.IsValid() {
		t.Fatalf("PacketTable %q is invalid for %s.", tname, typeString)
	}

	// write one packet to the packet table
	if err = table.Append(data); err != nil {
		t.Fatalf("Append failed with single packet for %s: %s", typeString, err)
	}

	// get the number of packets
	n, err := table.NumPackets()
	if err != nil {
		t.Fatalf("NumPackets failed for %s: %s", typeString, err)
	}
	if n != nrecords {
		t.Fatalf("Wrong number of packets reported for %s, expected %d but got %d", typeString, nrecords, n)
	}

	// iterate through packets
	for i := 0; i != n; i++ {
		p := make([]interface{}, 1)
		if err = table.Next(&p); err != nil {
			t.Fatalf("Next failed for %s: %s", typeString, err)
		}
	}

	// For now just conduct the "old" test case // FIXME(TacoVox)
	if reflect.TypeOf(data).String() != "[]hdf5.particle" {
		return
	}

	// reset index
	table.CreateIndex()

	readdata := make([]particle, nrecords)

	if err = table.ReadPackets(0, nrecords, &readdata); err != nil {
		t.Fatalf("ReadPackets failed for %s: %s", typeString, err)
	}

	if !reflect.DeepEqual(readdata, data) {
		t.Fatalf("Data differs for %s.\ngot= %#v\nwant=%#v\n", typeString, readdata, data)
	}
}

func TestPTStruct(t *testing.T) {
	// define an array of particles
	particles := []particle{
		{0, 0, 0, 0, 0, 0, 0, 0, 0.0, 0.0, false},
		{10, 10, 10, 10, 10, 10, 10, 10, 1.0, 10.0, false},
		{20, 20, 20, 20, 20, 20, 20, 20, 2.0, 20.0, false},
		{30, 30, 30, 30, 30, 30, 30, 30, 3.0, 30.0, false},
		{40, 40, 40, 40, 40, 40, 40, 40, 4.0, 40.0, true},
		{50, 50, 50, 50, 50, 50, 50, 50, 5.0, 50.0, true},
		{60, 60, 60, 60, 60, 60, 60, 60, 6.0, 60.0, true},
		{70, 70, 70, 70, 70, 70, 70, 70, 7.0, 70.0, true},
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

	testTable(t, particle{}, particles[0], 1)

	parts := particles[1:]
	testTable(t, particle{}, parts, 7)
}

func TestPTableBasic(t *testing.T) {
	// INT8
	i8s := []int8{-3, -2, -1, 0, 1, 2, 3, 4}
	testTable(t, T_NATIVE_INT8, i8s[0], 1)
	i8sub := i8s[1:]
	testTable(t, T_NATIVE_INT8, i8sub, 7)

	// INT16
	i16s := []int16{-3, -2, -1, 0, 1, 2, 3, 4}
	testTable(t, T_NATIVE_INT16, i16s[0], 1)
	i16sub := i16s[1:]
	testTable(t, T_NATIVE_INT16, i16sub, 7)

	// INT32
	i32s := []int32{-3, -2, -1, 0, 1, 2, 3, 4}
	testTable(t, T_NATIVE_INT32, i32s[0], 1)
	i32sub := i32s[1:]
	testTable(t, T_NATIVE_INT32, i32sub, 7)

	// INT64
	i64s := []int64{-3, -2, -1, 0, 1, 2, 3, 4}
	testTable(t, T_NATIVE_INT64, i64s[0], 1)
	i64sub := i64s[1:]
	testTable(t, T_NATIVE_INT64, i64sub, 7)

	// UINT8
	ui8s := []uint8{0, 1, 2, 3, 4, 5, 6, 7}
	testTable(t, T_NATIVE_UINT8, ui8s[0], 1)
	ui8sub := ui8s[1:]
	testTable(t, T_NATIVE_UINT8, ui8sub, 7)

	// UINT16
	ui16s := []uint16{0, 1, 2, 3, 4, 5, 6, 7}
	testTable(t, T_NATIVE_UINT16, ui16s[0], 1)
	ui16sub := ui16s[1:]
	testTable(t, T_NATIVE_UINT16, ui16sub, 7)

	// UINT32
	ui32s := []uint32{0, 1, 2, 3, 4, 5, 6, 7}
	testTable(t, T_NATIVE_UINT32, ui32s[0], 1)
	ui32sub := ui32s[1:]
	testTable(t, T_NATIVE_UINT32, ui32sub, 7)

	// UINT64
	ui64s := []uint64{0, 1, 2, 3, 4, 5, 6, 7}
	testTable(t, T_NATIVE_UINT64, ui64s[0], 1)
	ui64sub := ui64s[1:]
	testTable(t, T_NATIVE_UINT64, ui64sub, 7)

	// FLOAT32
	f32s := []float32{0., 1., 2., 3., 4., 5., 6., 7.}
	testTable(t, T_NATIVE_FLOAT, f32s[0], 1)
	f32sub := f32s[1:]
	testTable(t, T_NATIVE_FLOAT, f32sub, 7)

	// FLOAT64
	f64s := []float64{0., 1., 2., 3., 4., 5., 6., 7.}
	testTable(t, T_NATIVE_DOUBLE, f64s[0], 1)
	f64sub := f64s[1:]
	testTable(t, T_NATIVE_DOUBLE, f64sub, 7)

	// BOOL
	bs := []bool{false, true, false, true, false, true, false, true}
	testTable(t, T_NATIVE_HBOOL, bs[0], 1)
	bsub := bs[1:]
	testTable(t, T_NATIVE_HBOOL, bsub, 7)

	// STRING
	strs := []string{"zero", "one", "two", "three", "four", "five", "six", "seven"}
	testTable(t, T_GO_STRING, strs[0], 1)
	strsub := strs[1:]
	testTable(t, T_GO_STRING, strsub, 7)
}
