package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"index-indicator-apis/server/app/models"
	"index-indicator-apis/server/config"
)

// JSONError エラー情報を格納
type JSONError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

// APIError APIエラーを返す
func APIError(w http.ResponseWriter, errMessage string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonError, err := json.Marshal(JSONError{Error: errMessage, Code: code})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)
}

var apiValidPath = regexp.MustCompile("^/api/")

func apiMakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := apiValidPath.FindStringSubmatch(r.URL.Path)
		if len(m) == 0 {
			APIError(w, "Not found", http.StatusNotFound)
		}
		fn(w, r)
	}
}

func apiFgiHandler(w http.ResponseWriter, r *http.Request) {
	strLimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(strLimit)
	if strLimit == "" || err != nil || limit < 0 || limit > 100 {
		limit = 100
	}
	fgi := models.GetFgis(limit)
	js, err := json.Marshal(fgi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}

func apiLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login")
	fmt.Println("login 関数実行")
}

// StartWebServer webserver立ち上げ
func StartWebServer() error {
	fmt.Println("connecting...")
	http.HandleFunc("/api/fgi/", apiMakeHandler(apiFgiHandler))
	http.HandleFunc("/api/signup", apiMakeHandler(models.SignupHandler))
	http.HandleFunc("/api/login", apiMakeHandler(apiLoginHandler))
	fmt.Printf("connected port :%d\n", config.Config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
