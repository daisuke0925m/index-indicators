package models

import (
	"index-indicator-apis/server/app/entity"
	"testing"
)

func TestCreateUser(t *testing.T) {
	type args struct {
		user entity.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "case1",
			args: args{
				user: entity.User{
					UserName: "test2",
					Email:    "test2",
					Password: "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
