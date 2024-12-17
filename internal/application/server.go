package application

import (
	"encoding/json"
	"fmt"
	calc "github.com/voutoad/go_calc/pkg/go_calc"
	"log"
	"net/http"
)

type calculateRequest struct {
	Expression string `json:"expression"`
}

type calculateResponse struct {
	Result string `json:"result"`
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var reqData calculateRequest
		var resultStruct calculateResponse
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&reqData)
		//fmt.Printf("%+v", reqData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Internal server error"}`))
			return
		}
		resultData, err := calc.Calc(reqData.Expression)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(`{"error": "Expression is not valid"}`))
			return
		}
		resultStruct.Result = fmt.Sprintf("%.6f", resultData)
		result, err := json.Marshal(resultStruct)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error": "Internal server error"}`))
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error":"Method not allowed"}`))
	}
}

func (app *Application) RunServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", CalcHandler)
	handler := loggingMiddleware(mux)
	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		return err
	}
	return nil
}
