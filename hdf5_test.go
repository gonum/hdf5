package hdf5

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

func TestLibVersion(t *testing.T) {
	v, err := LibVersion()
	if err != nil {
		t.Fatalf("Could not get HDF5 library version: %s", err)
	}
	if v.Major < 1 || (v.Major == 1 && v.Minor < 8) {
		t.Fatalf("go-hdf5 requires HDF5 > 1.8.0, detected %s", v)
	}
}

func TestCpxCmpd(t *testing.T) {
	const fname = "SDScompound.h5"
	stdout := new(bytes.Buffer)
	cmd := exec.Command("test-go-cpxcmpd")
	cmd.Stdout = stdout
	cmd.Stderr = stdout
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		t.Fatalf("error: %v\n%s\n", err, string(stdout.Bytes()))
	}
	os.Remove(fname)
}

func TestTableRWCmd(t *testing.T) {
	const fname = "ex_table_01.h5"
	func() {
		stdout := new(bytes.Buffer)
		cmd := exec.Command("test-go-table-01")
		cmd.Stdout = stdout
		cmd.Stderr = stdout
		cmd.Stdin = os.Stdin

		err := cmd.Run()
		if err != nil {
			t.Fatalf("error: %v\n%s\n", err, string(stdout.Bytes()))
		}
	}()

	func() {
		stdout := new(bytes.Buffer)
		cmd := exec.Command("test-go-table-01-readback")
		cmd.Stdout = stdout
		cmd.Stderr = stdout
		cmd.Stdin = os.Stdin

		err := cmd.Run()
		if err != nil {
			t.Fatalf("error: %v\n%s\n", err, string(stdout.Bytes()))
		}
	}()
	os.Remove(fname)
}
