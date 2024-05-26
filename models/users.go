package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `form:"name" json:"name" binding:"required"`
	Aadhar      int    `form:"aadhar" json:"aadhar" binding:"required"`
	PhoneNumber int16  `form:"phone_number" json:"phonenumber" binding:"required"`
}
