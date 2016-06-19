package services

import (
    "github.com/vfreitas-/blog-api/models"
    "github.com/vfreitas-/blog-api/db"
)

type UserService struct {}

func (serv UserService) Create(u *models.User) {
    conn := db.GetDB()
    conn.Create(&u)
}
