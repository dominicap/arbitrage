package arbitrage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
)

// DirectedEdge represents a directed edge from vertex to vertex along with a weight.
type DirectedEdge struct {
	V      int
	W      int
	Weight float64
}

func createTable() map[string]map[string]float64 {
	codes, _ := values()
	sort.Strings(codes)

	table := make(map[string]map[string]float64)

	for _, code := range codes {
		url := latestExchangeRatesURL + "?app_id=" + openExchangeRatesKey + "&base=" + code

		response, err := http.Get(url)
		check(err)
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		check(err)

		var data LatestExchangeData
		json.Unmarshal(body, &data)

		table[data.Base] = data.Rates
	}

	return table
}
