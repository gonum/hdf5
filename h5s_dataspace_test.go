// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import (
	"testing"
)

func TestDataspace(t *testing.T) {
	// Creating this dataspace results in an error.
	if _, err := CreateDataspace(S_NO_CLASS); err == nil {
		t.Errorf("expected an error, but got nil")
	}

	// These dataspaces are legitimate.
	classes := []SpaceClass{S_SCALAR, S_SIMPLE, S_NULL}
	for _, class := range classes {
		// Create a new Dataspace
		ds, err := CreateDataspace(class)
		if err != nil {
			t.Fatal(err)
		}

		if ds.SimpleExtentType() != class {
			t.Errorf("Dataspace class mismatch: %q != %q", ds.SimpleExtentType(), class)
		}

		// Copy the Dataspace
		clone, err := ds.Copy()
		if err != nil {
			t.Fatal(err)
		}

		if ds.Name() != clone.Name() {
			t.Errorf("original dataspace name %q != clone name %q", ds.Name(), clone.Name())
		}
		if ds.IsSimple() != clone.IsSimple() {
			t.Errorf("original dataspace simplicity %v != clone simplicity: %v", ds.IsSimple(), clone.IsSimple())
		}
		// Close the Dataspace
		if err = ds.Close(); err != nil {
			t.Fatal(err)
		}
	}
}

func TestSimpleDataspace(t *testing.T) {
	dims := []uint{3, 3, 3}
	maxdims := []uint{9, 9, 9}
	ds, err := CreateSimpleDataspace(dims, maxdims)
	if err != nil {
		t.Fatal(err)
	}

	dsDims, dsMaxdims, err := ds.SimpleExtentDims()

	if err != nil {
		t.Fatal(err)
	}

	if !arrayEq(dims, dsDims) {
		t.Errorf("retrieved dims not equal: %v != %v", dims, dsDims)
	}

	if !arrayEq(maxdims, dsMaxdims) {
		t.Errorf("retrieved maxdims not equal: %v != %v", maxdims, dsMaxdims)
	}

	if ds.SimpleExtentNDims() != 3 {
		t.Errorf("wrong number of dimensions: got %d, want %d", ds.SimpleExtentNDims(), 3)
	}

	if ds.SimpleExtentType() != S_SIMPLE {
		t.Errorf("wrong extent type: got %d, want %d", ds.SimpleExtentType(), S_SIMPLE)
	}

	// npoints should be 3 * 3 * 3
	npoints := ds.SimpleExtentNPoints()
	if npoints != 27 {
		t.Errorf("wrong number of npoints: got %d, want %d", npoints, 27)
	}

	// SetOffset should only work for array a where len(a) == len(dims)
	if err = ds.SetOffset([]uint{1, 1, 1}); err != nil {
		t.Fatal(err)
	}

	if err = ds.SetOffset([]uint{1}); err == nil {
		t.Error("expected a non-nil error")
	}
}

func arrayEq(a, b []uint) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
