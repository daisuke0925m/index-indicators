package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"

	"index-indicators/server/app/entity"
	"index-indicators/server/app/models"
)

//App struct
type App struct {
	DB entity.DB
}

//NewApp return *APP
func NewApp(models entity.DB) *App {
	return &App{
		DB: models,
	}
}

// JSONResponse is a response mssage
type JSONResponse struct {
	Response string `json:"response"`
	Code     int    `json:"code"`
}

func (a *App) resposeStatusCode(w http.ResponseWriter, ResMessage string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_URL"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.WriteHeader(code)
	jsonError, err := json.Marshal(JSONResponse{Response: ResMessage, Code: code})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)
}

func (a *App) serveHTTPHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CORS_URL"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
}

func (a *App) tokenVerifyMiddleWare(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			a.serveHTTPHeaders(w)
			w.WriteHeader(http.StatusOK)
			return
		}
		accessDetails, err := models.ExtractTokenMetadata(r)
		if err != nil {
			a.resposeStatusCode(w, "token is invalid", http.StatusUnauthorized)
			return
		}

		// Redisからtokenを検索して見つからない場合はunauthorizedを返す。
		uid, authErr := models.FetchAuth(accessDetails)
		if authErr != nil {
			a.resposeStatusCode(w, "token is not found", http.StatusNotFound)
			return
		}

		// users/:idにマッチした場合、redisのuuidとパラメータのidが同じかチェック
		paths := strings.Split(r.URL.Path, "/")
		if paths[1] == "users" && regexp.MustCompile(`[0-9]`).Match([]byte(paths[2])) {
			id, err := strconv.Atoi(path.Base((paths[2])))
			if err != nil {
				a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
				return
			}
			if uid != id {
				a.resposeStatusCode(w, "token is invalid", http.StatusUnauthorized)
				return
			}
		}

		fn(w, r)
	})
}
