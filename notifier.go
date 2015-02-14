package main

import "github.com/ant0ine/go-json-rest/rest"

type Notifier interface {
	Notifications(*rest.Request) []string
}
