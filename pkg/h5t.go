package hdf5

/*
 #cgo LDFLAGS: -lhdf5
 #include "hdf5.h"

 #include <stdlib.h>
 #include <string.h>

 hid_t _go_hdf5_H5T_C_S1() { return H5T_C_S1; }
 hid_t _go_hdf5_H5T_FORTRAN_S1() { return H5T_FORTRAN_S1; }

 hid_t _go_hdf5_H5T_STD_I8BE() { return H5T_STD_I8BE; }
 hid_t _go_hdf5_H5T_STD_I8LE() { return H5T_STD_I8LE; }
 hid_t _go_hdf5_H5T_STD_I16BE() { return H5T_STD_I16BE; }
 hid_t _go_hdf5_H5T_STD_I16LE() { return H5T_STD_I16LE; }
 hid_t _go_hdf5_H5T_STD_I32BE() { return H5T_STD_I32BE; }
 hid_t _go_hdf5_H5T_STD_I32LE() { return H5T_STD_I32LE; }
 hid_t _go_hdf5_H5T_STD_I64BE() { return H5T_STD_I64BE; }
 hid_t _go_hdf5_H5T_STD_I64LE() { return H5T_STD_I64LE; }
 hid_t _go_hdf5_H5T_STD_U8BE() { return H5T_STD_U8BE; }
 hid_t _go_hdf5_H5T_STD_U8LE() { return H5T_STD_U8LE; }
 hid_t _go_hdf5_H5T_STD_U16BE() { return H5T_STD_U16BE; }
 hid_t _go_hdf5_H5T_STD_U16LE() { return H5T_STD_U16LE; }
 hid_t _go_hdf5_H5T_STD_U32BE() { return H5T_STD_U32BE; }
 hid_t _go_hdf5_H5T_STD_U32LE() { return H5T_STD_U32LE; }
 hid_t _go_hdf5_H5T_STD_U64BE() { return H5T_STD_U64BE; }
 hid_t _go_hdf5_H5T_STD_U64LE() { return H5T_STD_U64LE; }
 hid_t _go_hdf5_H5T_STD_B8BE() { return H5T_STD_B8BE; }
 hid_t _go_hdf5_H5T_STD_B8LE() { return H5T_STD_B8LE; }

 hid_t _go_hdf5_H5T_STD_B16BE() { return H5T_STD_B16BE; }
 hid_t _go_hdf5_H5T_STD_B16LE() { return H5T_STD_B16LE; }
 hid_t _go_hdf5_H5T_STD_B32BE() { return H5T_STD_B32BE; }
 hid_t _go_hdf5_H5T_STD_B32LE() { return H5T_STD_B32LE; }
 hid_t _go_hdf5_H5T_STD_B64BE() { return H5T_STD_B64BE; }
 hid_t _go_hdf5_H5T_STD_B64LE() { return H5T_STD_B64LE; }
 hid_t _go_hdf5_H5T_STD_REF_OBJ() { return H5T_STD_REF_OBJ; }
 hid_t _go_hdf5_H5T_STD_REF_DSETREG() { return H5T_STD_REF_DSETREG; }

 hid_t _go_hdf5_H5T_IEEE_F32BE() { return H5T_IEEE_F32BE; }
 hid_t _go_hdf5_H5T_IEEE_F32LE() { return H5T_IEEE_F32LE; }
 hid_t _go_hdf5_H5T_IEEE_F64BE() { return H5T_IEEE_F64BE; }
 hid_t _go_hdf5_H5T_IEEE_F64LE() { return H5T_IEEE_F64LE; }

 hid_t _go_hdf5_H5T_UNIX_D32BE() { return H5T_UNIX_D32BE; }
 hid_t _go_hdf5_H5T_UNIX_D32LE() { return H5T_UNIX_D32LE; }
 hid_t _go_hdf5_H5T_UNIX_D64BE() { return H5T_UNIX_D64BE; }
 hid_t _go_hdf5_H5T_UNIX_D64LE() { return H5T_UNIX_D64LE; }

 hid_t _go_hdf5_H5T_INTEL_I8() { return H5T_INTEL_I8; }
 hid_t _go_hdf5_H5T_INTEL_I16() { return H5T_INTEL_I16; }
 hid_t _go_hdf5_H5T_INTEL_I32() { return H5T_INTEL_I32; }
 hid_t _go_hdf5_H5T_INTEL_I64() { return H5T_INTEL_I64; }
 hid_t _go_hdf5_H5T_INTEL_U8() { return H5T_INTEL_U8; }
 hid_t _go_hdf5_H5T_INTEL_U16() { return H5T_INTEL_U16; }
 hid_t _go_hdf5_H5T_INTEL_U32() { return H5T_INTEL_U32; }
 hid_t _go_hdf5_H5T_INTEL_U64() { return H5T_INTEL_U64; }
 hid_t _go_hdf5_H5T_INTEL_B8() { return H5T_INTEL_B8; }
 hid_t _go_hdf5_H5T_INTEL_B16() { return H5T_INTEL_B16; }
 hid_t _go_hdf5_H5T_INTEL_B32() { return H5T_INTEL_B32; }
 hid_t _go_hdf5_H5T_INTEL_B64() { return H5T_INTEL_B64; }
 hid_t _go_hdf5_H5T_INTEL_F32() { return H5T_INTEL_F32; }
 hid_t _go_hdf5_H5T_INTEL_F64() { return H5T_INTEL_F64; }

 hid_t _go_hdf5_H5T_ALPHA_I8() { return H5T_ALPHA_I8; }
 hid_t _go_hdf5_H5T_ALPHA_I16() { return H5T_ALPHA_I16; }
 hid_t _go_hdf5_H5T_ALPHA_I32() { return H5T_ALPHA_I32; }
 hid_t _go_hdf5_H5T_ALPHA_I64() { return H5T_ALPHA_I64; }
 hid_t _go_hdf5_H5T_ALPHA_U8() { return H5T_ALPHA_U8; }
 hid_t _go_hdf5_H5T_ALPHA_U16() { return H5T_ALPHA_U16; }
 hid_t _go_hdf5_H5T_ALPHA_U32() { return H5T_ALPHA_U32; }
 hid_t _go_hdf5_H5T_ALPHA_U64() { return H5T_ALPHA_U64; }
 hid_t _go_hdf5_H5T_ALPHA_B8() { return H5T_ALPHA_B8; }
 hid_t _go_hdf5_H5T_ALPHA_B16() { return H5T_ALPHA_B16; }
 hid_t _go_hdf5_H5T_ALPHA_B32() { return H5T_ALPHA_B32; }
 hid_t _go_hdf5_H5T_ALPHA_B64() { return H5T_ALPHA_B64; }
 hid_t _go_hdf5_H5T_ALPHA_F32() { return H5T_ALPHA_F32; }
 hid_t _go_hdf5_H5T_ALPHA_F64() { return H5T_ALPHA_F64; }

 hid_t _go_hdf5_H5T_MIPS_I8() { return H5T_MIPS_I8; }
 hid_t _go_hdf5_H5T_MIPS_I16() { return H5T_MIPS_I16; }
 hid_t _go_hdf5_H5T_MIPS_I32() { return H5T_MIPS_I32; }
 hid_t _go_hdf5_H5T_MIPS_I64() { return H5T_MIPS_I64; }
 hid_t _go_hdf5_H5T_MIPS_U8() { return H5T_MIPS_U8; }
 hid_t _go_hdf5_H5T_MIPS_U16() { return H5T_MIPS_U16; }
 hid_t _go_hdf5_H5T_MIPS_U32() { return H5T_MIPS_U32; }
 hid_t _go_hdf5_H5T_MIPS_U64() { return H5T_MIPS_U64; }
 hid_t _go_hdf5_H5T_MIPS_B8() { return H5T_MIPS_B8; }
 hid_t _go_hdf5_H5T_MIPS_B16() { return H5T_MIPS_B16; }
 hid_t _go_hdf5_H5T_MIPS_B32() { return H5T_MIPS_B32; }
 hid_t _go_hdf5_H5T_MIPS_B64() { return H5T_MIPS_B64; }
 hid_t _go_hdf5_H5T_MIPS_F32() { return H5T_MIPS_F32; }
 hid_t _go_hdf5_H5T_MIPS_F64() { return H5T_MIPS_F64; }

 hid_t _go_hdf5_H5T_NATIVE_CHAR() { return H5T_NATIVE_CHAR; }
 hid_t _go_hdf5_H5T_NATIVE_INT() { return H5T_NATIVE_INT; }
 hid_t _go_hdf5_H5T_NATIVE_FLOAT() { return H5T_NATIVE_FLOAT; }
 hid_t _go_hdf5_H5T_NATIVE_SCHAR() { return H5T_NATIVE_SCHAR; }
 hid_t _go_hdf5_H5T_NATIVE_UCHAR() { return H5T_NATIVE_UCHAR; }
 hid_t _go_hdf5_H5T_NATIVE_SHORT() { return H5T_NATIVE_SHORT; }
 hid_t _go_hdf5_H5T_NATIVE_USHORT() { return H5T_NATIVE_USHORT; }
 hid_t _go_hdf5_H5T_NATIVE_UINT() { return H5T_NATIVE_UINT; }
 hid_t _go_hdf5_H5T_NATIVE_LONG() { return H5T_NATIVE_LONG; }
 hid_t _go_hdf5_H5T_NATIVE_ULONG() { return H5T_NATIVE_ULONG; }
 hid_t _go_hdf5_H5T_NATIVE_LLONG() { return H5T_NATIVE_LLONG; }
 hid_t _go_hdf5_H5T_NATIVE_ULLONG() { return H5T_NATIVE_ULLONG; }
 hid_t _go_hdf5_H5T_NATIVE_DOUBLE() { return H5T_NATIVE_DOUBLE; }
#if H5_SIZEOF_LONG_DOUBLE !=0
 hid_t _go_hdf5_H5T_NATIVE_LDOUBLE() { return H5T_NATIVE_LDOUBLE; }
#endif
 hid_t _go_hdf5_H5T_NATIVE_B8() { return H5T_NATIVE_B8; }
 hid_t _go_hdf5_H5T_NATIVE_B16() { return H5T_NATIVE_B16; }
 hid_t _go_hdf5_H5T_NATIVE_B32() { return H5T_NATIVE_B32; }
 hid_t _go_hdf5_H5T_NATIVE_B64() { return H5T_NATIVE_B64; }
 hid_t _go_hdf5_H5T_NATIVE_OPAQUE() { return H5T_NATIVE_OPAQUE; }
 hid_t _go_hdf5_H5T_NATIVE_HSIZE() { return H5T_NATIVE_HSIZE; }
 hid_t _go_hdf5_H5T_NATIVE_HSSIZE() { return H5T_NATIVE_HSSIZE; }
 hid_t _go_hdf5_H5T_NATIVE_HERR() { return H5T_NATIVE_HERR; }
 hid_t _go_hdf5_H5T_NATIVE_HBOOL() { return H5T_NATIVE_HBOOL; }

 hid_t _go_hdf5_H5T_NATIVE_INT8() { return H5T_NATIVE_INT8; }
 hid_t _go_hdf5_H5T_NATIVE_UINT8() { return H5T_NATIVE_UINT8; }
 hid_t _go_hdf5_H5T_NATIVE_INT16() { return H5T_NATIVE_INT16; }
 hid_t _go_hdf5_H5T_NATIVE_UINT16() { return H5T_NATIVE_UINT16; }
 hid_t _go_hdf5_H5T_NATIVE_INT32() { return H5T_NATIVE_INT32; }
 hid_t _go_hdf5_H5T_NATIVE_UINT32() { return H5T_NATIVE_UINT32; }
 hid_t _go_hdf5_H5T_NATIVE_INT64() { return H5T_NATIVE_INT64; }
 hid_t _go_hdf5_H5T_NATIVE_UINT64() { return H5T_NATIVE_UINT64; }

*/
import "C"

