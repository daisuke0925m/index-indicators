package controllers

import (
	"index-indicators/server/app/entity"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// ModelsMock test用のmock
type ModelsMock struct{}

func (m *ModelsMock) CreateUser(name string, email string, pass string) (err error) {
	return nil
}
func (m *ModelsMock) FindUserByID(id int) (entity.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte("testpass"), 10)
	if err != nil {
		log.Fatal(err)
	}
	return entity.User{
		ID:       1,
		UserName: "testuser",
		Email:    "test@test",
		Password: string(hash),
	}, nil
}
func (m *ModelsMock) FindUserByEmail(email string) (entity.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte("testpass"), 10)
	if err != nil {
		log.Fatal(err)
	}
	return entity.User{
		ID:       1,
		UserName: "testuser",
		Email:    "test@test",
		Password: string(hash),
	}, nil
}
func (m *ModelsMock) UpdateUser(foundUser entity.User) (err error) {
	return nil
}
func (m *ModelsMock) DeleteUser(id int, pass string) (err error) {
	return nil
}
func (m *ModelsMock) GetAllUsers() (users []entity.User, err error) {
	return users, nil
}
func (m *ModelsMock) CreateLike(user entity.User, symbol string) (err error) {
	return nil
}
func (m *ModelsMock) CheckLikesSymbol(userID int, symbol string) (entity.Like, error) {
	return entity.Like{
		ID:        1,
		UserID:    userID,
		Symbol:    "symbol1",
		CreatedAt: time.Now(),
	}, nil
}
func (m *ModelsMock) FetchSymbol(symbol string) (err error) {
	return nil
}

func (m *ModelsMock) FindUsersLikes(user entity.User) ([]entity.Like, error) {
	likes := []entity.Like{
		{
			ID:        1,
			UserID:    user.ID,
			Symbol:    "symbol1",
			CreatedAt: time.Now(),
		}, {
			ID:        2,
			UserID:    user.ID,
			Symbol:    "symbol2",
			CreatedAt: time.Now(),
		}}
	return likes, nil
}

func (m *ModelsMock) FindLikeByID(likeID int) (entity.Like, error) {
	return entity.Like{
		ID:        1,
		UserID:    1,
		Symbol:    "symbol1",
		CreatedAt: time.Now(),
	}, nil
}
func (m *ModelsMock) DeleteLike(like entity.Like) (err error) {
	return nil
}
func (m *ModelsMock) GetFgis(limit int) []entity.Fgi {
	return []entity.Fgi{{
		ID:        1,
		CreatedAt: time.Now(),
		NowValue:  1,
		NowText:   "fear",
		PcValue:   1,
		PcText:    "fear",
		OneWValue: 1,
		OneWText:  "fear",
		OneMValue: 1,
		OneMText:  "fear",
		OneYValue: 1,
		OneYText:  "fear",
	}}
}
func (m *ModelsMock) GetTickerAll(symbol string) ([]entity.Ticker, error) {
	return []entity.Ticker{{
		ID:        1,
		Symbol:    "symbol",
		Date:      time.Now(),
		Open:      1,
		High:      1,
		Low:       1,
		Close:     1,
		Volume:    1,
		CreatedAt: time.Now(),
	}}, nil
}
func (m *ModelsMock) CreateNewFgis() error {
	return nil
}
func (m *ModelsMock) SaveTickers() error {
	return nil
}
