#include "hdf5.h"
#include <stdlib.h>
#include <string.h>

herr_t
_go_hdf5_unsilence_errors(void)
{
  return H5Eset_auto(H5E_DEFAULT, (H5E_auto2_t)(H5Eprint), stderr);
}

herr_t
_go_hdf5_silence_errors(void)
{
  return H5Eset_auto(H5E_DEFAULT, NULL, NULL);
}

/* EOF */

