// Package arbitrage creates an opportunity to exploit the foreign exchange markets.
package arbitrage

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Constant URLs that provide information on currency and exchange rates.
const (
	ConvertExchangeRatesURL = "https://openexchangerates.org/api/convert/"
	CurrenciesURL           = "https://openexchangerates.org/api/currencies.json"
	LatestExchangeRatesURL  = "https://openexchangerates.org/api/latest.json"
)

var openExchangeRatesKey = keys()[0]

// CurrencyCode returns the 3 letter currency code for a given currency name.
func CurrencyCode(name string) string {
	name = strings.Title(name)
	for key, value := range CurrencyMap() {
		if name == value {
			return key
		}
	}
	panic("Currency name not found.")
}

// CurrencyName returns the country name for a given 3 letter currency code.
func CurrencyName(code string) string {
	code = strings.ToUpper(code)
	if CurrencyMap()[code] != "" {
		return CurrencyMap()[code]
	}
	panic("Currency code not found")
}

// CurrencyMap returns a map that contains a 3 letter currency code and the corresponding currency name.
func CurrencyMap() map[string]string {
	response, err := http.Get(CurrenciesURL)
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
