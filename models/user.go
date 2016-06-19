package models

import (
    "github.com/jinzhu/gorm"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    gorm.Model
    Login string `json:"login"`
    Password string `json:"password"`
    Name string `json:"name"`
}

func (u *User) BeforeSave() {
    // hash password
    hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

    if err != nil {
        panic(err)
    }

    u.Password = string(hash)
}
