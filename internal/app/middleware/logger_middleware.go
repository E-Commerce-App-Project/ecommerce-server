package middleware

import (
	"time"

	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/logger"
	"github.com/labstack/echo/v4"
)

func MiddlewareLogging(options commons.Options) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			c.Set("start", start)
			options.Logger = logger.NewLogger(c)
			res := next(c)
			options.Logger.Info("Incoming request")
			return res
		}
	}
}
