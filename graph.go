package arbitrage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
)

func createTable() {
	var codes []string
	for key := range currencyMap() {
		codes = append(codes, key)
	}
	sort.Strings(codes)

	file, err := os.Create("/var/temp/graph.json")
	check(err)
	defer file.Close()

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

	data, err := json.Marshal(table)
	check(err)

	file.WriteString(string(data))
}
