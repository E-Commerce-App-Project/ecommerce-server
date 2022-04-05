package database

import "gorm.io/gorm"

type TransactionEntity struct {
	gorm.Model
	TotalPrice int  `gorm:"type:int;not null" json:"total_price"`
	UserID     uint `gorm:"not null" json:"user_id"`

	ItemDetail []TransactionDetailEntity `gorm:"foreignkey:TransactionID" json:"items"`
}
