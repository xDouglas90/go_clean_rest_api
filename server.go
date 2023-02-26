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
	s := r.PathPrefix("/api/v1").Subrouter()
	s.HandleFunc("/", controllers.HomeHandler)
	s.HandleFunc("/posts", controllers.PostsHandler).Methods("GET")
	log.Println("Listening on port " + p)
	log.Fatalln(http.ListenAndServe(p, s))
	http.ListenAndServe(p, s)
}
