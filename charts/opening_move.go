package charts

import (
	"net/http"
	"sort"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"

	"COM6103-assessment/lichess"
)

func generateChart(data lichess.LichessData) *charts.Bar {
	bar := charts.NewBar()

	var moves []string
	var whiteRates []opts.BarData
	var blackRates []opts.BarData
	var drawRates []opts.BarData

	// Sort by white win rate descending
	sort.Slice(data.Moves, func(i, j int) bool {
		totalI := float64(data.Moves[i].White + data.Moves[i].Black + data.Moves[i].Draws)
		totalJ := float64(data.Moves[j].White + data.Moves[j].Black + data.Moves[j].Draws)

		if totalI == 0 || totalJ == 0 {
			return false
		}

		return float64(data.Moves[i].White)/totalI > float64(data.Moves[j].White)/totalJ
	})

	for _, move := range data.Moves {
		total := float64(move.White + move.Black + move.Draws)

		if total == 0 {
			continue
		}

		moves = append(moves, move.San)

		whiteRates = append(whiteRates, opts.BarData{
			Value: float64(move.White) / total * 100,
		})

		blackRates = append(blackRates, opts.BarData{
			Value: float64(move.Black) / total * 100,
		})

		drawRates = append(drawRates, opts.BarData{
			Value: float64(move.Draws) / total * 100,
		})
	}

	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1000px",
			Height: "600px",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Type: "value",
			Max:  100,
		}),
	)

	bar.SetXAxis(moves)

	bar.AddSeries("White wins", whiteRates,
		charts.WithBarChartOpts(opts.BarChart{
			Stack: "total",
		}),
	)

	bar.AddSeries("Black wins", blackRates,
		charts.WithBarChartOpts(opts.BarChart{
			Stack: "total",
		}),
	)

	bar.AddSeries("Draws", drawRates,
		charts.WithBarChartOpts(opts.BarChart{
			Stack: "total",
		}),
	)

	return bar
}

func RenderChart(data lichess.LichessData) {
	bar := generateChart(data)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bar.Render(w)
	})

	http.ListenAndServe(":8080", nil)
}
