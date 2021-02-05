package models

import (
	"fmt"
	"index-indicator-apis/server/app/entity"
	"time"
)

// CreateLike like作成
func (m *Models) CreateLike(userID int, symbol string) (err error) {
	newLike := &entity.Like{
		UserID:    userID,
		Symbol:    symbol,
		CreatedAt: time.Now(),
	}

	if err := m.DB.Create(&newLike).Error; err != nil {
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
