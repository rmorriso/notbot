package main

import (
	//	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
	//	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

type Push struct {
	Compare    string     `json:"compare"`
	Commits    []Commit   `json:"commits"`
	Repository Repository `json:"repository"`
}

type Commit struct {
	ID       string   `json:"id"`
	Message  string   `json:"message"`
	Added    []string `json:"added"`
	Removed  []string `json:"removed"`
	Modified []string `json:"modified"`
}

type Committer struct {
	Name  string `json:"name"`
	Email string `json:email"`
}

type Repository struct {
	URL string `json:"url"`
}

func PostGitlab(w rest.ResponseWriter, req *rest.Request) {
	push := &Push{}
	err := req.DecodeJsonPayload(&push)
	p, err := json.Marshal(push)
	if err != nil {
		log.Fatalf("Error %s\n", err)
	}
	log.Printf("Post: %s\n", p)
	ircNotify(push)
}

func ircNotify(push *Push) {
	notice := fmt.Sprintf("NOTICE #easyrtc : GitHub (%s): %v\n", push.Repository.URL, push.Commits)
	conn.Raw(notice)
}
