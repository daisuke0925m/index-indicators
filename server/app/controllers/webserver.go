package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	"index-indicator-apis/server/app/entity"
	"index-indicator-apis/server/app/models"

	"golang.org/x/crypto/bcrypt"
)

//App struct
type App struct {
	DB entity.DB
}

//NewApp return *APP
func NewApp(models *models.Models) *App {
	return &App{
		DB: models,
	}
}

// JSONError is returned when api return error
type JSONError struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func (a *App) apiError(w http.ResponseWriter, errMessage string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	jsonError, err := json.Marshal(JSONError{Error: errMessage, Code: code})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)
}

func (a *App) tokenVerifyMiddleWare(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessDetails, err := models.ExtractTokenMetadata(r)
		if err != nil {
			a.apiError(w, "unauthorized", http.StatusNotFound)
			return
		}

		// Redisからtokenを検索して見つからない場合はunauthorizedを返す。
		_, authErr := models.FetchAuth(accessDetails)
		if authErr != nil {
			a.apiError(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		fn(w, r)
	})
}

// ---------fgisHandlers---------
func (a *App) fgiHandler(w http.ResponseWriter, r *http.Request) {
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

// ---------usersHandlers---------
func (a *App) signupHandler(w http.ResponseWriter, r *http.Request) {
	var u entity.User
	json.NewDecoder(r.Body).Decode(&u)

	name := u.UserName
	email := u.Email
	pass := u.Password

	if name == "" {
		a.apiError(w, "UserName is required", http.StatusBadRequest)
		return
	}
	if email == "" {
		a.apiError(w, "Email is required", http.StatusBadRequest)
		return
	}
	if pass == "" {
		a.apiError(w, "Password is required", http.StatusBadRequest)
		return
	}

	if err := a.DB.CreateUser(name, email, pass); err != nil {
		a.apiError(w, "username or email are duplicated", http.StatusConflict)
		return
	}

	a.apiError(w, "success", http.StatusCreated)
	return
}

func (a *App) userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	var u entity.User
	json.NewDecoder(r.Body).Decode(&u)

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		a.apiError(w, "cloud not find user", http.StatusNotFound)
		return
	}

	err = a.DB.DeleteUser(id, u.Password)
	if err != nil {
		a.apiError(w, err.Error(), http.StatusBadRequest)
		return
	}
	a.apiError(w, "success", http.StatusOK)
	return
}

func (a *App) userUpdateHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		a.apiError(w, err.Error(), http.StatusBadRequest)
		return
	}

	foundUser, err := a.DB.FindUserByID(id)
	if err != nil {
		a.apiError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a.DB.UpdateUser(foundUser, r)
	if err := a.DB.UpdateUser(foundUser, r); err != nil {
		a.apiError(w, "success", http.StatusOK)
		return
	}
	a.apiError(w, "success", http.StatusOK)
	return
}

func (a *App) loginHandler(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)

	searchedUser, err := models.FindUser(user)
	if err != nil {
		a.apiError(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Println("compare the password")
	if err := bcrypt.CompareHashAndPassword([]byte(searchedUser.Password), []byte(user.Password)); err != nil {
		a.apiError(w, err.Error(), http.StatusUnauthorized)
		return
	}
	fmt.Println("password is be valid")

	token, err := models.CreateToken(searchedUser.ID)
	if err != nil {
		a.apiError(w, err.Error(), http.StatusUnauthorized)
		return
	}

	saveErr := models.CreateAuth(searchedUser.ID, token)
	if saveErr != nil {
		a.apiError(w, err.Error(), http.StatusUnauthorized)
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

func (a *App) logoutHandler(w http.ResponseWriter, r *http.Request) {
	accessDetails, err := models.ExtractTokenMetadata(r)
	if err != nil {
		a.apiError(w, "not found", http.StatusNotFound)
		return
	}

	deleted, delErr := models.DeleteAuth(accessDetails.AccessUUID)
	if delErr != nil || deleted == 0 {
		a.apiError(w, "not found", http.StatusNotFound)
		return
	}

	a.apiError(w, "success", http.StatusOK)
}

func (a *App) refreshTokenHandler(w http.ResponseWriter, r *http.Request) {

	mapToken := map[string]string{}
	json.NewDecoder(r.Body).Decode(&mapToken)
	refreshToken := mapToken["refresh_token"]

	tokens, errMsg := models.RefreshAuth(r, refreshToken)
	if errMsg != "" {
		a.apiError(w, errMsg, http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(tokens)
}
