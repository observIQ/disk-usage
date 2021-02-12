package alert

import (
    "github.com/BlueMedoraPublic/disk-usage/internal/alert/slack"
    "github.com/BlueMedoraPublic/disk-usage/internal/alert/stdout"
)

// Alert is an interface that sends alerts
type Alert interface{
    // Init returns an error if the confiugration is not valid
    Init() error

    // Send sends a message
    Send(message string) error
}

// Slack will return a new Slack type as an Alert interface
func Slack(channel, url string) (Alert, error) {
    a := slack.Slack{
        Channel: channel,
        HookURL: url,
    }
    return a, a.Init()
}

// Stdout will return a new Stdout type as an Alert interface
func Stdout() (Alert, error) {
    a := stdout.Stdout{}
    return a, a.Init()
}
