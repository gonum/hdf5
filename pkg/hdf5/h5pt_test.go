package hdf5

import "os"
import "testing"

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
	isthep      []int
	jmohep      [2][2]int
}

func TestTable(t *testing.T) {
	// define an array of particles
	p_data := []particle_t{
		{"zero", 0, 0, 0.0, 0., []int{0, 0}, [2][2]int{{0, 0}, {0, 0}}},
		{"one", 10, 10, 1.0, 10., []int{0, 0}, [2][2]int{{1, 0}, {0, 1}}},
		{"two", 20, 20, 2.0, 20., []int{0, 0}, [2][2]int{{2, 0}, {0, 2}}},
		{"three", 30, 30, 3.0, 30., []int{0, 0}, [2][2]int{{3, 0}, {0, 3}}},
		{"four", 40, 40, 4.0, 40., []int{0, 0}, [2][2]int{{4, 0}, {0, 4}}},
		{"five", 50, 50, 5.0, 50., []int{0, 0}, [2][2]int{{5, 0}, {0, 5}}},
		{"six", 60, 60, 6.0, 60., []int{0, 0}, [2][2]int{{6, 0}, {0, 6}}},
		{"seven", 70, 70, 7.0, 70., []int{0, 0}, [2][2]int{{7, 0}, {0, 7}}},
	}

	chunk_size := 10
	compress := 0

	// create a new file using default properties
	f, err := CreateFile(FNAME, F_ACC_TRUNC)
	if err != nil {
		t.Fatalf("CreateFile failed: %s", err)
	}
	defer os.Remove(FNAME)
	defer f.Close()

	// create a fixed-length packet table within the file
	table, err := f.CreateTableFrom(TABLE_NAME, particle_t{}, chunk_size, compress)
	if err != nil {
		t.Fatalf("CreateTableFrom failed: %s", err)
	}
	defer table.Close()

	if !table.IsValid() {
		t.Fatal("table is invalid")
	}

	// write one packet to the packet table
	if err = table.Append(p_data[0]); err != nil {
		t.Fatalf("Append failed with single packet: %s", err)
	}

	// write several packets
	if err = table.Append(p_data[1:]); err != nil {
		t.Fatalf("Append failed with multiple packets: %s", err)
	}

	// get the number of packets
	n, err := table.NumPackets()
	if err != nil {
		t.Fatalf("NumPackets failed: %s", err)
	}
	if n != NRECORDS {
		t.Fatalf("Wrong number of packets reported, expected %d but got %d", NRECORDS, n)
	}

	// iterate through packets
	for i := 0; i != n; i++ {
		p := make([]particle_t, 1)
		if err := table.Next(p); err != nil {
			t.Fatalf("Next failed: %s", err)
		}
	}

	// reset index
	table.CreateIndex()
	dst_buf := make([]particle_t, NRECORDS)
	if err = table.ReadPackets(0, NRECORDS, dst_buf); err != nil {
		t.Fatalf("ReadPackets failed: %s", err)
	}
}
