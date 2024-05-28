package router

import (
	"github.com/gorilla/mux"
	"github.com/skyakashh/UrlShortner/mongohelpers"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", mongohelpers.CreateMovie).Methods("POST")
	r.HandleFunc("/{shortURL}", mongohelpers.GetLongURL).Methods("GET")

	return r
}
