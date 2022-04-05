package database

import "gorm.io/gorm"

type ProductEntity struct {
	gorm.Model
	Name        string `gorm:"type:varchar(80);not null" json:"name"`
	Price       int    `gorm:"type:int;not null" json:"price"`
	Description string `gorm:"type:text;not null" json:"description"`
	Image       string `gorm:"type:varchar(255);not null" json:"image"`
	Stock       int    `gorm:"type:int;not null" json:"stock"`
	UserID      uint   `gorm:"not null" json:"user_id"`

	User UserEntity `gorm:"foreignkey:UserID" json:"user"`
}
