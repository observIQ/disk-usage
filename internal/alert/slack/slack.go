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

    u, err := url.Parse(a.HookURL)
    if err != nil {
        return errors.Wrap(err, "slack webhook url failed to parse")
    }

    if err := validateScheme(u.Scheme); err != nil {
        return err
    }

    return nil
}

func validateScheme(s string) error {
    for _, scheme := range []string{"http","https"} {
        if s == scheme {
            return nil
        }
    }
    return fmt.Errorf(fmt.Sprintf("slack webhook url should use scheme 'http' or 'https', got %s", s))
}
