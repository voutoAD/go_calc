package application_test

import (
	"bytes"
	"fmt"
	"github.com/voutoad/go_calc/internal/application"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testingCase struct {
	name               string
	expression         string
	expectedResult     string
	expectedStatusCode int
}

func TestCalculateHandlerSuccess(t *testing.T) {
	testCases := []testingCase{
		{
			name:               "simple",
			expression:         "2+2",
			expectedResult:     `{"result":"4.000000"}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "priority",
			expression:         "(2+2)*2",
			expectedResult:     `{"result":"8.000000"}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "priority",
			expression:         "2+2*2",
			expectedResult:     `{"result":"6.000000"}`,
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "division",
			expression:         "1/5",
			expectedResult:     `{"result":"0.200000"}`,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rawData := fmt.Sprintf(`{"expression":"%s"}`, testCase.expression)
			byts := bytes.NewBufferString(rawData)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", byts)
			application.CalcHandler(w, req)
			res := w.Result()
			defer res.Body.Close()
			data, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("failed to read response body: %v", err)
			}
			if res.StatusCode != testCase.expectedStatusCode {
				t.Errorf("expected status %d, got %d", testCase.expectedStatusCode, res.StatusCode)
			}
			if string(data) != testCase.expectedResult {
				t.Errorf("expected %s, got %s", testCase.expectedResult, string(data))
			}
		})
	}
}

func TestCalculateHandlerFailure(t *testing.T) {
	testCases := []testingCase{
		{
			name:               "simple",
			expression:         "2+2+",
			expectedResult:     `{"error": "Expression is not valid"}`,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		{
			name:               "division by zero",
			expression:         "1/0",
			expectedResult:     `{"error": "Expression is not valid"}`,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		{
			name:               "priority",
			expression:         "2+2**2",
			expectedResult:     `{"error": "Expression is not valid"}`,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		{
			name:               "priority",
			expression:         "((2+2-*(2",
			expectedResult:     `{"error": "Expression is not valid"}`,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
		{
			name:               "empty",
			expression:         "",
			expectedResult:     `{"error": "Expression is not valid"}`,
			expectedStatusCode: http.StatusUnprocessableEntity,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			rawData := fmt.Sprintf(`{"expression":"%s"}`, testCase.expression)
			byts := bytes.NewBufferString(rawData)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", byts)
			application.CalcHandler(w, req)
			res := w.Result()
			defer res.Body.Close()
			data, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("failed to read response body: %v", err)
			}
			if res.StatusCode != testCase.expectedStatusCode {
				t.Errorf("expected status %d, got %d", testCase.expectedStatusCode, res.StatusCode)
			}
			if string(data) != testCase.expectedResult {
				t.Errorf("expected %s, got %s", testCase.expectedResult, string(data))
			}
		})
	}
}
