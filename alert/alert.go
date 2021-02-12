package alert

type Alert interface{
    // Init returns an error if the confiugration is not valid
    Init() error

    // Send sends a message
    Send(message string) error
}