import (
	"unsafe"
	"os"
	//"runtime"
	"fmt"
	"reflect"
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

// list of predefined data types
var (
	T_C_S1 *DataType = new_dtype(C._go_hdf5_H5T_C_S1())
	T_FORTRAN_S1 *DataType = new_dtype(C._go_hdf5_H5T_FORTRAN_S1())

	T_STD_I8BE *DataType = new_dtype(C._go_hdf5_H5T_STD_I8BE())
	T_STD_I8LE *DataType = new_dtype(C._go_hdf5_H5T_STD_I8LE())
	T_STD_I16BE *DataType = new_dtype(C._go_hdf5_H5T_STD_I16BE())
	T_STD_I16LE *DataType = new_dtype(C._go_hdf5_H5T_STD_I16LE())
	T_STD_I32BE *DataType = new_dtype(C._go_hdf5_H5T_STD_I32BE())
	T_STD_I32LE *DataType = new_dtype(C._go_hdf5_H5T_STD_I32LE())
	T_STD_I64BE *DataType = new_dtype(C._go_hdf5_H5T_STD_I64BE())
	T_STD_I64LE *DataType = new_dtype(C._go_hdf5_H5T_STD_I64LE())
	T_STD_U8BE *DataType = new_dtype(C._go_hdf5_H5T_STD_U8BE())
	T_STD_U8LE *DataType = new_dtype(C._go_hdf5_H5T_STD_U8LE())
	T_STD_U16BE *DataType = new_dtype(C._go_hdf5_H5T_STD_U16BE())
	T_STD_U16LE *DataType = new_dtype(C._go_hdf5_H5T_STD_U16LE())
	T_STD_U32BE *DataType = new_dtype(C._go_hdf5_H5T_STD_U32BE())
	T_STD_U32LE *DataType = new_dtype(C._go_hdf5_H5T_STD_U32LE())
	T_STD_U64BE *DataType = new_dtype(C._go_hdf5_H5T_STD_U64BE())
	T_STD_U64LE *DataType = new_dtype(C._go_hdf5_H5T_STD_U64LE())
	T_STD_B8BE *DataType = new_dtype(C._go_hdf5_H5T_STD_B8BE())
	T_STD_B8LE *DataType = new_dtype(C._go_hdf5_H5T_STD_B8LE())

	T_STD_B16BE *DataType = new_dtype(C._go_hdf5_H5T_STD_B16BE())
	T_STD_B16LE *DataType = new_dtype(C._go_hdf5_H5T_STD_B16LE())
	T_STD_B32BE *DataType = new_dtype(C._go_hdf5_H5T_STD_B32BE())
	T_STD_B32LE *DataType = new_dtype(C._go_hdf5_H5T_STD_B32LE())
	T_STD_B64BE *DataType = new_dtype(C._go_hdf5_H5T_STD_B64BE())
	T_STD_B64LE *DataType = new_dtype(C._go_hdf5_H5T_STD_B64LE())
	T_STD_REF_OBJ *DataType = new_dtype(C._go_hdf5_H5T_STD_REF_OBJ())
	T_STD_REF_DSETREG *DataType = new_dtype(C._go_hdf5_H5T_STD_REF_DSETREG())

	T_IEEE_F32BE *DataType = new_dtype(C._go_hdf5_H5T_IEEE_F32BE())
	T_IEEE_F32LE *DataType = new_dtype(C._go_hdf5_H5T_IEEE_F32LE())
	T_IEEE_F64BE *DataType = new_dtype(C._go_hdf5_H5T_IEEE_F64BE())
	T_IEEE_F64LE *DataType = new_dtype(C._go_hdf5_H5T_IEEE_F64LE())

	T_UNIX_D32BE *DataType = new_dtype(C._go_hdf5_H5T_UNIX_D32BE())
	T_UNIX_D32LE *DataType = new_dtype(C._go_hdf5_H5T_UNIX_D32LE())
	T_UNIX_D64BE *DataType = new_dtype(C._go_hdf5_H5T_UNIX_D64BE())
	T_UNIX_D64LE *DataType = new_dtype(C._go_hdf5_H5T_UNIX_D64LE())

	T_INTEL_I8 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_I8())
	T_INTEL_I16 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_I16())
	T_INTEL_I32 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_I32())
	T_INTEL_I64 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_I64())
	T_INTEL_U8 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_U8())
	T_INTEL_U16 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_U16())
	T_INTEL_U32 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_U32())
	T_INTEL_U64 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_U64())
	T_INTEL_B8 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_B8())
	T_INTEL_B16 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_B16())
	T_INTEL_B32 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_B32())
	T_INTEL_B64 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_B64())
	T_INTEL_F32 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_F32())
	T_INTEL_F64 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_F64())

	T_ALPHA_I8 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_I8())
	T_ALPHA_I16 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_I16())
	T_ALPHA_I32 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_I32())
	T_ALPHA_I64 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_I64())
	T_ALPHA_U8 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_U8())
	T_ALPHA_U16 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_U16())
	T_ALPHA_U32 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_U32())
	T_ALPHA_U64 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_U64())
	T_ALPHA_B8 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_B8())
	T_ALPHA_B16 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_B16())
	T_ALPHA_B32 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_B32())
	T_ALPHA_B64 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_B64())
	T_ALPHA_F32 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_F32())
	T_ALPHA_F64 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_F64())

	T_MIPS_I8 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_I8())
	T_MIPS_I16 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_I16())
	T_MIPS_I32 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_I32())
	T_MIPS_I64 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_I64())
	T_MIPS_U8 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_U8())
	T_MIPS_U16 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_U16())
	T_MIPS_U32 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_U32())
	T_MIPS_U64 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_U64())
	T_MIPS_B8 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_B8())
	T_MIPS_B16 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_B16())
	T_MIPS_B32 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_B32())
	T_MIPS_B64 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_B64())
	T_MIPS_F32 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_F32())
	T_MIPS_F64 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_F64())

	T_NATIVE_CHAR *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_CHAR())
	T_NATIVE_INT *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_INT())
	T_NATIVE_FLOAT *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_FLOAT())
	T_NATIVE_SCHAR *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_SCHAR())
	T_NATIVE_UCHAR *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_UCHAR())
	T_NATIVE_SHORT *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_SHORT())
	T_NATIVE_USHORT *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_USHORT())
	T_NATIVE_UINT *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_UINT())
	T_NATIVE_LONG *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_LONG())
	T_NATIVE_ULONG *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_ULONG())
	T_NATIVE_LLONG *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_LLONG())
	T_NATIVE_ULLONG *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_ULLONG())
	T_NATIVE_DOUBLE *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_DOUBLE())
	//#if H5_SIZEOF_LONG_DOUBLE !=0
	T_NATIVE_LDOUBLE *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_LDOUBLE())
	//#endif
	T_NATIVE_B8 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_B8())
	T_NATIVE_B16 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_B16())
	T_NATIVE_B32 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_B32())
	T_NATIVE_B64 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_B64())
	T_NATIVE_OPAQUE *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_OPAQUE())
	T_NATIVE_HSIZE *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_HSIZE())
	T_NATIVE_HSSIZE *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_HSSIZE())
	T_NATIVE_HERR *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_HERR())
	T_NATIVE_HBOOL *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_HBOOL())

	T_NATIVE_INT8 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_INT8())
	T_NATIVE_UINT8 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_UINT8())
	T_NATIVE_INT16 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_INT16())
	T_NATIVE_UINT16 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_UINT16())
	T_NATIVE_INT32 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_INT32())
	T_NATIVE_UINT32 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_UINT32())
	T_NATIVE_INT64 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_INT64())
	T_NATIVE_UINT64 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_UINT64())
)

