package arbitrage

import (
	"io/ioutil"
	"net/http"
	"strings"
)

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

func main() {
	GetCurrencyName("usd")
}
