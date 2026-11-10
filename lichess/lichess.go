package lichess

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type LichessData struct {
	White int `json:"white"`
	Black int `json:"black"`
	Draws int `json:"draws"`

	Moves []struct {
		White         int    `json:"white"`
		Black         int    `json:"black"`
		Draws         int    `json:"draws"`
		AverageRating int    `json:"averageRating"`
		San           string `json:"san"`
		Uci           string `json:"uci"`
		Opening       struct {
			Eco  string `json:"eco"`
			Name string `json:"name"`
		}
	} `json:"moves"`
}

func GetLichessData(API_KEY string) LichessData {
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
	var jsonData LichessData
	err = json.NewDecoder(res.Body).Decode(&jsonData)
	if err != nil {
		slog.Error("Error unmarshalling JSON", "error", err)
	}

	return jsonData
}

// Pretty print the JSON data
func PrettyPrint(data LichessData) {
	prettyData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		slog.Error("Error pretty printing JSON", "error", err)
	}
	fmt.Println(string(prettyData))
}
