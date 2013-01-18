package hdf5

// #include "hdf5.h"
// #include "hdf5_hl.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"fmt"

	"reflect"
	"runtime"
	"unsafe"
)

// ---- HDF5 Packet Table ----

// a HDF5 packet table
type Table struct {
	id C.hid_t
}

func new_packet_table(id C.hid_t) *Table {
	t := &Table{id: id}
	runtime.SetFinalizer(t, (*Table).h5pt_finalizer)
	return t
}

func (t *Table) h5pt_finalizer() {
	err := t.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing packet table: %s", err))
	}
}

// Closes an open packet table.
// herr_t H5PTclose( hid_t table_id )
func (t *Table) Close() error {
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
	if o > 0 {
		return true
	}
	return false
}

func (t *Table) Id() int {
	return int(t.id)
}

// Reads a number of packets from a packet table.
// herr_t H5PTread_packets( hid_t table_id, hsize_t start, size_t nrecords, void* data)
func (t *Table) ReadPackets(start, nrecords int, data interface{}) error {
	c_start := C.hsize_t(start)
	c_nrecords := C.size_t(nrecords)
	rt := reflect.TypeOf(data)
	rv := reflect.ValueOf(data)
	c_data := unsafe.Pointer(nil)
	switch rt.Kind() {
	case reflect.Array:
		//fmt.Printf("--> array\n")
		if rv.Cap() < nrecords {
			panic(fmt.Sprintf("not enough capacity in array (cap=%d)", rv.Cap()))
		}
		c_data = unsafe.Pointer(rv.Index(0).UnsafeAddr())
		//c_nrecords = C.size_t(rv.Cap())

	case reflect.Slice:
		//fmt.Printf("--> slice\n")
		if rv.Cap() < nrecords {
			panic(fmt.Sprintf("not enough capacity in slice (cap=%d)", rv.Cap()))
			// buf_slice := reflect.MakeSlice(rt, nrecords, nrecords)
			// rv.Set(reflect.AppendSlice(rv, buf_slice))
		}
		slice := (*reflect.SliceHeader)(unsafe.Pointer(rv.UnsafeAddr()))
		c_data = unsafe.Pointer(slice.Data)
		//c_nrecords = C.size_t(rv.Cap())

	default:
		panic(fmt.Sprintf("unhandled kind (%s) need slice or array", rt.Kind()))
	}
	err := C.H5PTread_packets(t.id, c_start, c_nrecords, c_data)
	return togo_err(err)
}

// Appends packets to the end of a packet table.
// herr_t H5PTappend( hid_t table_id, size_t nrecords, const void *data)
func (t *Table) Append(data interface{}) error {
	rt := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	c_nrecords := C.size_t(0)
	c_data := unsafe.Pointer(nil)

	switch rt.Kind() {

	case reflect.Array:
		//fmt.Printf("-->array\n")
		c_nrecords = C.size_t(v.Len())
		c_data = unsafe.Pointer(v.UnsafeAddr())

	case reflect.Slice:
		fmt.Printf("-->slice\n")
		c_nrecords = C.size_t(v.Len())
		slice := (*reflect.SliceHeader)(unsafe.Pointer(v.UnsafeAddr()))
		c_data = unsafe.Pointer(slice.Data)
		//c_data = unsafe.Pointer(v.Index(0).Addr().UnsafeAddr())
		//c_data = unsafe.Pointer(&c_data)

	case reflect.String:
		c_nrecords = C.size_t(v.Len())
		str := (*reflect.StringHeader)(unsafe.Pointer(v.UnsafeAddr()))
		c_data = unsafe.Pointer(str.Data)

	case reflect.Ptr:
		//fmt.Printf("-->ptr %v\n", v.Elem())
		c_nrecords = C.size_t(1)
		c_data = unsafe.Pointer(v.Elem().UnsafeAddr())

	default:
		//fmt.Printf("-->\n")
		c_nrecords = C.size_t(1)
		c_data = unsafe.Pointer(v.UnsafeAddr())
	}

	fmt.Printf(":: append(%d, %d, %v)\n", c_nrecords, c_data, t.id)
	err := C.H5PTappend(t.id, c_nrecords, c_data)
	fmt.Printf(":: append(%d, %d) -> %v\n", c_nrecords, c_data, err)
	return togo_err(err)
}

// Reads packets from a packet table starting at the current index.
// herr_t H5PTget_next( hid_t table_id, size_t nrecords, void *data)
func (t *Table) Next(data interface{}) error {
	rt := reflect.TypeOf(data)
	rv := reflect.ValueOf(data)
	c_nrecords := C.size_t(0)
	c_data := unsafe.Pointer(nil)
	switch rt.Kind() {
	case reflect.Array:
		//fmt.Printf("--> array\n")
		if rv.Cap() <= 0 {
			panic(fmt.Sprintf("not enough capacity in array (cap=%d)", rv.Cap()))
		}
		c_data = unsafe.Pointer(rv.Index(0).UnsafeAddr())
		c_nrecords = C.size_t(rv.Cap())

	case reflect.Slice:
		//fmt.Printf("--> slice\n")
		if rv.Cap() <= 0 {
			panic(fmt.Sprintf("not enough capacity in slice (cap=%d)", rv.Cap()))
		}
		slice := (*reflect.SliceHeader)(unsafe.Pointer(rv.UnsafeAddr()))
		c_data = unsafe.Pointer(slice.Data)
		c_nrecords = C.size_t(rv.Cap())

	default:
		panic(fmt.Sprintf("unhandled kind (%s) need slice or array", rt.Kind()))
	}
	//fmt.Printf("--data: %v...\n", data)
	err := C.H5PTget_next(t.id, c_nrecords, c_data)
	//fmt.Printf("--data: err [%v]\n", err)
	//fmt.Printf("--data: %v... [%v]\n", data, err)
	return togo_err(err)
}

// Returns the number of packets in a packet table.
// herr_t H5PTget_num_packets( hid_t table_id, hsize_t * nrecords)
func (t *Table) NumPackets() (int, error) {
	c_nrecords := C.hsize_t(0)
	err := C.H5PTget_num_packets(t.id, &c_nrecords)
	return int(c_nrecords), togo_err(err)
}

// Resets a packet table's index to the first packet.
// herr_t H5PTcreate_index( hid_t table_id)
func (t *Table) CreateIndex() error {
	err := C.H5PTcreate_index(t.id)
	return togo_err(err)
}

// Sets a packet table's index.
// herr_t H5PTset_index( hid_t table_id, hsize_t pt_index)
func (t *Table) SetIndex(index int) error {
	c_idx := C.hsize_t(index)
	err := C.H5PTset_index(t.id, c_idx)
	return togo_err(err)
}

// Returns an identifier for a copy of the datatype for a dataset.
// hid_t H5Dget_type(hid_t dataset_id )
func (t *Table) Type() (*Datatype, error) {
	hid := C.H5Dget_type(t.id)
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	dt := new_dtype(hid, nil)
	return dt, err
}

// EOF
