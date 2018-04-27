// Copyright ©2017 The gonum Authors. All rights reserved.
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
	"runtime"
	"unsafe"
)

// appendData is a wrapper type for information necessary to create and
// subsequently write an in memory object to a PacketTable using Append().
type appendData struct {
	ptr        unsafe.Pointer
	memSize    C.size_t
	offset     C.size_t
	numRecords C.size_t
	strs       []unsafe.Pointer
}

func (ad *appendData) free() {
	C.free(ad.ptr)
	ad.ptr = nil

	for _, str := range ad.strs {
		C.free(str)
		str = nil
	}
}

// Table is an hdf5 packet-table.
type Table struct {
	Identifier
}

func newPacketTable(id C.hid_t) *Table {
	t := &Table{Identifier{id}}
	runtime.SetFinalizer(t, (*Table).finalizer)
	return t
}

func (t *Table) finalizer() {
	if err := t.Close(); err != nil {
		panic(fmt.Errorf("error closing packet table: %s", err))
	}
}

// Close closes an open packet table.
func (t *Table) Close() error {
	if t.id == 0 {
		return nil
	}
	err := h5err(C.H5PTclose(t.id))
	t.id = 0
	return err
}

// IsValid returns whether or not an indentifier points to a packet table.
func (t *Table) IsValid() bool {
	return C.H5PTis_valid(t.id) >= 0
}

func (t *Table) Id() int {
	return int(t.id)
}

// ReadPackets reads a number of packets from a packet table.
func (t *Table) ReadPackets(start, nrecords int, data interface{}) error {
	c_start := C.hsize_t(start)
	c_nrecords := C.size_t(nrecords)
	rv := reflect.Indirect(reflect.ValueOf(data))
	rt := rv.Type()
	c_data := unsafe.Pointer(nil)
	switch rt.Kind() {
	case reflect.Array:
		if rv.Len() < nrecords {
			panic(fmt.Errorf("not enough capacity in array (cap=%d)", rv.Len()))
		}
		c_data = unsafe.Pointer(rv.Index(0).UnsafeAddr())

	case reflect.Slice:
		if rv.Len() < nrecords {
			panic(fmt.Errorf("not enough capacity in slice (cap=%d)", rv.Len()))
		}
		slice := (*reflect.SliceHeader)(unsafe.Pointer(rv.UnsafeAddr()))
		c_data = unsafe.Pointer(slice.Data)

	default:
		panic(fmt.Errorf("unhandled kind (%s), need slice or array", rt.Kind()))
	}
	err := C.H5PTread_packets(t.id, c_start, c_nrecords, c_data)
	return h5err(err)
}

