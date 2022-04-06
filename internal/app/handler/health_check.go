package handler

import (
	"net/http"
	"time"

	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
	"github.com/labstack/echo/v4"
)

type HealthCheckHandler struct {
	HandlerOption
}

func (h HealthCheckHandler) HealthCheck(ctx echo.Context) (err error) {
	err = h.Services.HealthCheck.HealthCheckDbMysql()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, payload.ResponseFailedWithData("DB is unhealty", err))
		return
	}
	err = h.Services.HealthCheck.HealthCheckDbCache()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, payload.ResponseFailedWithData("Cache is unhealty", err))
		return
	}
	uptime := ctx.Get(commons.API_SERVER_TIME_KEY).(time.Time)
	data := payload.HealthCheckModel{
		Status:  "OK",
		Runtime: "go1.16",
		UpTime:  time.Since(uptime).String(),
		Version: "1.0.0",
	}

	ctx.JSON(http.StatusOK, payload.ResponseSuccess("Server is healthy", data))
	return
}
