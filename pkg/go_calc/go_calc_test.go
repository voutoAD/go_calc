package go_calc_test

import (
	go_calc "github.com/voutoAD/go_calc"
	"testing"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "simple",
			expression:     "1+1",
			expectedResult: 2,
		},
		{
			name:           "priority",
			expression:     "(2+2)*2",
			expectedResult: 8,
		},
		{
			name:           "priority",
			expression:     "2+2*2",
			expectedResult: 6,
		},
		{
			name:           "division",
			expression:     "1/5",
			expectedResult: 0.2,
		},
	}
	for _, test := range testCasesSuccess {
		t.Run(test.name, func(t *testing.T) {
			result, err := go_calc.Calc(test.expression)
			if err != nil {
				t.Fatalf("succesful case returns %serror", err)
			}
			if result != test.expectedResult {
				t.Fatalf("succesful case returns %f, want %f", result, test.expectedResult)
			}
		})
	}
}
