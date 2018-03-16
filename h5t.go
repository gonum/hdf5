// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

type Datatype struct {
	Identifier
}

type TypeClass C.H5T_class_t

const (
	T_NO_CLASS  TypeClass = -1 // Error
	T_INTEGER   TypeClass = 0  // integer types
	T_FLOAT     TypeClass = 1  // floating-point types
	T_TIME      TypeClass = 2  // date and time types
	T_STRING    TypeClass = 3  // character string types
	T_BITFIELD  TypeClass = 4  // bit field types
	T_OPAQUE    TypeClass = 5  // opaque types
	T_COMPOUND  TypeClass = 6  // compound types
	T_REFERENCE TypeClass = 7  // reference types
	T_ENUM      TypeClass = 8  // enumeration types
	T_VLEN      TypeClass = 9  // variable-length types
	T_ARRAY     TypeClass = 10 // array types
	T_NCLASSES  TypeClass = 11 // nbr of classes -- MUST BE LAST
)

// list of go types
var (
	_go_string_t reflect.Type = reflect.TypeOf(string(""))

	_go_boolean_t reflect.Type = reflect.TypeOf(bool(false))

	_go_int_t    reflect.Type = reflect.TypeOf(int(0))
	_go_int8_t   reflect.Type = reflect.TypeOf(int8(0))
	_go_int16_t  reflect.Type = reflect.TypeOf(int16(0))
	_go_int32_t  reflect.Type = reflect.TypeOf(int32(0))
	_go_int64_t  reflect.Type = reflect.TypeOf(int64(0))
	_go_uint_t   reflect.Type = reflect.TypeOf(uint(0))
	_go_uint8_t  reflect.Type = reflect.TypeOf(uint8(0))
	_go_uint16_t reflect.Type = reflect.TypeOf(uint16(0))
	_go_uint32_t reflect.Type = reflect.TypeOf(uint32(0))
	_go_uint64_t reflect.Type = reflect.TypeOf(uint64(0))

	_go_float32_t reflect.Type = reflect.TypeOf(float32(0))
	_go_float64_t reflect.Type = reflect.TypeOf(float64(0))

	_go_array_t reflect.Type = reflect.TypeOf([1]int{0})
	_go_slice_t reflect.Type = reflect.TypeOf([]int{0})

	_go_struct_t reflect.Type = reflect.TypeOf(struct{}{})

	_go_ptr_t reflect.Type = reflect.PtrTo(_go_int_t)
)

type typeMap map[TypeClass]reflect.Type

var (
	// Mapping of TypeClass to reflect.Type
	typeClassToGoType typeMap = typeMap{
		T_NO_CLASS:  nil,
		T_INTEGER:   _go_int_t,
		T_FLOAT:     _go_float32_t,
		T_TIME:      nil,
		T_STRING:    _go_string_t,
		T_BITFIELD:  nil,
		T_OPAQUE:    nil,
		T_COMPOUND:  _go_struct_t,
		T_REFERENCE: _go_ptr_t,
		T_ENUM:      _go_int_t,
		T_VLEN:      _go_slice_t,
		T_ARRAY:     _go_array_t,
	}

	parametricTypes typeMap = typeMap{
		// Only these types can be used with CreateDatatype
		T_COMPOUND: _go_struct_t,
		T_ENUM:     _go_int_t,
		T_OPAQUE:   nil,
		T_STRING:   _go_string_t,
	}
)

// OpenDatatype opens a named datatype.
func OpenDatatype(c CommonFG, name string, tapl_id int) (*Datatype, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	id := C.H5Topen2(C.hid_t(c.id), c_name, C.hid_t(tapl_id))
	if err := checkID(id); err != nil {
		return nil, err
	}
	return NewDatatype(id), nil
}

// NewDatatype creates a Datatype from an hdf5 id.
func NewDatatype(id C.hid_t) *Datatype {
	t := &Datatype{Identifier{id}}
	runtime.SetFinalizer(t, (*Datatype).finalizer)
	return t
}

// CreateDatatype creates a new datatype.
// class must be T_COMPUND, T_OPAQUE, T_ENUM or T_STRING.
// size is the size of the new datatype in bytes.
func CreateDatatype(class TypeClass, size int) (*Datatype, error) {
	_, ok := parametricTypes[class]
	if !ok {
		return nil,
			fmt.Errorf(
				"invalid TypeClass, want %v, %v, %v or %v, got %v",
				T_COMPOUND, T_OPAQUE, T_STRING, T_ENUM,
				class,
			)
	}

	hid := C.H5Tcreate(C.H5T_class_t(class), C.size_t(size))
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return NewDatatype(hid), nil
}

func (t *Datatype) finalizer() {
	if err := t.Close(); err != nil {
		panic(fmt.Errorf("error closing datatype: %s", err))
	}
}

// GoType returns the reflect.Type associated with the Datatype's TypeClass
func (t *Datatype) GoType() reflect.Type {
	return typeClassToGoType[t.Class()]
}

// Close releases a datatype.
func (t *Datatype) Close() error {
	if t.id == 0 {
		return nil
	}
	err := h5err(C.H5Tclose(t.id))
	t.id = 0
	return err
}

