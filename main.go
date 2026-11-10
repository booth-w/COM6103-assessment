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

type LichessData struct {
	White int `json:"white"`
	Black int `json:"black"`
	Draws int `json:"draws"`

	Moves []struct {
		White         int    `json:"white"`
		Black         int    `json:"black"`
		Draws         int    `json:"draws"`
		AvarageRating int    `json:"averageRating"`
		San           string `json:"san"`
		Uci           string `json:"uci"`
		Opening       struct {
			Eco  string `json:"eco"`
			Name string `json:"name"`
		}
	} `json:"moves"`
}

var jsonData LichessData
var API_KEY string

func initEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "error", err)
	}
	API_KEY = os.Getenv("LICHESS_API_KEY")
}

func getLichessData() {
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
}

// Pretty print the JSON data
func prettyPrint(data LichessData) {
	prettyData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		slog.Error("Error pretty printing JSON", "error", err)
	}
	fmt.Println(string(prettyData))
}

func main() {
	initEnv()
	getLichessData()

	prettyPrint(jsonData)
}
