package server

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/handler"
	_middleware "github.com/E-Commerce-App-Project/ecommerce-server/internal/app/middleware"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func intiRouter(e *echo.Echo, opt handler.HandlerOption) (err error) {
	e.HTTPErrorHandler = initErrorHandler(opt)
	healthCheckHandler := handler.HealthCheckHandler{}
	healthCheckHandler.HandlerOption = opt
	authHandler := handler.AuthHandler{}
	authHandler.HandlerOption = opt

	// global middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	e.Use(_middleware.MiddlewareLogging(opt.Options))
	jwtAuthGuard := []echo.MiddlewareFunc{_middleware.JTWAuthMiddleware(opt.Options), _middleware.CurrentUserMiddleware(opt)}
	// register all handler here
	apiV1 := e.Group("/api/v1")
	apiV1.Use(_middleware.MiddlewareUpTime())
	apiV1.GET("/health_check", healthCheckHandler.HealthCheck)
	apiV1.POST("/login", authHandler.Login)
	apiV1.POST("/register", authHandler.Register)
	apiV1.POST("/logout", authHandler.Logout, jwtAuthGuard...)
	return
}

func initErrorHandler(opt handler.HandlerOption) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if ok {
			isApiPath := regexp.MustCompile(`^/api`).MatchString(c.Request().URL.Path)
			message := report.Message
			opt.Logger.Error(fmt.Sprintf("http error %d - %v", report.Code, message))
			if isApiPath {
				if message == "missing or malformed jwt" {
					c.JSON(http.StatusUnauthorized, payload.ResponseFailedWithData("request not authorized", commons.ErrAuthorization))
					return
				}
				c.JSON(report.Code, payload.ResponseFailed(message.(string)))
			} else {
				title := ""
				switch report.Code {
				case http.StatusNotFound:
					title = "Not Found"
					message = "The requested URL was not found on the server. If you entered the URL manually please check your spelling and try again."
				case http.StatusInternalServerError:
					title = "Server Error"
					message = "The server encountered an internal error and was unable to complete your request.  Either the server is overloaded or there is an error in the application."
				default:
					title = "Unkown Error"
					message = "An unknown error occurred."
				}

				c.JSON(report.Code, payload.ResponseFailed(title+" | "+message.(string)))
			}

		} else {
			opt.Logger.Error(err.Error())
			return
		}

		return
	}
}
