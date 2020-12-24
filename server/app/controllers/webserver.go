package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"index-indicator-apis/server/app/entity"
	"index-indicator-apis/server/app/models"
	"index-indicator-apis/server/config"
)

// JSONError エラー情報を格納
type JSONError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func apiError(w http.ResponseWriter, errMessage string, code int) {
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
			apiError(w, "Not found", http.StatusNotFound)
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

func signupHandler(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)

	if user.UserName == "" {
		apiError(w, "Email is required", http.StatusBadRequest)
		return
	}
	if user.Email == "" {
		apiError(w, "Email is required", http.StatusBadRequest)
		return
	}
	if user.Password == "" {
		apiError(w, "Password is required", http.StatusBadRequest)
		return
	}

	// models.CreateUser(user)
	if err := models.CreateUser(user); err != nil {
		apiError(w, "username or email are duplicated", http.StatusConflict)
		return
	}

	apiError(w, "success", http.StatusCreated)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)

	searchedUser, err := models.Login(user)
	if err != nil {
		apiError(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Println(searchedUser)
	fmt.Println(user)
	// if user.Password != searchedUser.Password {
	// 	apiError(w, err.Error(), http.StatusUnauthorized)
	// 	return
	// }
	apiError(w, "success", http.StatusAccepted)
}

// StartWebServer webserver立ち上げ
func StartWebServer() error {
	fmt.Println("connecting...")
	http.HandleFunc("/api/fgi/", apiMakeHandler(apiFgiHandler))
	http.HandleFunc("/api/signup", apiMakeHandler(signupHandler))
	http.HandleFunc("/api/login", apiMakeHandler(loginHandler))
	fmt.Printf("connected port :%d\n", config.Config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
