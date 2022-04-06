package middleware

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JTWAuthMiddleware(options commons.Options) echo.MiddlewareFunc {
	appConfig := options.Config.GetAppConfig()
	JWTConfig := middleware.JWTConfig{
		Claims:     &payload.JWTCustomClaims{},
		SigningKey: []byte(appConfig.JWTSecret),
	}
	return middleware.JWTWithConfig(JWTConfig)
}
