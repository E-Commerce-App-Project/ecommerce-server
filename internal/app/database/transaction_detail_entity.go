package database

import "gorm.io/gorm"

type TransactionDetailEntity struct {
	gorm.Model
	TransactionID uint `gorm:"not null" json:"transaction_id"`
	ProductID     uint `gorm:"not null" json:"product_id"`
	Quantity      int  `gorm:"not null" json:"quantity"`
	Price         int  `gorm:"not null" json:"price"`

	Product     ProductEntity     `gorm:"foreignkey:ProductID" json:"product"`
	Transaction TransactionEntity `gorm:"foreignkey:TransactionID" json:"transaction"`
}
