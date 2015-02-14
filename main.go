package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/golang/glog"
	irc "github.com/rmorriso/goirc/client"
	"github.com/rmorriso/notbot/github"
	"github.com/rmorriso/notbot/gitlab"
	"github.com/rmorriso/notbot/sensu"
)

var (
	configFile string
	config     *Config
	conn       *irc.Conn
)

func init() {
	flag.StringVar(&configFile, "f", "/etc/notbot/notbot.yaml", "the notbot server config file")

}

func main() {
	flag.Parse()

	defer glog.Flush()

	// verify files exist
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		glog.Fatalf("notbot config file: %s\n", err)
	}

	args := flag.Args()
	glog.V(5).Infof("Args: %v\n", args)

	var err error
	config, err = Init(configFile)
	if err != nil {
		glog.Fatalf("Error in configuration: %s\n", err)
	}

	ircConfig := irc.NewConfig(config.Nick, config.Name)
	ircConfig.Pass = config.Password

	conn = irc.Client(ircConfig)
	conn.EnableStateTracking()

	conn.HandleFunc("connected",
		func(conn *irc.Conn, line *irc.Line) {
			channel := fmt.Sprintf("#%s", config.Channel)
			conn.Join(channel)
		})

	if err := conn.ConnectTo(config.Host); err != nil {
		log.Fatalf("Connection error: %s\n", err)
	}

	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		&rest.Route{"POST", "/github", Post},
		&rest.Route{"POST", "/gitlab", Post},
		&rest.Route{"POST", "/sensu", Post},
	)
	port := fmt.Sprintf(":%s", config.Port)
	if config.UseTLS {
		glog.Fatal(http.ListenAndServeTLS(port, config.CertFile, config.KeyFile, &handler))
	} else {
		glog.Fatal(http.ListenAndServe(port, &handler))
	}
}

func Post(w rest.ResponseWriter, req *rest.Request) {
	path := req.URL.Path

	log.Printf(path)
	var notifier Notifier
	switch path {
	case "/github":
		notifier = new(github.Push)
	case "/gitlab":
		notifier = new(gitlab.Push)
	case "/sensu":
		notifier = new(sensu.Alert)
	default:
		ircNotify(fmt.Sprintf("Invalid Notifier: %s", path))
		return
	}
	for _, notification := range notifier.Notifications(req) {
		ircNotify(notification)
		time.Sleep(2)
	}
	return
}

func ircNotify(notice string) {
	conn.Raw(fmt.Sprintf("NOTICE #easyrtc : %s\n", notice))
}
