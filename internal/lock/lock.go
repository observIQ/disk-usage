package lock

import (
    "fmt"

    "github.com/BlueMedoraPublic/disk-usage/internal/lock/file"
    "github.com/BlueMedoraPublic/disk-usage/internal/lock/null"
)

type Lock interface{
    Lock() error
    Unlock() error
    Exists() bool
    Path() string
}

func File(path string) (Lock, error) {
    if path == "" {
        return nil, fmt.Errorf("file lock path not set")
    }
    return file.New(path), nil
}

func Null() Lock {
    return null.Null{}
}
