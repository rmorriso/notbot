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

func (p *Push) String() string {
	var messages = ""
	for _, c := range p.Commits {
		messages = fmt.Sprintf("%s | %s", messages, c.Message)
	}
	return fmt.Sprintf("%s | %s %s", p.Repository.String(), p.Compare, messages)
}

type Commit struct {
	ID       string   `json:"id"`
	Message  string   `json:"message"`
	Added    []string `json:"added"`
	Removed  []string `json:"removed"`
	Modified []string `json:"modified"`
}

func (c *Commit) String() string {
	return c.Message
}

type Committer struct {
	Name  string `json:"name"`
	Email string `json:email"`
}

func (c *Committer) String() string {
	return c.Name
}

type Repository struct {
	Name         string `json:"name"`
	URL          string `json:"url"`
	MasterBranch string `jsona:"master_branch"`
}

func (r *Repository) String() string {
	return fmt.Sprintf("%s (%s)", r.Name, r.MasterBranch)
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
	notice := fmt.Sprintf("NOTICE #easyrtc : GitHub: %s\n", push)
	conn.Raw(notice)
}
