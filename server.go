package main

import (
    "github.com/gin-gonic/gin"

    "github.com/vfreitas-/blog-api/controllers"
    "github.com/vfreitas-/blog-api/db"
)

func main() {

    db.Init()

    r := gin.Default()

    post := new(controllers.PostController)
    r.POST("/post", post.Create)

    r.Run() // listen and server on 0.0.0.0:8080
}
