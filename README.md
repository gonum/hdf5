go-hdf5
=======
[![Build Status](https://travis-ci.org/sbinet/go-hdf5.png?branch=master)](https://travis-ci.org/binet/go-hdf5)

This is a fork of sbinet's ``cgo`` bindings for the ``C-API`` of ``hdf5``.
The aim is to provide a more object-like API similar to the C++ and Java APIs of HDF5.

It is a work in progress and the API has not been finalized yet, so expect breakage.

Documentation
-------------

http://godoc.org/github.com/sbinet/go-hdf5

Example
-------

- Hello world example: https://github.com/sbinet/go-hdf5/blob/master/cmd/test-go-hdf5/test-go-hdf5.go

- Writing/reading an ``hdf5`` with compound data: https://github.com/sbinet/go-hdf5/blob/master/cmd/test-go-cpxcmpd/test-go-cpxcmpd.go

Note
----

- *Only* version *1.8.x* of ``HDF5`` is supported.
- In order to use ``HDF5`` functions in more than one goroutine simultaneously, you must build the HDF5 library with threading support. Many binary distributions (RHEL/centos/Fedora packages, etc.) do not have this enabled. Therefore, you must be HDF5 yourself on these systems.


Known problems
--------------

- the ``h5pt`` packet table interface is broken.
