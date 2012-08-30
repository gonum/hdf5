package hdf5

// #include "hdf5.h"
// #include <stdlib.h>
// #include <string.h>
import "C"

import (
	"unsafe"

	//"runtime"
	"fmt"
	"reflect"
)

// ---- H5T: Datatype Interface ----

type DataType struct {
	id C.hid_t
	rt reflect.Type
}

type TypeClass C.H5T_class_t

const (
	// Error
	T_NO_CLASS TypeClass = -1

	// integer types
	T_INTEGER TypeClass = 0

	// floating-point types
	T_FLOAT TypeClass = 1

	// date and time types
	T_TIME TypeClass = 2

	// character string types
	T_STRING TypeClass = 3

	// bit field types
	T_BITFIELD TypeClass = 4

	// opaque types
	T_OPAQUE TypeClass = 5

	// compound types
	T_COMPOUND TypeClass = 6

	// reference types
	T_REFERENCE TypeClass = 7

	// enumeration types
	T_ENUM TypeClass = 8

	// variable-length types
	T_VLEN TypeClass = 9

	// array types
	T_ARRAY TypeClass = 10

	// nbr of classes -- MUST BE LAST
	T_NCLASSES TypeClass = 11
)

type dummy_struct struct{}

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

	_go_struct_t reflect.Type = reflect.TypeOf(dummy_struct{})

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

func new_dtype(id C.hid_t, rt reflect.Type) *DataType {
	t := &DataType{id: id, rt: rt}
	//runtime.SetFinalizer(t, (*DataType).h5t_finalizer)
	return t
}

// Creates a new datatype.
// hid_t H5Tcreate( H5T_class_t class, size_tsize ) 
func CreateDataType(class TypeClass, size int) (t *DataType, err error) {
	t = nil
	err = nil

	hid := C.H5Tcreate(C.H5T_class_t(class), C.size_t(size))
	err = togo_err(C.herr_t(int(hid)))
	if err != nil {
		return
	}
	t = new_dtype(hid, _type_cls_to_go_type[class])
	return
}

func (t *DataType) h5t_finalizer() {
	err := t.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing datatype: %s", err))
	}
}

// Releases a datatype.
// herr_t H5Tclose( hid_t dtype_id ) 
func (t *DataType) Close() error {
	if t.id > 0 {
		fmt.Printf("--- closing dtype [%d]...\n", t.id)
		err := togo_err(C.H5Tclose(t.id))
		t.id = 0
		return err
	}
	return nil
}

// Commits a transient datatype, linking it into the file and creating a new named datatype. 
// herr_t H5Tcommit( hid_t loc_id, const char *name, hid_t dtype_id, hid_t lcpl_id, hid_t tcpl_id, hid_t tapl_id ) 
//func (t *DataType) Commit()

// Determines whether a datatype is a named type or a transient type. 
// htri_tH5Tcommitted( hid_t dtype_id ) 
func (t *DataType) Committed() bool {
	o := int(C.H5Tcommitted(t.id))
	if o > 0 {
		return true
	}
	return false
}

// Copies an existing datatype.
// hid_t H5Tcopy( hid_t dtype_id ) 
func (t *DataType) Copy() (*DataType, error) {
	hid := C.H5Tcopy(t.id)
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	o := new_dtype(hid, t.rt)
	return o, err
}

// Determines whether two datatype identifiers refer to the same datatype. 
// htri_t H5Tequal( hid_t dtype_id1, hid_t dtype_id2 ) 
func (t *DataType) Equal(o *DataType) bool {
	v := int(C.H5Tequal(t.id, o.id))
	if v > 0 {
		return true
	}
	return false
}

// Locks a datatype. 
// herr_t H5Tlock( hid_t dtype_id ) 
func (t *DataType) Lock() error {
	return togo_err(C.H5Tlock(t.id))
}

// Returns the size of a datatype. 
// size_t H5Tget_size( hid_t dtype_id ) 
func (t *DataType) Size() int {
	return int(C.H5Tget_size(t.id))
}

// Sets the total size for an atomic datatype.
// herr_t H5Tset_size( hid_t dtype_id, size_tsize )
func (t *DataType) SetSize(sz int) error {
	err := C.H5Tset_size(t.id, C.size_t(sz))
	return togo_err(err)
}
// ---------------------------------------------------------------------------

// array data type
type ArrayType struct {
	DataType
}

func new_array_type(id C.hid_t) *ArrayType {
	t := &ArrayType{DataType{id: id}}
	//runtime.SetFinalizer(t, (*DataType).h5t_finalizer)
	return t
}

func NewArrayType(base_type *DataType, dims []int) (*ArrayType, error) {
	ndims := C.uint(len(dims))
	c_dims := (*C.hsize_t)(unsafe.Pointer(&dims[0]))

	hid := C.H5Tarray_create2(base_type.id, ndims, c_dims)
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	t := new_array_type(hid)
	return t, err
}

// Returns the rank of an array datatype.
// int H5Tget_array_ndims( hid_t adtype_id )
func (t *ArrayType) NDims() int {
	return int(C.H5Tget_array_ndims(t.id))
}

