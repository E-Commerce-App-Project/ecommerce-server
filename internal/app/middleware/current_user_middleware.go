package middleware

import (
	"net/http"

	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/handler"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
	"github.com/labstack/echo/v4"
)

func CurrentUserMiddleware(options handler.HandlerOption) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user")
			if user == nil {
				c.Response().WriteHeader(http.StatusUnauthorized)
				return c.JSON(http.StatusUnauthorized, payload.ResponseFailed(commons.ErrAuthorization.Error()))
			}
			return next(c)
		}
	}
}
