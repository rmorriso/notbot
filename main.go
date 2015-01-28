package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
	irc "github.com/rmorriso/goirc/client"
/*
	"github.com/rmorriso/notbot/github"
	"github.com/rmorriso/notbot/gitlab"
	"github.com/rmorriso/notbot/sensu"
*/
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
		&rest.Route{"POST", "/github", Post},
		&rest.Route{"POST", "/gitlab", Post},
		&rest.Route{"POST", "/sensu", Post},
	)
	http.ListenAndServe(":80", &handler)
}

func Post(w rest.ResponseWriter, req *rest.Request) {
	path := req.URL.Path

	fmt.Println(path)
	return

/*
	push := &Push{}
	err := req.DecodeJsonPayload(&push)
	p, err := json.Marshal(push)
	if err != nil {
		log.Fatalf("Error %s\n", err)
	}
	log.Printf("Post: %s\n", p)
	ircNotify(push)
*/
}

func ircNotify(notice string) {
	conn.Raw(fmt.Sprintf("NOTICE #easyrtc : GitHub: %s\n", notice))
}
