package main

import (
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

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
