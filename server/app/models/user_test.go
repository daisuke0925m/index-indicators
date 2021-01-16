package models

import (
	"fmt"
	"testing"
	"time"
)

type MockDB struct{}

type MockUser struct {
	DB        MockDB
	ID        int
	UserName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *MockUser) CreateUser(name, email, pass string) error {
	fmt.Println("create user. dummy")
	return nil
}

func TestUser_CreateUser(t *testing.T) {
	type fields struct {
		DB        MockDB
		ID        int
		UserName  string
		Email     string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	type args struct {
		name  string
		email string
		pass  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				name:  "testuser",
				email: "testuser@co.jp",
				pass:  "testpass",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := &MockUser{
				DB:        tt.fields.DB,
				ID:        tt.fields.ID,
				UserName:  tt.fields.UserName,
				Email:     tt.fields.Email,
				Password:  tt.fields.Password,
				CreatedAt: tt.fields.CreatedAt,
				UpdatedAt: tt.fields.UpdatedAt,
			}
			if err := user.CreateUser(tt.args.name, tt.args.email, tt.args.pass); (err != nil) != tt.wantErr {
				t.Errorf("User.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
