package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

func main(){
	if err := run(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(args []string) error {
	const dateLayout = "2006-Jan-02"

	if len(args) != 3 {
		return fmt.Errorf("incorrect number of arguments. 3 arguments are required: symbol, startDate, endDate")
	}

	symbol, startDate, endDate := args[0], args[1], args[2]

	if !validateSymbol(symbol){
		return fmt.Errorf("invalid symbol")
	}

	start, err := time.Parse(dateLayout, startDate)
	if err != nil {
		return fmt.Errorf("invalid start date. Dates should be in this layout %s", dateLayout)
	}

	end, err := time.Parse(dateLayout, endDate)
	if err != nil {
		return fmt.Errorf("invalid end date. Dates should be in this layout %s", dateLayout)
	}

	if end.Before(start) {
		return fmt.Errorf("end date cannot be before start date")
	}

	// fetch stock data
	// simpleReturn := calculateSimpleReturn()
	// maximumDrawdown := calculateMaximumDrawdown()
	// show calculated result

	return nil
}

func validateSymbol(symbol string) bool {
	valid := regexp.MustCompile(`^[A-Z]+$`)
	return valid.MatchString(symbol)
}

// func calculateSimpleReturn () int {}
// func calculateMaximumDrawdown () int {}
