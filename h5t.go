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
	id C.hid_t
	rt reflect.Type
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

type typeClassToType map[TypeClass]reflect.Type

var (
	// mapping of type-class to go-type
	_type_cls_to_go_type typeClassToType = typeClassToType{
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
)

func openDatatype(loc_id C.hid_t, name string, tapl_id int) (*Datatype, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	hid := C.H5Topen2(C.hid_t(loc_id), c_name, C.hid_t(tapl_id))
	err := h5err(C.herr_t(hid))
	if err != nil {
		return nil, err
	}
	dt := &Datatype{id: hid}
	runtime.SetFinalizer(dt, (*Datatype).finalizer)
	return dt, err
}

func NewDatatype(id C.hid_t, rt reflect.Type) *Datatype {
	t := &Datatype{id: id, rt: rt}
	//runtime.SetFinalizer(t, (*Datatype).finalizer)
	return t
}

// Creates a new datatype.
func CreateDatatype(class TypeClass, size int) (t *Datatype, err error) {
	t = nil
	err = nil

	hid := C.H5Tcreate(C.H5T_class_t(class), C.size_t(size))
	err = h5err(C.herr_t(int(hid)))
	if err != nil {
		return
	}
	t = NewDatatype(hid, _type_cls_to_go_type[class])
	return
}

func (t *Datatype) finalizer() {
	err := t.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing datatype: %s", err))
	}
}

// Releases a datatype.
func (t *Datatype) Close() error {
	if t.id > 0 {
		fmt.Printf("--- closing dtype [%d]...\n", t.id)
		err := h5err(C.H5Tclose(t.id))
		t.id = 0
		return err
	}
	return nil
}

// Determines whether a datatype is a named type or a transient type.
func (t *Datatype) Committed() bool {
	o := int(C.H5Tcommitted(t.id))
	if o > 0 {
		return true
	}
	return false
}

// Copies an existing datatype.
func (t *Datatype) Copy() (*Datatype, error) {
	hid := C.H5Tcopy(t.id)
	err := h5err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	o := NewDatatype(hid, t.rt)
	return o, err
}

// Determines whether two datatype identifiers refer to the same datatype.
func (t *Datatype) Equal(o *Datatype) bool {
	v := int(C.H5Tequal(t.id, o.id))
	if v > 0 {
		return true
	}
	return false
}

// Locks a datatype.
func (t *Datatype) Lock() error {
	return h5err(C.H5Tlock(t.id))
}

// Size returns the size of the Datatype.
func (t *Datatype) Size() uint {
	return uint(C.H5Tget_size(t.id))
}

// SetSize sets the total size of a Datatype.
func (t *Datatype) SetSize(sz uint) error {
	err := C.H5Tset_size(t.id, C.size_t(sz))
	return h5err(err)
}

type ArrayType struct {
	Datatype
}

func new_array_type(id C.hid_t) *ArrayType {
	t := &ArrayType{Datatype{id: id}}
	return t
}

func NewArrayType(base_type *Datatype, dims []int) (*ArrayType, error) {
	ndims := C.uint(len(dims))
	c_dims := (*C.hsize_t)(unsafe.Pointer(&dims[0]))

	hid := C.H5Tarray_create2(base_type.id, ndims, c_dims)
	err := h5err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	t := new_array_type(hid)
	return t, err
}

// Returns the rank of an array datatype.
func (t *ArrayType) NDims() int {
	return int(C.H5Tget_array_ndims(t.id))
}

// Retrieves sizes of array dimensions.
func (t *ArrayType) ArrayDims() []int {
	rank := t.NDims()
	dims := make([]int, rank)
	// fixme: int/hsize_t size!
	c_dims := (*C.hsize_t)(unsafe.Pointer(&dims[0]))
	c_rank := int(C.H5Tget_array_dims2(t.id, c_dims))
	if c_rank == rank {
		return dims
	}
	return nil
}

type VarLenType struct {
	Datatype
}

func NewVarLenType(base_type *Datatype) (*VarLenType, error) {
	hid := C.H5Tvlen_create(base_type.id)
	err := h5err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	dt := new_vltype(hid)
	return dt, err
}

func new_vltype(id C.hid_t) *VarLenType {
	t := &VarLenType{Datatype{id: id}}
	//runtime.SetFinalizer(t, (*Datatype).finalizer)
	return t
}

// Determines whether datatype is a variable-length string.
// htri_t H5Tis_variable_str( hid_t dtype_id )
func (vl *VarLenType) IsVariableStr() bool {
	o := int(C.H5Tis_variable_str(vl.id))
	if o > 0 {
		return true
	}
	return false
}

type CompoundType struct {
	Datatype
}

// Retrieves the number of elements in a compound or enumeration datatype.
func (t *CompoundType) NMembers() int {
	return int(C.H5Tget_nmembers(t.id))
}

// Returns datatype class of compound datatype member.
func (t *CompoundType) MemberClass(mbr_idx int) TypeClass {
	return TypeClass(C.H5Tget_member_class(t.id, C.uint(mbr_idx)))
}

