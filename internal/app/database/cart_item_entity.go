package database

type CartItemEntity struct {
	ID           uint64        `gorm:"primary_key" json:"id"`
	TotalProduct int           `gorm:"not null" json:"total_product"`
	TotalPrice   int           `gorm:"not null" json:"total_price"`
	UserID       uint          `gorm:"not null" json:"user_id"`
	ProductID    uint          `gorm:"not null" json:"product_id"`
	User         UserEntity    `gorm:"foreignkey:UserID" json:"user"`
	Product      ProductEntity `gorm:"foreignkey:ProductID" json:"product"`
}
