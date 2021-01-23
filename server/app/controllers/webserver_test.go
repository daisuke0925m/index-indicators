package controllers

import (
	"index-indicator-apis/server/app/entity"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// ModelsMock test用のmock
type ModelsMock struct{}

func (m *ModelsMock) CreateUser(name string, email string, pass string) (err error) {
	return nil
}
func (m *ModelsMock) FindUserByID(id int) (entity.User, error) {
	return entity.User{}, nil
}
func (m *ModelsMock) UpdateUser(foundUser entity.User) (err error) {
	return nil
}
func (m *ModelsMock) DeleteUser(id int, pass string) (err error) {
	return nil
}
func Test_signupHandler(t *testing.T) {
	app := NewApp(&ModelsMock{})
	mux := http.NewServeMux()
	mux.HandleFunc("/users", app.signupHandler)
	tests := []struct {
		name             string
		argRequestReader io.Reader
		wantStatusCode   int
	}{
		{name: "正常なリクエスト", argRequestReader: strings.NewReader(`{"user_name":"testuser","email":"test@test","password": "testpass"}`), wantStatusCode: http.StatusCreated},
		{name: "異常系(user_name)", argRequestReader: strings.NewReader(`{"user_name": "","email":"test@test","password": "testpass"}`), wantStatusCode: http.StatusBadRequest},
		{name: "異常系(email)", argRequestReader: strings.NewReader(`{"user_name": "testuser","email":"","password": "testpass"}`), wantStatusCode: http.StatusBadRequest},
		{name: "異常系(password)", argRequestReader: strings.NewReader(`{"user_name": "testuser","email":"test@test","password": ""}`), wantStatusCode: http.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			request, err := http.NewRequest("POST", "/users", tt.argRequestReader)
			if err != nil {
				t.Errorf("invalid Request reader %v", err)
			}
			mux.ServeHTTP(writer, request)
			if writer.Code != tt.wantStatusCode {
				t.Errorf("invalid status code want:%v, got:%v", tt.wantStatusCode, writer.Code)
			}
		})
	}
}

func Test_userDeleteHandler(t *testing.T) {
	app := NewApp(&ModelsMock{})
	mux := http.NewServeMux()
	mux.HandleFunc("/users/1", app.userDeleteHandler)
	tests := []struct {
		name             string
		id               string
		argRequestReader io.Reader
		wantStatusCode   int
	}{
		{name: "正常なリクエスト", id: "1", argRequestReader: strings.NewReader(`{"password": "testpass"}`), wantStatusCode: http.StatusOK},
		{name: "異常系(password)", id: "1", argRequestReader: strings.NewReader(`{"password": ""}`), wantStatusCode: http.StatusBadRequest},
		{name: "異常系(id)", id: "", argRequestReader: strings.NewReader(`{"password": "testpass"}`), wantStatusCode: http.StatusNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			request, err := http.NewRequest("DELETE", "/users/"+tt.id, tt.argRequestReader)
			if err != nil {
				t.Errorf("invalid Request reader %v", err)
			}
			mux.ServeHTTP(writer, request)
			if writer.Code != tt.wantStatusCode {
				t.Errorf("invalid status code want:%v, got:%v", tt.wantStatusCode, writer.Code)
			}
		})
	}
}
