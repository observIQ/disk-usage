package slack

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"
	"github.com/slack-go/slack"
)

// Slack type will send messages to slack
type Slack struct {
	Channel string
	HookURL string
}

// Send will send messages to slack
func (a Slack) Send(message string) error {
	s := slack.WebhookMessage{
		Channel: a.Channel,
		Text:    message,
	}
	return slack.PostWebhook(a.HookURL, &s)
}

// Init will configure slack
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
	for _, scheme := range []string{"http", "https"} {
		if s == scheme {
			return nil
		}
	}
	return fmt.Errorf(fmt.Sprintf("slack webhook url should use scheme 'http' or 'https', got %s", s))
}
