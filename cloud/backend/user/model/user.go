package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (u *User) TableName() string {
	return "user"
}
