package controllers

import (
    "github.com/gin-gonic/gin"

    "github.com/vfreitas-/blog-api/models"
    "github.com/vfreitas-/blog-api/services"
)

type PostController struct {}

var postService = new(services.PostService)

func (ctrl PostController) Create(c *gin.Context) {
    var post models.Post

    if c.Bind(&post) == nil {
        postService.Create(&post)
    }

    c.JSON(200, post)
}

func (ctrl PostController) Update(c *gin.Context) {
    var post models.Post

    id := c.Param("id")

    if c.Bind(&post) == nil {
        postService.Update(id, &post)
    }

    c.JSON(200, gin.H{
        "status": 1,
        "message": "Post updated with success",
        "post": post,
    })
}

func (ctrl PostController) Delete(c *gin.Context) {
    id := c.Param("id")

    postService.Delete(id)

    c.JSON(200, gin.H{
        "status": 1,
        "message": "Post deleted with success",
    })
}

func (ctrl PostController) List(c *gin.Context) {

    posts := postService.List()

    c.JSON(200, posts)
}
