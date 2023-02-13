package main

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name     string
		args     []string
		expected error
	}{
		{
			name:     "correct number of arguments",
			args:     []string{"AAPL", "2022-Jan-01", "2022-Dec-31"},
			expected: nil,
		},
		{
			name:     "incorrect number of arguments",
			args:     []string{"AAPL", "2022-Jan-01"},
			expected: fmt.Errorf("incorrect number of arguments. 3 arguments are required: symbol, startDate, endDate"),
		},
		{
			name:     "invalid symbol",
			args:     []string{"apple", "2022-Jan-01", "2022-Dec-31"},
			expected: fmt.Errorf("invalid symbol"),
		},
		{
			name:     "invalid start date",
			args:     []string{"AAPL", "2022-01-44", "2022-Dec-31"},
			expected: fmt.Errorf("invalid start date. Dates should be in this layout 2006-Jan-02"),
		},
		{
			name:     "invalid end date",
			args:     []string{"AAPL", "2022-01-01", "2022-12-33"},
			expected: fmt.Errorf("invalid start date. Dates should be in this layout 2006-Jan-02"),
		},
		{
			name:     "end date before start date",
			args:     []string{"AAPL", "2022-Dec-12", "2022-Jan-01"},
			expected: fmt.Errorf("end date cannot be before start date"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := run(testCase.args)
			if err != nil && (err.Error() != testCase.expected.Error()) {
				t.Errorf("run(%v) = %v, expected %v", testCase.args, err, testCase.expected)
			}
		})
	}
}
