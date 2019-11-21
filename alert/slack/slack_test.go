package slack

import (
    "testing"
)

func TestSend(t *testing.T) {
    a := Alert{
        Message: "fake",
        Channel: "fake",
        URL: "https://fake33slack11url.com",
    }
    if err := a.Send(); err == nil {
        t.Errorf("Expected an error when POSTing to " + a.URL)
    }
}
