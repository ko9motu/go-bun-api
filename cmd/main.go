package main

import (
	"github.com/ko9motu/go-bun-api/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUser)

	r.Run(":8080")
}
