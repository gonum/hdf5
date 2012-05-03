package hdf5

/*
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

 //#include "cgo_h5t_conv.h"
 */
import "C"

// list of predefined hdf5 data types
var (
	T_C_S1 *DataType = new_dtype(C._go_hdf5_H5T_C_S1(), _go_string_t)
	T_FORTRAN_S1 *DataType = new_dtype(C._go_hdf5_H5T_FORTRAN_S1(), _go_string_t)

	T_STD_I8BE *DataType = new_dtype(C._go_hdf5_H5T_STD_I8BE(), nil)
	T_STD_I8LE *DataType = new_dtype(C._go_hdf5_H5T_STD_I8LE(), nil)
	T_STD_I16BE *DataType = new_dtype(C._go_hdf5_H5T_STD_I16BE(), nil)
	T_STD_I16LE *DataType = new_dtype(C._go_hdf5_H5T_STD_I16LE(), nil)
	T_STD_I32BE *DataType = new_dtype(C._go_hdf5_H5T_STD_I32BE(), nil)
	T_STD_I32LE *DataType = new_dtype(C._go_hdf5_H5T_STD_I32LE(), nil)
	T_STD_I64BE *DataType = new_dtype(C._go_hdf5_H5T_STD_I64BE(), nil)
	T_STD_I64LE *DataType = new_dtype(C._go_hdf5_H5T_STD_I64LE(), nil)
	T_STD_U8BE *DataType = new_dtype(C._go_hdf5_H5T_STD_U8BE(), nil)
	T_STD_U8LE *DataType = new_dtype(C._go_hdf5_H5T_STD_U8LE(), nil)
	T_STD_U16BE *DataType = new_dtype(C._go_hdf5_H5T_STD_U16BE(), nil)
	T_STD_U16LE *DataType = new_dtype(C._go_hdf5_H5T_STD_U16LE(), nil)
	T_STD_U32BE *DataType = new_dtype(C._go_hdf5_H5T_STD_U32BE(), nil)
	T_STD_U32LE *DataType = new_dtype(C._go_hdf5_H5T_STD_U32LE(), nil)
	T_STD_U64BE *DataType = new_dtype(C._go_hdf5_H5T_STD_U64BE(), nil)
	T_STD_U64LE *DataType = new_dtype(C._go_hdf5_H5T_STD_U64LE(), nil)
	T_STD_B8BE *DataType = new_dtype(C._go_hdf5_H5T_STD_B8BE(), nil)
	T_STD_B8LE *DataType = new_dtype(C._go_hdf5_H5T_STD_B8LE(), nil)

	T_STD_B16BE *DataType = new_dtype(C._go_hdf5_H5T_STD_B16BE(), nil)
	T_STD_B16LE *DataType = new_dtype(C._go_hdf5_H5T_STD_B16LE(), nil)
	T_STD_B32BE *DataType = new_dtype(C._go_hdf5_H5T_STD_B32BE(), nil)
	T_STD_B32LE *DataType = new_dtype(C._go_hdf5_H5T_STD_B32LE(), nil)
	T_STD_B64BE *DataType = new_dtype(C._go_hdf5_H5T_STD_B64BE(), nil)
	T_STD_B64LE *DataType = new_dtype(C._go_hdf5_H5T_STD_B64LE(), nil)
	T_STD_REF_OBJ *DataType = new_dtype(C._go_hdf5_H5T_STD_REF_OBJ(), nil)
	T_STD_REF_DSETREG *DataType = new_dtype(C._go_hdf5_H5T_STD_REF_DSETREG(), nil)

	T_IEEE_F32BE *DataType = new_dtype(C._go_hdf5_H5T_IEEE_F32BE(), nil)
	T_IEEE_F32LE *DataType = new_dtype(C._go_hdf5_H5T_IEEE_F32LE(), nil)
	T_IEEE_F64BE *DataType = new_dtype(C._go_hdf5_H5T_IEEE_F64BE(), nil)
	T_IEEE_F64LE *DataType = new_dtype(C._go_hdf5_H5T_IEEE_F64LE(), nil)

	T_UNIX_D32BE *DataType = new_dtype(C._go_hdf5_H5T_UNIX_D32BE(), nil)
	T_UNIX_D32LE *DataType = new_dtype(C._go_hdf5_H5T_UNIX_D32LE(), nil)
	T_UNIX_D64BE *DataType = new_dtype(C._go_hdf5_H5T_UNIX_D64BE(), nil)
	T_UNIX_D64LE *DataType = new_dtype(C._go_hdf5_H5T_UNIX_D64LE(), nil)

	T_INTEL_I8 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_I8(), nil)
	T_INTEL_I16 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_I16(), nil)
	T_INTEL_I32 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_I32(), nil)
	T_INTEL_I64 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_I64(), nil)
	T_INTEL_U8 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_U8(), nil)
	T_INTEL_U16 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_U16(), nil)
	T_INTEL_U32 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_U32(), nil)
	T_INTEL_U64 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_U64(), nil)
	T_INTEL_B8 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_B8(), nil)
	T_INTEL_B16 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_B16(), nil)
	T_INTEL_B32 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_B32(), nil)
	T_INTEL_B64 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_B64(), nil)
	T_INTEL_F32 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_F32(), nil)
	T_INTEL_F64 *DataType = new_dtype(C._go_hdf5_H5T_INTEL_F64(), nil)

	T_ALPHA_I8 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_I8(), nil)
	T_ALPHA_I16 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_I16(), nil)
	T_ALPHA_I32 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_I32(), nil)
	T_ALPHA_I64 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_I64(), nil)
	T_ALPHA_U8 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_U8(), nil)
	T_ALPHA_U16 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_U16(), nil)
	T_ALPHA_U32 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_U32(), nil)
	T_ALPHA_U64 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_U64(), nil)
	T_ALPHA_B8 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_B8(), nil)
	T_ALPHA_B16 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_B16(), nil)
	T_ALPHA_B32 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_B32(), nil)
	T_ALPHA_B64 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_B64(), nil)
	T_ALPHA_F32 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_F32(), nil)
	T_ALPHA_F64 *DataType = new_dtype(C._go_hdf5_H5T_ALPHA_F64(), nil)

	T_MIPS_I8 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_I8(), nil)
	T_MIPS_I16 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_I16(), nil)
	T_MIPS_I32 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_I32(), nil)
	T_MIPS_I64 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_I64(), nil)
	T_MIPS_U8 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_U8(), nil)
	T_MIPS_U16 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_U16(), nil)
	T_MIPS_U32 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_U32(), nil)
	T_MIPS_U64 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_U64(), nil)
	T_MIPS_B8 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_B8(), nil)
	T_MIPS_B16 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_B16(), nil)
	T_MIPS_B32 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_B32(), nil)
	T_MIPS_B64 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_B64(), nil)
	T_MIPS_F32 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_F32(), nil)
	T_MIPS_F64 *DataType = new_dtype(C._go_hdf5_H5T_MIPS_F64(), nil)

	T_NATIVE_CHAR *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_CHAR(), nil)
	T_NATIVE_INT *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_INT(), _go_int_t)
	T_NATIVE_FLOAT *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_FLOAT(), _go_float32_t)
	T_NATIVE_SCHAR *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_SCHAR(), nil)
	T_NATIVE_UCHAR *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_UCHAR(), nil)
	T_NATIVE_SHORT *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_SHORT(), nil)
	T_NATIVE_USHORT *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_USHORT(), nil)
	T_NATIVE_UINT *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_UINT(), _go_uint_t)
	T_NATIVE_LONG *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_LONG(), nil)
	T_NATIVE_ULONG *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_ULONG(), nil)
	T_NATIVE_LLONG *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_LLONG(), nil)
	T_NATIVE_ULLONG *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_ULLONG(), nil)
	T_NATIVE_DOUBLE *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_DOUBLE(), _go_float64_t)
	//#if H5_SIZEOF_LONG_DOUBLE !=0
	T_NATIVE_LDOUBLE *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_LDOUBLE(), nil)
	//#endif
	T_NATIVE_B8 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_B8(), nil)
	T_NATIVE_B16 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_B16(), nil)
	T_NATIVE_B32 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_B32(), nil)
	T_NATIVE_B64 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_B64(), nil)
	T_NATIVE_OPAQUE *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_OPAQUE(), nil)
	T_NATIVE_HSIZE *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_HSIZE(), nil)
	T_NATIVE_HSSIZE *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_HSSIZE(), nil)
	T_NATIVE_HERR *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_HERR(), nil)
	T_NATIVE_HBOOL *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_HBOOL(), nil)

	T_NATIVE_INT8 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_INT8(), _go_int8_t)
	T_NATIVE_UINT8 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_UINT8(), _go_uint8_t)
	T_NATIVE_INT16 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_INT16(), _go_int16_t)
	T_NATIVE_UINT16 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_UINT16(), _go_uint16_t)
	T_NATIVE_INT32 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_INT32(), _go_int32_t)
	T_NATIVE_UINT32 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_UINT32(), _go_uint32_t)
	T_NATIVE_INT64 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_INT64(), _go_int64_t)
	T_NATIVE_UINT64 *DataType = new_dtype(C._go_hdf5_H5T_NATIVE_UINT64(), _go_uint64_t)

	T_GO_STRING *DataType = _make_go_string_datatype()
)

func _make_go_string_datatype() *DataType {
	dt, err := T_C_S1.Copy()
	if err != nil {
		panic(err)
	}
	dt.SetSize(C.H5T_VARIABLE)

	return dt
}

