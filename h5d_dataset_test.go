// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import (
	"os"
	"reflect"
	"testing"
)

func createDataset1(t *testing.T) error {
	// create a file with a single 5x20 dataset
	f, err := CreateFile(fname, F_ACC_TRUNC)
	if err != nil {
		t.Fatalf("CreateFile failed: %s", err)
		return err
	}
	defer f.Close()

	var data [100]uint16
	for i := range data {
		data[i] = uint16(i)
	}

	dims := []uint{20, 5}
	dspace, err := CreateSimpleDataspace(dims, dims)
	if err != nil {
		t.Fatal(err)
		return err
	}
	defer dspace.Close()

	dset, err := f.CreateDataset("dset", T_NATIVE_USHORT, dspace)
	if err != nil {
		t.Fatal(err)
		return err
	}
	defer dset.Close()

	err = dset.Write(&data[0])
	if err != nil {
		t.Fatal(err)
		return err
	}
	return err
}

/**
 * TestReadSubset based on the h5_subset.c sample with the HDF5 C library.
 * Original copyright notice:
 *
 * HDF5 (Hierarchical Data Format 5) Software Library and Utilities
 * Copyright 2006-2013 by The HDF Group.
 *
 * NCSA HDF5 (Hierarchical Data Format 5) Software Library and Utilities
 * Copyright 1998-2006 by the Board of Trustees of the University of Illinois.
 ****
 * Write some test data then read back a subset.
 */
func TestReadSubset(t *testing.T) {
	DisplayErrors(true)
	defer DisplayErrors(false)
	defer os.Remove(fname)
	err := createDataset1(t)
	if err != nil {
		return
	}

	// load a subset of the data
	f, err := OpenFile(fname, F_ACC_RDONLY)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	dset, err := f.OpenDataset("dset")
	if err != nil {
		t.Fatal(err)
	}
	defer dset.Close()

	// get the filespace and select the subset
	filespace := dset.Space()
	defer filespace.Close()
	offset, stride, count, block := [2]uint{5, 1}, [2]uint{1, 1}, [2]uint{5, 2}, [2]uint{1, 1}
	err = filespace.SelectHyperslab(offset[:], stride[:], count[:], block[:])
	if err != nil {
		t.Fatal(err)
	}

	// create the memory space for the subset
	dims, maxdims := [2]uint{2, 5}, [2]uint{2, 5}
	if err != nil {
		t.Fatal(err)
	}
	memspace, err := CreateSimpleDataspace(dims[:], maxdims[:])
	if err != nil {
		t.Fatal(err)
	}
	defer memspace.Close()

	expected := [10]uint16{26, 27, 31, 32, 36, 37, 41, 42, 46, 47}

	// test array
	{
		// create a buffer for the data
		data := [10]uint16{}

		// read the subset
		err = dset.ReadSubset(&data, memspace, filespace)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(data, expected) {
			t.Fatalf("ReadSubset-array error\ngot= %#v\nwant=%#v\n", data, expected)
		}
	}

	// test slice
	{
		// create a buffer for the data
		data := make([]uint16, 10)

		// read the subset
		err = dset.ReadSubset(&data, memspace, filespace)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(data, expected[:]) {
			t.Fatalf("ReadSubset-slice error\ngot= %#v\nwant=%#v\n", data, expected[:])
		}
	}
}

func TestWriteSubset(t *testing.T) {
	DisplayErrors(true)
	defer DisplayErrors(false)
	defer os.Remove(fname)

	fdims := []uint{12, 4, 6}
	fspace, err := CreateSimpleDataspace(fdims, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer fspace.Close()
	mdims := []uint{2, 6}
	mspace, err := CreateSimpleDataspace(mdims, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer mspace.Close()

	f, err := CreateFile(fname, F_ACC_TRUNC)
	if err != nil {
		t.Fatalf("CreateFile failed: %s\n", err)
	}
	defer f.Close()

	dset, err := f.CreateDataset("dset", T_NATIVE_USHORT, fspace)
	if err != nil {
		t.Fatal(err)
	}
	defer dset.Close()

	offset := []uint{6, 0, 0}
	stride := []uint{3, 1, 1}
	count := []uint{mdims[0], 1, mdims[1]}
	if err = fspace.SelectHyperslab(offset, stride, count, nil); err != nil {
		t.Fatal(err)
	}

	data := make([]uint16, mdims[0]*mdims[1])

	if err = dset.WriteSubset(&data, mspace, fspace); err != nil {
		t.Fatal(err)
	}
}

func TestSelectHyperslab(t *testing.T) {
	DisplayErrors(true)
	defer DisplayErrors(false)
	defer os.Remove(fname)

	dims := []uint{12, 4}
	dspace, err := CreateSimpleDataspace(dims, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer dspace.Close()
	offset, stride, count, block := []uint{1, 2}, []uint{2, 1}, []uint{4, 2}, []uint{1, 1}
	if err = dspace.SelectHyperslab(offset, stride, count, block); err != nil {
		t.Fatal(err)
	}
	if err = dspace.SelectHyperslab(offset, nil, count, block); err != nil {
		t.Fatal(err)
	}
	if err = dspace.SelectHyperslab(offset, stride, count, nil); err != nil {
		t.Fatal(err)
	}
	if err = dspace.SelectHyperslab(offset, nil, count, nil); err != nil {
		t.Fatal(err)
	}
}
