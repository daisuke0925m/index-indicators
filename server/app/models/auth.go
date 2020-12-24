package models

import (
	"index-indicator-apis/server/app/entity"
	"index-indicator-apis/server/mysql"
)

// Login login処理
func Login(u entity.User) (user entity.User, err error) {
	db, err := mysql.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	if err := db.Where("email = ?", u.Email).First(&u).Error; err != nil {
		return user, err
	}

	return u, nil
}
