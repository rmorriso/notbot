package main

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

type Gitlab struct {
	Username      string
	CommitMessage string
	URL           string
}

func GitlabHandle(w rest.ResponseWriter, req *rest.Request) {
	gitlab := Gitlab{
		Username: "Antoine",
	}
	w.WriteJson(&gitlab)
}

type Sensu struct {
	ID      string
	Message string
}

func SensuHandle(w rest.ResponseWriter, req *rest.Request) {
	sensu := Sensu{
		ID: "Antoine",
	}
	w.WriteJson(&sensu)
}

func main() {
	handler := rest.ResourceHandler{}
	handler.SetRoutes(
		&rest.Route{"POST", "/gitlab/commit", GitlabHandle},
		&rest.Route{"POST", "/sensu/alert", SensuHandle},
	)
	http.ListenAndServe(":8080", &handler)
}
