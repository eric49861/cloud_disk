package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID     string `json:"uuid" gorm:"column:uuid;not null;unique"`
	Name     string `json:"name" gorm:"column:name;not null;unique"`
	Password string `json:"password" gorm:"column:password;not null;"`
	Email    string `json:"email" gorm:"column:email;"`
}

func (*User) TableName() string {
	return "user"
}
