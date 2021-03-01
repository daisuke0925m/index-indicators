package controllers

import (
	"encoding/json"
	"index-indicators/server/app/entity"
	"net/http"
)

func (a *App) tickerHandler(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")

	if symbol == "" {
		a.resposeStatusCode(w, "symbol is required", http.StatusUnauthorized)
		return
	}

	tickers, err := a.DB.GetTickerAll(symbol)
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(tickers) == 0 {
		a.resposeStatusCode(w, "There is no brand you are looking for", http.StatusNoContent)
		return
	}

	type body struct {
		Daily []entity.Ticker `json:"daily,omitempty"`
	}

	tickerBody := body{
		Daily: tickers,
	}

	a.serveHTTPHeaders(w)
	json.NewEncoder(w).Encode(tickerBody)
}
