package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/api/hello",func(w http.ResponseWriter, r *http.Request) {
		if r.Method!=http.MethodGet{
			http.Error(w,"only GET allowed",http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Content-Type","application/json")
		resp:=Message{Message: "Hello , wolrd!"}
		json.NewEncoder(w).Encode(resp)
	})
	http.ListenAndServe(":8080",nil)
}