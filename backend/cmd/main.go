package main

import (
	"log"

	"github.com/ko9motu/go-bun-api/db"
	"github.com/ko9motu/go-bun-api/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	defer func() {
		if db.Db != nil {
			err := db.Db.Close()
			if err != nil {
				log.Fatalf("Error closing the database connection: %v", err)
			}
		}
	}()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Next()
	})

	r.POST("/users", handlers.CreateUser)
	r.GET("/users/:id", handlers.GetUser)
	r.POST("/users/update/:id", handlers.UpdateUser)
	r.POST("/users/delete/:id", handlers.DeleteUser)

	r.Run(":8080")
}
