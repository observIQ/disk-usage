package lockfile

import (
    "os"
    "testing"
)

func TestLock(t *testing.T) {
    if err := CreateLock(path); err != nil {
        t.Errorf(err.Error())
        return
    }

    if Exists(path) != true {
        t.Errorf("expected lock file " + path + " to exist, file not found")
    }

    if err := RemoveLock(path); err != nil {
        t.Errorf(err.Error())
    }

    if Exists(path) != false {
        t.Errorf("expected lock file to be not found after RemoveLock(), however, it exists")

        if err := os.Remove(path); err != nil {
            t.Errorf(err.Error(), "Failed to remove lockfile with os.Remove, something is likely wrong outside of this code.")
            return
        }
    }
}
