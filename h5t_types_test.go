// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import "testing"

func TestSimpleDatatypes(t *testing.T) {
	// Smoke tests for the simple datatypes
	tests := []struct {
		v             interface{}
		hasIllegalPtr bool
	}{
		{v: int(0), hasIllegalPtr: false},
		{v: int8(0), hasIllegalPtr: false},
		{v: int16(0), hasIllegalPtr: false},
		{v: int32(0), hasIllegalPtr: false},
		{v: int64(0), hasIllegalPtr: false},
		{v: uint(0), hasIllegalPtr: false},
		{v: uint8(0), hasIllegalPtr: false},
		{v: uint16(0), hasIllegalPtr: false},
		{v: uint32(0), hasIllegalPtr: false},
		{v: uint64(0), hasIllegalPtr: false},
		{v: float32(0), hasIllegalPtr: false},
		{v: float64(0), hasIllegalPtr: false},
		{v: string(""), hasIllegalPtr: false},
		{v: ([]int)(nil), hasIllegalPtr: false},
		{v: [1]int{0}, hasIllegalPtr: false},
		{v: bool(true), hasIllegalPtr: false},
		{v: (*int)(nil), hasIllegalPtr: false},
		{v: (*int8)(nil), hasIllegalPtr: false},
		{v: (*int16)(nil), hasIllegalPtr: false},
		{v: (*int32)(nil), hasIllegalPtr: false},
		{v: (*int64)(nil), hasIllegalPtr: false},
		{v: (*uint)(nil), hasIllegalPtr: false},
		{v: (*uint8)(nil), hasIllegalPtr: false},
		{v: (*uint16)(nil), hasIllegalPtr: false},
		{v: (*uint32)(nil), hasIllegalPtr: false},
		{v: (*uint64)(nil), hasIllegalPtr: false},
		{v: (*float32)(nil), hasIllegalPtr: false},
		{v: (*float64)(nil), hasIllegalPtr: false},
		{v: (*string)(nil), hasIllegalPtr: true},
		{v: (*[]int)(nil), hasIllegalPtr: true},
		{v: (*[1]int)(nil), hasIllegalPtr: false},
		{v: (*bool)(nil), hasIllegalPtr: false},
		{v: (**int)(nil), hasIllegalPtr: true},
		{v: (**int8)(nil), hasIllegalPtr: true},
		{v: (**int16)(nil), hasIllegalPtr: true},
		{v: (**int32)(nil), hasIllegalPtr: true},
		{v: (**int64)(nil), hasIllegalPtr: true},
		{v: (**uint)(nil), hasIllegalPtr: true},
		{v: (**uint8)(nil), hasIllegalPtr: true},
		{v: (**uint16)(nil), hasIllegalPtr: true},
		{v: (**uint32)(nil), hasIllegalPtr: true},
		{v: (**uint64)(nil), hasIllegalPtr: true},
		{v: (**float32)(nil), hasIllegalPtr: true},
		{v: (**float64)(nil), hasIllegalPtr: true},
		{v: (**string)(nil), hasIllegalPtr: true},
		{v: (**[]int)(nil), hasIllegalPtr: true},
		{v: (**[1]int)(nil), hasIllegalPtr: true},
		{v: (**bool)(nil), hasIllegalPtr: true},
	}

	for _, test := range tests {
		dt, err := NewDatatypeFromValue(test.v)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			continue
		}
		gotIllegalPtr := dt.hasIllegalGoPointer()
		if gotIllegalPtr != test.hasIllegalPtr {
			t.Errorf("unexpected illegal pointer status for %T: got:%t want:%t", test.v, gotIllegalPtr, test.hasIllegalPtr)
		}
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
		if dt.hasIllegalGoPointer() {
			t.Errorf("unexpected illegal pointer for %T", val)
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
	if dtype.hasIllegalGoPointer() {
		t.Errorf("unexpected illegal pointer for %T", test)
	}
	dtypes = append(dtypes, dtype)

	// pointer to value
	dtype, err = NewDatatypeFromValue(&test)
	if err != nil {
		t.Fatal(err)
	}
	if !dtype.hasIllegalGoPointer() {
		t.Errorf("expected illegal pointer for %T", &test)
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
}
