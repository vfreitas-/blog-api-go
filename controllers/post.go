package controllers

import (
    "github.com/gin-gonic/gin"

    "github.com/vfreitas-/blog-api/models"
    "github.com/vfreitas-/blog-api/services"
    "github.com/vfreitas-/blog-api/db"
)

type PostController struct {}

var conn = db.GetDB()

var service = new(services.PostService)

func (ctrl PostController) Create(c *gin.Context) {
    var post models.Post

    c.Bind(&post)
    conn.Create(&post)
    // service.Create(&post)

    c.JSON(200, post)
}
