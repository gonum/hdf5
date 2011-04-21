package hdf5

/*
 #cgo LDFLAGS: -lhdf5
 #include "hdf5.h"

 #include <stdlib.h>
 #include <string.h>
 */
import "C"

import (
	"unsafe"
	"os"
	"runtime"
	"fmt"
)

// ---- H5T: Datatype Interface ----

type DataType struct {
	id C.hid_t
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

// Creates a new datatype.
// hid_t H5Tcreate( H5T_class_t class, size_tsize ) 
func CreateDataType(class TypeClass, size int) (t *DataType, err os.Error) {
	t = nil
	err = nil

	hid := C.H5Tcreate(C.H5T_class_t(class), C.size_t(size))
	err = togo_err(C.herr_t(int(hid)))
	if err != nil {
		return
	}
	t = &DataType{id:hid}
	runtime.SetFinalizer(t, (*DataType).h5t_finalizer)
	return
}

func (t *DataType) h5t_finalizer() {
	err := t.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing file: %s",err))
	}
}

// Releases a datatype.
// herr_t H5Tclose( hid_t dtype_id ) 
func (t *DataType) Close() os.Error {
	return togo_err(C.H5Tclose(t.id))
}

// Commits a transient datatype, linking it into the file and creating a new named datatype. 
// herr_t H5Tcommit( hid_t loc_id, const char *name, hid_t dtype_id, hid_t lcpl_id, hid_t tcpl_id, hid_t tapl_id ) 
//func (t *DataType) Commit()

// Determines whether a datatype is a named type or a transient type. 
// htri_tH5Tcommitted( hid_t dtype_id ) 
func (t *DataType) Committed() bool {
	o := int(C.H5Tcommitted(t.id))
	if o> 0 {
		return true
	} 
	return false
}

// Copies an existing datatype.
// hid_t H5Tcopy( hid_t dtype_id ) 
func (t *DataType) Copy() (*DataType, os.Error) {
	hid := C.H5Tcopy(t.id)
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	o := &DataType{id:hid}
	runtime.SetFinalizer(o, (*DataType).h5t_finalizer)
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
func (t *DataType) Lock() os.Error {
	return togo_err(C.H5Tlock(t.id))
}

// Returns the size of a datatype. 
// size_t H5Tget_size( hid_t dtype_id ) 
func (t *DataType) Size() int {
	return int(C.H5Tget_size(t.id))
}

// compound data type
type CompoundDataType struct {
	DataType
}

// Retrieves the number of elements in a compound or enumeration datatype. 
// int H5Tget_nmembers( hid_t dtype_id ) 
func (t *CompoundDataType) NMembers() int {
	return int(C.H5Tget_nmembers(t.id))
}

// Returns datatype class of compound datatype member. 
// H5T_class_t H5Tget_member_class( hid_t cdtype_id, unsigned member_no ) 
func (t *CompoundDataType) MemberClass(mbr_idx int) TypeClass {
	return TypeClass(C.H5Tget_member_class(t.id, C.uint(mbr_idx)))
}

// Retrieves the name of a compound or enumeration datatype member. 
// char * H5Tget_member_name( hid_t dtype_id, unsigned field_idx ) 
func (t *CompoundDataType) MemberName(mbr_idx int) string {
	c_name := C.H5Tget_member_name(t.id, C.uint(mbr_idx))
	return C.GoString(c_name)
}

// Retrieves the index of a compound or enumeration datatype member. 
// int H5Tget_member_index( hid_t dtype_id, const char * field_name ) 
func (t *CompoundDataType) MemberIndex(name string) int {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	return int(C.H5Tget_member_index(t.id, c_name))
}

// Retrieves the offset of a field of a compound datatype. 
// size_t H5Tget_member_offset( hid_t dtype_id, unsigned memb_no ) 
func (t *CompoundDataType) MemberOffset(mbr_idx int) int {
	return int(C.H5Tget_member_offset(t.id, C.uint(mbr_idx)))
}

// Returns the datatype of the specified member. 
// hid_t H5Tget_member_type( hid_t dtype_id, unsigned field_idx ) 
func (t *CompoundDataType) MemberType(mbr_idx int) (*DataType, os.Error) {
	hid := C.H5Tget_member_type(t.id, C.uint(mbr_idx))
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	dt := &DataType{id:hid}
	runtime.SetFinalizer(dt, (*DataType).h5t_finalizer)
	return dt, nil
}

// Adds a new member to a compound datatype. 
// herr_t H5Tinsert( hid_t dtype_id, const char * name, size_t offset, hid_t field_id ) 
func (t *CompoundDataType) Insert(name string, offset int, field *DataType) os.Error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	err := C.H5Tinsert(t.id, c_name, C.size_t(offset), field.id)
	return togo_err(err)
}

// Recursively removes padding from within a compound datatype. 
// herr_t H5Tpack( hid_t dtype_id ) 
func (t *CompoundDataType) Pack() os.Error {
	err := C.H5Tpack(t.id)
	return togo_err(err)
}

// --- opaque type ---
type OpaqueDataType struct {
	DataType
}

// Tags an opaque datatype. 
// herr_t H5Tset_tag( hid_t dtype_id const char *tag ) 
func (t *OpaqueDataType) SetTag(tag string) os.Error {
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

// EOF
