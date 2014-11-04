package main_test

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

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
