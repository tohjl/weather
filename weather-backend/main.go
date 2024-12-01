package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response header to JSON
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Hello, World!",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Handle the /api/hello endpoint
	http.HandleFunc("/api/hello", helloWorldHandler)

	// Enable CORS for all origins (or specify allowed origins if needed)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(http.DefaultServeMux))
}
