package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Student struct {
	Name string `json:"name"`
}

func GetName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application.json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("arman")
}

func serevr() {
	r := chi.NewRouter()
	r.Get("/", GetName)
	fmt.Println("gfdhfgh on : 8080")
	http.ListenAndServe(":8080", r)

}
