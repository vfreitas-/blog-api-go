package services

import (
    "github.com/vfreitas-/blog-api/models"
    "github.com/vfreitas-/blog-api/db"
)

var conn = db.GetDB()

type PostService struct {}

func (serv PostService) Create(p *models.Post) {
    conn.Create(p)
}