// extractValues extracts the values to be appended to a packet table and adds
// them to the appendData structure.
//
// Struct values must only have exported fields, otherwise extractValues will
// panic.
func extractValues(ad *appendData, data interface{}) error {
	rv := reflect.Indirect(reflect.ValueOf(data))
	rt := rv.Type()

	var dataPtr unsafe.Pointer
	var dataSize C.size_t

	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			if err := extractValues(ad, rv.Index(i).Interface()); err != nil {
				return err
			}
		}
		ad.numRecords = C.size_t(rv.Len())
		return nil

	case reflect.Struct:
		offset := ad.offset
		for i := 0; i < rv.NumField(); i++ {
			sfv := rv.Field(i).Interface()
			// In order to keep the struct offset always correct
			ad.offset = offset + C.size_t(rt.Field(i).Offset)
			if err := extractValues(ad, sfv); err != nil {
				return err
			}
			// Reset the offset to the correct array sized
			ad.offset = offset + C.size_t(rt.Size())
		}
		ad.numRecords = 1
		return nil

	case reflect.String:
		ad.numRecords = 1
		stringData := C.CString(rv.String())
		ad.strs = append(ad.strs, unsafe.Pointer(stringData))
		dataPtr = unsafe.Pointer(&stringData)
		dataSize = C.size_t(unsafe.Sizeof(dataPtr))

	case reflect.Ptr:
		ad.numRecords = 1
		return extractValues(ad, rv.Elem())

	case reflect.Int8:
		ad.numRecords = 1
		val := C.int8_t(rv.Int())
		dataPtr = unsafe.Pointer(&val)
		dataSize = 1

	case reflect.Uint8:
		ad.numRecords = 1
		val := C.uint8_t(rv.Uint())
		dataPtr = unsafe.Pointer(&val)
		dataSize = 1

	case reflect.Int16:
		ad.numRecords = 1
		val := C.int16_t(rv.Int())
		dataPtr = unsafe.Pointer(&val)
		dataSize = 2

	case reflect.Uint16:
		ad.numRecords = 1
		val := C.uint16_t(rv.Uint())
		dataPtr = unsafe.Pointer(&val)
		dataSize = 2

	case reflect.Int32:
		ad.numRecords = 1
		val := C.int32_t(rv.Int())
		dataPtr = unsafe.Pointer(&val)
		dataSize = 4

	case reflect.Uint32:
		ad.numRecords = 1
		val := C.uint32_t(rv.Uint())
		dataPtr = unsafe.Pointer(&val)
		dataSize = 4

	case reflect.Int64:
		ad.numRecords = 1
		val := C.int64_t(rv.Int())
		dataPtr = unsafe.Pointer(&val)
		dataSize = 8

	case reflect.Uint64:
		ad.numRecords = 1
		val := C.uint64_t(rv.Uint())
		dataPtr = unsafe.Pointer(&val)
		dataSize = 8

	case reflect.Float32:
		ad.numRecords = 1
		val := C.float(rv.Float())
		dataPtr = unsafe.Pointer(&val)
		dataSize = 4

	case reflect.Float64:
		ad.numRecords = 1
		val := C.double(rv.Float())
		dataPtr = unsafe.Pointer(&val)
		dataSize = 8

	case reflect.Bool:
		ad.numRecords = 1
		val := C.uchar(0)
		if rv.Bool() {
			val = 1
		}
		dataPtr = unsafe.Pointer(&val)
		dataSize = 1

	default:
		return fmt.Errorf("hdf5: PT Append does not support datatype (%s).", rt.Kind())
	}

	ad.memSize = dataSize + ad.offset
	ad.ptr = C.realloc(ad.ptr, ad.memSize)
	C.memcpy(unsafe.Pointer(uintptr(ad.ptr)+uintptr(ad.offset)), dataPtr, dataSize)
	ad.offset += dataSize

	return nil
}

// Append appends packets to the end of a packet table.
//
// Struct values must only have exported fields, otherwise Append will panic.
func (t *Table) Append(data interface{}) error {
	var ad appendData
	defer ad.free()

	if err := extractValues(&ad, data); err != nil {
		return err
	}

	return h5err(C.H5PTappend(t.id, ad.numRecords, ad.ptr))
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
	c_nrecords := C.hsize_t(0)
	err := C.H5PTget_num_packets(t.id, &c_nrecords)
	return int(c_nrecords), h5err(err)
}

// CreateIndex resets a packet table's index to the first packet.
func (t *Table) CreateIndex() error {
	err := C.H5PTcreate_index(t.id)
	return h5err(err)
}

// SetIndex sets a packet table's index.
func (t *Table) SetIndex(index int) error {
	c_idx := C.hsize_t(index)
	err := C.H5PTset_index(t.id, c_idx)
	return h5err(err)
}

// Type returns an identifier for a copy of the datatype for a dataset.
func (t *Table) Type() (*Datatype, error) {
	hid := C.H5Dget_type(t.id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return NewDatatype(hid), nil
}

func createTable(id C.hid_t, name string, dtype *Datatype, chunkSize, compression int) (*Table, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	chunk := C.hsize_t(chunkSize)
	compr := C.int(compression)
	hid := C.H5PTcreate_fl(id, c_name, dtype.id, chunk, compr)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newPacketTable(hid), nil
}

func createTableFrom(id C.hid_t, name string, dtype interface{}, chunkSize, compression int) (*Table, error) {
	var err error
	switch dt := dtype.(type) {
	case reflect.Type:
		if hdfDtype, err := NewDataTypeFromType(dt); err == nil {
			return createTable(id, name, hdfDtype, chunkSize, compression)
		}
	case *Datatype:
		return createTable(id, name, dt, chunkSize, compression)
	default:
		if hdfDtype, err := NewDataTypeFromType(reflect.TypeOf(dtype)); err == nil {
			return createTable(id, name, hdfDtype, chunkSize, compression)
		}
	}
	return nil, err
}

func openTable(id C.hid_t, name string) (*Table, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5PTopen(id, c_name)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return newPacketTable(hid), nil
}
