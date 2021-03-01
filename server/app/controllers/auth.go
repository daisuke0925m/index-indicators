package controllers

import (
	"encoding/json"
	"index-indicators/server/app/entity"
	"index-indicators/server/app/models"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (a *App) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		a.serveHTTPHeaders(w)
		w.WriteHeader(http.StatusOK)
		return
	}
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)

	foundUser, err := a.DB.FindUserByEmail(user.Email)
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password)); err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusUnauthorized)
		return
	}

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

	at := &http.Cookie{
		Name:     "at",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}

	http.SetCookie(w, at)
	rt := &http.Cookie{
		Name:     "rt",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	}
	http.SetCookie(w, rt)

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

	type resBody struct {
		ID int `json:"id,omitempty"`
	}

	id, err := strconv.Atoi(tokens["user_id"])
	body := &resBody{id}

	js, err := json.Marshal(body)
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(js)
}
