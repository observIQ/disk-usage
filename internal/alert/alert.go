package alert

import (
    "github.com/BlueMedoraPublic/disk-usage/internal/alert/slack"
    "github.com/BlueMedoraPublic/disk-usage/internal/alert/stdout"
)

type Alert interface{
    // Init returns an error if the confiugration is not valid
    Init() error

    // Send sends a message
    Send(message string) error
}

func Slack(channel, url string) (Alert, error) {
    a := slack.Slack{
        Channel: channel,
        HookURL: url,
    }
    return a, a.Init()
}

func Stdout() (Alert, error) {
    a := stdout.Stdout{}
    return a, a.Init()
}
