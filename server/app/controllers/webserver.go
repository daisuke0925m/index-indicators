package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"index-indicator-apis/server/app/entity"
	"index-indicator-apis/server/app/models"
	"index-indicator-apis/server/config"
	"index-indicator-apis/server/db"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
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

func tokenVerifyMiddleWare(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessDetails, err := models.ExtractTokenMetadata(r)
		if err != nil {
			apiError(w, "unauthorized", http.StatusNotFound)
			return
		}

		// Redisからtokenを検索して見つからない場合はunauthorizedを返す。
		_, authErr := models.FetchAuth(accessDetails)
		if authErr != nil {
			apiError(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		fn(w, r)
	})
}

func fgiHandler(w http.ResponseWriter, r *http.Request) {
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

func signupHandler(u models.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)

		name := user.UserName
		email := user.Email
		pass := user.Password

		if name == "" {
			apiError(w, "UserName is required", http.StatusBadRequest)
			return
		}
		if email == "" {
			apiError(w, "Email is required", http.StatusBadRequest)
			return
		}
		if pass == "" {
			apiError(w, "Password is required", http.StatusBadRequest)
			return
		}

		if err := u.CreateUser(name, email, pass); err != nil {
			apiError(w, "username or email are duplicated", http.StatusConflict)
			return
		}

		apiError(w, "success", http.StatusCreated)
		return
	}
}

func userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	foundUser, err := models.FindUserByID(r)
	if err != nil {
		apiError(w, err.Error(), http.StatusInternalServerError)
	}

	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		apiError(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	db, err := db.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	if err := db.Delete(&foundUser).Error; err != nil {
		apiError(w, err.Error(), http.StatusInternalServerError)
	}

	apiError(w, "success", http.StatusOK)
}

func userUpdateHandler(w http.ResponseWriter, r *http.Request) {
	foundUser, err := models.FindUserByID(r)
	if err != nil {
		apiError(w, err.Error(), http.StatusInternalServerError)
	}

	models.UpdateUser(foundUser, r)
	apiError(w, "success", http.StatusOK)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)

	searchedUser, err := models.FindUser(user)
	if err != nil {
		apiError(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Println("compare the password")
	if err := bcrypt.CompareHashAndPassword([]byte(searchedUser.Password), []byte(user.Password)); err != nil {
		apiError(w, err.Error(), http.StatusUnauthorized)
		return
	}
	fmt.Println("password is be valid")

	token, err := models.CreateToken(searchedUser.ID)
	if err != nil {
		apiError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	saveErr := models.CreateAuth(searchedUser.ID, token)
	if saveErr != nil {
		apiError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(tokens)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	accessDetails, err := models.ExtractTokenMetadata(r)
	if err != nil {
		apiError(w, "not found", http.StatusNotFound)
		return
	}

	deleted, delErr := models.DeleteAuth(accessDetails.AccessUUID)
	if delErr != nil || deleted == 0 {
		apiError(w, "not found", http.StatusNotFound)
		return
	}

	apiError(w, "success", http.StatusOK)
}

func refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	mapToken := map[string]string{}
	json.NewDecoder(r.Body).Decode(&mapToken)
	refreshToken := mapToken["refresh_token"]

	tokens, errMsg := models.RefreshAuth(r, refreshToken)
	if errMsg != "" {
		apiError(w, errMsg, http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(tokens)
}

// StartWebServer webserver立ち上げ
func StartWebServer() error {
	db, err := db.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	r := mux.NewRouter()
	http.Handle("/", r)

	fmt.Println("connecting...")
	// users
	r.HandleFunc("/users", signupHandler(&models.User{DB: db})).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", userDeleteHandler).Methods("DELETE")
	r.HandleFunc("/users/{id:[0-9]+}", userUpdateHandler).Methods("PUT")
	// auth
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/logout", tokenVerifyMiddleWare(logoutHandler)).Methods("POST")
	r.HandleFunc("/refresh_token", refreshTokenHandler).Methods("POST")
	// fgi
	r.HandleFunc("/fgi", tokenVerifyMiddleWare(fgiHandler)).Methods("GET")
	fmt.Printf("connected port :%d\n", config.Config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), nil)
}
