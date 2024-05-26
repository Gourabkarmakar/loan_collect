package initializer

import (
	"log"

	"text.com/models"
)

func SyscDb() {

	// Sync Database with User and Loan models
	err := DB.AutoMigrate(&models.User{}, &models.Loan{})

	// Check if Got any error
	if err != nil {
		log.Println("Error syncing database:", err)
	} else {
		log.Println("Database synced successfully with User and Loan models")
	}

}
