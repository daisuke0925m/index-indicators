package controllers

import (
	"encoding/json"
	"index-indicators/server/app/entity"
	"net/http"
	"strconv"
)

func (a *App) fgiHandler(w http.ResponseWriter, r *http.Request) {
	strLimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(strLimit)
	if strLimit == "" || err != nil || limit < 0 || limit > 100 {
		limit = 100
	}
	fgis := a.DB.GetFgis(limit)

	type body struct {
		Fgis []entity.Fgi `json:"fgis,omitempty"`
	}

	fgisBody := body{
		Fgis: fgis,
	}

	a.serveHTTPHeaders(w)
	json.NewEncoder(w).Encode(fgisBody)
}
