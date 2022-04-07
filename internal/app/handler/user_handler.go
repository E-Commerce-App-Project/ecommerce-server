package handler

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	HandlerOption
}

func (uh *UserHandler) GetAllHandler(c echo.Context) error {

	users, err := uh.Services.User.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, payload.ResponseFailed("Failed"))
	}
	return c.JSON(http.StatusOK, payload.ResponseSuccess("Success", users))
}

func (uh *UserHandler) GetUserProfile(c echo.Context) error {
	user := c.Get(commons.CTX_USER_KEY).(*jwt.Token)
	claims := user.Claims.(*payload.JWTCustomClaims)
	id := claims.UserID

	users, err := uh.Services.User.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, payload.ResponseFailed("Failed"))
	}
	return c.JSON(http.StatusOK, payload.ResponseSuccess("Success", users))
}

func (uh *UserHandler) CreateUser(c echo.Context) error {

	var user payload.RegisterPayload
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, payload.ResponseFailed("Failed"))
	}
	userModel, err := uh.Services.User.CreateUser(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, payload.ResponseFailed("Failed"))
	}
	return c.JSON(http.StatusOK, payload.ResponseSuccess("Succes create user", userModel))
}

func (uh *UserHandler) DeleteUser(c echo.Context) error {

	user := c.Get(commons.CTX_USER_KEY).(*jwt.Token)
	claims := user.Claims.(*payload.JWTCustomClaims)
	id := claims.UserID

	err := uh.Services.User.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, payload.ResponseFailed("Failed"))
	}
	return c.JSON(http.StatusOK, payload.ResponseSuccessWithoutData("succes Delete User"))
}

func (uh *UserHandler) UpdateUser(c echo.Context) error {

	user := c.Get(commons.CTX_USER_KEY).(*jwt.Token)
	claims := user.Claims.(*payload.JWTCustomClaims)
	id := claims.UserID
	var userPlyd payload.RegisterPayload

	if err := c.Bind(&userPlyd); err != nil {
		return c.JSON(http.StatusInternalServerError, payload.ResponseFailed("Failed"))
	}
	userModel, err := uh.Services.User.UpdateUser(id, userPlyd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, payload.ResponseFailed("Failed"))
	}
	return c.JSON(http.StatusOK, payload.ResponseSuccess("success update user by id", userModel))
}
