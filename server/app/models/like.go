package models

import (
	"fmt"
	"index-indicators/server/app/entity"
	"time"
)

// CreateLike like作成
func (m *Models) CreateLike(user entity.User, symbol string) (err error) {
	if err := m.DB.Model(&user).Association("Likes").Append(&entity.Like{Symbol: symbol, CreatedAt: time.Now()}).Error; err != nil {
		return err
	}
	return nil
}

// CheckLikesSymbol 同じシンボルを登録していないかチェック
func (m *Models) CheckLikesSymbol(userID int, symbol string) (entity.Like, error) {
	var like entity.Like
	if err := m.DB.Where("user_id = ?", userID).Where("symbol = ?", symbol).First(&like).Error; err != nil {
		return like, err
	}
	return like, nil
}

// FetchSymbol symbolが有効かチェック
func (m *Models) FetchSymbol(symbol string) (err error) {
	var tickers []entity.Ticker
	nonTickersSymbols := []string{"fgi"}
	checked := false

	if err := m.DB.Model(&entity.Ticker{}).Select("symbol").Group("symbol").Find(&tickers).Error; err != nil {
		return err
	}

	for _, ticker := range tickers {
		if symbol == ticker.Symbol {
			checked = true
		}
	}

	for _, nonTS := range nonTickersSymbols {
		if symbol == nonTS {
			checked = true
		}
	}

	if !checked {
		return fmt.Errorf("symbol is invalid")
	}

	return nil
}

// FindUsersLikes Userの全てのLikesを取得
func (m *Models) FindUsersLikes(user entity.User) ([]entity.Like, error) {
	var likes []entity.Like
	if err := m.DB.Model(&user).Association("Likes").Find(&likes).Error; err != nil {
		return likes, err
	}
	return likes, nil
}

// FindLikeByID IDからLIKEを取得
func (m *Models) FindLikeByID(likeID int) (entity.Like, error) {
	var like entity.Like
	if err := m.DB.Where("id = ?", likeID).Find(&like).Error; err != nil {
		return like, err
	}
	return like, nil
}

// DeleteLike LIKEを削除
func (m *Models) DeleteLike(like entity.Like) (err error) {
	if err := m.DB.Delete(&like).Error; err != nil {
		return err
	}
	return nil
}
