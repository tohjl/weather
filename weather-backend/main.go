package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"io/ioutil"
	"net/http"
)

// Struct to handle the API response (modify according to the API structure if needed)
type ForecastResponse struct {
	Data interface{} `json:"data"`
}

func weatherForecastHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response header to JSON
	w.Header().Set("Content-Type", "application/json")

	// Call the external API
	apiURL := "https://api.data.gov.my/weather/forecast"
	resp, err := http.Get(apiURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch weather forecast"})
		return
	}
	defer resp.Body.Close()

	// Check if the response status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Unexpected status code from API"})
		return
	}

	// Read the API response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to read API response"})
		return
	}

	// Forward the API response directly to the frontend
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func main() {
	// Handle the /api/weather endpoint
	http.HandleFunc("/api/weather", weatherForecastHandler)

	// Enable CORS for all origins
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(http.DefaultServeMux))
}
