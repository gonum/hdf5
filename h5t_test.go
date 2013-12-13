package hdf5

import (
	"testing"
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
	}

	for test := range tests {
		NewDatatypeFromValue(test)
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
		dt := NewDatatypeFromValue(val)
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
	dt := CompoundType{*NewDatatypeFromValue(test)}
	if dt.NMembers() != 3 {
		t.Errorf("wrong number of members: got %d, want %d", dt.NMembers(), 3)
	}

	member_classes := []TypeClass{
		T_INTEGER,
		T_STRING,
		T_COMPOUND,
	}
	// Test the member classes, and also test that they can be constructed
	for idx, class := range member_classes {
		if dt.MemberClass(idx) != class {
			t.Errorf("wrong TypeClass: got %v, want %v", dt.MemberClass(idx), class)
		}
		_, err := dt.MemberType(idx)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Test the member names
	member_names := []string{"a", "b", "c"}
	for idx, name := range member_names {
		if dt.MemberName(idx) != name {
			t.Errorf("wrong name: got %q, want %q", dt.MemberName(idx), name)
		}
		if dt.MemberIndex(name) != idx {
			t.Errorf("wrong index: got %d, want %d", dt.MemberIndex(name), idx)
		}
	}

	// Pack the datatype, otherwise offsets are implementation defined
	dt.Pack()
	member_offsets := []int{0, 4, 12}
	for idx, offset := range member_offsets {
		if dt.MemberOffset(idx) != offset {
			t.Errorf("wrong offset: got %d, want %d", dt.MemberOffset(idx), offset)
		}
	}
}
