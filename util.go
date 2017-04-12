package arbitrage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	currenciesURL = "https://openexchangerates.org/api/currencies.json"
	latestURL     = "http://api.fixer.io/latest"
)

// LatestExchangeData defines the data received from Open Exchange Rates
type LatestExchangeData struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

var err error

func currencyCode(name string) (string, error) {
	name = strings.Title(name)
	for key, value := range currencyMap() {
		if name == value {
			return key, nil
		}
	}
	return "", errors.New("currency name not found")
}

func currencyName(code string) (string, error) {
	code = strings.ToUpper(code)
	if currencyMap()[code] != "" {
		return currencyMap()[code], nil
	}
	return "", errors.New("currency name not found")
}

func currencyMap() map[string]string {
	response, err := http.Get(currenciesURL)
	check(err)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	check(err)

	curMap := make(map[string]string)
	json.Unmarshal(body, &curMap)

	response, err = http.Get(latestURL)
	check(err)
	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)
	check(err)

	var latest LatestExchangeData
	json.Unmarshal(body, &latest)

	relMap := make(map[string]string)
	for code, _ := range latest.Rates {
		relMap[code] = curMap[code]
	}

	return relMap
}

func values() ([]string, []string) {
	var codes, names []string
	for key, value := range currencyMap() {
		codes = append(codes, key)
		names = append(names, value)
	}
	return codes, names
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}
