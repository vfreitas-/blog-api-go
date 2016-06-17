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

    conn := db.GetDB()

    if c.Bind(&post) == nil {
        conn.Create(&post)
    }
    // service.Create(&post)
    c.JSON(200, post)
}

func (ctrl PostController) Update(c *gin.Context) {
    var post models.Post

    conn := db.GetDB()

    id := c.Param("id")
    conn.First(&post, id)

    if c.Bind(&post) == nil {
        conn.Save(&post)
    }

    c.JSON(200, gin.H{
        "status": 1,
        "message": "Post updated with success",
        "post": post,
    })
}

func (ctrl PostController) Delete(c *gin.Context) {
    var post models.Post

    conn := db.GetDB()

    id := c.Param("id")
    conn.First(&post, id)
    conn.Delete(&post)

    c.JSON(200, gin.H{
        "status": 1,
        "message": "Post deleted with success",
    })
}

func (ctrl PostController) List(c *gin.Context) {
    var posts = []models.Post{}

    conn := db.GetDB()
    conn.Find(&posts)

    c.JSON(200, posts)
}
