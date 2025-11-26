package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type message struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/hello",func(w http.ResponseWriter, r *http.Request) {
		if r.Method !=http.MethodGet{
			http.Error(w,"only GET allowed",http.StatusMethodNotAllowed)
			return
		}
		fmt.Fprintln(w,"Hello! this is a GET response.")
	})
	http.HandleFunc("/submit",func(w http.ResponseWriter, r *http.Request) {
		if r.Method!=http.MethodPost{
			http.Error(w,"only POST allowed",http.StatusMethodNotAllowed)
			return
		}
		var msg message
		err:=json.NewDecoder(r.Body).Decode(&msg)
		if err!=nil{
			http.Error(w,"invalid JSON",http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w,"Received: %s\n",msg.Text)
	})
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080",nil)
}