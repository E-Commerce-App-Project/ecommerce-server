package database

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Name         string              `gorm:"type:varchar(80);not null" json:"name"`
	Email        string              `gorm:"type:varchar(60);not null;unique" json:"email"`
	Password     string              `gorm:"type:varchar(255);not null" json:"password"`
	PhoneNumber  string              `gorm:"type:varchar(20);not null;unique" json:"phone_number"`
	Address      string              `gorm:"type:varchar(255);not null" json:"address"`
	CartItems    []CartItemEntity    `gorm:"foreignkey:UserID" json:"cart_items"`
	Transactions []TransactionEntity `gorm:"foreignkey:UserID" json:"transactions"`
}
