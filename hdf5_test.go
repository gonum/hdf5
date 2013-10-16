package hdf5

import "testing"

func TestLibVersion(t *testing.T) {
	v, err := LibVersion()
	if err != nil {
		t.Fatalf("Could not get HDF5 library version: %s", err)
	}
	if v.Major < 1 || (v.Major == 1 && v.Minor < 8) {
		t.Fatalf("go-hdf5 requires HDF5 > 1.8.0, detected %s", v)
	}
}
