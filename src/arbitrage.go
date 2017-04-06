package arbitrage

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	ConvertExchangeRatesURL = "https://openexchangerates.org/api/convert/"
	CurrenciesURL           = "https://openexchangerates.org/api/currencies.json"
	LatestExchangeRatesURL  = "https://openexchangerates.org/api/latest.json"
)

var APIKey string = getAPIKey()

func GetCurrencyCode(name string) string { return name }

func GetCurrencyName(code string) string {
	code = strings.ToUpper(code)

	response, err := http.Get("https://openexchangerates.org/api/currencies.json")
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

func getAPIKey() string {
	contents, err := ioutil.ReadFile(".keys/keys")
	if err != nil {
		panic(err.Error())
	}
	return string(contents)
}
