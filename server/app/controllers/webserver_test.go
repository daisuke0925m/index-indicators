package controllers

import (
	"index-indicator-apis/server/app/entity"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type ModelsMock struct{}

func NewAppMock(models *ModelsMock) *App {
	return &App{
		DB: models,
	}
}

func (a *ModelsMock) CreateUser(name, email, pass string) (err error) {
	return nil
}

func (a *ModelsMock) FindUserByID(id int) (user entity.User, err error) {
	return entity.User{
		ID:       1,
		UserName: "testuser",
		Email:    "test@co.jp",
		Password: "testpass",
	}, nil
}

func (a *ModelsMock) UpdateUser(foundUser entity.User) (err error) {
	return nil
}

func (a *ModelsMock) DeleteUser(id int, pass string) (err error) {
	return nil
}

func TestApp_signupHandler(t *testing.T) {
	var models ModelsMock
	app := NewAppMock(&models)
	mux := http.NewServeMux()
	mux.HandleFunc("/users", app.signupHandler)

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"user_name":"testuser","email":"test@test","password": "testpass"}`)
	request, _ := http.NewRequest("POST", "/users", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 201 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
