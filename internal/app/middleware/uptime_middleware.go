package middleware

import (
	"time"

	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/labstack/echo/v4"
)

func MiddlewareUpTime() echo.MiddlewareFunc {
	uptime := time.Now()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(commons.API_SERVER_TIME_KEY, uptime)
			return next(c)
		}
	}
}
