package controllers

import (
	"encoding/json"
	"index-indicators/server/app/entity"
	"net/http"
	"path"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

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
