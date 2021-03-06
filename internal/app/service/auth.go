package service

import (
	"fmt"
	"time"

	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/database"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
	"github.com/golang-jwt/jwt"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Login(data payload.LoginPayload) (payload.AuthModel, error)
	Register(data payload.RegisterPayload) (payload.AuthModel, error)
	ClaimToken(user database.UserEntity) (payload.JWTClaimResult, error)
	Logout(token string) bool
	CheckValidToken(token string, userID uint) bool
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
		return payload.AuthModel{}, commons.ErrInvalidCredential
	}

	return au.GetToken(user)
}

func (au *authService) GetToken(user database.UserEntity) (payload.AuthModel, error) {
	authResult := payload.AuthModel{}
	payload, err := au.ClaimToken(user)
	if err != nil {
		return authResult, commons.ErrJWTGenerate
	}

	authResult.Email = user.Email
	authResult.UserID = user.ID
	authResult.Payload = payload.AuthToken
	authResult.TokenType = "Bearer"
	err = au.opt.Cache.WriteCache(authResult.Payload, fmt.Sprint(authResult.UserID), commons.EXIRED_TOKEN_TIME)
	if err != nil {
		return authResult, commons.ErrCacheConn
	}

	return authResult, nil
}

func (au *authService) ClaimToken(user database.UserEntity) (payload.JWTClaimResult, error) {
	expiredDuration := time.Now().Add(commons.EXIRED_TOKEN_TIME)
	claims := &payload.JWTCustomClaims{
		Email:  user.Email,
		UserID: int(user.ID),
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

func (au *authService) Logout(token string) bool {
	// remove token from redis
	err := au.opt.Cache.DeleteCache(token)

	return err == nil
}

func (au *authService) CheckValidToken(token string, userID uint) bool {
	// check token in redis
	data, err := au.opt.Cache.ReadCache(token)
	result, err := redis.Uint64(data, err)
	if err != nil {
		return false
	}
	return uint(result) == userID
}
