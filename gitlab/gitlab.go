package gitlab

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

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

	var commits = ""
	for _, c := range push.Commits {
		commits = fmt.Sprintf("%s | %s (%s)", commits, strings.TrimSpace(c.Message), c.URL)
	}
	return fmt.Sprintf("GitLab: %s (%s) [ %s ]", push.Repository.Name, push.Ref, commits)
}

type Push struct {
	Before     string     `json:"before"`
	After      string     `json:"after"`
	Ref        string     `json:"ref"`
	UserID     int        `json:"user_id"`
	Username   string     `json:"user_name"`
	ProjectID  int        `json:"project_id"`
	Commits    []Commit   `json:"commits"`
	Repository Repository `json:"repository"`
}

type Commit struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	URL       string `json:"url"`
	Author    Author `json:"author"`
}

func (c *Commit) String() string {
	return c.Message
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (a *Author) String() string {
	return a.Name
}

type Repository struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
}

func (r *Repository) String() string {
	return r.Name
}
