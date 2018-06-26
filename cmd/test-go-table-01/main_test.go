// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

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
