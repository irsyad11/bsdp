package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func hellWorld(w http.ResponseWriter, r *http.Request) {
	response := Response{Message: "Hello, Dunia !"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("Server running on port 4012")
	http.HandleFunc("/hello", hellWorld)
	http.ListenAndServe(":4012", nil)
}
