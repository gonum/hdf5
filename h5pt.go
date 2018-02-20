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

func extractStructValues(rv reflect.Value, rt reflect.Type) (ptr unsafe.Pointer, err error) {
	// Initially the memory will be allocated to 64 bytes. If more is required
	// later, this will be increased.
	memSize := C.size_t(64)
	ptr = C.malloc(memSize)

	for i := 0; i < rv.NumField(); i++ {
		var dataPtr unsafe.Pointer
		var dataSize C.size_t
		f := rv.Field(i)
		ft := f.Type().Kind()
		offset := C.size_t(rt.Field(i).Offset)

		switch ft {
		case reflect.Int8:
			val := C.int8_t(int8(f.Int()))
			dataPtr = unsafe.Pointer(&val)
			dataSize = C.size_t(unsafe.Sizeof(val))
		case reflect.Uint8:
			val := C.uint8_t(uint8(f.Uint()))
			dataPtr = unsafe.Pointer(&val)
			dataSize = C.size_t(unsafe.Sizeof(val))
		case reflect.Int16:
			val := C.uint8_t(int16(f.Int()))
			dataPtr = unsafe.Pointer(&val)
			dataSize = C.size_t(unsafe.Sizeof(val))
		case reflect.Uint16:
			val := C.uint8_t(uint16(f.Uint()))
			dataPtr = unsafe.Pointer(&val)
			dataSize = C.size_t(unsafe.Sizeof(val))
		case reflect.Int32:
			val := C.uint8_t(int32(f.Int()))
			dataPtr = unsafe.Pointer(&val)
			dataSize = C.size_t(unsafe.Sizeof(val))
		case reflect.Uint32:
			val := C.uint8_t(uint32(f.Uint()))
			dataPtr = unsafe.Pointer(&val)
			dataSize = C.size_t(unsafe.Sizeof(val))
		case reflect.Int64:
			val := C.int64_t(int64(f.Int()))
			dataPtr = unsafe.Pointer(&val)
			dataSize = C.size_t(unsafe.Sizeof(val))
		case reflect.Uint64:
			val := C.uint64_t(uint64(f.Uint()))
			dataPtr = unsafe.Pointer(&val)
			dataSize = C.size_t(unsafe.Sizeof(val))
		case reflect.Float32:
			val := C.float(float32(f.Float()))
			dataPtr = unsafe.Pointer(&val)
			dataSize = C.size_t(unsafe.Sizeof(val))
		case reflect.Float64:
			val := C.double(float64(f.Float()))
			dataPtr = unsafe.Pointer(&val)
			dataSize = C.size_t(unsafe.Sizeof(val))
		case reflect.Bool:
			val := C.uchar(0)
			if f.Bool() {
				val = 1
			}
			dataPtr = unsafe.Pointer(&val)
			dataSize = C.size_t(unsafe.Sizeof(val))
		default:
			err = fmt.Errorf("hdf5: Could not append struct member %s "+
				"to PacketTable", ft)
			return nil, err
		}

		// If the earlier allocated memory is not enough, we will increase it
		// by another 64 bytes.
		if memSize < (dataSize + offset) {
			memSize += 64
			ptr = C.realloc(ptr, memSize)
			C.memset(unsafe.Pointer(uintptr(ptr)+uintptr(offset)), 0, 64)
		}
		C.memcpy(unsafe.Pointer(uintptr(ptr)+uintptr(offset)), dataPtr, dataSize)
	}

	return ptr, nil
}

// Append appends packets to the end of a packet table.
func (t *Table) Append(data interface{}) (err error) {
	rv := reflect.Indirect(reflect.ValueOf(data))
	rp := reflect.Indirect(reflect.ValueOf(&data))
	rt := rv.Type()
	cNrecords := C.size_t(1)
	cData := unsafe.Pointer(nil)

	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			if err = t.Append(rv.Index(i).Interface()); err != nil {
				return err
			}
		}
		return nil

	case reflect.Struct:
		if cData, err = extractStructValues(rv, rt); err != nil {
			return err
		}
		defer C.free(cData)

	case reflect.String:
		stringData := C.CString(rv.String())
		defer C.free(unsafe.Pointer(stringData))
		cData = unsafe.Pointer(&stringData)

	case reflect.Ptr:
		ptrVal := rp.Elem()
		cData = unsafe.Pointer(&ptrVal)

	case reflect.Bool:
		val := C.uchar(0)
		if data.(bool) {
			val = 1
		}
		cData = unsafe.Pointer(&val)

	case reflect.Int8:
		val := C.int8_t(data.(int8))
		cData = unsafe.Pointer(&val)

	case reflect.Uint8:
		val := C.uint8_t(data.(uint8))
		cData = unsafe.Pointer(&val)

	case reflect.Uint16:
		val := C.uint16_t(data.(uint16))
		cData = unsafe.Pointer(&val)

	case reflect.Int16:
		val := C.int16_t(data.(int16))
		cData = unsafe.Pointer(&val)

	case reflect.Int32:
		val := C.int32_t(data.(int32))
		cData = unsafe.Pointer(&val)

	case reflect.Uint32:
		val := C.uint32_t(data.(uint32))
		cData = unsafe.Pointer(&val)

	case reflect.Int64:
		val := C.int64_t(data.(int64))
		cData = unsafe.Pointer(&val)

	case reflect.Uint64:
		val := C.uint64_t(data.(uint64))
		cData = unsafe.Pointer(&val)

	case reflect.Float32:
		val := C.float(data.(float32))
		cData = unsafe.Pointer(&val)

	case reflect.Float64:
		val := C.double(data.(float64))
		cData = unsafe.Pointer(&val)

	default:
		return fmt.Errorf("hdf5: PT Append does not support datatype (%s).", rt.Kind())
	}

	return h5err(C.H5PTappend(t.id, cNrecords, cData))
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
