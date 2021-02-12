package file

import (
    "os"
    "testing"
)

func TestLock(t *testing.T) {
    f := File{path}

    if err := f.Lock(); err != nil {
        t.Errorf(err.Error())
        return
    }

    if f.Exists() != true {
        t.Errorf("expected lock file " + f.Path() + " to exist, file not found")
    }

    if err := f.Unlock(); err != nil {
        t.Errorf(err.Error())
    }

    if f.Exists() != false {
        t.Errorf("expected lock file to be not found after RemoveLock(), however, it exists")

        if err := os.Remove(f.Path()); err != nil {
            t.Errorf(err.Error(), "Failed to remove lockfile with os.Remove, something is likely wrong outside of this code.")
            return
        }
    }
}

func TestPath(t *testing.T) {
    f := New(path)
    if f.Path() != path {
        t.Errorf("expected f.Path() to return " + path + ", got " + f.Path())
    }
}
