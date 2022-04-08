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

func (ch CartHandler) AddToCart(ctx echo.Context) error {
	user := ctx.Get(commons.CTX_USER_KEY).(*jwt.Token)
	claims := user.Claims.(*payload.JWTCustomClaims)
	id := claims.UserID

	addToCartPayload := payload.AddProductToCartPayload{}
	if err := ctx.Bind(&addToCartPayload); err != nil {
		return ctx.JSON(http.StatusBadRequest, payload.ResponseFailedWithData("Failed to bind payload", err))
	}

	addToCartPayload.UserID = id
	cart, err := ch.Services.Cart.AddToCart(addToCartPayload)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, payload.ResponseFailedWithData("Failed to add to cart", err))
	}

	return ctx.JSON(http.StatusOK, payload.ResponseSuccess("Successfully add to cart", cart))
}

func (ch CartHandler) UpdateCart(ctx echo.Context) error {
	user := ctx.Get(commons.CTX_USER_KEY).(*jwt.Token)
	claims := user.Claims.(*payload.JWTCustomClaims)
	id := claims.UserID

	updateCartPayload := payload.UpdateCartPayload{}
	if err := ctx.Bind(&updateCartPayload); err != nil {
		return ctx.JSON(http.StatusBadRequest, payload.ResponseFailedWithData("Failed to bind payload", err))
	}

	updateCartPayload.UserID = id
	cart, err := ch.Services.Cart.UpdateCart(updateCartPayload)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, payload.ResponseFailedWithData("Failed to update cart", err))
	}

	return ctx.JSON(http.StatusOK, payload.ResponseSuccess("Successfully update cart", cart))
}

func (ch CartHandler) DeleteCart(ctx echo.Context) error {
	user := ctx.Get(commons.CTX_USER_KEY).(*jwt.Token)
	claims := user.Claims.(*payload.JWTCustomClaims)
	id := claims.UserID

	deleteCartPayload := payload.DeleteUserCartPayload{}
	if err := ctx.Bind(&deleteCartPayload); err != nil {
		return ctx.JSON(http.StatusBadRequest, payload.ResponseFailedWithData("Failed to bind payload", err))
	}

	deleteCartPayload.UserID = id
	deleted := ch.Services.Cart.DeleteCart(deleteCartPayload)
	if deleted {
		return ctx.JSON(http.StatusOK, payload.ResponseSuccess("Successfully delete cart", deleted))
	}

	return ctx.JSON(http.StatusInternalServerError, payload.ResponseFailedWithData("Failed to delete cart", deleted))
}

func (ch CartHandler) Checkout(ctx echo.Context) error {
	user := ctx.Get(commons.CTX_USER_KEY).(*jwt.Token)
	claims := user.Claims.(*payload.JWTCustomClaims)
	id := claims.UserID

	checkoutPayload := payload.CheckoutPayload{}
	if err := ctx.Bind(&checkoutPayload); err != nil {
		return ctx.JSON(http.StatusBadRequest, payload.ResponseFailedWithData("Failed to bind payload", err))
	}

	checkoutPayload.UserID = id
	checkout, err := ch.Services.Cart.Checkout(checkoutPayload)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, payload.ResponseFailedWithData("Failed to checkout", err))
	}

	return ctx.JSON(http.StatusOK, payload.ResponseSuccess("Successfully checkout", checkout))
}
