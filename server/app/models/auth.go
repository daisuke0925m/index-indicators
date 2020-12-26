package models

import (
	"fmt"
	"index-indicator-apis/server/app/entity"
	"index-indicator-apis/server/config"
	"index-indicator-apis/server/mysql"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
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
func CreateToken(userid int) (*entity.TokenDetails, error) {
	td := &entity.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUUID = uuid.NewV4().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["authorized"] = td.AccessUUID
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(config.Config.JwtSecret))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(config.Config.JwtSecret))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// ExtractToken cookieからjwtを取得
func ExtractToken(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// var client *redis.Client

// // InitRedis redis起動
// func InitRedis() {
// 	dsn := os.Getenv("REDIS_DSN")
// 	if len(dsn) == 0 {
// 		dsn = "localhost:6379"
// 	}
// 	client = redis.NewClient(&redis.Options{
// 		Addr: dsn, //redis port
// 	})
// 	_, err := client.Ping().Result()
// 	if err != nil {
// 		panic(err)
// 	}
// }