// Retrieves the name of a compound or enumeration datatype member.
func (t *CompoundType) MemberName(mbr_idx int) string {
	c_name := C.H5Tget_member_name(t.id, C.uint(mbr_idx))
	return C.GoString(c_name)
}

// Retrieves the index of a compound or enumeration datatype member.
func (t *CompoundType) MemberIndex(name string) int {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	return int(C.H5Tget_member_index(t.id, c_name))
}

// Retrieves the offset of a field of a compound datatype.
func (t *CompoundType) MemberOffset(mbr_idx int) int {
	return int(C.H5Tget_member_offset(t.id, C.uint(mbr_idx)))
}

// Returns the datatype of the specified member.
func (t *CompoundType) MemberType(mbr_idx int) (*Datatype, error) {
	hid := C.H5Tget_member_type(t.id, C.uint(mbr_idx))
	err := h5err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	dt := NewDatatype(hid, t.rt.Field(mbr_idx).Type)
	return dt, nil
}

// Adds a new member to a compound datatype.
func (t *CompoundType) Insert(name string, offset int, field *Datatype) error {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return h5err(C.H5Tinsert(t.id, cname, C.size_t(offset), field.id))
}

// Recursively removes padding from within a compound datatype.
func (t *CompoundType) Pack() error {
	return h5err(C.H5Tpack(t.id))
}

type OpaqueDatatype struct {
	Datatype
}

// Tags an opaque datatype.
func (t *OpaqueDatatype) SetTag(tag string) error {
	ctag := C.CString(tag)
	defer C.free(unsafe.Pointer(ctag))
	return h5err(C.H5Tset_tag(t.id, ctag))
}

// Gets the tag associated with an opaque datatype.
func (t *OpaqueDatatype) Tag() string {
	cname := C.H5Tget_tag(t.id)
	if cname != nil {
		return C.GoString(cname)
	}
	return ""
}

// NewDatatypeFromValue creates  a datatype from a value in an interface.
func NewDatatypeFromValue(v interface{}) *Datatype {
	t := reflect.TypeOf(v)
	return newDataTypeFromType(t)
}

func newDataTypeFromType(t reflect.Type) *Datatype {

	var dt *Datatype = nil

	switch t.Kind() {

	case reflect.Int:
		dt = T_NATIVE_INT // FIXME: .Copy() instead ?

	case reflect.Int8:
		dt = T_NATIVE_INT8

	case reflect.Int16:
		dt = T_NATIVE_INT16

	case reflect.Int32:
		dt = T_NATIVE_INT32

	case reflect.Int64:
		dt = T_NATIVE_INT64

	case reflect.Uint:
		dt = T_NATIVE_UINT // FIXME: .Copy() instead ?

	case reflect.Uint8:
		dt = T_NATIVE_UINT8

	case reflect.Uint16:
		dt = T_NATIVE_UINT16

	case reflect.Uint32:
		dt = T_NATIVE_UINT32

	case reflect.Uint64:
		dt = T_NATIVE_UINT64

	case reflect.Float32:
		dt = T_NATIVE_FLOAT

	case reflect.Float64:
		dt = T_NATIVE_DOUBLE

	case reflect.String:
		dt = T_GO_STRING
		//dt = T_C_S1

	case reflect.Array:
		elem_type := newDataTypeFromType(t.Elem())
		dims := getArrayDims(t)
		adt, err := NewArrayType(elem_type, dims)
		if err != nil {
			panic(err)
		}
		dt, err = adt.Copy()
		if err != nil {
			panic(err)
		}

	case reflect.Slice:
		elem_type := newDataTypeFromType(t.Elem())
		vlen_dt, err := NewVarLenType(elem_type)
		if err != nil {
			panic(err)
		}
		dt, err = vlen_dt.Copy()
		if err != nil {
			panic(err)
		}

	case reflect.Struct:
		sz := int(t.Size())
		hdf_dt, err := CreateDatatype(T_COMPOUND, sz)
		if err != nil {
			panic(err)
		}
		cdt := &CompoundType{*hdf_dt}
		n := t.NumField()
		for i := 0; i < n; i++ {
			f := t.Field(i)
			var field_dt *Datatype = nil
			field_dt = newDataTypeFromType(f.Type)
			offset := int(f.Offset + 0)
			if field_dt == nil {
				panic(fmt.Sprintf("pb with field [%d-%s]", i, f.Name))
			}
			field_name := string(f.Tag)
			if len(field_name) == 0 {
				field_name = f.Name
			}
			err = cdt.Insert(field_name, offset, field_dt)
			if err != nil {
				panic(fmt.Sprintf("pb with field [%d-%s]: %s", i, f.Name, err))
			}
		}
		cdt.Lock()
		dt, err = cdt.Copy()
		if err != nil {
			panic(err)
		}

	case reflect.Ptr:
		panic("sorry, pointers not yet supported")

	default:
		panic(fmt.Sprintf("unhandled kind (%d)", t.Kind()))
	}

	return dt
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
