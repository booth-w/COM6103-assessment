package charts

import (
	"slices"
	"sort"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"

	"COM6103-assessment/lichess"
)

func generateWinRateChart(data lichess.LichessData) *charts.Bar {
	bar := charts.NewBar()

	var sans []string
	var whiteRates []opts.BarData
	var blackRates []opts.BarData
	var drawRates []opts.BarData

	moves := slices.Clone(data.Moves)

	// Sort by white win rate descending
	sort.Slice(moves, func(i, j int) bool {
		totalI := float64(moves[i].White + moves[i].Black + moves[i].Draws)
		totalJ := float64(moves[j].White + moves[j].Black + moves[j].Draws)

		if totalI == 0 || totalJ == 0 {
			return false
		}

		return float64(moves[i].White)/totalI > float64(moves[j].White)/totalJ
	})

	for _, move := range moves {
		total := float64(move.White + move.Black + move.Draws)

		if total == 0 {
			continue
		}

		sans = append(sans, move.San)

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

	bar.SetXAxis(sans)

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

func generatePopularityChart(data lichess.LichessData) *charts.Bar {
	bar := charts.NewBar()

	var sans []string
	var whiteRates []opts.BarData

	moves := slices.Clone(data.Moves)

	// Sort by white popularity descending
	sort.Slice(moves, func(i, j int) bool {
		return moves[i].White > moves[j].White
	})

	for _, move := range moves {
		sans = append(sans, move.San)
		whiteRates = append(whiteRates, opts.BarData{
			Value: move.White,
		})
	}

	bar.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "1000px",
			Height: "600px",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Type: "value",
		}),
	)

	bar.SetXAxis(sans)

	bar.AddSeries("Popularity of white's first moves", whiteRates,
		charts.WithBarChartOpts(opts.BarChart{
			Stack: "total",
		}),
	)

	return bar
}
