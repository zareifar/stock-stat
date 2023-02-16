package main

import (
	"testing"
)

func TestCalculateMaxDrawdown(t *testing.T) {
	data := JSONData{
		Datatable: Datatable{
			Columns: []Column{
				{Name: "date", Type: "date"},
				{Name: "adj_close", Type: "double"},
			},
			Data: [][]interface{}{
				{"2022-01-01", 100.0},
				{"2022-01-02", 110.0},
				{"2022-01-03", 90.0},
				{"2022-01-04", 95.0},
				{"2022-01-05", 80.0},
				{"2022-01-06", 85.0},
			},
		},
	}

	expectedMaxDrawdown := float64(0.3)

	maxDrawdown, err := CalculateMaxDrawdown(data)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if maxDrawdown != expectedMaxDrawdown {
		t.Fatalf("expected max drawdown: %f, but got: %f", expectedMaxDrawdown, maxDrawdown)
	}
}

func TestCalculateMaxDrawdownWithInsufficientData(t *testing.T) {
	data := JSONData{
		Datatable: Datatable{
			Columns: []Column{
				{Name: "date", Type: "date"},
				{Name: "adj_close", Type: "double"},
			},
			Data: [][]interface{}{
				{"2022-01-01", 100},
			},
		},
	}

	_, err := CalculateMaxDrawdown(data)
	if err == nil {
		t.Fatalf("expected error but got none")
	}
}

func TestCalculateMaxDrawdownWithNoAdjCloseColumn(t *testing.T) {
	data := JSONData{
		Datatable: Datatable{
			Columns: []Column{
				{Name: "date", Type: "date"},
			},
			Data: [][]interface{}{
				{"2022-01-01"},
				{"2022-01-02"},
				{"2022-01-03"},
				{"2022-01-04"},
				{"2022-01-05"},
				{"2022-01-06"},
			},
		},
	}

	_, err := CalculateMaxDrawdown(data)
	if err == nil {
		t.Fatalf("expected error but got none")
	}
}

func TestCalculateSimpleReturn(t *testing.T) {
	data := JSONData{
		Datatable: Datatable{
			Columns: []Column{
				{Name: "date", Type: "date"},
				{Name: "adj_close", Type: "double"},
			},
			Data: [][]interface{}{
				{"2022-01-01", 100.0},
				{"2022-01-02", 110.0},
			},
		},
	}

	simpleReturn, err := CalculateSimpleReturn(data)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedReturn := float64(0.1)

	if simpleReturn != expectedReturn {
		t.Errorf("Expected simple return %v, but got %v", expectedReturn, simpleReturn)
	}
}

func TestCalculateSimpleReturnWithNoAdjCloseColumn(t *testing.T) {
	data := JSONData{
		Datatable: Datatable{
			Columns: []Column{
				{Name: "date", Type: "date"},
			},
			Data: [][]interface{}{
				{"2022-01-01", 100.0},
				{"2022-01-02", 110.0},
			},
		},
	}

	_, err := CalculateSimpleReturn(data)

	if err == nil {
		t.Error("Expected an error due to missing adj_close column, but got nil")
	}
}

func TestCalculateSimpleReturn_WithInsufficientData(t *testing.T) {
	data := JSONData{
		Datatable: Datatable{
			Columns: []Column{
				{Name: "date", Type: "date"},
				{Name: "adj_close", Type: "double"},
			},
			Data: [][]interface{}{
				{"2022-01-01", 100},
			},
		},
	}

	_, err := CalculateSimpleReturn(data)

	if err == nil {
		t.Error("Expected an error due to insufficient data, but got nil")
	}
}

func TestCalculateSimpleReturn_WithZeroStartPrice(t *testing.T) {
	data := JSONData{
		Datatable: Datatable{
			Columns: []Column{
				{Name: "date", Type: "date"},
				{Name: "adj_close", Type: "double"},
			},
			Data: [][]interface{}{
				{"2022-01-01", 0.0},
				{"2022-01-02", 110.0},
			},
		},
	}

	_, err := CalculateSimpleReturn(data)

	if err == nil {
		t.Errorf("Expected an error due to starting price of zero, but got nil")
	}
}
