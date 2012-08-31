go-hdf5
=======

Naive ``cgo`` bindings for the ``C-API`` of ``hdf5``.

Documentation
-------------

http://go.pkgdoc.org/github.com/sbinet/go-hdf5/pkg/hdf5

Example
-------

- Hello world example: https://github.com/sbinet/go-hdf5/blob/master/cmd/test-go-hdf5/test-go-hdf5.go

- Writing/reading an ``hdf5`` with compound data: https://github.com/sbinet/go-hdf5/blob/master/cmd/test-go-cpxcmpd/test-go-cpxcmpd.go

Note
----

- *Only* version *1.8.x* of ``HDF5`` is supported.


Known problems
--------------

- the ``h5pt`` packet table interface is broken.
