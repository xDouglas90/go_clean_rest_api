package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xdouglas90/gomux-rest-api/models"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Up and running...")
}

func PostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := json.Marshal(models.Posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
