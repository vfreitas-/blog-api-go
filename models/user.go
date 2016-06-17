package models

import (
    "github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Login string `json:"login"`
    Password string `json:"password"`
    Name string `json:"name"`
}

func (u *User) BeforeSave() {
    // hash password
}
