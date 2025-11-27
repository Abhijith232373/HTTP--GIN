package main

import (
	"encoding/json"
	"net/http"
)


type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		var user User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "success",
			"user":   user,
		})
	})

	http.ListenAndServe(":8080", nil)
}
