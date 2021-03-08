# hdf5

[![Build status](https://github.com/gonum/hdf5/workflows/CI/badge.svg)](https://github.com/gonum/hdf5/actions)
[![Build Status](https://secure.travis-ci.org/gonum/hdf5.png)](http://travis-ci.org/gonum/hdf5)
[![GoDoc](https://godoc.org/gonum.org/v1/hdf5?status.svg)](https://godoc.org/gonum.org/v1/hdf5)

Naive ``cgo`` bindings for the ``C-API`` of ``hdf5``.

**WIP: No stable API for this package yet.**

**NOTE** that starting with Go >= 1.6, one needs to run with `GODEBUG=cgocheck=0` to disable the new stricter `CGo` rules.

## Example

- Hello world example: [cmd/test-go-hdf5/main.go](https://github.com/gonum/hdf5/blob/master/cmd/test-go-hdf5/main.go)

- Writing/reading an ``hdf5`` with compound data: [cmd/test-go-cpxcmpd/main.go](https://github.com/gonum/hdf5/blob/master/cmd/test-go-cpxcmpd/main.go)

## Note

- *Only* version *1.8.x* of ``HDF5`` is supported.
- In order to use ``HDF5`` functions in more than one goroutine simultaneously, you must build the HDF5 library with threading support. Many binary distributions (RHEL/centos/Fedora packages, etc.) do not have this enabled. Therefore, you must build HDF5 yourself on these systems.


## Known problems

- the ``h5pt`` packet table interface is broken.
- support for structs with slices and strings as fields is broken

## License

Please see github.com/gonum/gonum for general license information, contributors, authors, etc on the Gonum suite of packages.
