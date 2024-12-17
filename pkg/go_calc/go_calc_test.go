package go_calc_test

import (
	"github.com/voutoad/go_calc/pkg/go_calc"
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

	testCasesFail := []struct {
		name       string
		expression string
	}{
		{
			name:       "simple",
			expression: "2+2+",
		},
		{
			name:       "division by zero",
			expression: "1/0",
		},
		{
			name:       "priority",
			expression: "2+2**2",
		},
		{
			name:       "priority",
			expression: "((2+2-*(2",
		},
		{
			name:       "empty",
			expression: "",
		},
	}
	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := go_calc.Calc(testCase.expression)
			if err == nil {
				t.Fatalf("expression %s is invalid, but given result %f", testCase.expression, result)
			}
		})
	}
}
