package stdout

import (
    "fmt"
)

type Stdout struct {

}

func (s Stdout) Init() error {
    return nil
}

func (s Stdout) Send(m string) error {
    fmt.Println("m")
    return nil
}
