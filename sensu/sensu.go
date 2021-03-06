package sensu

import (
	"fmt"

	"github.com/ant0ine/go-json-rest/rest"
)

type Alert struct {
	ID      string
	Message string
}

func (alert *Alert) Notifications(req *rest.Request) []string {
	err := req.DecodeJsonPayload(alert)
	if err != nil {
		return []string{fmt.Sprintf("Sensu Notification Error: %s", err)}
	}
	return []string{fmt.Sprintf("Sensu: %s", alert.Message)}
}
