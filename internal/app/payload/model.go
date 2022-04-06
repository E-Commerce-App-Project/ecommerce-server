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
	Email string `json:"email"`
	jwt.StandardClaims
}

type JWTClaimResult struct {
	Expired   time.Time
	AuthToken string
}
