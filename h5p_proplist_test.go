// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import (
	"errors"
	"math"
	"os"
	"testing"
)

/**
 * These test cases are based on the h5_cmprss.c by The HDF Group.
 * https://support.hdfgroup.org/HDF5/examples/intro.html#c
 */

func TestDeflate(t *testing.T) {
	DisplayErrors(true)
	fn, dsn, dims := "cmprss_deflate.h5", "dset_cmpress", []uint{1000, 1000}
	defer DisplayErrors(false)
	defer os.Remove(fn)

	dclp, err := NewPropList(H5P_DATASET_CREATE)
	if err != nil {
		t.Fatal(err)
	}
	defer dclp.Close()
	err = dclp.SetChunk(2, []uint{100, 100})
	if err != nil {
		t.Fatal(err)
	}
	err = dclp.SetDeflate(7)
	if err != nil {
		t.Fatal(err)
	}

	data0, err := save(fn, dsn, dims, dclp)
	if err != nil {
		t.Fatal(err)
	}

	data1, err := load(fn, dsn)
	if err != nil {
		t.Fatal(err)
	}

	if err := compare(data0, data1); err != nil {
		t.Fatal(err)
	}
}

func save(fn, dsn string, dims []uint, dclp *PropList) ([]float64, error) {
	f, err := CreateFile(fn, F_ACC_TRUNC)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dspace, err := CreateSimpleDataspace(dims, dims)
	if err != nil {
		return nil, err
	}

	dset, err := f.CreateDatasetWith(dsn, T_NATIVE_DOUBLE, dspace, dclp)
	if err != nil {
		return nil, err
	}
	defer dset.Close()

	n := dims[0] * dims[1]
	data := make([]float64, n)
	for i := uint(0); i < n; i++ {
		data[i] = float64((i*i*i + 13) % 8191)
	}
	err = dset.Write(&data[0])
	if err != nil {
		return nil, err
	}
	return data, nil
}

func load(fn, dsn string) ([]float64, error) {
	f, err := OpenFile(fn, F_ACC_RDONLY)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dset, _ := f.OpenDataset(dsn)
	if err != nil {
		return nil, err
	}
	defer dset.Close()

	dims, _, err := dset.Space().SimpleExtentDims()
	if err != nil {
		return nil, err
	}

	data := make([]float64, dims[0]*dims[1])
	dset.Read(&data[0])
	return data, nil
}

func compare(ds0, ds1 []float64) error {
	n0, n1 := len(ds0), len(ds1)
	if n0 != n1 {
		return errors.New("Dimensions not meet")
	}
	for i := 0; i < n0; i++ {
		d := math.Abs(ds0[i] - ds1[i])
		if d > 1e-7 {
			return errors.New("Values not meet")
		}
	}
	return nil
}