// Retrieves sizes of array dimensions.
// int H5Tget_array_dims2( hid_t adtype_id, hsize_t dims[] )
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
// ---------------------------------------------------------------------------

// variable length array data type
type VarLenType struct {
	DataType
}

func NewVarLenType(base_type *DataType) (*VarLenType, error) {
	hid := C.H5Tvlen_create(base_type.id)
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	dt := new_vltype(hid)
	return dt, err
}

func new_vltype(id C.hid_t) *VarLenType {
	t := &VarLenType{DataType{id: id}}
	//runtime.SetFinalizer(t, (*DataType).h5t_finalizer)
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

// ---------------------------------------------------------------------------

// compound data type
type CompType struct {
	DataType
}

// Retrieves the number of elements in a compound or enumeration datatype. 
// int H5Tget_nmembers( hid_t dtype_id ) 
func (t *CompType) NMembers() int {
	return int(C.H5Tget_nmembers(t.id))
}

// Returns datatype class of compound datatype member. 
// H5T_class_t H5Tget_member_class( hid_t cdtype_id, unsigned member_no ) 
func (t *CompType) MemberClass(mbr_idx int) TypeClass {
	return TypeClass(C.H5Tget_member_class(t.id, C.uint(mbr_idx)))
}

// Retrieves the name of a compound or enumeration datatype member. 
// char * H5Tget_member_name( hid_t dtype_id, unsigned field_idx ) 
func (t *CompType) MemberName(mbr_idx int) string {
	c_name := C.H5Tget_member_name(t.id, C.uint(mbr_idx))
	return C.GoString(c_name)
}

// Retrieves the index of a compound or enumeration datatype member. 
// int H5Tget_member_index( hid_t dtype_id, const char * field_name ) 
func (t *CompType) MemberIndex(name string) int {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	return int(C.H5Tget_member_index(t.id, c_name))
}

// Retrieves the offset of a field of a compound datatype. 
// size_t H5Tget_member_offset( hid_t dtype_id, unsigned memb_no ) 
func (t *CompType) MemberOffset(mbr_idx int) int {
	return int(C.H5Tget_member_offset(t.id, C.uint(mbr_idx)))
}

// Returns the datatype of the specified member. 
// hid_t H5Tget_member_type( hid_t dtype_id, unsigned field_idx ) 
func (t *CompType) MemberType(mbr_idx int) (*DataType, error) {
	hid := C.H5Tget_member_type(t.id, C.uint(mbr_idx))
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	dt := new_dtype(hid, t.rt.Field(mbr_idx).Type)
	return dt, nil
}

// Adds a new member to a compound datatype. 
// herr_t H5Tinsert( hid_t dtype_id, const char * name, size_t offset, hid_t field_id ) 
func (t *CompType) Insert(name string, offset int, field *DataType) error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	//fmt.Printf("inserting [%s] at offset:%d (id=%d)...\n", name, offset, field.id)
	err := C.H5Tinsert(t.id, c_name, C.size_t(offset), field.id)
	return togo_err(err)
}

// Recursively removes padding from within a compound datatype. 
// herr_t H5Tpack( hid_t dtype_id ) 
func (t *CompType) Pack() error {
	err := C.H5Tpack(t.id)
	return togo_err(err)
}

// --- opaque type ---
type OpaqueDataType struct {
	DataType
}

// Tags an opaque datatype. 
// herr_t H5Tset_tag( hid_t dtype_id const char *tag ) 
func (t *OpaqueDataType) SetTag(tag string) error {
	c_tag := C.CString(tag)
	defer C.free(unsafe.Pointer(c_tag))

	err := C.H5Tset_tag(t.id, c_tag)
	return togo_err(err)
}

// Gets the tag associated with an opaque datatype. 
// char *H5Tget_tag( hid_t dtype_id ) 
func (t *OpaqueDataType) Tag() string {
	c_name := C.H5Tget_tag(t.id)
	if c_name != nil {
		return C.GoString(c_name)
	}
	return ""
}

// -----------------------------------------

// create a data-type from a golang value
func NewDataTypeFromValue(v interface{}) *DataType {
	t := reflect.TypeOf(v)
	return new_dataTypeFromType(t)
}

func new_dataTypeFromType(t reflect.Type) *DataType {

	var dt *DataType = nil

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
		elem_type := new_dataTypeFromType(t.Elem())
		n := t.Len()
		dims := []int{n}
		adt, err := NewArrayType(elem_type, dims)
		if err != nil {
			panic(err)
		}
		dt, err = adt.Copy()
		if err != nil {
			panic(err)
		}

	case reflect.Slice:
		elem_type := new_dataTypeFromType(t.Elem())
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
		hdf_dt, err := CreateDataType(T_COMPOUND, sz)
		if err != nil {
			panic(err)
		}
		cdt := &CompType{*hdf_dt}
		n := t.NumField()
		for i := 0; i < n; i++ {
			f := t.Field(i)
			var field_dt *DataType = nil
			field_dt = new_dataTypeFromType(f.Type)
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
// EOF
