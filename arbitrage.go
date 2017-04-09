package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"runtime"
	"sort"
	"strings"
)

const (
	convertExchangeRatesURL = "https://openexchangerates.org/api/convert/"
	currenciesURL           = "https://openexchangerates.org/api/currencies.json"
	latestExchangeRatesURL  = "https://openexchangerates.org/api/latest.json"
)

var openExchangeRatesKey = keys()[0]

var (
	code, name string
	value      float64
)

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

func createTable() {
	var codes []string
	for key := range currencyMap() {
		codes = append(codes, key)
	}
	sort.Strings(codes)

	file, err := os.Create("/var/tmp/table")
	check(err)
	defer file.Close()

	for _, code := range codes {
		_, err := file.WriteString(fmt.Sprintf("%10s", code))
		check(err)
	}

	file.WriteString("\n")
	for _, code := range codes {
		_, err := file.WriteString(code + "\n\n")
		check(err)
	}
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

func init() {
    var name string
}

func main() {

}
