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

func calcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var reqData calculateRequest
		var resultJSON calculateResponse
		data := r.Body
		buf := make([]byte, 100)
		n, err := data.Read(buf)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		buf = buf[:n]
		err = json.Unmarshal(buf, &reqData)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		expr := reqData.Expression
		resultData, err := calc.Calc(expr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		resultJSON.Result = fmt.Sprintf("%.6f", resultData)
		resultBytes, err := json.Marshal(resultJSON)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(resultBytes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error":"Method not allowed"}`))
	}
}

func (app *Application) RunServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", calcHandler)
	handler := loggingMiddleware(mux)
	err := http.ListenAndServe(":3000", handler)
	if err != nil {
		return err
	}
	return nil
}
