package main

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"time"
)

func CalculateMaxDrawdown(data JSONData) (float64, error) {
	// Calculate maximum drawdown on the given dataset
	// Parameters:
	// 		data (JSONData): stock data over a certain period of time
	// Retruns:
	// 		maxDrawdown (float64): The calculated maximum drawdown
	// 		error (error): error if encountered any
	adjCloseIndex := findColIndex("adj_close", data.Datatable.Columns)
	dateIndex := findColIndex("date", data.Datatable.Columns)
	stockData := data.Datatable.Data

	err := validateStockData(adjCloseIndex, dateIndex, stockData)
	if err != nil {
		return 0, err
	}

	sortDataChronologically(stockData, dateIndex)

	// Flatten array
	adjClosePrices := make([]float64, len(stockData))
	for i, data := range stockData {
		adjClosePrices[i] = data[adjCloseIndex].(float64)
	}

	cumReturns := calculateCumulativeReturns(adjClosePrices)

	rollingMax := calculateRollingMax(cumReturns)

	drawdowns := calculateDrawdowns(cumReturns, rollingMax)

	maxDrawdown := drawdowns[0]
	for i := 1; i < len(drawdowns); i++ {
		if drawdowns[i] > maxDrawdown {
			maxDrawdown = drawdowns[i]
		}
	}

	truncatedMaxDrawdown := float64(int(maxDrawdown*10000)) / 10000.0

	return truncatedMaxDrawdown, nil
}

func CalculateSimpleReturn(data JSONData) (float64, error) {
	// Calculate simple return based on the given dataset.
	// Parameters:
	// 		data (JSONData): stock data over a certain period of time
	// Returns:
	// 		simpleReturn (float64): The calculated simple return
	// 		error (error): error if encountered any
	adjCloseIndex := findColIndex("adj_close", data.Datatable.Columns)
	dateIndex := findColIndex("date", data.Datatable.Columns)
	stockData := data.Datatable.Data

	err := validateStockData(adjCloseIndex, dateIndex, stockData)
	if err != nil {
		return 0, err
	}

	sortDataChronologically(data.Datatable.Data, dateIndex)

	startPrice := stockData[0][adjCloseIndex].(float64)
	endPrice := stockData[len(stockData)-1][adjCloseIndex].(float64)

	if startPrice == 0 {
		return 0, fmt.Errorf("start price is zero, abort calculation")
	}

	simpleReturn := (endPrice - startPrice) / startPrice

	truncatedSimpleReturn := float64(int(simpleReturn*10000)) / 10000.0

	return truncatedSimpleReturn, nil
}

func calculateCumulativeReturns(adjustedClose []float64) []float64 {
	// Calculate cumulative returns from a slice of adjusted close values
	cumReturns := make([]float64, len(adjustedClose))
	cumReturns[0] = 0.0
	for i := 1; i < len(adjustedClose); i++ {
		cumReturns[i] = (adjustedClose[i] / adjustedClose[0]) - 1
	}
	return cumReturns
}

func calculateDrawdowns(cumReturns, rollingMax []float64) []float64 {
	// Calculate drawdowns from a slice of cumulative returns and a slice of rolling maximum values
	drawdowns := make([]float64, len(cumReturns))
	for i := 0; i < len(cumReturns); i++ {
		drawdowns[i] = rollingMax[i] - cumReturns[i]
	}
	return drawdowns
}

func calculateRollingMax(numbers []float64) []float64 {
	// Calculate rolling maximum value of a slice of numbers
	rollingMax := make([]float64, len(numbers))
	rollingMax[0] = numbers[0]
	for i := 1; i < len(numbers); i++ {
		rollingMax[i] = math.Max(rollingMax[i-1], numbers[i])
	}
	return rollingMax
}

func validateStockData(adjCloseIndex int, dateIndex int, stockData [][]interface{}) error {
	if adjCloseIndex == -1 {
		return errors.New("adj_close column was not found")
	}

	if dateIndex == -1 {
		return errors.New("date column was not found")
	}

	if len(stockData) < 2 {
		return errors.New("not enough data were provided to perform calculation")
	}

	return nil
}

func findColIndex(colName string, columns []Column) int {
	// Find the index of the first column that matches the name with given colName
	// Patameters:
	// 		colName (string): The name of the column to match
	// 		columns ([]Columns): Array of columns to search in
	// Returns:
	// 		colIndex (int): Index of the matched column or -1 if no match was found
	colIndex := -1

	for i, col := range columns {
		if col.Name == colName {
			colIndex = i
			break
		}
	}

	return colIndex
}

func sortDataChronologically(arr [][]interface{}, dateIndex int) {
	// Sort stock data according to the date in ascending order
	// Parameters:
	// 	arr ([][]interface{}): stock data containing a date field
	//  dateIndex (int): index of the date column in stock data array
	config := GetConfig()
	sort.Slice(arr, func(i, j int) bool {
		currDate, _ := time.Parse(config.Core.DateInputLayout, arr[i][dateIndex].(string))
		nextDate, _ := time.Parse(config.Core.DateInputLayout, arr[j][dateIndex].(string))
		return currDate.Before(nextDate)
	})
}
