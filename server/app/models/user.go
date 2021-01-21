package models

import (
	"log"
	"time"

	"index-indicator-apis/server/app/entity"

	"golang.org/x/crypto/bcrypt"
)

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
func (m *Models) FindUserByID(id int) (entity.User, error) {
	var user entity.User
	if err := m.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// UpdateUser userアップデート
func (m *Models) UpdateUser(foundUser entity.User) (err error) {
	if err := m.DB.Save(&foundUser).Error; err != nil {
		return err
	}
	return nil
}
