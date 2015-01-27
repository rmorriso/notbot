package main

import (
	//	"io/ioutil"
	"encoding/json"
	"log"
	"net/http"

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
	/*	content, err := ioutil.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			rest.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println(string(content))
	*/

	push := &Push{}
	err := req.DecodeJsonPayload(&push)
	p, err := json.Marshal(push)
	if err != nil {
		log.Fatalf("Error %s\n", err)
	}

	log.Printf("Push: %s\n", p)

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteJson(&push)
}
