// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import (
	"image"
	"image/color"
	"image/jpeg"
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
	if *g1.File() != *f {
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
	if *g2.File() != *f {
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

	g3, err = g2.OpenGroup("baz")
	if err != nil {
		t.Fatal(err)
	}

	_, err = g3.OpenGroup("bs")
	if err == nil {
		t.Fatal("expected error on opening invalid group")
	}

	data := 5

	dtype, err := NewDatatypeFromValue(data)
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

func TestImage(t *testing.T) {
	f, err := CreateFile(fname, F_ACC_TRUNC)
	if err != nil {
		t.Fatalf("CreateFile failed: %v", err)
	}
	defer os.Remove(fname)
	defer f.Close()
	img := image.NewRGBA(image.Rect(0, 0, 1000, 500))
	for y := 200; y < 300; y++ {
		for x := 400; x < 600; x++ {
			img.Set(x, y, color.RGBA{255, 0, 255, 255})
		}
	}
	if err != nil {
		t.Fatalf("image decoding failed: %v", err)
	}
	g1, err := f.CreateGroup("foo")
	if err != nil {
		t.Fatalf("couldn't create group: %v", err)
	}
	defer g1.Close()
	err = g1.CreateImage("image", img)
	if err != nil {
		t.Fatalf("image saving failed: %v", err)
	}
	imgRead, err := g1.ReadImage("image")
	if err != nil {
		t.Fatalf("image reading failed: %v", err)
	}
	gotWidth := imgRead.Bounds().Max.X
	gotHeight := imgRead.Bounds().Max.Y
	if gotWidth != 1000 || gotHeight != 500 {
		t.Errorf("image dimension mismatch: got %dx%d, want:1000x500", gotWidth, gotHeight)
	}

	imgfile, err := os.Create("img.jpg")
	if err != nil {
		t.Fatalf("image file creation failed: %v", err)
	}
	defer os.Remove("img.jpg")
	defer imgfile.Close()

	err = jpeg.Encode(imgfile, imgRead, nil)
	if err != nil {
		t.Errorf("unexpected error saving image: %v", err)
	}
}
