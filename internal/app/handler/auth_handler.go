package handler

import (
	"net/http"

	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	HandlerOption
}

func (ah AuthHandler) Login(ctx echo.Context) (err error) {
	pyld := payload.LoginPayload{}

	if err = ctx.Bind(&pyld); err != nil {
		return ctx.JSON(http.StatusBadRequest, payload.ResponseFailedWithData("Invalid payload", commons.ErrParsingBody))
	}

	authResult, err := ah.Services.Auth.Login(pyld)

	if err != nil {
		if err == commons.ErrInvalidCredential {
			return ctx.JSON(http.StatusBadRequest, payload.ResponseFailedWithData("Invalid email or password", err))
		}
		return ctx.JSON(http.StatusInternalServerError, payload.ResponseFailedWithData("Internal server error", err))
	}
	return ctx.JSON(http.StatusOK, payload.ResponseSuccess("Login success", authResult))
}

func (ah AuthHandler) Register(ctx echo.Context) (err error) {
	pyld := payload.RegisterPayload{}

	if err = ctx.Bind(&pyld); err != nil {
		return ctx.JSON(http.StatusBadRequest, payload.ResponseFailedWithData("Invalid payload", commons.ErrParsingBody))
	}

	authResult, err := ah.Services.Auth.Register(pyld)

	if err != nil {
		if err == commons.ErrEmailExists {
			return ctx.JSON(http.StatusBadRequest, payload.ResponseFailedWithData("Email already exists", err))
		}
		if err == commons.ErrPhoneExistError {
			return ctx.JSON(http.StatusBadRequest, payload.ResponseFailedWithData("Phone already exists", err))
		}
		return ctx.JSON(http.StatusInternalServerError, payload.ResponseFailedWithData("Internal server error", err))
	}
	return ctx.JSON(http.StatusOK, payload.ResponseSuccess("Register success", authResult))
}
