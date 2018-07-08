// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

// #include "hdf5.h"
// #include "hdf5_hl.h"
// #include <stdint.h>
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"

	"gonum.org/v1/hdf5/cmem"
)

// Table is an hdf5 packet-table.
type Table struct {
	Identifier
}

func newPacketTable(id C.hid_t) *Table {
	return &Table{Identifier{id}}
}

// Close closes an open packet table.
func (t *Table) Close() error {
	return t.closeWith(h5ptclose)
}

func h5ptclose(id C.hid_t) C.herr_t {
	return C.H5PTclose(id)
}

// IsValid returns whether or not an indentifier points to a packet table.
func (t *Table) IsValid() bool {
	return C.H5PTis_valid(t.id) >= 0
}

// ReadPackets reads a number of packets from a packet table.
func (t *Table) ReadPackets(start, nRecords int, data interface{}) error {
	cStart := C.hsize_t(start)
	cNRecords := C.size_t(nRecords)
	rv := reflect.Indirect(reflect.ValueOf(data))
	rt := rv.Type()
	var cData unsafe.Pointer
	switch rt.Kind() {
	case reflect.Array:
		if rv.Len() < nRecords {
			panic(fmt.Errorf("not enough capacity in array (cap=%d)", rv.Len()))
		}
		cData = unsafe.Pointer(rv.Index(0).UnsafeAddr())

	case reflect.Slice:
		if rv.Len() < nRecords {
			panic(fmt.Errorf("not enough capacity in slice (cap=%d)", rv.Len()))
		}
		slice := (*reflect.SliceHeader)(unsafe.Pointer(rv.UnsafeAddr()))
		cData = unsafe.Pointer(slice.Data)

	default:
		panic(fmt.Errorf("unhandled kind (%s), need slice or array", rt.Kind()))
	}
	err := C.H5PTread_packets(t.id, cStart, cNRecords, cData)
	return h5err(err)
}

// Append appends packets to the end of a packet table.
//
// Struct values must only have exported fields, otherwise Append will panic.
func (t *Table) Append(args ...interface{}) error {
	if len(args) == 0 {
		return fmt.Errorf("hdf5: no arguments passed to packet table append.")
	}

	var enc cmem.Encoder
	for _, arg := range args {
		if err := enc.Encode(arg); err != nil {
			return err
		}
	}

	if len(enc.Buf) <= 0 {
		return fmt.Errorf("hdf5: invalid empty buffer")
	}

	return h5err(C.H5PTappend(t.id, C.size_t(len(args)), unsafe.Pointer(&enc.Buf[0])))
}

// Next reads packets from a packet table starting at the current index into the value pointed at by data.
// i.e. data is a pointer to an array or a slice.
func (t *Table) Next(data interface{}) error {
	rt := reflect.TypeOf(data)
	if rt.Kind() != reflect.Ptr {
		return fmt.Errorf("hdf5: invalid value type. got=%v, want pointer", rt.Kind())
	}
	rt = rt.Elem()
	rv := reflect.Indirect(reflect.ValueOf(data))

	n := C.size_t(0)
	cdata := unsafe.Pointer(nil)
	switch rt.Kind() {
	case reflect.Array:
		if rv.Cap() <= 0 {
			panic(fmt.Errorf("not enough capacity in array (cap=%d)", rv.Cap()))
		}
		cdata = unsafe.Pointer(rv.UnsafeAddr())
		n = C.size_t(rv.Cap())

	case reflect.Slice:
		if rv.Cap() <= 0 {
			panic(fmt.Errorf("not enough capacity in slice (cap=%d)", rv.Cap()))
		}
		slice := (*reflect.SliceHeader)(unsafe.Pointer(rv.UnsafeAddr()))
		cdata = unsafe.Pointer(slice.Data)
		n = C.size_t(rv.Cap())

	default:
		panic(fmt.Errorf("unsupported kind (%s), need slice or array", rt.Kind()))
	}
	err := C.H5PTget_next(t.id, n, cdata)
	return h5err(err)
}

// NumPackets returns the number of packets in a packet table.
func (t *Table) NumPackets() (int, error) {
	var n C.hsize_t
	err := C.H5PTget_num_packets(t.id, &n)
	return int(n), h5err(err)
}

// CreateIndex resets a packet table's index to the first packet.
func (t *Table) CreateIndex() error {
	err := C.H5PTcreate_index(t.id)
	return h5err(err)
}

// SetIndex sets a packet table's index.
func (t *Table) SetIndex(index int) error {
	cIdx := C.hsize_t(index)
	err := C.H5PTset_index(t.id, cIdx)
	return h5err(err)
}

// Type returns an identifier for a copy of the datatype for a dataset. The returned
// datatype must be closed by the user when it is no longer needed.
func (t *Table) Type() (*Datatype, error) {
	hid := C.H5Dget_type(t.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return NewDatatype(hid), nil
}

func createTable(id C.hid_t, name string, typ *Datatype, chunkSize, compression int) (*Table, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	chunk := C.hsize_t(chunkSize)
	compr := C.int(compression)
	hid := C.H5PTcreate_fl(id, cName, typ.id, chunk, compr)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newPacketTable(hid), nil
}

func createTableFrom(id C.hid_t, name string, typ interface{}, chunkSize, compression int) (*Table, error) {
	var err error
	switch dt := typ.(type) {
	case reflect.Type:
		if typ, err := NewDataTypeFromType(dt); err == nil {
			return createTable(id, name, typ, chunkSize, compression)
		}
	case *Datatype:
		return createTable(id, name, dt, chunkSize, compression)
	default:
		if typ, err := NewDataTypeFromType(reflect.TypeOf(typ)); err == nil {
			return createTable(id, name, typ, chunkSize, compression)
		}
	}
	return nil, err
}

func openTable(id C.hid_t, name string) (*Table, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	hid := C.H5PTopen(id, cName)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newPacketTable(hid), nil
}
