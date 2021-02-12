package slack

import (
    "github.com/slack-go/slack"
)

type Alert struct {
    Message string
    Channel string
    URL string
}

func (a Alert) Send() error {
    s := slack.WebhookMessage{
        Channel: a.Channel,
        Text:    a.Message,
    }
    return slack.PostWebhook(a.URL, &s)
}