// Committed determines whether a datatype is a named type or a transient type.
func (t *Datatype) Committed() bool {
	return C.H5Tcommitted(t.id) > 0
}

// Copy copies an existing datatype.
func (t *Datatype) Copy() (*Datatype, error) {
	return copyDatatype(t.id)
}

// copyDatatype should be called by any function wishing to return
// an existing Datatype from a Dataset or Attribute.
func copyDatatype(id C.hid_t) (*Datatype, error) {
	hid := C.H5Tcopy(id)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return NewDatatype(hid), nil
}

// Equal determines whether two datatype identifiers refer to the same datatype.
func (t *Datatype) Equal(o *Datatype) bool {
	return C.H5Tequal(t.id, o.id) > 0
}

// Lock locks a datatype.
func (t *Datatype) Lock() error {
	return h5err(C.H5Tlock(t.id))
}

// Size returns the size of the Datatype.
func (t *Datatype) Size() uint {
	return uint(C.H5Tget_size(t.id))
}

// SetSize sets the total size of a Datatype.
func (t *Datatype) SetSize(sz int) error {
	err := C.H5Tset_size(t.id, C.size_t(sz))
	return h5err(err)
}

type ArrayType struct {
	Datatype
}

// NewArrayType creates a new ArrayType.
// base_type specifies the element type of the array.
// dims specify the dimensions of the array.
func NewArrayType(base_type *Datatype, dims []int) (*ArrayType, error) {
	ndims := C.uint(len(dims))
	c_dims := (*C.hsize_t)(unsafe.Pointer(&dims[0]))

	hid := C.H5Tarray_create2(base_type.id, ndims, c_dims)
	if err := checkID(hid); err != nil {
		return nil, err
	}
	t := &ArrayType{Datatype{Identifier{hid}}}
	runtime.SetFinalizer(t, (*ArrayType).finalizer)
	return t, nil
}

// NDims returns the rank of an ArrayType.
func (t *ArrayType) NDims() int {
	return int(C.H5Tget_array_ndims(t.id))
}

// ArrayDims returns the array dimensions.
func (t *ArrayType) ArrayDims() []int {
	rank := t.NDims()
	dims := make([]int, rank)
	hdims := make([]C.hsize_t, rank)
	slice := (*reflect.SliceHeader)(unsafe.Pointer(&hdims))
	c_dims := (*C.hsize_t)(unsafe.Pointer(slice.Data))
	c_rank := int(C.H5Tget_array_dims2(t.id, c_dims))
	if c_rank != rank {
		return nil
	}
	for i, n := range hdims {
		dims[i] = int(n)
	}
	return dims
}

type VarLenType struct {
	Datatype
}

// NewVarLenType creates a new VarLenType.
// base_type specifies the element type of the VarLenType.
func NewVarLenType(base_type *Datatype) (*VarLenType, error) {
	id := C.H5Tvlen_create(base_type.id)
	if err := checkID(id); err != nil {
		return nil, err
	}
	t := &VarLenType{Datatype{Identifier{id}}}
	runtime.SetFinalizer(t, (*VarLenType).finalizer)
	return t, nil
}

// IsVariableStr determines whether the VarLenType is a string.
func (vl *VarLenType) IsVariableStr() bool {
	return C.H5Tis_variable_str(vl.id) > 0
}

type CompoundType struct {
	Datatype
}

// NewCompoundType creates a new CompoundType.
// size is the size in bytes of the compound datatype.
func NewCompoundType(size int) (*CompoundType, error) {
	id := C.H5Tcreate(C.H5T_class_t(T_COMPOUND), C.size_t(size))
	if err := checkID(id); err != nil {
		return nil, err
	}
	t := &CompoundType{Datatype{Identifier{id}}}
	runtime.SetFinalizer(t, (*CompoundType).finalizer)
	return t, nil
}

// NMembers returns the number of elements in a compound or enumeration datatype.
func (t *CompoundType) NMembers() int {
	return int(C.H5Tget_nmembers(t.id))
}

// Class returns the TypeClass of the DataType
func (t *Datatype) Class() TypeClass {
	return TypeClass(C.H5Tget_class(t.id))
}

// MemberClass returns datatype class of compound datatype member.
func (t *CompoundType) MemberClass(mbr_idx int) TypeClass {
	return TypeClass(C.H5Tget_member_class(t.id, C.uint(mbr_idx)))
}

// MemberName returns the name of a compound or enumeration datatype member.
func (t *CompoundType) MemberName(mbr_idx int) string {
	c_name := C.H5Tget_member_name(t.id, C.uint(mbr_idx))
	defer C.free(unsafe.Pointer(c_name))
	return C.GoString(c_name)
}

// MemberIndex returns the index of a compound or enumeration datatype member.
func (t *CompoundType) MemberIndex(name string) int {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	return int(C.H5Tget_member_index(t.id, c_name))
}

