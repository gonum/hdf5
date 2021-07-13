// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import (
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	f, err := CreateFile(fname, F_ACC_TRUNC)
	if err != nil {
		t.Fatalf("CreateFile failed: %s", err)
	}
	defer os.Remove(fname)
	defer f.Close()

	if fileName := f.FileName(); fileName != fname {
		t.Fatalf("FileName() have %v, want %v", fileName, fname)
	}
	// The file is also the root group
	if name := f.Name(); name != "/" {
		t.Fatalf("Name() have %v, want %v", name, fname)
	}
	if err := f.Flush(F_SCOPE_GLOBAL); err != nil {
		t.Fatalf("Flush() failed: %s", err)
	}
	if !IsHDF5(fname) {
		t.Fatalf("IsHDF5 returned false")
	}
	if n, err := f.NumObjects(); err != nil {
		t.Fatalf("NumObjects failed: %s", err)
	} else if n != 0 {
		t.Fatalf("empty file had %d objects", n)
	}

	if _, err := f.ObjectNameByIndex(0); err == nil {
		t.Fatalf("expected error")
	}

	f2 := f.File()
	defer f2.Close()
	fName := f.FileName()
	f2Name := f2.FileName()
	if fName != f2Name {
		t.Fatalf("f2 FileName() have %v, want %v", f2Name, fName)
	}

	// Test a Group
	groupName := "test"
	g, err := f.CreateGroup(groupName)
	if err != nil {
		t.Fatalf("CreateGroup() failed: %s", err)
	}
	defer g.Close()
	if name := g.Name(); name != "/"+groupName {
		t.Fatalf("Group Name() have %v, want /%v", name, groupName)
	}

	g2, err := f.OpenGroup(groupName)
	if err != nil {
		t.Fatalf("OpenGroup() failed: %s", err)
	}
	defer g2.Close()
	if name := g2.Name(); name != "/"+groupName {
		t.Fatalf("Group Name() have %v, want /%v", name, groupName)
	}

	if n, err := f.NumObjects(); err != nil {
		t.Fatalf("NumObjects failed: %s", err)
	} else if n != 1 {
		t.Fatalf("NumObjects: got %d, want %d", n, 1)
	}

	if name, err := f.ObjectNameByIndex(0); err != nil {
		t.Fatalf("ObjectNameByIndex failed: %s", err)
	} else if name != groupName {
		t.Fatalf("ObjectNameByIndex: got %q, want %q", name, groupName)
	}
	if _, err := f.ObjectNameByIndex(1); err == nil {
		t.Fatalf("expected error")
	}

	// Test a Dataset
	ds, err := CreateDataspace(S_SCALAR)
	if err != nil {
		t.Fatalf("CreateDataspace failed: %s", err)
	}
	defer ds.Close()
	dsetName := "test_dataset"
	dset, err := f.CreateDataset(dsetName, T_NATIVE_INT, ds)
	if err != nil {
		t.Fatalf("CreateDataset failed: %s", err)
	}
	defer dset.Close()
	if name := dset.Name(); name != "/"+dsetName {
		t.Fatalf("Dataset Name() have %v, want /%v", name, dsetName)
	}
	dFile := dset.File()
	if dFile.Name() != f.Name() {
		t.Fatalf("Dataset File() have %v, want %v", dFile.Name(), f.Name())
	}
	defer dFile.Close()

	if n, err := f.NumObjects(); err != nil {
		t.Fatalf("NumObjects failed: %s", err)
	} else if n != 2 {
		t.Fatalf("NumObjects: got %d, want %d", n, 1)
	}

	for i, n := range []string{groupName, dsetName} {
		if name, err := f.ObjectNameByIndex(uint(i)); err != nil {
			t.Fatalf("ObjectNameByIndex failed: %s", err)
		} else if name != n {
			t.Fatalf("ObjectNameByIndex: got %q, want %q", name, groupName)
		}
	}
}

func TestClosedFile(t *testing.T) {
	f, err := CreateFile(fname, F_ACC_TRUNC)
	if err != nil {
		t.Fatalf("CreateFile failed: %s", err)
	}
	fName := f.Name()
	f2 := f.File()
	f.Close()

	f2Name := f2.FileName()
	if f2Name != fname {
		t.Fatalf("f2 FileName() have %v, want %v", f2Name, fName)
	}
	f2.Close()

	os.Remove(fname)
	f3 := f.File()
	if f3 != nil {
		t.Fatalf("expected file to be nil")
	}

}
