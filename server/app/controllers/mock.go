package controllers

import (
	"index-indicators/server/app/entity"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// ModelsMock test用のmock
type ModelsMock struct{}

// CreateUser mock
func (m *ModelsMock) CreateUser(name string, email string, pass string) (err error) {
	return nil
}

// FindUserByID mock
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

// FindUserByEmail mock
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

// UpdateUser mock
func (m *ModelsMock) UpdateUser(foundUser entity.User) (err error) {
	return nil
}

// DeleteUser mock
func (m *ModelsMock) DeleteUser(id int, pass string) (err error) {
	return nil
}

// GetAllUsers mock
func (m *ModelsMock) GetAllUsers() (users []entity.User, err error) {
	return users, nil
}

// CreateLike mock
func (m *ModelsMock) CreateLike(user entity.User, symbol string) (err error) {
	return nil
}

// CheckLikesSymbol mock
func (m *ModelsMock) CheckLikesSymbol(userID int, symbol string) (entity.Like, error) {
	return entity.Like{
		ID:        1,
		UserID:    userID,
		Symbol:    "symbol1",
		CreatedAt: time.Now(),
	}, nil
}

// FetchSymbol mock
func (m *ModelsMock) FetchSymbol(symbol string) (err error) {
	return nil
}

// FindUsersLikes mock
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

// FindLikeByID mock
func (m *ModelsMock) FindLikeByID(likeID int) (entity.Like, error) {
	return entity.Like{
		ID:        1,
		UserID:    1,
		Symbol:    "symbol1",
		CreatedAt: time.Now(),
	}, nil
}

// DeleteLike mock
func (m *ModelsMock) DeleteLike(like entity.Like) (err error) {
	return nil
}

// GetFgis mock
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

// GetTickerAll mock
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

// CreateNewFgis mock
func (m *ModelsMock) CreateNewFgis() error {
	return nil
}

// SaveTickers mock
func (m *ModelsMock) SaveTickers() error {
	return nil
}
