// Copyright ©2019 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import (
	"fmt"
	"math"
	"os"
	"testing"
)

/**
 * These test cases are based on the h5_cmprss.c by The HDF Group.
 * https://support.hdfgroup.org/HDF5/examples/intro.html#c
 */

func TestChunk(t *testing.T) {
	DisplayErrors(true)
	defer DisplayErrors(false)
	var (
		fn    = "test_chunk.h5"
		dsn   = "dset_chunk"
		dims  = []uint{1000, 1000}
		cdims = []uint{100, 100}
	)
	defer os.Remove(fn)

	dcpl, err := NewPropList(P_DATASET_CREATE)
	if err != nil {
		t.Fatal(err)
	}
	defer dcpl.Close()
	err = dcpl.SetChunk(cdims)
	if err != nil {
		t.Fatal(err)
	}

	cdimsChunk, err := dcpl.GetChunk(len(cdims))
	if err != nil {
		t.Fatal(err)
	}
	for i, cdim := range cdimsChunk {
		if cdim != cdims[i] {
			t.Fatalf("chunked dimensions mismatch: %d != %d", cdims[i], cdim)
		}
	}

	data0, err := save(fn, dsn, dims, dcpl)
	if err != nil {
		t.Fatal(err)
	}

	data1, err := load(fn, dsn, nil)
	if err != nil {
		t.Fatal(err)
	}

	if err := compare(data0, data1); err != nil {
		t.Fatal(err)
	}
}

func TestDeflate(t *testing.T) {
	DisplayErrors(true)
	defer DisplayErrors(false)
	var (
		fn    = "test_cmprss_deflate.h5"
		dsn   = "dset_cmpress"
		dims  = []uint{1000, 1000}
		cdims = []uint{100, 100}
	)
	defer os.Remove(fn)

	dcpl, err := NewPropList(P_DATASET_CREATE)
	if err != nil {
		t.Fatal(err)
	}
	defer dcpl.Close()
	err = dcpl.SetChunk(cdims)
	if err != nil {
		t.Fatal(err)
	}
	err = dcpl.SetDeflate(DefaultCompression)
	if err != nil {
		t.Fatal(err)
	}

	data0, err := save(fn, dsn, dims, dcpl)
	if err != nil {
		t.Fatal(err)
	}

	data1, err := load(fn, dsn, nil)
	if err != nil {
		t.Fatal(err)
	}

	if err := compare(data0, data1); err != nil {
		t.Fatal(err)
	}
}

func TestChunkCache(t *testing.T) {
	DisplayErrors(true)
	defer DisplayErrors(false)
	var (
		fn    = "test_chunk_cache.h5"
		dsn   = "dset_chunk_cache"
		dims  = []uint{1000, 1000}
		cdims = []uint{100, 100}
	)
	defer os.Remove(fn)

	dcpl, err := NewPropList(P_DATASET_CREATE)
	if err != nil {
		t.Fatal(err)
	}
	defer dcpl.Close()
	err = dcpl.SetChunk(cdims)
	if err != nil {
		t.Fatal(err)
	}

	cdimsChunk, err := dcpl.GetChunk(len(cdims))
	if err != nil {
		t.Fatal(err)
	}
	for i, cdim := range cdimsChunk {
		if cdim != cdims[i] {
			t.Fatalf("chunked dimensions mismatch: %d != %d", cdims[i], cdim)
		}
	}

	data0, err := save(fn, dsn, dims, dcpl)
	if err != nil {
		t.Fatal(err)
	}

	dapl, err := NewPropList(P_DATASET_ACCESS)
	if err != nil {
		t.Fatal(err)
	}
	defer dapl.Close()

	nslots, nbytes, w0, err := dapl.GetChunkCache()
	if err != nil {
		t.Fatal(err)
	}

	nslotsNew, nbytesNew, w0New := nslots*4, nbytes*2, w0/3
	if err := dapl.SetChunkCache(nslotsNew, nbytesNew, w0New); err != nil {
		t.Fatal(err)
	}
	if err := checkChunkCache(nslotsNew, nbytesNew, w0New, dapl); err != nil {
		t.Fatal(err)
	}

	data1, err := load(fn, dsn, dapl)
	if err != nil {
		t.Fatal(err)
	}
	if err := compare(data0, data1); err != nil {
		t.Fatal(err)
	}

	if err := dapl.SetChunkCache(D_CHUNK_CACHE_NSLOTS_DEFAULT, D_CHUNK_CACHE_NBYTES_DEFAULT, D_CHUNK_CACHE_W0_DEFAULT); err != nil {
		t.Fatal(err)
	}
	if err := checkChunkCache(nslots, nbytes, w0, dapl); err != nil {
		t.Fatal(err)
	}
}

func save(fn, dsn string, dims []uint, dcpl *PropList) ([]float64, error) {
	f, err := CreateFile(fn, F_ACC_TRUNC)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dspace, err := CreateSimpleDataspace(dims, dims)
	if err != nil {
		return nil, err
	}
	defer dspace.Close()

	dset, err := f.CreateDatasetWith(dsn, T_NATIVE_DOUBLE, dspace, dcpl)
	if err != nil {
		return nil, err
	}
	defer dset.Close()

	n := dims[0] * dims[1]
	data := make([]float64, n)
	for i := range data {
		data[i] = float64((i*i*i + 13) % 8191)
	}
	err = dset.Write(&data[0])
	if err != nil {
		return nil, err
	}
	return data, nil
}

func load(fn, dsn string, dapl *PropList) ([]float64, error) {
	f, err := OpenFile(fn, F_ACC_RDONLY)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var dset *Dataset
	if dapl == nil {
		dset, err = f.OpenDataset(dsn)
	} else {
		dset, err = f.OpenDatasetWith(dsn, dapl)
	}
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
		return fmt.Errorf("dimensions mismatch: %d != %d", n0, n1)
	}
	for i := 0; i < n0; i++ {
		d := math.Abs(ds0[i] - ds1[i])
		if d > 1e-7 {
			return fmt.Errorf("values at index %d differ: %f != %f", i, ds0[i], ds1[i])
		}
	}
	return nil
}

func checkChunkCache(nslots, nbytes int, w0 float64, dapl *PropList) error {
	nslotsCache, nbytesCache, w0Cache, err := dapl.GetChunkCache()
	if err != nil {
		return err
	}

	if nslotsCache != nslots {
		return fmt.Errorf("`nslots` mismatch: %d != %d", nslots, nslotsCache)
	}
	if nbytesCache != nbytes {
		return fmt.Errorf("`nbytes` mismatch: %d != %d", nbytes, nbytesCache)
	}
	if math.Abs(w0Cache-w0) > 1e-5 {
		return fmt.Errorf("`w0` mismatch: %.6f != %.6f", w0, w0Cache)
	}
	return nil
}
