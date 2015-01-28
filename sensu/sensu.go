package sensu

import (
	"fmt"

	"github.com/ant0ine/go-json-rest/rest"
)

type Sensu rest.Request

func (s *Sensu) Notification() string {
	alert := &Alert{}
	err := (*rest.Request)(s).DecodeJsonPayload(&alert)
	if err != nil {
		return fmt.Sprintf("Sensu Notification Error: %s", err)
	}
	return "sensu"
}

type Alert struct {
	ID      string
	Message string
}

