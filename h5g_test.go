// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import (
	"os"
	"testing"
)

func TestGroup(t *testing.T) {
	f, err := CreateFile(fname, F_ACC_TRUNC)
	if err != nil {
		t.Fatalf("CreateFile failed: %s", err)
	}
	defer os.Remove(fname)
	defer f.Close()

	g1, err := f.CreateGroup("foo")
	if err != nil {
		t.Fatalf("couldn't create group: %s", err)
	}
	if *g1.File() != *f {
		t.Fatal("wrong file for group")
	}
	if g1.Name() != "/foo" {
		t.Errorf("wrong Name for group: want %q, got %q", "/foo", g1.Name())
	}

	g2, err := g1.CreateGroup("bar")
	if err != nil {
		t.Fatalf("couldn't create group: %s", err)
	}
	if *g2.File() != *f {
		t.Fatal("wrong file for group")
	}
	if g2.Name() != "/foo/bar" {
		t.Errorf("wrong Name for group: want %q, got %q", "/foo/bar", g1.Name())
	}

	g3, err := g2.CreateGroup("baz")
	if err != nil {
		t.Fatalf("couldn't create group: %s", err)
	}
	if *g3.File() != *f {
		t.Fatal("wrong file for group")
	}
	if g3.Name() != "/foo/bar/baz" {
		t.Errorf("wrong Name for group: want %q, got %q", "/foo/bar/bar", g1.Name())
	}

	if nObjs, err := g2.NumObjects(); err != nil {
		t.Fatal(err)
	} else if nObjs != 1 {
		t.Errorf("wrong number of objects in group: want 1, got %d", nObjs)
	}

	err = g1.Close()
	if err != nil {
		t.Error(err)
	}
	err = g2.Close()
	if err != nil {
		t.Error(err)
	}
	err = g3.Close()
	if err != nil {
		t.Error(err)
	}

	g2, err = f.OpenGroup("/foo/bar")
	if err != nil {
		t.Fatal(err)
	}

	g3, err = g2.OpenGroup("baz")
	if err != nil {
		t.Fatal(err)
	}

	_, err = g3.OpenGroup("bs")
	if err == nil {
		t.Fatal("expected error on opening invalid group")
	}

	data := 5

	dtype, err := NewDataTypeFromValue(data)
	if err != nil {
		t.Fatal(err)
	}

	dims := []uint{1}
	dspace, err := CreateSimpleDataspace(dims, dims)
	if err != nil {
		t.Fatal(err)
	}

	dset, err := g3.CreateDataset("dset", dtype, dspace)
	if err != nil {
		t.Fatal(err)
	}

	dset2, err := g3.OpenDataset("dset")
	if dset.Name() != dset2.Name() {
		t.Error("expected dataset names to be equal")
	}

	dset2, err = g3.OpenDataset("bs")
	if err == nil {
		t.Errorf("opened dataset that was never created: %v", dset2)
	}

}
