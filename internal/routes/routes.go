package routes

import (
	"net/http"

	"github.com/N30A/fondhav/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux, fundHandler *handlers.FundHandler) {
	mux.HandleFunc("GET /funds", fundHandler.GetFunds)
	mux.HandleFunc("GET /funds/{isin}", fundHandler.GetFund)
	mux.HandleFunc("GET /funds/{isin}/holdings", fundHandler.GetFundHoldings)
}
