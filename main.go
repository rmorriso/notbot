package main

import (
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	irc "github.com/rmorriso/goirc/client"
)

var (
	host    = "irc.priologic.com"
	channel = "#easyrtc"
	conn    *irc.Conn
)

func init() {
	config := irc.NewConfig("notbot", "notbot")
	config.Pass = "k@b00dle"

	conn = irc.Client(config)
	conn.EnableStateTracking()

	conn.HandleFunc("connected",
		func(conn *irc.Conn, line *irc.Line) {
			conn.Join(channel)
		})
}

func main() {
	if err := conn.ConnectTo(host); err != nil {
		log.Fatalf("Connection error: %s\n", err)
	}

	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		&rest.Route{"POST", "/gitlab", PostGitlab},
		&rest.Route{"POST", "/sensu", PostSensu},
	)
	http.ListenAndServe(":80", &handler)
}
