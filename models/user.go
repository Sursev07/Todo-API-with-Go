package models

import (
	"todo-API-with-go/helpers"
	"gorm.io/gorm"
)

type User struct {
	OwnModel
	FullName string `gorm:"not null;" json:"full_name" form:"full_name" `
	Email    string `gorm:"not type:"varchar(255)" null;uniqueIndex" json:"email" valid:"required~your email is required,email~invalid email format" `
	Password string `gorm:"not null;" json:"password" form:"password" `
	Todos    []Todo `gorm:"foreignKey:AuthorId"`
}

type ResultErr struct {
	Code    uint         `json:"code"`
	Error 	string      `json:"error"`
	Message string      `json:"message"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}