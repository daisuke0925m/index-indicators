package entity

import "net/http"

type DB interface {
	CreateUser(name, email, pass string) (err error)
	FindUserByID(r *http.Request) (User, error)
	UpdateUser(foundUser User, r *http.Request) (err error)
	DeleteUser(id int, pass string) (err error)
}
