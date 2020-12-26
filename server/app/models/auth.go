package models

import (
	"fmt"
	"index-indicator-apis/server/app/entity"
	"index-indicator-apis/server/config"
	"index-indicator-apis/server/mysql"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// FindUser 検索処理
func FindUser(u entity.User) (user entity.User, err error) {
	fmt.Println("start find user")
	db, err := mysql.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	if err := db.Where("email = ?", u.Email).First(&u).Error; err != nil {
		fmt.Println("error!")
		return user, err
	}

	fmt.Println("finish! find a user")

	return u, nil
}

// CreateToken jwtToken作成
func CreateToken(userid int) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(config.Config.JwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

// SaveTokenToCookie jwtTokenをCookieに保存
func SaveTokenToCookie(token string, w http.ResponseWriter) (err error) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	return err
}
