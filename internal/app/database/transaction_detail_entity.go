package database

type TransactionDetailEntity struct {
	ID            uint `gorm:"primary_key" json:"id"`
	TransactionID uint `gorm:"not null" json:"transaction_id"`
	ProductID     uint `gorm:"not null" json:"product_id"`
	Quantity      int  `gorm:"not null" json:"quantity"`
	Price         int  `gorm:"not null" json:"price"`

	Product     ProductEntity     `gorm:"foreignkey:ProductID" json:"product"`
	Transaction TransactionEntity `gorm:"foreignkey:TransactionID" json:"transaction"`
}
