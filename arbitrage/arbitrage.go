package main

import (
	"fmt"

	"github.com/dominicap/arbitrage"
)

func main() {
	var value float64
	fmt.Printf("Enter the numerical value of your starting value: ")
	fmt.Scan(&value)

	var code string
	fmt.Printf("Enter the ISO code of your starting currency: ")
	fmt.Scan(&code)

	arbitrage.Arbitrage(value, code)
}
