// Copyright ©2017 The Gonum Authors. All rights reserved.
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

 size_t size_t_H5T_VARIABLE() { return H5T_VARIABLE; }
*/
import "C"

// list of predefined hdf5 data types
var (
	T_C_S1       *Datatype = NewDatatype(C._H5T_C_S1())
	T_FORTRAN_S1 *Datatype = NewDatatype(C._H5T_FORTRAN_S1())

	T_STD_I8BE  *Datatype = NewDatatype(C._H5T_STD_I8BE())
	T_STD_I8LE  *Datatype = NewDatatype(C._H5T_STD_I8LE())
	T_STD_I16BE *Datatype = NewDatatype(C._H5T_STD_I16BE())
	T_STD_I16LE *Datatype = NewDatatype(C._H5T_STD_I16LE())
	T_STD_I32BE *Datatype = NewDatatype(C._H5T_STD_I32BE())
	T_STD_I32LE *Datatype = NewDatatype(C._H5T_STD_I32LE())
	T_STD_I64BE *Datatype = NewDatatype(C._H5T_STD_I64BE())
	T_STD_I64LE *Datatype = NewDatatype(C._H5T_STD_I64LE())
	T_STD_U8BE  *Datatype = NewDatatype(C._H5T_STD_U8BE())
	T_STD_U8LE  *Datatype = NewDatatype(C._H5T_STD_U8LE())
	T_STD_U16BE *Datatype = NewDatatype(C._H5T_STD_U16BE())
	T_STD_U16LE *Datatype = NewDatatype(C._H5T_STD_U16LE())
	T_STD_U32BE *Datatype = NewDatatype(C._H5T_STD_U32BE())
	T_STD_U32LE *Datatype = NewDatatype(C._H5T_STD_U32LE())
	T_STD_U64BE *Datatype = NewDatatype(C._H5T_STD_U64BE())
	T_STD_U64LE *Datatype = NewDatatype(C._H5T_STD_U64LE())
	T_STD_B8BE  *Datatype = NewDatatype(C._H5T_STD_B8BE())
	T_STD_B8LE  *Datatype = NewDatatype(C._H5T_STD_B8LE())

	T_STD_B16BE       *Datatype = NewDatatype(C._H5T_STD_B16BE())
	T_STD_B16LE       *Datatype = NewDatatype(C._H5T_STD_B16LE())
	T_STD_B32BE       *Datatype = NewDatatype(C._H5T_STD_B32BE())
	T_STD_B32LE       *Datatype = NewDatatype(C._H5T_STD_B32LE())
	T_STD_B64BE       *Datatype = NewDatatype(C._H5T_STD_B64BE())
	T_STD_B64LE       *Datatype = NewDatatype(C._H5T_STD_B64LE())
	T_STD_REF_OBJ     *Datatype = NewDatatype(C._H5T_STD_REF_OBJ())
	T_STD_REF_DSETREG *Datatype = NewDatatype(C._H5T_STD_REF_DSETREG())

	T_IEEE_F32BE *Datatype = NewDatatype(C._H5T_IEEE_F32BE())
	T_IEEE_F32LE *Datatype = NewDatatype(C._H5T_IEEE_F32LE())
	T_IEEE_F64BE *Datatype = NewDatatype(C._H5T_IEEE_F64BE())
	T_IEEE_F64LE *Datatype = NewDatatype(C._H5T_IEEE_F64LE())

	T_UNIX_D32BE *Datatype = NewDatatype(C._H5T_UNIX_D32BE())
	T_UNIX_D32LE *Datatype = NewDatatype(C._H5T_UNIX_D32LE())
	T_UNIX_D64BE *Datatype = NewDatatype(C._H5T_UNIX_D64BE())
	T_UNIX_D64LE *Datatype = NewDatatype(C._H5T_UNIX_D64LE())

	T_INTEL_I8  *Datatype = NewDatatype(C._H5T_INTEL_I8())
	T_INTEL_I16 *Datatype = NewDatatype(C._H5T_INTEL_I16())
	T_INTEL_I32 *Datatype = NewDatatype(C._H5T_INTEL_I32())
	T_INTEL_I64 *Datatype = NewDatatype(C._H5T_INTEL_I64())
	T_INTEL_U8  *Datatype = NewDatatype(C._H5T_INTEL_U8())
	T_INTEL_U16 *Datatype = NewDatatype(C._H5T_INTEL_U16())
	T_INTEL_U32 *Datatype = NewDatatype(C._H5T_INTEL_U32())
	T_INTEL_U64 *Datatype = NewDatatype(C._H5T_INTEL_U64())
	T_INTEL_B8  *Datatype = NewDatatype(C._H5T_INTEL_B8())
	T_INTEL_B16 *Datatype = NewDatatype(C._H5T_INTEL_B16())
	T_INTEL_B32 *Datatype = NewDatatype(C._H5T_INTEL_B32())
	T_INTEL_B64 *Datatype = NewDatatype(C._H5T_INTEL_B64())
	T_INTEL_F32 *Datatype = NewDatatype(C._H5T_INTEL_F32())
	T_INTEL_F64 *Datatype = NewDatatype(C._H5T_INTEL_F64())

	T_ALPHA_I8  *Datatype = NewDatatype(C._H5T_ALPHA_I8())
	T_ALPHA_I16 *Datatype = NewDatatype(C._H5T_ALPHA_I16())
	T_ALPHA_I32 *Datatype = NewDatatype(C._H5T_ALPHA_I32())
	T_ALPHA_I64 *Datatype = NewDatatype(C._H5T_ALPHA_I64())
	T_ALPHA_U8  *Datatype = NewDatatype(C._H5T_ALPHA_U8())
	T_ALPHA_U16 *Datatype = NewDatatype(C._H5T_ALPHA_U16())
	T_ALPHA_U32 *Datatype = NewDatatype(C._H5T_ALPHA_U32())
	T_ALPHA_U64 *Datatype = NewDatatype(C._H5T_ALPHA_U64())
	T_ALPHA_B8  *Datatype = NewDatatype(C._H5T_ALPHA_B8())
	T_ALPHA_B16 *Datatype = NewDatatype(C._H5T_ALPHA_B16())
	T_ALPHA_B32 *Datatype = NewDatatype(C._H5T_ALPHA_B32())
	T_ALPHA_B64 *Datatype = NewDatatype(C._H5T_ALPHA_B64())
	T_ALPHA_F32 *Datatype = NewDatatype(C._H5T_ALPHA_F32())
	T_ALPHA_F64 *Datatype = NewDatatype(C._H5T_ALPHA_F64())

	T_MIPS_I8  *Datatype = NewDatatype(C._H5T_MIPS_I8())
	T_MIPS_I16 *Datatype = NewDatatype(C._H5T_MIPS_I16())
	T_MIPS_I32 *Datatype = NewDatatype(C._H5T_MIPS_I32())
	T_MIPS_I64 *Datatype = NewDatatype(C._H5T_MIPS_I64())
	T_MIPS_U8  *Datatype = NewDatatype(C._H5T_MIPS_U8())
	T_MIPS_U16 *Datatype = NewDatatype(C._H5T_MIPS_U16())
	T_MIPS_U32 *Datatype = NewDatatype(C._H5T_MIPS_U32())
	T_MIPS_U64 *Datatype = NewDatatype(C._H5T_MIPS_U64())
	T_MIPS_B8  *Datatype = NewDatatype(C._H5T_MIPS_B8())
	T_MIPS_B16 *Datatype = NewDatatype(C._H5T_MIPS_B16())
	T_MIPS_B32 *Datatype = NewDatatype(C._H5T_MIPS_B32())
	T_MIPS_B64 *Datatype = NewDatatype(C._H5T_MIPS_B64())
	T_MIPS_F32 *Datatype = NewDatatype(C._H5T_MIPS_F32())
	T_MIPS_F64 *Datatype = NewDatatype(C._H5T_MIPS_F64())

	T_NATIVE_CHAR   *Datatype = NewDatatype(C._H5T_NATIVE_CHAR())
	T_NATIVE_INT    *Datatype = NewDatatype(C._H5T_NATIVE_INT())
	T_NATIVE_FLOAT  *Datatype = NewDatatype(C._H5T_NATIVE_FLOAT())
	T_NATIVE_SCHAR  *Datatype = NewDatatype(C._H5T_NATIVE_SCHAR())
	T_NATIVE_UCHAR  *Datatype = NewDatatype(C._H5T_NATIVE_UCHAR())
	T_NATIVE_SHORT  *Datatype = NewDatatype(C._H5T_NATIVE_SHORT())
	T_NATIVE_USHORT *Datatype = NewDatatype(C._H5T_NATIVE_USHORT())
	T_NATIVE_UINT   *Datatype = NewDatatype(C._H5T_NATIVE_UINT())
	T_NATIVE_LONG   *Datatype = NewDatatype(C._H5T_NATIVE_LONG())
	T_NATIVE_ULONG  *Datatype = NewDatatype(C._H5T_NATIVE_ULONG())
	T_NATIVE_LLONG  *Datatype = NewDatatype(C._H5T_NATIVE_LLONG())
	T_NATIVE_ULLONG *Datatype = NewDatatype(C._H5T_NATIVE_ULLONG())
	T_NATIVE_DOUBLE *Datatype = NewDatatype(C._H5T_NATIVE_DOUBLE())
	//#if H5_SIZEOF_LONG_DOUBLE !=0
	T_NATIVE_LDOUBLE *Datatype = NewDatatype(C._H5T_NATIVE_LDOUBLE())
	//#endif
	T_NATIVE_B8     *Datatype = NewDatatype(C._H5T_NATIVE_B8())
	T_NATIVE_B16    *Datatype = NewDatatype(C._H5T_NATIVE_B16())
	T_NATIVE_B32    *Datatype = NewDatatype(C._H5T_NATIVE_B32())
	T_NATIVE_B64    *Datatype = NewDatatype(C._H5T_NATIVE_B64())
	T_NATIVE_OPAQUE *Datatype = NewDatatype(C._H5T_NATIVE_OPAQUE())
	T_NATIVE_HSIZE  *Datatype = NewDatatype(C._H5T_NATIVE_HSIZE())
	T_NATIVE_HSSIZE *Datatype = NewDatatype(C._H5T_NATIVE_HSSIZE())
	T_NATIVE_HERR   *Datatype = NewDatatype(C._H5T_NATIVE_HERR())
	T_NATIVE_HBOOL  *Datatype = NewDatatype(C._H5T_NATIVE_HBOOL())

	T_NATIVE_INT8   *Datatype = NewDatatype(C._H5T_NATIVE_INT8())
	T_NATIVE_UINT8  *Datatype = NewDatatype(C._H5T_NATIVE_UINT8())
	T_NATIVE_INT16  *Datatype = NewDatatype(C._H5T_NATIVE_INT16())
	T_NATIVE_UINT16 *Datatype = NewDatatype(C._H5T_NATIVE_UINT16())
	T_NATIVE_INT32  *Datatype = NewDatatype(C._H5T_NATIVE_INT32())
	T_NATIVE_UINT32 *Datatype = NewDatatype(C._H5T_NATIVE_UINT32())
	T_NATIVE_INT64  *Datatype = NewDatatype(C._H5T_NATIVE_INT64())
	T_NATIVE_UINT64 *Datatype = NewDatatype(C._H5T_NATIVE_UINT64())

	T_GO_STRING *Datatype = makeGoStringDatatype()
)

//
var h5t_VARIABLE = int(C.size_t_H5T_VARIABLE())

func makeGoStringDatatype() *Datatype {
	dt, err := T_C_S1.Copy()
	if err != nil {
		panic(err)
	}
	err = dt.SetSize(h5t_VARIABLE)
	if err != nil {
		panic(err)
	}
	dt.goPtrPathLen = 1 // This is the first field of the string header.
	return dt
}
