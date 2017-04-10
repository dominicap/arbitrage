package arbitrage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"
	"sort"
	"strings"
)

// Constant Open Exchange Rates base URLs for receiving latest exchange rates.
const (
	ConvertExchangeRatesURL = "https://openexchangerates.org/api/convert/"
	CurrenciesURL           = "https://openexchangerates.org/api/currencies.json"
	LatestExchangeRatesURL  = "https://openexchangerates.org/api/latest.json"
)

// LatestExchangeRate is the struct that defines the data received from Open Exchange Rates.
type LatestExchangeRate struct {
	disclaimer string
	license    string
	timestamp  int
	base       string
	rates      map[string]float64
}

var openExchangeRatesKey = keys()[0]

var err error

// CurrencyCode returns the ISO code of a currency given the exact name.
func CurrencyCode(name string) (string, error) {
	name = strings.Title(name)
	for key, value := range CurrencyMap() {
		if name == value {
			return key, nil
		}
	}
	return "", errors.New("currency name not found")
}

// CurrencyName returns the exact name of a currency given the ISO Code.
func CurrencyName(code string) (string, error) {
	code = strings.ToUpper(code)
	if CurrencyMap()[code] != "" {
		return CurrencyMap()[code], nil
	}
	return "", errors.New("currency name not found")
}

// CurrencyMap returns a map of the name of the currency to its ISO code.
func CurrencyMap() map[string]string {
	response, err := http.Get(CurrenciesURL)
	check(err)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	check(err)

	var curMap map[string]string
	json.Unmarshal(body, &curMap)

	return curMap
}

// CreateTable makes a JSON graph of exchange rates and writes it to a file.
func CreateTable() {
	var codes []string
	for key := range CurrencyMap() {
		codes = append(codes, key)
	}
	sort.Strings(codes)

	file, err := os.Create("data/json/graph.json")
	check(err)
	defer file.Close()

	var table map[string]map[string]float64

	for _, code := range codes {
		url := LatestExchangeRatesURL + "?app_id=" + openExchangeRatesKey + "&base=" + code

		response, err := http.Get(url)
		check(err)
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		check(err)

		var data LatestExchangeRate
		json.Unmarshal(body, &data)

		table[data.base] = data.rates
	}

	data, err := json.Marshal(table)
	check(err)

	file.WriteString(string(data))
}

func keys() []string {
	_, fileName, _, _ := runtime.Caller(1)
	filePath := path.Join(path.Dir(fileName), "/.keys/keys")

	contents, err := ioutil.ReadFile(filePath)
	check(err)

	return strings.Split(string(contents), "\n")
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
