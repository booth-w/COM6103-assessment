package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var jsonData map[string]any

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "error", err)
	}
	API_KEY := os.Getenv("LICHESS_API_KEY")

	url := "https://explorer.lichess.org/lichess"

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("Error creating request", "error", err)
	}

	// Add auth header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", API_KEY))

	// Send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("Error sending request", "req", req, "error", err)
	}
	defer res.Body.Close()

	// Read and parse the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Error reading response body", "error", err)
	}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		slog.Error("Error unmarshalling JSON", "error", err)
	}

	// Remove top and recent games
	delete(jsonData, "topGames")
	delete(jsonData, "recentGames")

	// Pretty print the JSON data
	prettyData, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(prettyData))
}
