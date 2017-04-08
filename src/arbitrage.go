// Package arbitrage provides an oppurtunity to exploit the foreign exchange markets.
package arbitrage

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// URLs used to receive information on currency and exchange rates.
const (
	ConvertExchangeRatesURL = "https://openexchangerates.org/api/convert/"
	CurrenciesURL           = "https://openexchangerates.org/api/currencies.json"
	LatestExchangeRatesURL  = "https://openexchangerates.org/api/latest.json"
)

var openExchangeRatesKey = keys()[0]

// CurrencyCode returns the 3 letter currency code for a given country name.
func CurrencyCode(name string) string { return name }

// CurrencyName returns the country name for a given 3 letter currency code.
func CurrencyName(code string) string {
	code = strings.ToUpper(code)

	response, err := http.Get(CurrenciesURL)
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err.Error())
	}

	return code
}

func keys() []string {
	contents, err := ioutil.ReadFile(".keys/keys")
	if err != nil {
		panic(err.Error())
	}
	return strings.Split(string(contents), "\n")
}
