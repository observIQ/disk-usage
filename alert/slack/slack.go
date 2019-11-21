package slack

import (
    "net/http"
    "strconv"
    "bytes"
    "github.com/pkg/errors"
)

type Alert struct {
    Message string
    Channel string
    URL string
}

func (a Alert) Send() error {
	var json []byte = []byte(
		`
		{
			"channel": "` + a.Channel + `",
			"text":"` + a.Message + `"
		}
		`)

	req, err := http.NewRequest("POST", a.URL, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

    if err != nil {
        return errors.Wrap(err, "failed to send slack alert!")
    }

    if resp.StatusCode != 200 {
        return errors.New("Slack returned status " + strconv.Itoa(resp.StatusCode))
    }

    return nil
}
