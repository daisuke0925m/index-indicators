package entity

import (
	"time"
)

// User user情報
type User struct {
	ID        int       `json:"id,omitempty" gorm:"primaryKey,unique"`
	UserName  string    `json:"user_name,omitempty" gorm:"unique" validate:"required,min=5,max=10,excludesall=!()#@{}"`
	Email     string    `json:"email,omitempty" gorm:"unique" validate:"required"`
	Password  string    `json:"password,omitempty" validate:"required,min=5,max=10,excludesall=!()#@{}"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Likes     []Like    `gorm:"foreignkey:UserID" json:"likes,omitempty"`
}
