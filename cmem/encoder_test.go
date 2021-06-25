// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmem

import (
	"bytes"
	"testing"
	"encoding/binary"
)

func TestEncode(t *testing.T) {
	v := struct {
		V1 uint8
		V2 uint64
		V3 uint8
		V4 uint16
	}{1, 2, 3, 4}
	want := []byte{
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // |........|
		0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // |........|
		0x03, 0x00, 0x04, 0x00, /*                   */ // |....|
	}
	var enc Encoder
	err := enc.Encode(v)
	if err != nil {
		t.Fatalf("could not encode: %v", err)
	}
	if !bytes.Equal(enc.Buf, want) {
		t.Fatalf("encoding error:\ngot = %v\nwant= %v", enc.Buf, want)
	}
}

func TestEncodeSlice(t *testing.T) {
	v := []int32{1, 20, 300, 4000, 50000, 600000, 7000000, 8000000}
	want := []byte{
		0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // |........|
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // |,.......|
	}
	var enc Encoder
	err := enc.Encode(v)
	if err != nil {
		t.Fatalf("could not encode: %v", err)
	}
	pointer := enc.pointerSlice[0]
	bs := make([]byte, 8)
        binary.LittleEndian.PutUint64(bs, uint64(uintptr(pointer)))
	copy(want[8:16], bs[:])

	if !bytes.Equal(enc.Buf, want) {
		t.Fatalf("encoding error:\ngot = %v\nwant= %v", enc.Buf, want)
	}

}
