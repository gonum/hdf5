package hdf5

/*
 #cgo LDFLAGS: -lhdf5 -lhdf5_hl
 #include "hdf5.h"
 #include "hdf5_hl.h"

 #include <stdlib.h>
 #include <string.h>
 */
import "C"

import (
	"os"
	//"io"
	"reflect"
)

type Encoder interface {
	Encode(v interface{}) os.Error
	Close() os.Error
}

func NewEncoder(f *File) Encoder {
	e := &encoder{f:f, t:nil}
	return e
}

type encoder struct {
	f *File
	t *Table
}

func (e *encoder) Close() os.Error {
	e.t.Close()
	return e.f.Close()
}

func (e *encoder) Encode(v interface{}) os.Error {
	rt := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)
	if rt.Kind() == reflect.Ptr {
		rv = rv.Elem()
		rt = rt.Elem()
	}
	
	if e.t == nil {
		chunk_size := 10
		compress := 3
		t, err := e.f.CreateTableFrom("table", rt, chunk_size, compress)
		if err != nil {
			return err
		}
		e.t = t
	}
	return e.t.Append(rv.Interface())
}

// EOF
