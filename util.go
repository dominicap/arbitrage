package arbitrage

import (
    "encoding/json"
    "errors"
	"io/ioutil"
    "net/http"
	"path"
	"runtime"
	"strings"
)

const (
	convertExchangeRatesURL = "https://openexchangerates.org/api/convert/"
	currenciesURL           = "https://openexchangerates.org/api/currencies.json"
	latestExchangeRatesURL  = "https://openexchangerates.org/api/latest.json"
)

type latestExchangeRate struct {
	disclaimer string
	license    string
	timestamp  int
	base       string
	rates      map[string]float64
}

var openExchangeRatesKey = keys()[0]

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

	var curMap map[string]string
	json.Unmarshal(body, &curMap)

	return curMap
}

func check(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func keys() []string {
	_, fileName, _, _ := runtime.Caller(1)
	filePath := path.Join(path.Dir(fileName), "/.keys/keys")

	contents, err := ioutil.ReadFile(filePath)
	check(err)

	return strings.Split(string(contents), "\n")
}
