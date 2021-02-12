package slack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSend(t *testing.T) {
	a := Slack{
		Channel: "fake",
		HookURL: "https://fake33slack11url.com",
	}
	if err := a.Send("some message"); err == nil {
		t.Errorf("Expected an error when POSTing to " + a.HookURL)
	}
}

func TestInit(t *testing.T) {
	cases := []struct {
		name        string
		config      Slack
		expectedErr bool
	}{
		{
			name: "channel is not required",
			config: Slack{
				HookURL: "http://validurl.com",
			},
			expectedErr: false,
		},
		{
			name: "scheme is missing",
			config: Slack{
				HookURL: "validurl.com",
			},
			expectedErr: true,
		},
		{
			name: "invalid scheme",
			config: Slack{
				HookURL: "ftp://validurl.com",
			},
			expectedErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.config.Init()
			if tc.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
