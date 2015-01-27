package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

type Gitlab struct {
	Username      string
	CommitMessage string
	URL           string
}

func PostGitlab(w rest.ResponseWriter, req *rest.Request) {
	content, err := ioutil.ReadAll(req.Body)
	req.Body.Close()
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(string(content))

	gitlab := Gitlab{
		Username: "Antoine",
	}
	w.WriteJson(&gitlab)
}

type Sensu struct {
	ID      string
	Message string
}

func PostSensu(w rest.ResponseWriter, req *rest.Request) {
	sensu := Sensu{
		ID: "Antoine",
	}
	w.WriteJson(&sensu)
}

func main() {
	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		&rest.Route{"POST", "/gitlab", PostGitlab},
		&rest.Route{"POST", "/sensu/alert", PostSensu},
	)
	http.ListenAndServe(":80", &handler)
}
