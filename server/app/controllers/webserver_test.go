package controllers

import (
	"index-indicator-apis/server/app/entity"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// ModelsMock test用のmock
type ModelsMock struct{}

func (m *ModelsMock) CreateUser(name string, email string, pass string) (err error) {
	return nil
}
func (m *ModelsMock) FindUserByID(id int) (entity.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte("testpass"), 10)
	if err != nil {
		log.Fatal(err)
	}
	return entity.User{
		ID:       1,
		UserName: "testuser",
		Email:    "test@test",
		Password: string(hash),
	}, nil
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

func Test_userUpdateHandler(t *testing.T) {
	app := NewApp(&ModelsMock{})
	mux := http.NewServeMux()
	mux.HandleFunc("/users/1", app.userUpdateHandler)
	tests := []struct {
		name             string
		id               string
		argRequestReader io.Reader
		wantStatusCode   int
	}{
		{
			name: "正常なリクエスト",
			id:   "1",
			argRequestReader: strings.NewReader(`{
				"user": {
					"password": "testpass"
				},
				"new_user": {
					"user_name": "newuser",
					"email": "test@test",
					"password": "newpass"
				}
			}`),
			wantStatusCode: http.StatusOK,
		},
		{
			name: "正常なリクエスト(new user_name)",
			id:   "1",
			argRequestReader: strings.NewReader(`{
				"user": {
					"password": "testpass"
				},
				"new_user": {
					"user_name": "",
					"email": "new@test",
					"password": "newpass"
				}
			}`),
			wantStatusCode: http.StatusOK,
		},
		{
			name: "正常なリクエスト(new email)",
			id:   "1",
			argRequestReader: strings.NewReader(`{
				"user": {
					"password": "testpass"
				},
				"new_user": {
					"user_name": "newuser",
					"email": "",
					"password": "newpass"
				}
			}`),
			wantStatusCode: http.StatusOK,
		},
		{
			name: "正常なリクエスト(new password)",
			id:   "1",
			argRequestReader: strings.NewReader(`{
				"user": {
					"password": "testpass"
				},
				"new_user": {
					"user_name": "newuser",
					"email": "new@test",
					"password": ""
				}
			}`),
			wantStatusCode: http.StatusOK,
		},
		{
			name: "異常系(password)",
			id:   "1",
			argRequestReader: strings.NewReader(`{
				"user": {
					"password": "testpassss"
				},
				"new_user": {
					"user_name": "newuser",
					"email": "new@test",
					"password": "newpass"
				}
			}`),
			wantStatusCode: http.StatusNotAcceptable,
		},
		{
			name: "異常系(id)",
			id:   "2",
			argRequestReader: strings.NewReader(`{
				"user": {
					"password": "testpass"
				},
				"new_user": {
					"user_name": "newuser",
					"email": "new@test",
					"password": "newpass"
				}
			}`),
			wantStatusCode: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			request, err := http.NewRequest("PUT", "/users/"+tt.id, tt.argRequestReader)
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
