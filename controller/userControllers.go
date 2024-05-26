package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"text.com/initializer"
	"text.com/models"
)

func RegisterUser(c *gin.Context) {
	var newUser models.User

	err := c.ShouldBindJSON(&newUser)

	log.Println(newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Got someting Error",
		})
		return

	}

	// Register an User
	result := initializer.DB.Create(&newUser)


	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "account created",
		"user":    newUser,
	})

}
