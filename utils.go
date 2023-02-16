package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"time"
)

type Column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Datatable struct {
	Data    [][]interface{} `json:"data"`
	Columns []Column        `json:"columns"`
}

type JSONData struct {
	Datatable Datatable `json:"datatable"`
	Meta      struct {
		NextCursorID interface{} `json:"next_cursor_id"`
	} `json:"meta"`
}

func CleanInput(args []string) (string, time.Time, time.Time, error) {
	// Take a slice of string arguments and return a validated symbol string,
	// start date, end date, and an error.
	// Parameters:
	//   - args (]string): A slice of string arguments containing the symbol, start date, and end date.
	//
	// Returns:
	//   - symbol (string): The stock symbol.
	//   - startDate (time.Time): The parsed start date.
	//   - endDate (time.Time): The parsed end date.
	//   - error: An error indicating any validation errors or other issues that may have occurred.
	config := GetConfig()

	if len(args) != 3 {
		return "", time.Time{}, time.Time{}, errors.New("incorrect number of arguments. 3 arguments are required: symbol, startDate, endDate")
	}

	symbol, startDateString, endDateString := args[0], args[1], args[2]

	if !validateSymbol(symbol) {
		return "", time.Time{}, time.Time{}, errors.New("invalid symbol")
	}

	startDate, err := time.Parse(config.Core.DateInputLayout, startDateString)
	if err != nil {
		return "", time.Time{}, time.Time{}, fmt.Errorf("invalid start date. Dates should be in this layout %s", config.Core.DateInputLayout)
	}

	endDate, err := time.Parse(config.Core.DateInputLayout, endDateString)
	if err != nil {
		return "", time.Time{}, time.Time{}, fmt.Errorf("invalid end date. Dates should be in this layout %s", config.Core.DateInputLayout)
	}

	if endDate.Before(startDate) {
		return "", time.Time{}, time.Time{}, errors.New("end date cannot be before start date")
	}

	return symbol, startDate, endDate, nil
}

func DisplayResults(
	symbol string,
	startDate time.Time,
	endDate time.Time,
	simpleRetrunValue float64,
	maxDrawdownValue float64,
	){
		config := GetConfig()
		resultMessage := fmt.Sprintf(
			"Calculated result for %s from %s to %s:\n",
			symbol,
			startDate.Format(config.Api.Nasdaq.DateFormat),
			endDate.Format(config.Api.Nasdaq.DateFormat),
		)
		resultMessage += fmt.Sprintf("Simple return: %.2f%%\n", (simpleRetrunValue*100))
		resultMessage += fmt.Sprintf("Maximum drawdown: %.2f%%\n", (maxDrawdownValue*100))

		fmt.Println(resultMessage)
		PostMessageToTelegramChannel(resultMessage)
}

func ParseJSONData(jsonData []byte) JSONData {
	var data JSONData
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		panic(err)
	}

	return data
}

func SortChronologically(arr [][]interface{}, dateIndex int){
	config := GetConfig()
	sort.Slice(arr, func(i, j int) bool {
		currDate, _ := time.Parse(config.Core.DateInputLayout, arr[i][dateIndex].(string))
		nextDate, _ := time.Parse(config.Core.DateInputLayout, arr[j][dateIndex].(string))
		return currDate.Before(nextDate)
	})
}

func validateSymbol(symbol string) bool {
	valid := regexp.MustCompile(`^[A-Z.]{1,6}$`)
	return valid.MatchString(symbol)
}
