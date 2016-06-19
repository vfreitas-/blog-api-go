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

    v1 := r.Group("/v1")
    {
        v1.GET("/posts", post.List)
        v1.POST("/posts", post.Create)
        v1.PUT("/posts/:id", post.Update)
        v1.DELETE("/posts/:id", post.Delete)
    }

    r.Run() // listen and server on 0.0.0.0:8080
}
