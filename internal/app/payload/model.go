package payload

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type HealthCheckModel struct {
	Status  string `json:"status" example:"success"`
	Runtime string `json:"runtime" example:"go1.13.4"`
	UpTime  string `json:"uptime" example:"1h"`
	Version string `json:"version" example:"1.0.0"`
} //@name HealthCheckModel

type AuthModel struct {
	UserID    uint   `json:"user_id" example:"1"`
	Email     string `json:"email" example:"foo@bar.com"`
	Payload   string `json:"payload" example:"Bearer"`
	TokenType string `json:"token_type" example:"aa3f97cec3f342bc9b11e0592bbce319"`
} //@name AuthModel

type JWTCustomClaims struct {
	Email  string `json:"email"`
	UserID int    `json:"user_id"`

	jwt.StandardClaims
}

type JWTClaimResult struct {
	Expired   time.Time
	AuthToken string
}

type CartModel struct {
	ID        uint   `json:"id" example:"1"`
	UserID    uint   `json:"user_id" example:"1"`
	ProductID uint   `json:"product_id" example:"1"`
	Image     string `json:"image" example:"https://example.com/image.jpg"`
	Name      string `json:"name" example:"Product Name"`
	Price     int    `json:"price" example:"10000"`
	Quantity  uint   `json:"quantity" example:"1"`
}

type UserModel struct {
	UserID      uint   `json:"user_id" example:"1"`
	Name        string `json:"name" example:"Budi"`
	Email       string `json:"email" example:"foo@bar.com"`
	PhoneNumber string `json:"phone_number" example:"089123123"`
	Address     string `json:"address" example:"Bandung"`
}

type ProductModel struct {
	Name        string `json:"name" example:"sepatu"`
	Price       int    `json:"price" example:"500000"`
	Description string `json:"description" example:"Sepatu lari"`
	Image       string `json:"image" example:"image"`
	Stock       int    `json:"stock" example:"10"`
	UserID      uint   `json:"user_id" example:"1"`
}
