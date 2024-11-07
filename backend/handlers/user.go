package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ko9motu/go-bun-api/db"
	"github.com/ko9motu/go-bun-api/models"
)

/*
	fix me !

- DBの接続効率をいい塩梅にする
*/
func CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.Db.NewInsert().Model(&user).Exec(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func GetUser(c *gin.Context) {
	ctx := c.Request.Context()
	var user models.User

	id := c.Param("id")

	err := db.Db.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	var user models.User

	id := c.Param("id")

	err := db.Db.NewUpdate().Model(&user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()
	var user models.User

	id := c.Param("id")
	err := db.Db.NewDelete().Model(&user).Where("id = ?", id).Scan(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)

}
