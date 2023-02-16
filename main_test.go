package main

import (
	"errors"
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
			args:     []string{"AAPL", "2017-01-01", "2018-01-01"},
			expected: nil,
		},
		{
			name:     "incorrect number of arguments",
			args:     []string{"AAPL", "2017-01-01"},
			expected: errors.New("incorrect number of arguments. 3 arguments are required: symbol, startDate, endDate"),
		},
		{
			name:     "invalid symbol",
			args:     []string{"apple", "2017-01-01", "2018-01-01"},
			expected: errors.New("invalid symbol"),
		},
		{
			name:     "invalid start date",
			args:     []string{"AAPL", "2017-01-44", "2018-01-01"},
			expected: errors.New("invalid start date. Dates should be in this layout 2006-01-02"),
		},
		{
			name:     "invalid end date",
			args:     []string{"AAPL", "2017-01-01", "2018-22-01"},
			expected: errors.New("invalid end date. Dates should be in this layout 2006-01-02"),
		},
		{
			name:     "end date before start date",
			args:     []string{"AAPL", "2018-01-01", "2017-01-01"},
			expected: errors.New("end date cannot be before start date"),
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