// MemberOffset returns the offset of a field of a compound datatype.
func (t *CompoundType) MemberOffset(mbr_idx int) int {
	return int(C.H5Tget_member_offset(t.id, C.uint(mbr_idx)))
}

// MemberType returns the datatype of the specified member.
func (t *CompoundType) MemberType(mbr_idx int) (*Datatype, error) {
	hid := C.H5Tget_member_type(t.id, C.uint(mbr_idx))
	if err := checkID(hid); err != nil {
		return nil, err
	}
	return NewDatatype(hid), nil
}

// Insert adds a new member to a compound datatype.
func (t *CompoundType) Insert(name string, offset int, field *Datatype) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return h5err(C.H5Tinsert(t.id, cname, C.size_t(offset), field.id))
}

// Pack recursively removes padding from within a compound datatype.
// This is analogous to C struct packing and will give a space-efficient
// type on the disk. However, using this may require type conversions
// on more machines, so may be a worse option.
func (t *CompoundType) Pack() error {
	return h5err(C.H5Tpack(t.id))
}

type OpaqueDatatype struct {
	Datatype
}

// SetTag tags an opaque datatype.
func (t *OpaqueDatatype) SetTag(tag string) error {
	ctag := C.CString(tag)
	defer C.free(unsafe.Pointer(ctag))
	return h5err(C.H5Tset_tag(t.id, ctag))
}

// Tag returns the tag associated with an opaque datatype.
func (t *OpaqueDatatype) Tag() string {
	cname := C.H5Tget_tag(t.id)
	if cname != nil {
		defer C.free(unsafe.Pointer(cname))
		return C.GoString(cname)
	}
	return ""
}

// NewDatatypeFromValue creates  a datatype from a value in an interface.
func NewDatatypeFromValue(v interface{}) (*Datatype, error) {
	return NewDataTypeFromType(reflect.TypeOf(v))
}

// NewDatatypeFromType creates a new Datatype from a reflect.Type.
func NewDataTypeFromType(t reflect.Type) (*Datatype, error) {

	var dt *Datatype = nil
	var err error

	switch t.Kind() {

	case reflect.Int:
		dt, err = T_NATIVE_INT.Copy()

	case reflect.Int8:
		dt, err = T_NATIVE_INT8.Copy()

	case reflect.Int16:
		dt, err = T_NATIVE_INT16.Copy()

	case reflect.Int32:
		dt, err = T_NATIVE_INT32.Copy()

	case reflect.Int64:
		dt, err = T_NATIVE_INT64.Copy()

	case reflect.Uint:
		dt, err = T_NATIVE_UINT.Copy()

	case reflect.Uint8:
		dt, err = T_NATIVE_UINT8.Copy()

	case reflect.Uint16:
		dt, err = T_NATIVE_UINT16.Copy()

	case reflect.Uint32:
		dt, err = T_NATIVE_UINT32.Copy()

	case reflect.Uint64:
		dt, err = T_NATIVE_UINT64.Copy()

	case reflect.Float32:
		dt, err = T_NATIVE_FLOAT.Copy()

	case reflect.Float64:
		dt, err = T_NATIVE_DOUBLE.Copy()

	case reflect.String:
		dt, err = T_GO_STRING.Copy()

	case reflect.Bool:
		dt, err = T_NATIVE_HBOOL.Copy()

	case reflect.Array:
		elem_type, err := NewDataTypeFromType(t.Elem())
		if err != nil {
			return nil, err
		}

		dims := getArrayDims(t)

		adt, err := NewArrayType(elem_type, dims)
		if err != nil {
			return nil, err
		}

		dt = &adt.Datatype

	case reflect.Slice:
		elem_type, err := NewDataTypeFromType(t.Elem())
		if err != nil {
			return nil, err
		}

		sdt, err := NewVarLenType(elem_type)
		if err != nil {
			return nil, err
		}

		dt = &sdt.Datatype

	case reflect.Struct:
		sz := int(t.Size())
		cdt, err := NewCompoundType(sz)
		if err != nil {
			return nil, err
		}
		n := t.NumField()
		for i := 0; i < n; i++ {
			f := t.Field(i)
			var field_dt *Datatype = nil
			field_dt, err = NewDataTypeFromType(f.Type)
			if err != nil {
				return nil, err
			}
			offset := int(f.Offset + 0)
			if field_dt == nil {
				return nil, fmt.Errorf("pb with field [%d-%s]", i, f.Name)
			}
			field_name := string(f.Tag)
			if len(field_name) == 0 {
				field_name = f.Name
			}
			err = cdt.Insert(field_name, offset, field_dt)
			if err != nil {
				return nil, fmt.Errorf("pb with field [%d-%s]: %s", i, f.Name, err)
			}
		}
		dt = &cdt.Datatype

	default:
		// Should never happen.
		panic(fmt.Errorf("unhandled kind (%v)", t.Kind()))
	}

	return dt, err
}

func getArrayDims(dt reflect.Type) []int {
	result := []int{}
	if dt.Kind() == reflect.Array {
		result = append(result, dt.Len())
		for _, dim := range getArrayDims(dt.Elem()) {
			result = append(result, dim)
		}
	}
	return result
}
