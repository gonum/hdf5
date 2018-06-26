// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import (
	"runtime"
	"testing"
	"time"
)

func TestSimpleDatatypes(t *testing.T) {
	// Smoke tests for the simple datatypes
	tests := []interface{}{
		int(0),
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		uint(0),
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),
		float32(0),
		float64(0),
		string(""),
		bool(true),
	}

	for test := range tests {
		NewDatatypeFromValue(test)
		// Test again for usage with ptrs
		NewDatatypeFromValue(&test)
	}
}

// Test for array datatypes. Checks that the number of dimensions is correct.
func TestArrayDatatype(t *testing.T) {
	tests := map[int]interface{}{
		1: [8]int{1, 1, 2, 3, 5, 8, 13},
		2: [2][2]int{{1, 2}, {3, 4}},
		3: [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
	}

	for dims, val := range tests {
		dt, err := NewDatatypeFromValue(val)
		if err != nil {
			t.Fatal(err)
		}
		adt := ArrayType{*dt}
		if adt.NDims() != dims {
			t.Errorf("wrong number of dimensions: got %d, want %d", adt.NDims(), dims)
		}
	}
}

func TestStructDatatype(t *testing.T) {
	test := struct {
		a int32
		b string
		c struct {
			a int32
			b string
		}
	}{}

	// Test that the type can be constructed and that the number of
	// members is as expected.
	var dtypes []*Datatype

	// "Regular" value
	dtype, err := NewDatatypeFromValue(test)
	if err != nil {
		t.Fatal(err)
	}
	dtypes = append(dtypes, dtype)

	// pointer to value
	dtype, err = NewDatatypeFromValue(test)
	if err != nil {
		t.Fatal(err)
	}
	dtypes = append(dtypes, dtype)

	for _, dtype := range dtypes {
		dt := CompoundType{*dtype}
		if dt.NMembers() != 3 {
			t.Errorf("wrong number of members: got %d, want %d", dt.NMembers(), 3)
		}

		memberClasses := []TypeClass{
			T_INTEGER,
			T_STRING,
			T_COMPOUND,
		}
		// Test the member classes, and also test that they can be constructed
		for idx, class := range memberClasses {
			if dt.MemberClass(idx) != class {
				t.Errorf("wrong TypeClass: got %v, want %v", dt.MemberClass(idx), class)
			}
			_, err := dt.MemberType(idx)
			if err != nil {
				t.Fatal(err)
			}
		}

		// Test the member names
		memberNames := []string{"a", "b", "c"}
		for idx, name := range memberNames {
			if dt.MemberName(idx) != name {
				t.Errorf("wrong name: got %q, want %q", dt.MemberName(idx), name)
			}
			if dt.MemberIndex(name) != idx {
				t.Errorf("wrong index: got %d, want %d", dt.MemberIndex(name), idx)
			}
		}

		// Pack the datatype, otherwise offsets are implementation defined
		dt.Pack()
		memberOffsets := []int{0, 4, 12}
		for idx, offset := range memberOffsets {
			if dt.MemberOffset(idx) != offset {
				t.Errorf("wrong offset: got %d, want %d", dt.MemberOffset(idx), offset)
			}
		}
	}
}

func TestCloseBehavior(t *testing.T) {
	var s struct {
		a int
		b float64
	}
	dtype, err := NewDatatypeFromValue(s)
	if err != nil {
		t.Fatal(err)
	}
	defer dtype.Close()

	// Sleep to ensure GC runs before returning
	runtime.GC()
	time.Sleep(100 * time.Millisecond)
}
