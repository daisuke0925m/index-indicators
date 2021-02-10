package models

import (
	"fmt"
	"index-indicators/server/app/entity"
	"index-indicators/server/config"
	"index-indicators/server/db"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

// ---SignUp処理の関数-------------

// FindUser 検索処理
func FindUser(u entity.User) (user entity.User, err error) {
	fmt.Println("start find user")
	db, err := db.SQLConnect()
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

// ---Login処理の関数-------------

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
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(config.Config.JwtAccess))
	if err != nil {
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(config.Config.JwtRefresh))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// CreateAuth redisにTokenDetailを保存する
func CreateAuth(userid int, td *entity.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := db.Redis.Set(td.AccessUUID, strconv.Itoa(userid), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := db.Redis.Set(td.RefreshUUID, strconv.Itoa(userid), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

// ---認証処理の関数-------------

// ExtractToken Cookieからjwtを取得
func ExtractToken(r *http.Request) string {
	cookieAt, err := r.Cookie("at")
	if err != nil {
		return ""
	}
	accessToken := cookieAt.Value
	return accessToken
}

// VerifyToken tokenが正しいか検証
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config.JwtAccess), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// TokenValid tokenの有効期限を検証
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// ExtractTokenMetadata  tokenをAccessDetailsで返す
func ExtractTokenMetadata(r *http.Request) (*entity.AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &entity.AccessDetails{
			AccessUUID: accessUUID,
			UserID:     userID,
		}, nil
	}
	return nil, err
}

// FetchAuth tokenを元にredisからuserIDを抽出
func FetchAuth(authD *entity.AccessDetails) (int, error) {
	userid, err := db.Redis.Get(authD.AccessUUID).Result()
	if err != nil {
		return 0, err
	}
	userID, err := strconv.Atoi(userid)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

// DeleteAuth Redisのjwtメタデータを削除
func DeleteAuth(givenUUID string) (int64, error) {
	deleted, err := db.Redis.Del(givenUUID).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

// RefreshAuth accessTokenが切れたているとき,Tokenを再生成する。
func RefreshAuth(r *http.Request, refreshToken string) (map[string]string, string) {

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config.JwtRefresh), nil
	})

	// token期限チェック
	if err != nil {
		return nil, "unauthorized"
	}

	// token正否チェック
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return nil, "unauthorized"
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		refreshUUID, ok := claims["refresh_uuid"].(string)
		if !ok {
			return nil, "unauthorized"
		}
		userID, err := strconv.Atoi(fmt.Sprintf("%.f", claims["user_id"]))
		if err != nil {
			return nil, "unauthorized"
		}

		deleted, delErr := DeleteAuth(refreshUUID)
		if delErr != nil || deleted == 0 {
			return nil, "unauthorized"
		}

		ts, createErr := CreateToken(userID)
		if createErr != nil {
			return nil, "unauthorized"
		}

		saveErr := CreateAuth(userID, ts)
		if saveErr != nil {
			return nil, "unauthorized"
		}

		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
			"user_id":       strconv.Itoa(userID),
		}

		return tokens, ""
	}
	{
		return nil, "unauthorized"
	}

}
