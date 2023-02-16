package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Print disclaimer
	fmt.Println(`
		This program is using the Nasdaq API in "none-subscriber" mode,
		and as a result you can view a subset of the data.
		The sample includes data from 2017-09-01 to 2017-10-31 for the following tickers:
		MMM, AXP, AAPL, BA, CAT, CVX, CSCO, KO, DIS, XOM, GE, GS, HD, IBM, INTC, JNJ, JPM, MCD, MRK, MSFT, NKE, PFE, PG, TRV, UTX, UNH, VZ, V, WMT
	`)

	if err := run(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(args []string) error {
	symbol, startDate, endDate, err := CleanInput(args)

	if err != nil {
		return err
	}

	jsonData := FetchStockData(symbol, startDate, endDate)
	data := ParseJSONData(jsonData)

	simpleReturn, err := CalculateSimpleReturn(data)
	if err != nil {
		log.Println("simple return calculation was unsuccessful:", err)
	}

	maximumDrawdown, err := CalculateMaxDrawdown(data)
	if err != nil {
		log.Println("maximum drawdown calculation was unsuccessful:", err)
	}

	DisplayResults(symbol, startDate, endDate, simpleReturn, maximumDrawdown)
	
	return nil
}
