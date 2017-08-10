// Copyright ©2017 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

/*
 #include "hdf5.h"

 #include <stdlib.h>
 #include <string.h>

 hid_t _H5T_C_S1() { return H5T_C_S1; }
 hid_t _H5T_FORTRAN_S1() { return H5T_FORTRAN_S1; }

 hid_t _H5T_STD_I8BE() { return H5T_STD_I8BE; }
 hid_t _H5T_STD_I8LE() { return H5T_STD_I8LE; }
 hid_t _H5T_STD_I16BE() { return H5T_STD_I16BE; }
 hid_t _H5T_STD_I16LE() { return H5T_STD_I16LE; }
 hid_t _H5T_STD_I32BE() { return H5T_STD_I32BE; }
 hid_t _H5T_STD_I32LE() { return H5T_STD_I32LE; }
 hid_t _H5T_STD_I64BE() { return H5T_STD_I64BE; }
 hid_t _H5T_STD_I64LE() { return H5T_STD_I64LE; }
 hid_t _H5T_STD_U8BE() { return H5T_STD_U8BE; }
 hid_t _H5T_STD_U8LE() { return H5T_STD_U8LE; }
 hid_t _H5T_STD_U16BE() { return H5T_STD_U16BE; }
 hid_t _H5T_STD_U16LE() { return H5T_STD_U16LE; }
 hid_t _H5T_STD_U32BE() { return H5T_STD_U32BE; }
 hid_t _H5T_STD_U32LE() { return H5T_STD_U32LE; }
 hid_t _H5T_STD_U64BE() { return H5T_STD_U64BE; }
 hid_t _H5T_STD_U64LE() { return H5T_STD_U64LE; }
 hid_t _H5T_STD_B8BE() { return H5T_STD_B8BE; }
 hid_t _H5T_STD_B8LE() { return H5T_STD_B8LE; }

 hid_t _H5T_STD_B16BE() { return H5T_STD_B16BE; }
 hid_t _H5T_STD_B16LE() { return H5T_STD_B16LE; }
 hid_t _H5T_STD_B32BE() { return H5T_STD_B32BE; }
 hid_t _H5T_STD_B32LE() { return H5T_STD_B32LE; }
 hid_t _H5T_STD_B64BE() { return H5T_STD_B64BE; }
 hid_t _H5T_STD_B64LE() { return H5T_STD_B64LE; }
 hid_t _H5T_STD_REF_OBJ() { return H5T_STD_REF_OBJ; }
 hid_t _H5T_STD_REF_DSETREG() { return H5T_STD_REF_DSETREG; }

 hid_t _H5T_IEEE_F32BE() { return H5T_IEEE_F32BE; }
 hid_t _H5T_IEEE_F32LE() { return H5T_IEEE_F32LE; }
 hid_t _H5T_IEEE_F64BE() { return H5T_IEEE_F64BE; }
 hid_t _H5T_IEEE_F64LE() { return H5T_IEEE_F64LE; }

 hid_t _H5T_UNIX_D32BE() { return H5T_UNIX_D32BE; }
 hid_t _H5T_UNIX_D32LE() { return H5T_UNIX_D32LE; }
 hid_t _H5T_UNIX_D64BE() { return H5T_UNIX_D64BE; }
 hid_t _H5T_UNIX_D64LE() { return H5T_UNIX_D64LE; }

 hid_t _H5T_INTEL_I8() { return H5T_INTEL_I8; }
 hid_t _H5T_INTEL_I16() { return H5T_INTEL_I16; }
 hid_t _H5T_INTEL_I32() { return H5T_INTEL_I32; }
 hid_t _H5T_INTEL_I64() { return H5T_INTEL_I64; }
 hid_t _H5T_INTEL_U8() { return H5T_INTEL_U8; }
 hid_t _H5T_INTEL_U16() { return H5T_INTEL_U16; }
 hid_t _H5T_INTEL_U32() { return H5T_INTEL_U32; }
 hid_t _H5T_INTEL_U64() { return H5T_INTEL_U64; }
 hid_t _H5T_INTEL_B8() { return H5T_INTEL_B8; }
 hid_t _H5T_INTEL_B16() { return H5T_INTEL_B16; }
 hid_t _H5T_INTEL_B32() { return H5T_INTEL_B32; }
 hid_t _H5T_INTEL_B64() { return H5T_INTEL_B64; }
 hid_t _H5T_INTEL_F32() { return H5T_INTEL_F32; }
 hid_t _H5T_INTEL_F64() { return H5T_INTEL_F64; }

 hid_t _H5T_ALPHA_I8() { return H5T_ALPHA_I8; }
 hid_t _H5T_ALPHA_I16() { return H5T_ALPHA_I16; }
 hid_t _H5T_ALPHA_I32() { return H5T_ALPHA_I32; }
 hid_t _H5T_ALPHA_I64() { return H5T_ALPHA_I64; }
 hid_t _H5T_ALPHA_U8() { return H5T_ALPHA_U8; }
 hid_t _H5T_ALPHA_U16() { return H5T_ALPHA_U16; }
 hid_t _H5T_ALPHA_U32() { return H5T_ALPHA_U32; }
 hid_t _H5T_ALPHA_U64() { return H5T_ALPHA_U64; }
 hid_t _H5T_ALPHA_B8() { return H5T_ALPHA_B8; }
 hid_t _H5T_ALPHA_B16() { return H5T_ALPHA_B16; }
 hid_t _H5T_ALPHA_B32() { return H5T_ALPHA_B32; }
 hid_t _H5T_ALPHA_B64() { return H5T_ALPHA_B64; }
 hid_t _H5T_ALPHA_F32() { return H5T_ALPHA_F32; }
 hid_t _H5T_ALPHA_F64() { return H5T_ALPHA_F64; }

 hid_t _H5T_MIPS_I8() { return H5T_MIPS_I8; }
 hid_t _H5T_MIPS_I16() { return H5T_MIPS_I16; }
 hid_t _H5T_MIPS_I32() { return H5T_MIPS_I32; }
 hid_t _H5T_MIPS_I64() { return H5T_MIPS_I64; }
 hid_t _H5T_MIPS_U8() { return H5T_MIPS_U8; }
 hid_t _H5T_MIPS_U16() { return H5T_MIPS_U16; }
 hid_t _H5T_MIPS_U32() { return H5T_MIPS_U32; }
 hid_t _H5T_MIPS_U64() { return H5T_MIPS_U64; }
 hid_t _H5T_MIPS_B8() { return H5T_MIPS_B8; }
 hid_t _H5T_MIPS_B16() { return H5T_MIPS_B16; }
 hid_t _H5T_MIPS_B32() { return H5T_MIPS_B32; }
 hid_t _H5T_MIPS_B64() { return H5T_MIPS_B64; }
 hid_t _H5T_MIPS_F32() { return H5T_MIPS_F32; }
 hid_t _H5T_MIPS_F64() { return H5T_MIPS_F64; }

 hid_t _H5T_NATIVE_CHAR() { return H5T_NATIVE_CHAR; }
 hid_t _H5T_NATIVE_INT() { return H5T_NATIVE_INT; }
 hid_t _H5T_NATIVE_FLOAT() { return H5T_NATIVE_FLOAT; }
 hid_t _H5T_NATIVE_SCHAR() { return H5T_NATIVE_SCHAR; }
 hid_t _H5T_NATIVE_UCHAR() { return H5T_NATIVE_UCHAR; }
 hid_t _H5T_NATIVE_SHORT() { return H5T_NATIVE_SHORT; }
 hid_t _H5T_NATIVE_USHORT() { return H5T_NATIVE_USHORT; }
 hid_t _H5T_NATIVE_UINT() { return H5T_NATIVE_UINT; }
 hid_t _H5T_NATIVE_LONG() { return H5T_NATIVE_LONG; }
 hid_t _H5T_NATIVE_ULONG() { return H5T_NATIVE_ULONG; }
 hid_t _H5T_NATIVE_LLONG() { return H5T_NATIVE_LLONG; }
 hid_t _H5T_NATIVE_ULLONG() { return H5T_NATIVE_ULLONG; }
 hid_t _H5T_NATIVE_DOUBLE() { return H5T_NATIVE_DOUBLE; }
#if H5_SIZEOF_LONG_DOUBLE !=0
 hid_t _H5T_NATIVE_LDOUBLE() { return H5T_NATIVE_LDOUBLE; }
#endif
 hid_t _H5T_NATIVE_B8() { return H5T_NATIVE_B8; }
 hid_t _H5T_NATIVE_B16() { return H5T_NATIVE_B16; }
 hid_t _H5T_NATIVE_B32() { return H5T_NATIVE_B32; }
 hid_t _H5T_NATIVE_B64() { return H5T_NATIVE_B64; }
 hid_t _H5T_NATIVE_OPAQUE() { return H5T_NATIVE_OPAQUE; }
 hid_t _H5T_NATIVE_HSIZE() { return H5T_NATIVE_HSIZE; }
 hid_t _H5T_NATIVE_HSSIZE() { return H5T_NATIVE_HSSIZE; }
 hid_t _H5T_NATIVE_HERR() { return H5T_NATIVE_HERR; }
 hid_t _H5T_NATIVE_HBOOL() { return H5T_NATIVE_HBOOL; }

 hid_t _H5T_NATIVE_INT8() { return H5T_NATIVE_INT8; }
 hid_t _H5T_NATIVE_UINT8() { return H5T_NATIVE_UINT8; }
 hid_t _H5T_NATIVE_INT16() { return H5T_NATIVE_INT16; }
 hid_t _H5T_NATIVE_UINT16() { return H5T_NATIVE_UINT16; }
 hid_t _H5T_NATIVE_INT32() { return H5T_NATIVE_INT32; }
 hid_t _H5T_NATIVE_UINT32() { return H5T_NATIVE_UINT32; }
 hid_t _H5T_NATIVE_INT64() { return H5T_NATIVE_INT64; }
 hid_t _H5T_NATIVE_UINT64() { return H5T_NATIVE_UINT64; }
*/
import "C"

