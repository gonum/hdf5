// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmem

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"testing"
)

func TestEncode(t *testing.T) {
	for i, tc := range []struct {
		v    interface{}
		want []byte
	}{
		{
			v: struct {
				V1 uint8
				V2 uint64
				V3 uint8
				V4 uint16
			}{1, 2, 3, 4},
			want: []byte{
				0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // |........|
				0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // |........|
				0x03, 0x00, 0x04, 0x00, /*                   */ // |....|
			},
		},
		{
			v: []int32{1, 20, 300, 4000, 50000, 600000, 7000000, 8000000},
			want: []byte{
				0x01, 0x00, 0x00, 0x00, 0x14, 0x00, 0x00, 0x00, // |........|
				0x2c, 0x01, 0x00, 0x00, 0xa0, 0x0f, 0x00, 0x00, // |,.......|
				0x50, 0xc3, 0x00, 0x00, 0xc0, 0x27, 0x09, 0x00, // |P....'..|
				0xc0, 0xcf, 0x6a, 0x00, 0x00, 0x12, 0x7a, 0x00, // |..j...z.|
			},
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			oldEndian := nativeEndian
			nativeEndian = binary.LittleEndian
			defer func() { nativeEndian = oldEndian }()
			var enc Encoder
			err := enc.Encode(tc.v)
			if err != nil {
				t.Fatalf("could not encode: %v", err)
			}
			if !bytes.Equal(enc.Buf, tc.want) {
				t.Fatalf("encoding error:\ngot = %v\nwant= %v", enc.Buf, tc.want)
			}
		})
	}
}
