package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response struct for JSON output
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Handler is the entry point for Vercel Go functions
func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request:", r.Method, r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := Response{
		Status:  "OK",
		Message: "API is running smoothly.",
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("Error encoding JSON:", err)
	}
}

