package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/xdouglas90/gomux-rest-api/controllers"
)

func main() {
	r := mux.NewRouter()
	p := ":8080"
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/", controllers.HomeHandler)
	log.Println("Listening on port " + p)
	log.Fatalln(http.ListenAndServe(p, s))
	http.ListenAndServe(p, s)
}
