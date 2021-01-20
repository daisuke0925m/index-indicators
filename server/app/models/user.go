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

	"golang.org/x/crypto/bcrypt"
)

// Fetch user fetch
func (m *Models) Fetch(id int) (err error) {
	var user entity.User
	if err := m.DB.First(&user, id).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser user登録
func (m *Models) CreateUser(name, email, pass string) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		log.Fatal(err)
	}
	newUser := &entity.User{
		UserName:  name,
		Email:     email,
		Password:  string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := m.DB.Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}

// DeleteUser user削除
func (m *Models) DeleteUser(id int, pass string) (err error) {
	var user entity.User
	if err := m.DB.First(&user, id).Error; err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return err
	}

	if err := m.DB.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

// FindUserByID idからuserを検索
func (m *Models) FindUserByID(r *http.Request) (entity.User, error) {
	var user entity.User
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return entity.User{}, err
	}

	if err := m.DB.Where("id = ?", id).First(&user).Error; err != nil {
		fmt.Println("error!")
		return user, err
	}
	return user, nil
}

// UpdateUser userアップデート
func (m *Models) UpdateUser(foundUser entity.User, r *http.Request) (err error) {
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

	if err := m.DB.Save(&foundUser).Error; err != nil {
		return err
	}

	return nil
}
