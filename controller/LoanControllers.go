package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"text.com/initializer"
	"text.com/models"
)

func RequestLoan(c *gin.Context) {
	var newLoan models.Loan
	// Req Loan
	log.Println("Request The Loan")

	err := c.ShouldBindJSON(&newLoan)
	if err != nil {

		var message models.SendMessage = models.SendMessage{

			StatusCode: http.StatusBadRequest,
			Message:    "Loan Does not approve",
		}

		log.Println(err)

		c.JSON(http.StatusBadRequest, message)
		return
	}

	var user models.User
	if result := initializer.DB.First(&user, newLoan.UserID); result.Error != nil {
		var message models.SendMessage = models.SendMessage{
			StatusCode: http.StatusBadRequest,
			Message:    "User not found",
		}
		log.Println(result.Error)
		c.JSON(http.StatusBadRequest, message)
		return
	}

	// Associate the fetched user with the new loan
	newLoan.User = user
	newLoan.Paid = false // Ensure the Paid field defaults to false

	result := initializer.DB.Create(&newLoan)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{

			"message": "Loan not get approved",
		})
		return
	}

	// Approved Loan
	c.JSON(http.StatusAccepted, newLoan)
	log.Println("Loan Approved user : " + newLoan.User.Name + ", id : " + strconv.Itoa(int(user.ID)))

}

func GetAllLoans(c *gin.Context) {
	fmt.Println("Get All The Loan Details")

	var loans []models.Loan
	if result := initializer.DB.Find(&loans); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, loans)
}

func PayLoan(c *gin.Context) {
	fmt.Println("Pay The Loan")

	id := c.Param("id")
	var loan models.Loan
	if result := initializer.DB.First(&loan, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	// Assuming you have a Paid field in the Loan model
	loan.Paid = true // Or whatever logic is appropriate
	if result := initializer.DB.Save(&loan); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, loan)
}

func DeleteLoan(c *gin.Context) {
	fmt.Println("Loan Complete")

	id := c.Param("id")
	if result := initializer.DB.Delete(&models.Loan{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan successfully deleted"})
}

func GetLoanInfo(c *gin.Context) {
	fmt.Println("Get Loan Info")

	id := c.Param("id")
	var loan models.Loan
	if result := initializer.DB.First(&loan, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Loan not found"})
		return
	}

	c.JSON(http.StatusOK, loan)
}
