package disk

import (
	"encoding/json"
)

type message struct {
	Host     System `json:"host"`
	Message  string `json:"message"`
	Severity string `json:"severity"`
}

func (c Config) message(msg, sev string) error {
	m := message{
		Host:     c.Host,
		Message:  msg,
		Severity: sev,
	}

	b, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return err
	}

	return c.Alert.Send(string(b))
}
