package entity

import "time"

// Like type Like
type Like struct {
	ID        int       `json:"id,omitempty" gorm:"primaryKey,unique"`
	UserID    int       `json:"user_id,omitempty"`
	User      User      `gorm:"foreignkey:UserID" json:"user,omitempty"`
	Symbol    string    `json:"symbol,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
