// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmem

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
	"unsafe"
)

var nativeEndian binary.ByteOrder

func init() {
	one := uint16(0x1)
	nativeEndian = [...]binary.ByteOrder{
		binary.BigEndian,
		binary.LittleEndian,
	}[(*[2]byte)(unsafe.Pointer(&one))[0]]
}

// Encoder is a wrapper type for information necessary to create and
// subsequently write an in memory object to a PacketTable using Append().
type Encoder struct {
	// Buf contains the encoded data.
	Buf    []byte
	offset int
}

// Encode encodes the value passed as data to binary form stored in []Buf. This
// buffer is a Go representation of a C in-memory object that can be appended
// to e.g. a HDF5 PacketTable.
//
// Struct values must only have exported fields, otherwise Encode will panic.
func (enc *Encoder) Encode(data interface{}) error {
	padding := enc.offset - len(enc.Buf)

	if padding > 0 {
		enc.Buf = append(enc.Buf, make([]byte, padding)...)
	}

	if data, ok := data.(CMarshaler); ok {
		raw, err := data.MarshalC()
		if err != nil {
			return err
		}

		enc.Buf = append(enc.Buf, raw...)

		return nil
	}

	rv := reflect.Indirect(reflect.ValueOf(data))
	if !rv.IsValid() {
		return fmt.Errorf("cmem: reflect.ValueOf returned invalid value for type %T", data)
	}

	rt := rv.Type()

	switch rt.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			if err := enc.Encode(rv.Index(i).Interface()); err != nil {
				return err
			}
		}
		return nil

	case reflect.Struct:
		offset := enc.offset
		for i := 0; i < rv.NumField(); i++ {
			sfv := rv.Field(i).Interface()
			// In order to keep the memory offset always correct, we use the
			// structs offset.
			enc.offset = offset + int(rt.Field(i).Offset)
			if err := enc.Encode(sfv); err != nil {
				return err
			}
			// Reset the offset to the correct array size.
			enc.offset = offset + int(rt.Size())
		}
		return nil

	case reflect.String:
		str := append([]byte(rv.String()), 0)

		// This direct machine conversion is only used
		// because HDF5 uses machine endianism.
		//
		// DO NOT DO THIS AT HOME.
		//
		raw := (*[unsafe.Sizeof(uintptr(0))]byte)(unsafe.Pointer(&str))
		enc.Buf = append(enc.Buf, raw[:]...)
		enc.offset += len(raw)

	case reflect.Ptr:
		return enc.Encode(rv.Elem())

	case reflect.Int8:
		enc.Buf = append(enc.Buf, byte(rv.Int()))
		enc.offset++

	case reflect.Uint8:
		enc.Buf = append(enc.Buf, byte(rv.Uint()))
		enc.offset++

	case reflect.Int16:
		var raw [2]byte
		nativeEndian.PutUint16(raw[:], uint16(rv.Int()))
		enc.Buf = append(enc.Buf, raw[:]...)
		enc.offset += 2

	case reflect.Uint16:
		var raw [2]byte
		nativeEndian.PutUint16(raw[:], uint16(rv.Uint()))
		enc.Buf = append(enc.Buf, raw[:]...)
		enc.offset += 2

	case reflect.Int32:
		var raw [4]byte
		nativeEndian.PutUint32(raw[:], uint32(rv.Int()))
		enc.Buf = append(enc.Buf, raw[:]...)
		enc.offset += 4

	case reflect.Uint32:
		var raw [4]byte
		nativeEndian.PutUint32(raw[:], uint32(rv.Uint()))
		enc.Buf = append(enc.Buf, raw[:]...)
		enc.offset += 4

	case reflect.Int64:
		var raw [8]byte
		nativeEndian.PutUint64(raw[:], uint64(rv.Int()))
		enc.Buf = append(enc.Buf, raw[:]...)
		enc.offset += 8

	case reflect.Uint64:
		var raw [8]byte
		nativeEndian.PutUint64(raw[:], uint64(rv.Uint()))
		enc.Buf = append(enc.Buf, raw[:]...)
		enc.offset += 8

	case reflect.Float32:
		var raw [4]byte
		nativeEndian.PutUint32(raw[:], math.Float32bits(float32(rv.Float())))
		enc.Buf = append(enc.Buf, raw[:]...)
		enc.offset += 4

	case reflect.Float64:
		var raw [8]byte
		nativeEndian.PutUint64(raw[:], math.Float64bits(rv.Float()))
		enc.Buf = append(enc.Buf, raw[:]...)
		enc.offset += 8

	case reflect.Bool:
		val := byte(0)
		if rv.Bool() {
			val = 1
		}
		enc.Buf = append(enc.Buf, val)
		enc.offset++

	default:
		return fmt.Errorf("hdf5: PT Append does not support datatype (%s)", rt.Kind())
	}

	return nil
}
