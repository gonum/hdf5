// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import (
	"os"
	"reflect"
	"testing"
)

const (
	fname     string = "ex_table_01.h5"
	tname     string = "table"
	chunkSize int    = 10
	compress  int    = 0
)

type particle struct {
	// Name        string  `hdf5:"Name"` // FIXME(TacoVox): ReadPackets seems to need an update
	Vehicle_no  uint8   `hdf5:"Vehicle Number"`
	Satellites  int8    `hdf5:"Satellites"`
	Cars_no     int16   `hdf5:"Number of Cars"`
	Trucks_no   int16   `hdf5:"Number of Trucks"`
	Min_speed   uint32  `hdf5:"Minimum Speed"`
	Lati        int32   `hdf5:"Latitude"`
	Max_speed   uint64  `hdf5:"Maximum Speed"`
	Longi       int64   `hdf5:"Longitude"`
	Pressure    float32 `hdf5:"Pressure"`
	Temperature float64 `hdf5:"Temperature"`
	Accurate    bool    `hdf5:"Accurate"`
	// isthep      []int       // FIXME(sbinet)
	// jmohep      [2][2]int64 // FIXME(sbinet)
}

func testTable(t *testing.T, dType interface{}, data ...interface{}) {
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
	if err = table.Append(data...); err != nil {
		t.Fatalf("Append failed with single packet for %s: %s", typeString, err)
	}

	// get the number of packets
	n, err := table.NumPackets()
	if err != nil {
		t.Fatalf("NumPackets failed for %s: %s", typeString, err)
	}
	if n != len(data) {
		t.Fatalf("Wrong number of packets reported for %s, expected %d but got %d", typeString, len(data), n)
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

	readdata := make([]particle, len(data))

	if err = table.ReadPackets(0, len(data), &readdata); err != nil {
		t.Fatalf("ReadPackets failed for %s: %s", typeString, err)
	}

	if !reflect.DeepEqual(readdata, data) {
		t.Fatalf("Data differs for %s.\ngot= %#v\nwant=%#v\n", typeString, readdata, data)
	}
}

func TestPTStruct(t *testing.T) {
	// particles := []particle{
	// 	{"zero", 0, 0, 0.0, 0., []int{0, 0}, [2][2]int{{0, 0}, {0, 0}}},
	// 	{"one", 10, 10, 1.0, 10., []int{0, 0}, [2][2]int{{1, 0}, {0, 1}}},
	// 	{"two", 20, 20, 2.0, 20., []int{0, 0}, [2][2]int{{2, 0}, {0, 2}}},
	// 	{"three", 30, 30, 3.0, 30., []int{0, 0}, [2][2]int{{3, 0}, {0, 3}}},
	// 	{"four", 40, 40, 4.0, 40., []int{0, 0}, [2][2]int{{4, 0}, {0, 4}}},
	// 	{"five", 50, 50, 5.0, 50., []int{0, 0}, [2][2]int{{5, 0}, {0, 5}}},
	// 	{"six", 60, 60, 6.0, 60., []int{0, 0}, [2][2]int{{6, 0}, {0, 6}}},
	// 	{"seven", 70, 70, 7.0, 70., []int{0, 0}, [2][2]int{{7, 0}, {0, 7}}},
	// } // TODO use array and strings when read is fixed.

	testTable(t, particle{}, particle{0, 0, 0, 0, 0, 0, 0, 0, 0.0, 0.0, false})
	testTable(t, particle{},
		particle{10, 10, 10, 10, 10, 10, 10, 10, 1.0, 10.0, false},
		particle{20, 20, 20, 20, 20, 20, 20, 20, 2.0, 20.0, false},
		particle{30, 30, 30, 30, 30, 30, 30, 30, 3.0, 30.0, false},
		particle{40, 40, 40, 40, 40, 40, 40, 40, 4.0, 40.0, true},
		particle{50, 50, 50, 50, 50, 50, 50, 50, 5.0, 50.0, true},
		particle{60, 60, 60, 60, 60, 60, 60, 60, 6.0, 60.0, true},
		particle{70, 70, 70, 70, 70, 70, 70, 70, 7.0, 70.0, true},
	)
}

func TestPTableBasic(t *testing.T) {
	// INT8
	testTable(t, T_NATIVE_INT8, int8(-3))
	testTable(t, T_NATIVE_INT8,
		int8(-2),
		int8(-1),
		int8(0),
		int8(1),
		int8(2),
		int8(3),
		int8(4),
	)

	// INT16
	testTable(t, T_NATIVE_INT16, int16(-3))
	testTable(t, T_NATIVE_INT16,
		int16(-2),
		int16(-1),
		int16(0),
		int16(1),
		int16(2),
		int16(3),
		int16(4),
	)

	// INT32
	testTable(t, T_NATIVE_INT32, int32(-3))
	testTable(t, T_NATIVE_INT32,
		int32(-2),
		int32(-1),
		int32(0),
		int32(1),
		int32(2),
		int32(3),
		int32(4),
	)

	// INT64
	testTable(t, T_NATIVE_INT64, int64(-3))
	testTable(t, T_NATIVE_INT64,
		int64(-2),
		int64(-1),
		int64(0),
		int64(1),
		int64(2),
		int64(3),
		int64(4),
	)

	// UINT8
	testTable(t, T_NATIVE_UINT8, uint8(0))
	testTable(t, T_NATIVE_UINT8,
		uint8(1),
		uint8(2),
		uint8(3),
		uint8(4),
		uint8(5),
		uint8(6),
		uint8(7),
	)

	// UINT16
	testTable(t, T_NATIVE_UINT16, uint16(0))
	testTable(t, T_NATIVE_UINT16,
		uint16(1),
		uint16(2),
		uint16(3),
		uint16(4),
		uint16(5),
		uint16(6),
		uint16(7),
	)

	// UINT32
	testTable(t, T_NATIVE_UINT32, uint32(0))
	testTable(t, T_NATIVE_UINT32,
		uint32(1),
		uint32(2),
		uint32(3),
		uint32(4),
		uint32(5),
		uint32(6),
		uint32(7),
	)

	// UINT64
	testTable(t, T_NATIVE_UINT64, uint64(0))
	testTable(t, T_NATIVE_UINT64,
		uint64(1),
		uint64(2),
		uint64(3),
		uint64(4),
		uint64(5),
		uint64(6),
		uint64(7),
	)

	// FLOAT32
	testTable(t, T_NATIVE_FLOAT, float32(0))
	testTable(t, T_NATIVE_FLOAT,
		float32(1),
		float32(2),
		float32(3),
		float32(4),
		float32(5),
		float32(6),
		float32(7),
	)

	// FLOAT64
	testTable(t, T_NATIVE_DOUBLE, float64(0))
	testTable(t, T_NATIVE_DOUBLE,
		float64(1),
		float64(2),
		float64(3),
		float64(4),
		float64(5),
		float64(6),
		float64(7),
	)

	// BOOL
	testTable(t, T_NATIVE_HBOOL, false)
	testTable(t, T_NATIVE_HBOOL,
		true,
		false,
		true,
		false,
		true,
		false,
		true,
	)

	// STRING
	testTable(t, T_GO_STRING, "zero")
	testTable(t, T_GO_STRING,
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
	)
}
