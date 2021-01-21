package controllers

import (
	"github.com/gorilla/mux"
)

// Route return API routing
func Route(app *App) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", app.signupHandler).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", app.userDeleteHandler).Methods("DELETE")
	r.HandleFunc("/users/{id:[0-9]+}", app.userUpdateHandler).Methods("PUT")
	// auth
	r.HandleFunc("/login", app.loginHandler).Methods("POST")
	r.HandleFunc("/logout", app.tokenVerifyMiddleWare(app.logoutHandler)).Methods("POST")
	r.HandleFunc("/refresh_token", app.refreshTokenHandler).Methods("POST")
	// fgi
	r.HandleFunc("/fgi", app.tokenVerifyMiddleWare(app.fgiHandler)).Methods("GET")
	return r
}
