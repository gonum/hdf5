// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

import (
	"os"
	"reflect"
	"testing"
)

func TestWriteAttribute(t *testing.T) {
	DisplayErrors(true)
	defer DisplayErrors(false)
	defer os.Remove(fname)

	f, err := CreateFile(fname, F_ACC_TRUNC)
	if err != nil {
		t.Fatalf("CreateFile failed: %s\n", err)
	}
	defer f.Close()

	scalar, err := CreateDataspace(S_SCALAR)
	if err != nil {
		t.Fatalf("CreateDataspace failed: %s\n", err)
	}
	defer scalar.Close()

	dset, err := f.CreateDataset("dset", T_NATIVE_USHORT, scalar)
	if err != nil {
		t.Fatalf("CreateDataset failed: %s\n", err)
	}
	defer dset.Close()

	strVal := "I am a string attribute"
	intVal := 42
	fltVal := 1.234
	arrVal := [3]byte{128, 0, 255}

	attrs := map[string]struct {
		Value interface{}
		Type  reflect.Type
	}{
		"My string attribute": {&strVal, reflect.TypeOf(strVal)},
		"My int attribute":    {&intVal, reflect.TypeOf(intVal)},
		"My float attribute":  {&fltVal, reflect.TypeOf(fltVal)},
		"My array attribute":  {&arrVal, reflect.TypeOf(arrVal)},
	}

	for name, v := range attrs {
		t.Run(name, func(t *testing.T) {
			dtype, err := NewDataTypeFromType(v.Type)
			if err != nil {
				t.Fatalf("NewDatatypeFromValue failed: %v", err)
			}
			defer dtype.Close()

			attr, err := dset.CreateAttribute(name, dtype, scalar)
			if err != nil {
				t.Fatalf("CreateAttribute failed: %v", err)
			}
			defer attr.Close()

			if err := attr.Write(v.Value, dtype); err != nil {
				t.Fatalf("Attribute write failed: %v", err)
			}

			got := reflect.New(v.Type).Elem()
			if err := attr.Read(got.Addr().Interface(), dtype); err != nil {
				t.Fatalf("Attribute read failed: %v", err)
			}

			want := reflect.ValueOf(v.Value).Elem().Interface()
			if !reflect.DeepEqual(want, got.Interface()) {
				t.Fatalf("round trip failed:\ngot = %v\nwant= %v\n", got.Interface(), want)
			}
		})
	}
}
