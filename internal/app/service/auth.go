package service

import (
	"fmt"
	"time"

	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/database"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(data payload.LoginPayload) (payload.AuthModel, error)
	Register(data payload.RegisterPayload) (payload.AuthModel, error)
	ClaimToken(user database.UserEntity) (payload.JWTClaimResult, error)
}

type authService struct {
	opt Option
}

func NewAuthService(opt Option) IAuthService {
	return &authService{
		opt: opt,
	}
}

func (au *authService) Login(data payload.LoginPayload) (payload.AuthModel, error) {

	user, err := au.opt.Auth.GetUserByEmail(data.Email)
	if err != nil {
		return payload.AuthModel{}, commons.ErrInvalidCredential
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		fmt.Print(user.Password, " ", data.Password)
		return payload.AuthModel{}, commons.ErrInvalidCredential
	}

	return au.GetToken(user)
}

func (au *authService) GetToken(user database.UserEntity) (payload.AuthModel, error) {
	authResult := payload.AuthModel{}
	payload, err := au.ClaimToken(user)
	if err != nil {
		return authResult, err
	}

	authResult.Email = user.Email
	authResult.UserID = user.ID
	authResult.Payload = payload.AuthToken
	authResult.TokenType = "Bearer"

	if err != nil {
		return authResult, err
	}

	return authResult, nil
}

func (au *authService) ClaimToken(user database.UserEntity) (payload.JWTClaimResult, error) {
	expiredDuration := time.Now().Add(time.Hour * 72)
	claims := &payload.JWTCustomClaims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredDuration.Unix(),
			Subject:   fmt.Sprint(user.ID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	appConfig := au.opt.Config.GetAppConfig()
	tokenString, err := token.SignedString([]byte(appConfig.JWTSecret))

	if err != nil {
		return payload.JWTClaimResult{}, err
	}

	return payload.JWTClaimResult{
		Expired:   expiredDuration,
		AuthToken: tokenString,
	}, nil
}

func (au authService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (au *authService) Register(data payload.RegisterPayload) (payload.AuthModel, error) {
	user := database.UserEntity{}
	hashedPassword, err := au.hashPassword(data.Password)

	if err != nil {
		return payload.AuthModel{}, commons.ErrHashPassword
	}

	userData := database.UserEntity{}
	userData.Email = data.Email
	userData.Password = hashedPassword
	userData.Name = data.Name
	userData.PhoneNumber = data.Phone
	userData.Address = data.Address
	user, err = au.opt.Auth.RegisterUser(userData)
	if err != nil {
		return payload.AuthModel{}, err
	}
	return au.GetToken(user)
}
