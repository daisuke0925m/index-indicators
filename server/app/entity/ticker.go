package entity

import "time"

// Ticker type Ticker
type Ticker struct {
	ID        int       `json:"id,omitempty" gorm:"primaryKey,unique"`
	Symbol    string    `json:"symbol,omitempty"`
	Date      time.Time `json:"date,omitempty"`
	Open      float64   `json:"open,omitempty"`
	High      float64   `json:"high,omitempty"`
	Low       float64   `json:"low,omitempty"`
	Close     float64   `json:"close,omitempty"`
	Volume    float64   `json:"volume,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