func new_dtype(id C.hid_t) *DataType {
	t := &DataType{id:id}
	//runtime.SetFinalizer(t, (*DataType).h5t_finalizer)
	return t
}

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
	t = new_dtype(hid)
	return
}

func (t *DataType) h5t_finalizer() {
	err := t.Close()
	if err != nil {
		panic(fmt.Sprintf("error closing datatype: %s",err))
	}
}

// Releases a datatype.
// herr_t H5Tclose( hid_t dtype_id ) 
func (t *DataType) Close() os.Error {
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
	o := new_dtype(hid)
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

// ---------------------------------------------------------------------------

// array data type
type ArrayType struct {
	DataType
}

func new_array_type(id C.hid_t) *ArrayType {
	t := &ArrayType{DataType{id:id}}
	//runtime.SetFinalizer(t, (*DataType).h5t_finalizer)
	return t
}

func NewArrayType(base_type *DataType, dims []int) (*ArrayType, os.Error) {
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
	c_rank := int(C.H5Tget_array_dims(t.id, c_dims))
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

func NewVarLenType(base_type *DataType) (*VarLenType, os.Error) {
	hid := C.H5Tvlen_create(base_type.id)
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	dt := new_vltype(hid)
	return dt, err
}

func new_vltype(id C.hid_t) *VarLenType {
	t := &VarLenType{DataType{id:id}}
	//runtime.SetFinalizer(t, (*DataType).h5t_finalizer)
	return t
}

// Determines whether datatype is a variable-length string.
// htri_t H5Tis_variable_str( hid_t dtype_id )
func (vl *VarLenType) IsVariableStr() bool {
	o := int(C.H5Tis_variable_str(vl.id))
	if o> 0 {
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
func (t *CompType) MemberType(mbr_idx int) (*DataType, os.Error) {
	hid := C.H5Tget_member_type(t.id, C.uint(mbr_idx))
	err := togo_err(C.herr_t(int(hid)))
	if err != nil {
		return nil, err
	}
	dt := new_dtype(hid)
	return dt, nil
}

// Adds a new member to a compound datatype. 
// herr_t H5Tinsert( hid_t dtype_id, const char * name, size_t offset, hid_t field_id ) 
func (t *CompType) Insert(name string, offset int, field *DataType) os.Error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	fmt.Printf("inserting [%s] at offset:%d...\n", name, offset)
	err := C.H5Tinsert(t.id, c_name, C.size_t(offset), field.id)
	return togo_err(err)
}

// Recursively removes padding from within a compound datatype. 
// herr_t H5Tpack( hid_t dtype_id ) 
func (t *CompType) Pack() os.Error {
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

// -----------------------------------------

// create a data-type from a golang value
func NewDataTypeFromValue(v interface{}) *DataType {
	t := reflect.Typeof(v)
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
		dt = T_C_S1

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
		//panic("sorry, arrays not yet supported")

	case  reflect.Slice:
		elem_type := new_dataTypeFromType(t.Elem())
		vlen_dt, err := NewVarLenType(elem_type)
		if err != nil {
			panic(err)
		}
		dt, err = vlen_dt.Copy()
		if err != nil {
			panic(err)
		}
		//panic("sorry, arrays and slices not yet supported")

	case reflect.Struct:
		sz := int(t.Size())
		fmt.Printf("==> struct [%s] (sz: %d)...\n", t.Name(), sz)
		hdf_dt, err := CreateDataType(T_COMPOUND, sz)
		if err != nil {
			panic(err)
		}
		cdt := &CompType{*hdf_dt}
		n := t.NumField()
		for i := 0; i < n; i++ {
			f := t.Field(i)
			field_dt := new_dataTypeFromType(f.Type)
			if field_dt == nil {
				panic(fmt.Sprintf("pb with field [%d-%s]",i,f.Name))
			}
			err = cdt.Insert(f.Name, int(f.Offset), field_dt)
			if err != nil {
				panic(fmt.Sprintf("pb with field [%d-%s]: %s",i,f.Name,err))
			}
		}
		cdt.Lock()
		dt, err = cdt.Copy()
		if err != nil {
			panic(err)
		}
		fmt.Printf("==> struct [%s] (sz: %d)... [done]\n", t.Name(), sz)

	case reflect.Ptr:
		panic("sorry, pointers not yet supported")

	default:
		panic(fmt.Sprintf("unhandled kind (%d)", t.Kind()))
	}

	return dt
}
// EOF
