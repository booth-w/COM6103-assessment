package charts

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/components"

	"COM6103-assessment/lichess"
)

func RenderAll(data lichess.LichessData) {
	page := components.NewPage()
	page.AddCharts(
		generateWinRateChart(data),
		generatePopularityChart(data),
	)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page.Render(w)
	})

	http.ListenAndServe(":8080", nil)
}
