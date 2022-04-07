package handler

import (
	"net/http"

	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	HandlerOption
}

func (ch CartHandler) GetCart(ctx echo.Context) error {
	user := ctx.Get(commons.CTX_USER_KEY).(*jwt.Token)
	claims := user.Claims.(*payload.JWTCustomClaims)
	id := claims.UserID
	getCartPyload := payload.GetUserCartPayload{}
	getCartPyload.UserID = id
	cart, err := ch.Services.Cart.GetCart(getCartPyload)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, payload.ResponseFailedWithData("Failed to get cart", err))
	}

	return ctx.JSON(http.StatusOK, payload.ResponseSuccess("Successfully get cart", cart))
}
