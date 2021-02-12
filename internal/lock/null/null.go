package null

type Null struct {

}

func (n Null) Lock() error {
    return nil
}

func (n Null) Unlock() error {
    return nil
}

func (n Null) Exists() bool {
    return false
}

func (n Null) Path() string {
    return ""
}
