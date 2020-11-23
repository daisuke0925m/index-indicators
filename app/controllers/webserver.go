package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"index-indicator-apis/app/models"
	"index-indicator-apis/config"
)

var templates = template.Must(template.ParseFiles("app/views/fgi.html"))

// viewFgiHandler
func viewFgiHandler(w http.ResponseWriter, r *http.Request) {
	result := models.GetLatestFgi()
	err := templates.ExecuteTemplate(w, "fgi.html", result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// StartWebServer webserver立ち上げ
func StartWebServer() error {
	http.HandleFunc("/chart/", viewFgiHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
