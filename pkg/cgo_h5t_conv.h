#ifndef CGO_H5T_CONV_H
#define CGO_H5T_CONV_H 1

#include <stdint.h>
#include "hdf5.h"
#include "hdf5_hl.h"

#include <string.h>
//void *memcpy(void *dest, const void *src, size_t n);

#include <stdlib.h>
//void *calloc(size_t nmemb, size_t size);

#include <stdio.h>

extern void *H5I_object(hid_t id);

typedef struct { char *p; int n; } _my_go_GoString;
// GoString -> CString conversion function
herr_t
H5T_conv_gostring_cstring(hid_t src_id, hid_t dst_id, H5T_cdata_t *cdata, size_t nelmts, size_t buf_stride, size_t bkg_stride, void *buf, void* bkg, hid_t dxpl_id)
{
  unsigned char *src = NULL; // source datatype
  unsigned char *dst = NULL; // destination datatype
  size_t elmtno;             // element number
  size_t nchars = 0;         // nbr of characters copied
  uint8_t *s, *sp, *d, *dp;  // src and dst traversal pointers
  herr_t ret_value = 0;      // return value

  //_my_go_GoString p;

  fprintf(stderr, "--h5t-conv-reg--(%d -> %d)\n", (int)src_id, (int)dst_id);
 
 switch (cdata->command) {
  case H5T_CONV_INIT: {
    fprintf(stderr, "--h5t-conv-init--(%d -> %d)\n", (int)src_id, (int)dst_id);
    src = (unsigned char*)H5I_object(src_id);
    dst = (unsigned char*)H5I_object(dst_id);
    if (NULL == src || NULL == dst) {
      ret_value = -1;
      goto done;
      break;
    }
    cdata->need_bkg = H5T_BKG_NO;
    break;
  }

  case H5T_CONV_FREE: {
    fprintf(stderr, "--h5t-conv-free--(%d -> %d)\n", (int)src_id, (int)dst_id);
    break;
  }

  case H5T_CONV_CONV: {
    fprintf(stderr, "--h5t-conv-conv--(%d -> %d)\n", (int)src_id, (int)dst_id);
    // get the datatypes
    src = (unsigned char*)H5I_object(src_id);
    dst = (unsigned char*)H5I_object(dst_id);
    if (NULL == src || NULL == dst) {
      ret_value = -1;
      break;
    }
    sp = dp = (uint8_t*)buf;
    /* The conversion loop. */
    fprintf(stderr, "--h5t-conv-conv-- nelemts: %d\n", (int)nelmts);
    for (elmtno=0; elmtno<nelmts; elmtno++) {
      s = sp;
      d = dp;
      _my_go_GoString *go_src = (_my_go_GoString*)s;
      nchars = go_src->n;
      d = (unsigned char*)calloc(nchars+1, sizeof(unsigned char*)*nchars+1);
      memcpy(d, go_src->p, nchars);
      d[nchars] = 0;
    } /* conversion loop */

    break;
  }
  }
 done:

  return ret_value;
}

// herr_t H5Tregister(H5T_pers_t pers, const char *name, hid_t src_type, hid_t dest_type, H5T_conv_t func)
 
herr_t h5t_my_register(hid_t src_id, hid_t dst_id)
{
  //return -1;
  fprintf(stderr, "--> registering... %d -> %d\n", (int)src_id, (int)dst_id);
  return H5Tregister(H5T_PERS_SOFT, "gostring->cstring", src_id, dst_id, H5T_conv_gostring_cstring);
}

#endif /* !CGO_H5T_CONV_H */
