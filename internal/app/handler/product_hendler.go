package handler

import (
	"fmt"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	HandlerOption
}

func (ph *ProductHandler) GetAllProduct(c echo.Context) error {

	products, err := ph.Services.Product.GetAllProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, payload.ResponseFailed("Failed"))
	}
	return c.JSON(http.StatusOK, payload.ResponseSuccess("Success", products))
}

func (ph *ProductHandler) GetProductById(c echo.Context) error {

	var id, _ = strconv.Atoi(c.Param("id"))

	products, err := ph.Services.Product.GetProductById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, payload.ResponseFailed("Failed"))
	}
	return c.JSON(http.StatusOK, payload.ResponseSuccess("Success", products))
}

func (ph *ProductHandler) GetProductByIdUser(c echo.Context) error {

	products := c.Get(commons.CTX_USER_KEY).(*jwt.Token)
	claims := products.Claims.(*payload.JWTCustomClaims)
	UserID := int(claims.UserID)

	product, err := ph.Services.Product.GetProductByIdUser(UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, payload.ResponseFailed("Failed"))
	}
	return c.JSON(http.StatusOK, payload.ResponseSuccess("Success", product))
}

func (ph *ProductHandler) CreateProduct(c echo.Context) error {

	var product payload.CreateProductPayload
	if err := c.Bind(&product); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, payload.ResponseFailed("Failed"))
	}

	products := c.Get(commons.CTX_USER_KEY).(*jwt.Token)
	claims := products.Claims.(*payload.JWTCustomClaims)
	product.UserID = uint(claims.UserID)

	productModel, err := ph.Services.Product.CreateProduct(product)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, payload.ResponseFailed("Failed"))
	}

	return c.JSON(http.StatusOK, payload.ResponseSuccess("Succes create user", productModel))

}

func (ph *ProductHandler) DeleteProduct(c echo.Context) error {

	product := c.Get(commons.CTX_USER_KEY).(*jwt.Token)
	claims := product.Claims.(*payload.JWTCustomClaims)
	userID := claims.UserID
	var id, _ = strconv.Atoi(c.Param("id"))

	err := ph.Services.Product.DeleteProduct(id, userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, payload.ResponseFailed("Failed"))
	}
	return c.JSON(http.StatusOK, payload.ResponseSuccessWithoutData("succes Delete Product"))
}

func (ph *ProductHandler) UpdateProduct(c echo.Context) error {

	product := c.Get(commons.CTX_USER_KEY).(*jwt.Token)
	claims := product.Claims.(*payload.JWTCustomClaims)
	userID := claims.UserID

	var id, _ = strconv.Atoi(c.Param("id"))
	var productplyd payload.CreateProductPayload
	if err := c.Bind(&productplyd); err != nil {
		return c.JSON(http.StatusInternalServerError, payload.ResponseFailed("Failed"))
	}
	productModel, err := ph.Services.Product.UpdateProduct(id, userID, productplyd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, payload.ResponseFailed("Failed"))
	}
	return c.JSON(http.StatusOK, payload.ResponseSuccess("success update product by id", productModel))
}
