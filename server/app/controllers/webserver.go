package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	"index-indicator-apis/server/app/entity"
	"index-indicator-apis/server/app/models"

	"golang.org/x/crypto/bcrypt"
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
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
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
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
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
		_, authErr := models.FetchAuth(accessDetails)
		if authErr != nil {
			a.resposeStatusCode(w, "token is not found", http.StatusNotFound)
			return
		}
		fn(w, r)
	})
}

// ---------usersHandlers---------
func (a *App) userGetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}

	foundUser, err := a.DB.FindUserByID(id)
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type user struct {
		ID       int    `json:"id,omitempty"`
		UserName string `json:"user_name,omitempty"`
		Email    string `json:"email,omitempty"`
	}

	body := &user{
		ID:       foundUser.ID,
		UserName: foundUser.UserName,
		Email:    foundUser.Email,
	}

	a.serveHTTPHeaders(w)
	js, err := json.Marshal(body)
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(js)
	return
}

func (a *App) signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		a.serveHTTPHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}
	var u entity.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := u.UserName
	email := u.Email
	pass := u.Password

	if name == "" {
		a.resposeStatusCode(w, "UserName is required", http.StatusBadRequest)
		return
	}
	if email == "" {
		a.resposeStatusCode(w, "Email is required", http.StatusBadRequest)
		return
	}
	if pass == "" {
		a.resposeStatusCode(w, "Password is required", http.StatusBadRequest)
		return
	}

	if err := a.DB.CreateUser(name, email, pass); err != nil {
		a.resposeStatusCode(w, "username or email are duplicated", http.StatusConflict)
		return
	}

	a.resposeStatusCode(w, "success", http.StatusCreated)
	return
}

func (a *App) userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	a.serveHTTPHeaders(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	var u entity.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		a.resposeStatusCode(w, "cloud not find user", http.StatusNotFound)
		return
	}

	if u.Password == "" {
		a.resposeStatusCode(w, "Password is required", http.StatusBadRequest)
		return
	}

	err = a.DB.DeleteUser(id, u.Password)
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}

	at := &http.Cookie{
		Name:     "at",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		// Secure:   true,TODO
	}
	http.SetCookie(w, at)
	rt := &http.Cookie{
		Name:     "rt",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		// Secure:   true,TODO
	}
	http.SetCookie(w, rt)

	a.resposeStatusCode(w, "success", http.StatusOK)
	return
}

func (a *App) userUpdateHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}

	foundUser, err := a.DB.FindUserByID(id)
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type body struct {
		User struct {
			Password string `json:"password,omitempty"`
		} `json:"user,omitempty"`
		NewUser struct {
			UserName string `json:"user_name,omitempty"`
			Email    string `json:"email,omitempty"`
			Password string `json:"password,omitempty"`
		} `json:"new_user,omitempty"`
	}

	var updateUser body
	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(updateUser.User.Password)); err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	if updateUser.NewUser.UserName != "" {
		foundUser.UserName = updateUser.NewUser.UserName
	}
	if updateUser.NewUser.Email != "" {
		foundUser.Email = updateUser.NewUser.Email
	}
	if updateUser.NewUser.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(updateUser.NewUser.Password), 10)
		if err != nil {
			a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
			return
		}
		foundUser.Password = string(hash)
	}

	if err := a.DB.UpdateUser(foundUser); err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.resposeStatusCode(w, "success", http.StatusOK)
	return
}

// ---------authHandlers---------
func (a *App) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		a.serveHTTPHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)

	foundUser, err := models.FindUser(user)
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Println("compare the password")
	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusUnauthorized)
		return
	}
	fmt.Println("password is be valid")

	token, err := models.CreateToken(foundUser.ID)
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusUnauthorized)
		return
	}

	saveErr := models.CreateAuth(foundUser.ID, token)
	if saveErr != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusUnauthorized)
		return
	}

	accessCookie := &http.Cookie{
		Name:     "at",
		Value:    token.AccessToken,
		HttpOnly: true,
		// Secure:   true,TODO
	}
	http.SetCookie(w, accessCookie)
	refreshCookie := &http.Cookie{
		Name:     "rt",
		Value:    token.RefreshToken,
		HttpOnly: true,
		// Secure:   true,TODO
	}

	a.serveHTTPHeaders(w)
	http.SetCookie(w, refreshCookie)

	type resBody struct {
		ID       int    `json:"id,omitempty"`
		UserName string `json:"user_name,omitempty"`
		Email    string `json:"email,omitempty"`
	}

	body := &resBody{
		ID:       foundUser.ID,
		UserName: foundUser.UserName,
		Email:    foundUser.Email,
	}

	js, err := json.Marshal(body)
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(js)
	return
}

func (a *App) logoutHandler(w http.ResponseWriter, r *http.Request) {
	accessDetails, err := models.ExtractTokenMetadata(r)
	if err != nil {
		a.resposeStatusCode(w, "not found", http.StatusNotFound)
		return
	}

	deleted, delErr := models.DeleteAuth(accessDetails.AccessUUID)
	if delErr != nil || deleted == 0 {
		a.resposeStatusCode(w, "not found", http.StatusNotFound)
		return
	}

	a.resposeStatusCode(w, "success", http.StatusOK)
}

func (a *App) refreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	cookieRt, err := r.Cookie("rt")
	if err != nil {
		a.resposeStatusCode(w, "can't read cookie", http.StatusBadRequest)
		return
	}
	refreshToken := cookieRt.Value

	tokens, errMsg := models.RefreshAuth(r, refreshToken)
	if errMsg != "" {
		a.resposeStatusCode(w, errMsg, http.StatusUnauthorized)
		return
	}

	accessCookie := &http.Cookie{
		Name:     "at",
		Value:    tokens["access_token"],
		HttpOnly: true,
		// Secure:   true,TODO
	}
	http.SetCookie(w, accessCookie)
	refreshCookie := &http.Cookie{
		Name:     "rt",
		Value:    tokens["refresh_token"],
		HttpOnly: true,
		// Secure:   true,TODO
	}
	http.SetCookie(w, refreshCookie)

	a.serveHTTPHeaders(w)
}

// ---------fgisHandlers---------
func (a *App) fgiHandler(w http.ResponseWriter, r *http.Request) {
	strLimit := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(strLimit)
	if strLimit == "" || err != nil || limit < 0 || limit > 100 {
		limit = 100
	}
	fgis := models.GetFgis(limit)

	type body struct {
		Fgis []entity.Fgi `json:"fgis,omitempty"`
	}

	fgisBody := body{
		Fgis: fgis,
	}

	a.serveHTTPHeaders(w)
	json.NewEncoder(w).Encode(fgisBody)
}

// ---------ticker---------
func (a *App) tickerHandler(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")

	if symbol == "" {
		a.resposeStatusCode(w, "symbol is required", http.StatusUnauthorized)
		return
	}

	tickers, err := models.GetTickerAll(symbol)
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
