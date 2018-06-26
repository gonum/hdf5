// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cmem provides helper functionality for accessing memory in a C
// compatible fashion.
package cmem

// CMarshaler is an interface for types that can marshal their data into a
// C compatible binary layout.
type CMarshaler interface {
	MarshalC() ([]byte, error)
}