// list of predefined hdf5 data types
var (
	T_C_S1       *DataType = NewDataType(C._H5T_C_S1())
	T_FORTRAN_S1 *DataType = NewDataType(C._H5T_FORTRAN_S1())

	T_STD_I8BE  *DataType = NewDataType(C._H5T_STD_I8BE())
	T_STD_I8LE  *DataType = NewDataType(C._H5T_STD_I8LE())
	T_STD_I16BE *DataType = NewDataType(C._H5T_STD_I16BE())
	T_STD_I16LE *DataType = NewDataType(C._H5T_STD_I16LE())
	T_STD_I32BE *DataType = NewDataType(C._H5T_STD_I32BE())
	T_STD_I32LE *DataType = NewDataType(C._H5T_STD_I32LE())
	T_STD_I64BE *DataType = NewDataType(C._H5T_STD_I64BE())
	T_STD_I64LE *DataType = NewDataType(C._H5T_STD_I64LE())
	T_STD_U8BE  *DataType = NewDataType(C._H5T_STD_U8BE())
	T_STD_U8LE  *DataType = NewDataType(C._H5T_STD_U8LE())
	T_STD_U16BE *DataType = NewDataType(C._H5T_STD_U16BE())
	T_STD_U16LE *DataType = NewDataType(C._H5T_STD_U16LE())
	T_STD_U32BE *DataType = NewDataType(C._H5T_STD_U32BE())
	T_STD_U32LE *DataType = NewDataType(C._H5T_STD_U32LE())
	T_STD_U64BE *DataType = NewDataType(C._H5T_STD_U64BE())
	T_STD_U64LE *DataType = NewDataType(C._H5T_STD_U64LE())
	T_STD_B8BE  *DataType = NewDataType(C._H5T_STD_B8BE())
	T_STD_B8LE  *DataType = NewDataType(C._H5T_STD_B8LE())

	T_STD_B16BE       *DataType = NewDataType(C._H5T_STD_B16BE())
	T_STD_B16LE       *DataType = NewDataType(C._H5T_STD_B16LE())
	T_STD_B32BE       *DataType = NewDataType(C._H5T_STD_B32BE())
	T_STD_B32LE       *DataType = NewDataType(C._H5T_STD_B32LE())
	T_STD_B64BE       *DataType = NewDataType(C._H5T_STD_B64BE())
	T_STD_B64LE       *DataType = NewDataType(C._H5T_STD_B64LE())
	T_STD_REF_OBJ     *DataType = NewDataType(C._H5T_STD_REF_OBJ())
	T_STD_REF_DSETREG *DataType = NewDataType(C._H5T_STD_REF_DSETREG())

	T_IEEE_F32BE *DataType = NewDataType(C._H5T_IEEE_F32BE())
	T_IEEE_F32LE *DataType = NewDataType(C._H5T_IEEE_F32LE())
	T_IEEE_F64BE *DataType = NewDataType(C._H5T_IEEE_F64BE())
	T_IEEE_F64LE *DataType = NewDataType(C._H5T_IEEE_F64LE())

	T_UNIX_D32BE *DataType = NewDataType(C._H5T_UNIX_D32BE())
	T_UNIX_D32LE *DataType = NewDataType(C._H5T_UNIX_D32LE())
	T_UNIX_D64BE *DataType = NewDataType(C._H5T_UNIX_D64BE())
	T_UNIX_D64LE *DataType = NewDataType(C._H5T_UNIX_D64LE())

	T_INTEL_I8  *DataType = NewDataType(C._H5T_INTEL_I8())
	T_INTEL_I16 *DataType = NewDataType(C._H5T_INTEL_I16())
	T_INTEL_I32 *DataType = NewDataType(C._H5T_INTEL_I32())
	T_INTEL_I64 *DataType = NewDataType(C._H5T_INTEL_I64())
	T_INTEL_U8  *DataType = NewDataType(C._H5T_INTEL_U8())
	T_INTEL_U16 *DataType = NewDataType(C._H5T_INTEL_U16())
	T_INTEL_U32 *DataType = NewDataType(C._H5T_INTEL_U32())
	T_INTEL_U64 *DataType = NewDataType(C._H5T_INTEL_U64())
	T_INTEL_B8  *DataType = NewDataType(C._H5T_INTEL_B8())
	T_INTEL_B16 *DataType = NewDataType(C._H5T_INTEL_B16())
	T_INTEL_B32 *DataType = NewDataType(C._H5T_INTEL_B32())
	T_INTEL_B64 *DataType = NewDataType(C._H5T_INTEL_B64())
	T_INTEL_F32 *DataType = NewDataType(C._H5T_INTEL_F32())
	T_INTEL_F64 *DataType = NewDataType(C._H5T_INTEL_F64())

	T_ALPHA_I8  *DataType = NewDataType(C._H5T_ALPHA_I8())
	T_ALPHA_I16 *DataType = NewDataType(C._H5T_ALPHA_I16())
	T_ALPHA_I32 *DataType = NewDataType(C._H5T_ALPHA_I32())
	T_ALPHA_I64 *DataType = NewDataType(C._H5T_ALPHA_I64())
	T_ALPHA_U8  *DataType = NewDataType(C._H5T_ALPHA_U8())
	T_ALPHA_U16 *DataType = NewDataType(C._H5T_ALPHA_U16())
	T_ALPHA_U32 *DataType = NewDataType(C._H5T_ALPHA_U32())
	T_ALPHA_U64 *DataType = NewDataType(C._H5T_ALPHA_U64())
	T_ALPHA_B8  *DataType = NewDataType(C._H5T_ALPHA_B8())
	T_ALPHA_B16 *DataType = NewDataType(C._H5T_ALPHA_B16())
	T_ALPHA_B32 *DataType = NewDataType(C._H5T_ALPHA_B32())
	T_ALPHA_B64 *DataType = NewDataType(C._H5T_ALPHA_B64())
	T_ALPHA_F32 *DataType = NewDataType(C._H5T_ALPHA_F32())
	T_ALPHA_F64 *DataType = NewDataType(C._H5T_ALPHA_F64())

	T_MIPS_I8  *DataType = NewDataType(C._H5T_MIPS_I8())
	T_MIPS_I16 *DataType = NewDataType(C._H5T_MIPS_I16())
	T_MIPS_I32 *DataType = NewDataType(C._H5T_MIPS_I32())
	T_MIPS_I64 *DataType = NewDataType(C._H5T_MIPS_I64())
	T_MIPS_U8  *DataType = NewDataType(C._H5T_MIPS_U8())
	T_MIPS_U16 *DataType = NewDataType(C._H5T_MIPS_U16())
	T_MIPS_U32 *DataType = NewDataType(C._H5T_MIPS_U32())
	T_MIPS_U64 *DataType = NewDataType(C._H5T_MIPS_U64())
	T_MIPS_B8  *DataType = NewDataType(C._H5T_MIPS_B8())
	T_MIPS_B16 *DataType = NewDataType(C._H5T_MIPS_B16())
	T_MIPS_B32 *DataType = NewDataType(C._H5T_MIPS_B32())
	T_MIPS_B64 *DataType = NewDataType(C._H5T_MIPS_B64())
	T_MIPS_F32 *DataType = NewDataType(C._H5T_MIPS_F32())
	T_MIPS_F64 *DataType = NewDataType(C._H5T_MIPS_F64())

	T_NATIVE_CHAR   *DataType = NewDataType(C._H5T_NATIVE_CHAR())
	T_NATIVE_INT    *DataType = NewDataType(C._H5T_NATIVE_INT())
	T_NATIVE_FLOAT  *DataType = NewDataType(C._H5T_NATIVE_FLOAT())
	T_NATIVE_SCHAR  *DataType = NewDataType(C._H5T_NATIVE_SCHAR())
	T_NATIVE_UCHAR  *DataType = NewDataType(C._H5T_NATIVE_UCHAR())
	T_NATIVE_SHORT  *DataType = NewDataType(C._H5T_NATIVE_SHORT())
	T_NATIVE_USHORT *DataType = NewDataType(C._H5T_NATIVE_USHORT())
	T_NATIVE_UINT   *DataType = NewDataType(C._H5T_NATIVE_UINT())
	T_NATIVE_LONG   *DataType = NewDataType(C._H5T_NATIVE_LONG())
	T_NATIVE_ULONG  *DataType = NewDataType(C._H5T_NATIVE_ULONG())
	T_NATIVE_LLONG  *DataType = NewDataType(C._H5T_NATIVE_LLONG())
	T_NATIVE_ULLONG *DataType = NewDataType(C._H5T_NATIVE_ULLONG())
	T_NATIVE_DOUBLE *DataType = NewDataType(C._H5T_NATIVE_DOUBLE())
	//#if H5_SIZEOF_LONG_DOUBLE !=0
	T_NATIVE_LDOUBLE *DataType = NewDataType(C._H5T_NATIVE_LDOUBLE())
	//#endif
	T_NATIVE_B8     *DataType = NewDataType(C._H5T_NATIVE_B8())
	T_NATIVE_B16    *DataType = NewDataType(C._H5T_NATIVE_B16())
	T_NATIVE_B32    *DataType = NewDataType(C._H5T_NATIVE_B32())
	T_NATIVE_B64    *DataType = NewDataType(C._H5T_NATIVE_B64())
	T_NATIVE_OPAQUE *DataType = NewDataType(C._H5T_NATIVE_OPAQUE())
	T_NATIVE_HSIZE  *DataType = NewDataType(C._H5T_NATIVE_HSIZE())
	T_NATIVE_HSSIZE *DataType = NewDataType(C._H5T_NATIVE_HSSIZE())
	T_NATIVE_HERR   *DataType = NewDataType(C._H5T_NATIVE_HERR())
	T_NATIVE_HBOOL  *DataType = NewDataType(C._H5T_NATIVE_HBOOL())

	T_NATIVE_INT8   *DataType = NewDataType(C._H5T_NATIVE_INT8())
	T_NATIVE_UINT8  *DataType = NewDataType(C._H5T_NATIVE_UINT8())
	T_NATIVE_INT16  *DataType = NewDataType(C._H5T_NATIVE_INT16())
	T_NATIVE_UINT16 *DataType = NewDataType(C._H5T_NATIVE_UINT16())
	T_NATIVE_INT32  *DataType = NewDataType(C._H5T_NATIVE_INT32())
	T_NATIVE_UINT32 *DataType = NewDataType(C._H5T_NATIVE_UINT32())
	T_NATIVE_INT64  *DataType = NewDataType(C._H5T_NATIVE_INT64())
	T_NATIVE_UINT64 *DataType = NewDataType(C._H5T_NATIVE_UINT64())

	T_GO_STRING *DataType = makeGoStringDataType()
)

//
var h5t_VARIABLE int64 = C.H5T_VARIABLE

func makeGoStringDataType() *DataType {
	dt, err := T_C_S1.Copy()
	if err != nil {
		panic(err)
	}
	err = dt.SetSize(uint(h5t_VARIABLE))
	if err != nil {
		panic(err)
	}
	return dt
}
