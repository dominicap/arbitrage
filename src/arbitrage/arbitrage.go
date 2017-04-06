package arbitrage

import (
	"encoding/json"
	"net/http"
)

func GetFullCurrencyName(identifier string) string {
	response, err := http.Get("https://openexchangerates.org/api/currencies.json")
	if err != nil {
		panic(err)
	}
    defer response.Body.Close()

}

func main() {

}
