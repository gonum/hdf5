package hdf5

import (
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	f, err := CreateFile(FNAME, F_ACC_TRUNC)
	if err != nil {
		t.Fatalf("CreateFile failed: %s", err)
	}
	defer os.Remove(FNAME)
	defer f.Close()

	if fileName := f.FileName(); fileName != FNAME {
		t.Fatalf("FileName() have %v, want %v", fileName, FNAME)
	}

	// The file is also the root group
	if name := f.Name(); name != "/" {
		t.Fatalf("Name() have %v, want %v", name, FNAME)
	}

	if err := f.Flush(F_SCOPE_GLOBAL); err != nil {
		t.Fatalf("Flush() failed: %s", err)
	}

	if !IsHdf5(FNAME) {
		t.Fatalf("IsHdf5 returned false")
	}

	groupName := "test"
	g, err := f.CreateGroup(groupName)
	if err != nil {
		t.Fatalf("CreateGroup() failed: %s", err)
	}
	if name := g.Name(); name != "/"+groupName {
		t.Fatalf("Group Name() have %v, want %v", name, groupName)
	}

	g2, err := f.OpenGroup(groupName)
	if err != nil {
		t.Fatalf("OpenGroup() failed: %s", err)
	}
	if name := g2.Name(); name != "/"+groupName {
		t.Fatalf("Group Name() have %v, want %v", name, groupName)
	}
}
