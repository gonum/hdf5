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
		t.Fatalf("err: %s", "/foo shouldn't exist at this time")
	}

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

	if !f.LinkExists("/foo") {
		t.Fatalf("err: %s", "/foo should exist at this time")
	}

	if g1.LinkExists("bar") {
		t.Fatalf("err: %s", "/foo shouldn't have bar as a child bar at this time")
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
		t.Fatalf("err: %s", "/foo shouldn't have bar as a child bar at this time")
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
		t.Fatalf("CreateFile failed: %s", err)
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
		t.Fatalf("image decoding failed: %s", err)
	}
	g1, err := f.CreateGroup("foo")
	if err != nil {
		t.Fatalf("couldn't create group: %s", err)
	}
	defer g1.Close()
	err = g1.CreateTrueImage("image", img)
	if err != nil {
		t.Fatalf("image saving failed: %s", err)
	}
	imgRead, err := g1.ReadTrueImage("image")
	if err != nil {
		t.Fatalf("image reading failed: %s", err)
	}
	widthGot := imgRead.Bounds().Max.X
	heightGot := imgRead.Bounds().Max.Y
	if widthGot != 1000 || heightGot != 500 {
		t.Fatalf("image dimension miss match: Got %d * %d, suppose to be 1000x500", widthGot, heightGot)
	}

	imgfile, err := os.Create("img.jpg")
	if err != nil {
		t.Fatalf("image file creation failed: %s", err)
	}
	defer os.Remove("img.jpg")
	defer imgfile.Close()

	jpeg.Encode(imgfile, imgRead, nil)
}
