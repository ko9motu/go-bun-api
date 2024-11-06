package handlers

import (
	"net/http"

	"github.com/ko9motu/go-bun-api/db"
	"github.com/ko9motu/go-bun-api/models"

	"github.com/gin-gonic/gin"
)

/*
	fix me !

- DBの接続効率をいい塩梅にする
*/
func CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	db := db.SetupDB()
	defer db.Close()
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.NewInsert().Model(&user).Exec(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	db := db.SetupDB()
	defer db.Close()
	id := c.Param("id")
	var user models.User

	err := db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
