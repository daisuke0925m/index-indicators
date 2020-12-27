package models

import (
	"fmt"
	"log"
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
