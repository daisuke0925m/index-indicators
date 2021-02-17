package controllers

import (
	"github.com/gorilla/mux"
)

// Route return API routing
func Route(app *App) *mux.Router {
	r := mux.NewRouter()
	// user
	r.HandleFunc("/users/{id:[0-9]+}", app.userGetHandler).Methods("GET")
	r.HandleFunc("/users", app.signupHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/users/{id:[0-9]+}", app.userDeleteHandler).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/users/{id:[0-9]+}", app.userUpdateHandler).Methods("PUT")
	// like
	r.HandleFunc("/users/{id:[0-9]+}/likes", app.tokenVerifyMiddleWare(app.likeGetALLHandler)).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}/likes", app.tokenVerifyMiddleWare(app.likePostHandler)).Methods("POST", "OPTIONS")
	r.HandleFunc("/users/{id:[0-9]+}/likes/{id:[0-9]+}", app.tokenVerifyMiddleWare(app.likeDeleteHandler)).Methods("DELETE", "OPTIONS")
	// auth
	r.HandleFunc("/login", app.loginHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/logout", app.tokenVerifyMiddleWare(app.logoutHandler)).Methods("POST")
	r.HandleFunc("/refresh_token", app.tokenVerifyMiddleWare(app.refreshTokenHandler)).Methods("POST")
	// fgi
	r.HandleFunc("/fgi", app.fgiHandler).Methods("GET")
	// ticker
	r.HandleFunc("/ticker", app.tickerHandler).Methods("GET")
	return r
}
