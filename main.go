package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"

	"COM6103-assessment/charts"
	"COM6103-assessment/lichess"
)

var API_KEY string

func initEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "error", err)
	}
	API_KEY = os.Getenv("LICHESS_API_KEY")
}

func main() {
	initEnv()
	lichessData := lichess.GetLichessData(API_KEY, "lichess")
	mastersData := lichess.GetLichessData(API_KEY, "masters")

	charts.RenderAll(lichessData, mastersData)
}
