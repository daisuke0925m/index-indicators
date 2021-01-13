package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type FakeUser struct {
	ID        int       `json:"id,omitempty" gorm:"primaryKey,unique"`
	UserName  string    `json:"user_name,omitempty" gorm:"unique"`
	Email     string    `json:"email,omitempty" gorm:"unique"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (user *FakeUser) CreateUser(name string, email string, pass string) (err error) {
	user.UserName = name
	user.Email = email
	user.Password = pass
	return
}

func Test_signupHandler(t *testing.T) {
	mux := http.NewServeMux()
	user := &FakeUser{}
	mux.HandleFunc("/users", signupHandler(user))

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"user_name":"testuser","email":"test@test","password": "testpass"}`)
	request, _ := http.NewRequest("POST", "/users", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 201 {
		t.Errorf("Response code is %v", writer.Code)
	}

	if user.UserName != "testuser" {
		t.Error("UserName is not correct", user.UserName)
	}

	if user.Email != "test@test" {
		t.Error("Email is not correct", user.Email)
	}

	if user.Password != "testpass" {
		t.Error("Password is not correct", user.Password)
	}
}
