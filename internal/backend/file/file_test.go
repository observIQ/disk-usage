package file

import (
	"testing"
)

func TestPath(t *testing.T) {
	f := New(path)
	if f.Path() != path {
		t.Errorf("expected f.Path() to return " + path + ", got " + f.Path())
	}
}
