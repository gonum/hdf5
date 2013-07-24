go-hdf5
=======

This is a fork of sbinet's ``cgo`` bindings for the ``C-API`` of ``hdf5``.
The aim is to provide a more object-like API similar to the C++ and Java APIs of HDF5.

It is a work in progress and the API has not been finalized yet, so expect breakage.

Documentation
-------------

http://godoc.org/github.com/kisielk/go-hdf5/lib/

Example
-------

- Hello world example: https://github.com/kisielk/go-hdf5/blob/master/cmd/test-go-hdf5/test-go-hdf5.go

- Writing/reading an ``hdf5`` with compound data: https://github.com/kisielk/go-hdf5/blob/master/cmd/test-go-cpxcmpd/test-go-cpxcmpd.go

Note
----

- *Only* version *1.8.x* of ``HDF5`` is supported.


Known problems
--------------

- the ``h5pt`` packet table interface is broken.
