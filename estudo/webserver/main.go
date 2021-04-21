package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8888", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	user := User{1, "fulano"}

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Panicln(err)
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(userJson)
}
