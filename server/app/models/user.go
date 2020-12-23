package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"index-indicator-apis/server/app/entity"
	"index-indicator-apis/server/mysql"

	"golang.org/x/crypto/bcrypt"
)

// SignupHandler user登録
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("start signup\n")
	fmt.Fprintf(w, "signup")

	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		// エラーハンドリング
		fmt.Println("email error")
		return
	}

	if user.Password == "" {
		// エラーハンドリング
		fmt.Println("pass error")
		return
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Fatal(err)
	}
	user.Password = string(hash)

	db, err := mysql.SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.Create(&user)

	fmt.Printf("%v\n", user)
	fmt.Println("finish! created a user")

}
