// Copyright ©2017 The Gonum Authors. All rights reserved.
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

	if f.LinkExists("/foo") {
		t.Error("unexpected /foo link present")
	}

	g1, err := f.CreateGroup("foo")
	if err != nil {
		t.Fatalf("couldn't create group: %s", err)
	}
	g1f := g1.File()
	defer g1f.Close()
	if *g1f != *f {
		t.Fatal("wrong file for group")
	}
	if g1.Name() != "/foo" {
		t.Errorf(`wrong name for group: want:"/foo", got:%q`, g1.Name())
	}

	if !f.LinkExists("/foo") {
		t.Error(`unexpected "/foo" group link not present`)
	}

	if g1.LinkExists("bar") {
		t.Error(`unexpected "bar" child link for "/foo" group`)
	}

	g2, err := g1.CreateGroup("bar")
	if err != nil {
		t.Fatalf("couldn't create group: %s", err)
	}
	g2f := g2.File()
	defer g2f.Close()
	if *g2f != *f {
		t.Fatal("wrong file for group")
	}
	if g2.Name() != "/foo/bar" {
		t.Errorf("wrong Name for group: want %q, got %q", "/foo/bar", g1.Name())
	}
	if !g1.LinkExists("bar") {
		t.Error(`expected "bar" child link for "/foo" group not present`)
	}

	g3, err := g2.CreateGroup("baz")
	if err != nil {
		t.Fatalf("couldn't create group: %s", err)
	}
	g3f := g3.File()
	defer g3f.Close()
	if *g3f != *f {
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

	if name, err := g2.ObjectNameByIndex(0); err != nil {
		t.Fatalf("could not retrieve object name idx=%d: %+v", 0, err)
	} else if got, want := name, "baz"; got != want {
		t.Errorf("invalid name for object idx=%d: got=%q, want=%q", 0, got, want)
	}

	if typ, err := g2.ObjectTypeByIndex(0); err != nil {
		t.Fatalf("could not retrieve object type idx=%d: %+v", 0, err)
	} else if got, want := typ, H5G_GROUP; got != want {
		t.Errorf("invalid type for object idx=%d: got=%v, want=%v", 0, got, want)
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
	defer g2.Close()

	g3, err = g2.OpenGroup("baz")
	if err != nil {
		t.Fatal(err)
	}
	defer g3.Close()

	_, err = g3.OpenGroup("bs")
	if err == nil {
		t.Fatal("expected error on opening invalid group")
	}

	data := 5

	dtype, err := NewDatatypeFromValue(data)
	if err != nil {
		t.Fatal(err)
	}
	defer dtype.Close()

	dims := []uint{1}
	dspace, err := CreateSimpleDataspace(dims, dims)
	if err != nil {
		t.Fatal(err)
	}
	defer dspace.Close()

	dset, err := g3.CreateDataset("dset", dtype, dspace)
	if err != nil {
		t.Fatal(err)
	}
	defer dset.Close()

	dset2, err := g3.OpenDataset("dset")
	if dset.Name() != dset2.Name() {
		t.Error("expected dataset names to be equal")
	}
	defer dset2.Close()

	dset2, err = g3.OpenDataset("bs")
	if err == nil {
		t.Errorf("opened dataset that was never created: %v", dset2)
	}

}
