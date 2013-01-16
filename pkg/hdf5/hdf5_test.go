package hdf5

import "testing"

func TestLibVersion(t *testing.T) {
	m, n, r, err := GetLibVersion()
	if err != nil {
		t.Fatalf("Could not get HDF5 library version: %s", err)
	}
	if m < 1 || (m == 1 && n < 8) {
		t.Fatalf("go-hdf5 requires HDF5 > 1.8.0, detected %d.%d.%d", m, n, r)
	}
}
