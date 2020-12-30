package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"

	"index-indicator-apis/server/app/entity"
	"index-indicator-apis/server/db"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser user登録
func CreateUser(user entity.User) (err error) {
	fmt.Printf("start signup\n")

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(hash)

	db, err := db.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if err := db.Create(&user).Error; err != nil {
		return err
	}

	fmt.Printf("%v\n", user)
	fmt.Println("finish! created a user")
	return
}

// FindUserByID idからuserを検索
func FindUserByID(r *http.Request) (entity.User, error) {
	var user entity.User
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return entity.User{}, err
	}

	db, err := db.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		fmt.Println("error!")
		return user, err
	}
	return user, nil
}

func UpdateUser(foundUser entity.User, r *http.Request) (err error) {
	type body struct {
		User struct {
			Password string `json:"password,omitempty"`
		} `json:"user,omitempty"`
		NewUser struct {
			UserName string `json:"user_name,omitempty"`
			Email    string `json:"email,omitempty"`
			Password string `json:"password,omitempty"`
		} `json:"new_user,omitempty"`
	}

	var updateUser body
	json.NewDecoder(r.Body).Decode(&updateUser)

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(updateUser.User.Password)); err != nil {
		return err
	}

	if updateUser.NewUser.UserName != "" {
		foundUser.UserName = updateUser.NewUser.UserName
	}
	if updateUser.NewUser.Email != "" {
		foundUser.Email = updateUser.NewUser.Email
	}
	if updateUser.NewUser.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(updateUser.NewUser.Password), 10)
		if err != nil {
			log.Fatal(err)
		}
		foundUser.Password = string(hash)
	}

	db, err := db.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	if err := db.Save(&foundUser).Error; err != nil {
		return err
	}

	return nil
}
