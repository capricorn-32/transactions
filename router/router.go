package router

import (
	"net/http"
	"transactions/config"
	"transactions/handler"

	"github.com/gorilla/mux"
)

func loggingMiddleware(next http.Handler) http.Handler {
	logger := config.GetLogger()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("%s %s", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func NewRouter(h *handler.Handler) http.Handler {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/accounts", h.Account.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{account_id}", h.Account.GetAccount).Methods("GET")
	r.HandleFunc("/transactions", h.Transaction.SubmitTransaction).Methods("POST")
	return r
}
