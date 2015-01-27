package main

import (
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

