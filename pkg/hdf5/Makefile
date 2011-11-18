# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=bitbucket.org/binet/go-hdf5/pkg/hdf5

CGO_LDFLAGS=-lhdf5 -lhdf5_hl

CGOFILES=\
	hdf5.go\
    hdf5_hl.go\
    h5d.go\
    h5f.go\
    h5g.go\
    h5p.go\
    h5s.go\
    h5t.go\
    h5t_shim.go\
    h5pt.go\



include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
