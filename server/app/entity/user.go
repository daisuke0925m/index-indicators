package entity

import (
	"time"
)

// User user情報
type User struct {
	ID        int       `json:"id,omitempty" gorm:"primaryKey,unique"`
	UserName  string    `json:"user_name,omitempty" gorm:"unique"`
	Email     string    `json:"email,omitempty" gorm:"unique"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
