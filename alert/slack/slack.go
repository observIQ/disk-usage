package slack

import (
    "fmt"
    "net/url"

    "github.com/slack-go/slack"
    "github.com/pkg/errors"
)

type Slack struct {
    Channel string
    HookURL string
}

func (a Slack) Send(message string) error {
    s := slack.WebhookMessage{
        Channel: a.Channel,
        Text:    message,
    }
    return slack.PostWebhook(a.HookURL, &s)
}

func (a Slack) Init() error {
    if a.HookURL == "" {
        return fmt.Errorf("slack webhook url is not set")
    }

    if _, err := url.Parse(a.HookURL); err != nil {
        return errors.Wrap(err, "slack webhook url failed to parse")
    }

    return nil
}
