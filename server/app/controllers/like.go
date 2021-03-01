package controllers

import (
	"encoding/json"
	"index-indicators/server/app/entity"
	"net/http"
	"regexp"
	"strconv"
)

func (a *App) likeGetALLHandler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`[\d\-]+`)
	values := re.FindStringSubmatch(r.URL.Path)
	userID, err := strconv.Atoi(values[0])
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := a.DB.FindUserByID(userID)
	if err != nil {
		a.resposeStatusCode(w, "User not found", http.StatusNotFound)
		return
	}

	likes, err := a.DB.FindUsersLikes(user)
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}

	type body struct {
		Likes []entity.Like `json:"likes,omitempty"`
	}

	likesBody := body{
		Likes: likes,
	}

	a.serveHTTPHeaders(w)
	json.NewEncoder(w).Encode(likesBody)
}

func (a *App) likePostHandler(w http.ResponseWriter, r *http.Request) {
	a.serveHTTPHeaders(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	re := regexp.MustCompile(`[\d\-]+`)
	values := re.FindStringSubmatch(r.URL.Path)
	userID, err := strconv.Atoi(values[0])
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := a.DB.FindUserByID(userID)
	if err != nil {
		a.resposeStatusCode(w, "User not found", http.StatusNotFound)
		return
	}

	type reqBody struct {
		Symbol string `json:"symbol,omitempty"`
	}
	var body reqBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		a.resposeStatusCode(w, "request body is invalid", http.StatusBadRequest)
		return
	}
	symbol := body.Symbol
	// bodyチェック
	if symbol == "" {
		a.resposeStatusCode(w, "Symbol is required", http.StatusBadRequest)
		return
	}

	// symbolが有効か
	if err := a.DB.FetchSymbol(symbol); err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusNotFound)
		return
	}

	// すでに存在している場合はエラーを返す
	like, err := a.DB.CheckLikesSymbol(userID, symbol)
	if err != err {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}
	if like.Symbol != "" {
		a.resposeStatusCode(w, "Symbol is already registered", http.StatusConflict)
		return
	}

	if err := a.DB.CreateLike(user, symbol); err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.resposeStatusCode(w, "success", http.StatusCreated)
	return
}

func (a *App) likeDeleteHandler(w http.ResponseWriter, r *http.Request) {
	re, err := regexp.Compile(`[\d\-]+`)
	if err != nil {
		a.resposeStatusCode(w, "could not parse path", http.StatusBadRequest)
		return
	}
	values := re.FindAllString(r.URL.Path, -1)
	userID, err := strconv.Atoi(values[0])
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}
	likeID, err := strconv.Atoi(values[1])
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = a.DB.FindUserByID(userID)
	if err != nil {
		a.resposeStatusCode(w, "User not found", http.StatusNotFound)
		return
	}

	like, err := a.DB.FindLikeByID(likeID)
	if err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := a.DB.DeleteLike(like); err != nil {
		a.resposeStatusCode(w, err.Error(), http.StatusBadRequest)
		return
	}
	a.resposeStatusCode(w, "success", http.StatusCreated)
	return
}
