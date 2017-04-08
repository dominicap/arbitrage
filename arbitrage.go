package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
)

const (
	convertExchangeRatesURL = "https://openexchangerates.org/api/convert/"
	currenciesURL           = "https://openexchangerates.org/api/currencies.json"
	latestExchangeRatesURL  = "https://openexchangerates.org/api/latest.json"
)

var openExchangeRatesKey = keys()[0]

var (
	code  string
	name  string
	value float64
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
	_, fileName, _, _ := runtime.Caller(1)
	filePath := path.Join(path.Dir(fileName), "/.keys/keys")

	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err.Error())
	}

	return strings.Split(string(contents), "\n")
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func usage() {
	fmt.Fprintf(os.Stderr, "OVERVIEW: An exploitation of the discrepancies in the foreign exchange markets\n\n")
	fmt.Fprintf(os.Stderr, "USAGE: arbitrage [value] [currency code]\n\n")
	fmt.Fprintf(os.Stderr, "OPTIONS:\n")
	flag.PrintDefaults()
}

var (
	helpFlag = flag.Bool("help", false, "Displays the help message.")
	versionFlag = flag.Bool("version", false, "Displays the version number.")
)

func init() {
	flag.Usage = usage()
	flag.Parse()

	if flag.NArg() == 1 {
		panic("must provide arguments")
	}
	value, err = strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		panic("monetary value must be a number")
	}
	if value == 0 {
		panic("value must be greater than 0")
	}
	value = toFixed(value, 2)

	code = os.Args[2]
	name, err = currencyName(code)
	if err != nil {
		panic(err.Error())
	}
}

func main() {

}
