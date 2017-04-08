package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	convertExchangeRatesURL = "https://openexchangerates.org/api/convert/"
	currenciesURL           = "https://openexchangerates.org/api/currencies.json"
	latestExchangeRatesURL  = "https://openexchangerates.org/api/latest.json"
)

var openExchangeRatesKey = keys()[0]

func currencyCode(name string) string {
	name = strings.Title(name)
	for key, value := range currencyMap() {
		if name == value {
			return key
		}
	}
	panic("Currency name not found.")
}

func currencyName(code string) string {
	code = strings.ToUpper(code)
	if currencyMap()[code] != "" {
		return currencyMap()[code]
	}
	panic("Currency code not found")
}

func currencyMap() map[string]string {
	response, err := http.Get(currenciesURL)
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	var curMap map[string]string
	json.Unmarshal(body, &curMap)

	return curMap
}

func keys() []string {
	contents, err := ioutil.ReadFile(".keys/keys")
	if err != nil {
		panic(err.Error())
	}
	return strings.Split(string(contents), "\n")
}

func main() {

}
