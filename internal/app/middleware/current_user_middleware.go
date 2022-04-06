package middleware

import (
	"net/http"
	"strconv"

	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/handler"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CurrentUserMiddleware(options handler.HandlerOption) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get(commons.CTX_USER_KEY).(*jwt.Token)
			claims := user.Claims.(*payload.JWTCustomClaims)
			options.Logger.Info("Current user: ", claims.StandardClaims.Subject)
			if id, err := strconv.ParseUint(claims.StandardClaims.Subject, 10, 32); err == nil {
				hasLogged := options.Auth.CheckValidToken(user.Raw, uint(id))
				options.Logger.Info("Parsing: ", user.Raw, " has logged: ", hasLogged)
				if !hasLogged {
					c.Response().WriteHeader(http.StatusUnauthorized)
					return c.JSON(http.StatusUnauthorized, payload.ResponseFailed("Unauthorized"))
				}
				return next(c)
			}

			c.Response().WriteHeader(http.StatusUnauthorized)
			return c.JSON(http.StatusUnauthorized, payload.ResponseFailed("Unauthorized"))
		}
	}
}
