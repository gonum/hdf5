package hdf5

/*
 #cgo LDFLAGS: -lhdf5 -lhdf5_hl
 #include "hdf5.h"
 #include "hdf5_hl.h"

 #include <stdlib.h>
 #include <string.h>
 */
import "C"

import (
	"unsafe"
	"os"
	"runtime"
	"fmt"
)

// ---- HDF5 Packet Table ----

// a HDF5 packet table
type Table struct {
	id C.hid_t
}

func new_packet_table(id C.hid_t) *Table {
	t := &Table{id:id}
	runtime.SetFinalizer(t, (*Table).h5pt_finalizer)
	return t
}

func (t *Table) h5pt_finalizer() {
	err := t.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing packet table: %s",err))
	}
}

// Closes an open packet table.
// herr_t H5PTclose( hid_t table_id )
func (t *Table) Close() os.Error {
	if t.id > 0 {
		err := togo_err(C.H5PTclose(t.id))
		if err != nil {
			t.id = 0
		}
		return err
	}
	return nil
}

// Determines whether an indentifier points to a packet table.
// herr_t H5PTis_valid( hid_t table_id)
func (t *Table) IsValid() bool {
	o := int(C.H5PTis_valid(t.id))
	if o> 0 {
		return true
	} 
	return false
}

// Reads a number of packets from a packet table.
// herr_t H5PTread_packets( hid_t table_id, hsize_t start, size_t nrecords, void* data)
func (t *Table) ReadPackets(start, nrecords int) ([]interface{}, os.Error) {
	data := make([]interface{}, nrecords)
	c_data := unsafe.Pointer(&data[0])
	c_start := C.hsize_t(start)
	c_nrecords := C.size_t(nrecords)
	err := C.H5PTread_packets(t.id, c_start, c_nrecords, c_data)
	return data, togo_err(err)
}

// EOF

