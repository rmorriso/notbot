package gitlab

import (
	//	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
	//	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

type GitLab rest.Request

func (g *GitLab) Notification() string {
	push := &Push{}
	err := (*rest.Request)(g).DecodeJsonPayload(&push)
	pstring, err := json.Marshal(push)
	if err != nil {
		return fmt.Sprintf("GitLab Notification Error %s", err)
	}
	log.Printf("Post: %s\n", pstring)

	var messages = ""
	for _, c := range push.Commits {
		messages = fmt.Sprintf("%s | %s", messages, c.Message)
	}
	return fmt.Sprintf("%s | %s [ %s ]", push.Repository.String(), push.Compare, messages)
}

type Push struct {
	Commits    []Commit   `json:"commits"`
	Repository Repository `json:"repository"`
}

type Commit struct {
	ID       string   `json:"id"`
	Message  string   `json:"message"`
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
	Branch string `json:"branch"`
}

func (r *Repository) String() string {
	return fmt.Sprintf("%s (%s)", r.Name, r.Branch)
}
