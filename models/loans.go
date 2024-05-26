package models

import "gorm.io/gorm"

type Loan struct {
	gorm.Model
	UserID uint   `json:"user_id" binding:"required"`    // Foreign key field
	User   User   `gorm:"foreignKey:UserID" json:"user"` // Define the foreign key
	Loan   int32  `json:"loan" binding:"required"`
	Period string `json:"period" binding:"required"`
	Rate   int16  `json:"rate" binding:"required"`
	Paid   bool   `json:"paid" gorm:"default:false" `
}
