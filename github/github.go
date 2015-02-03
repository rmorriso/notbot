package github

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ant0ine/go-json-rest/rest"
)

type Push struct {
	Compare    string     `json:"compare"`
	Commits    []Commit   `json:"commits"`
	Repository Repository `json:"repository"`
}

func (push *Push) Notification(req *rest.Request) string {
	err := req.DecodeJsonPayload(push)
	pstring, err := json.Marshal(push)
	if err != nil {
		return fmt.Sprintf("GitHub Notification Error %s", err)
	}
	log.Printf("Post: %s\n", pstring)

	var messages = ""
	for _, c := range push.Commits {
		messages = fmt.Sprintf("%s | %s", messages, c.Message)
	}
	return fmt.Sprintf("GitHub: %s | %s [ %s ]", push.Repository.String(), push.Compare, messages)
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
	MasterBranch string `json:"master_branch"`
}

func (r *Repository) String() string {
	return fmt.Sprintf("%s (%s)", r.Name, r.MasterBranch)
}

/*
func PostGitHub(w rest.ResponseWriter, req *rest.Request) {
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
*/
