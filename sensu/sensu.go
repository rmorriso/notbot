package sensu

import (
	"github.com/ant0ine/go-json-rest/rest"
)

type Sensu struct {
	ID      string
	Message string
}

func Notification(req *rest.Request) string {
	return "sensu"
}

