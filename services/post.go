package services

import (
    "github.com/vfreitas-/blog-api/models"
    "github.com/vfreitas-/blog-api/db"
)

type PostService struct {}

func (serv PostService) Create(p *models.Post) {
    conn := db.GetDB()
    conn.Create(&p)
}

func (serv PostService) Update(id string, p *models.Post) {
    conn := db.GetDB()

    conn.Where("id = ?", id).Save(&p)
}

func (serv PostService) Delete(id string) {
    var p models.Post
    conn := db.GetDB()

    conn.Where("id = ?", id).Delete(&p)
}

func (serv PostService) List() ([]models.Post) {
    var p = []models.Post{}
    conn := db.GetDB()

    conn.Find(&p)

    return p
}
