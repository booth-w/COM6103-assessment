package charts

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"

	"COM6103-assessment/lichess"
)

func RenderAll(lichessData, mastersData lichess.LichessData) {
	page := components.NewPage()
	page.AddCustomizedCSSAssets("style.css")

	chartFunctions := []func(data lichess.LichessData) *charts.Bar{
		generateWinRateChart,
		generatePopularityChart,
	}

	for _, chart := range chartFunctions {
		page.AddCharts(
			chart(lichessData),
			chart(mastersData),
		)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page.Render(w)
	})

	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "charts/style.css")
	})

	http.ListenAndServe(":8080", nil)
}
