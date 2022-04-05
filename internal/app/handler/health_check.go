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
	uptime := ctx.Get(commons.API_SERVER_TIME_KEY).(time.Time)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, payload.ResponseFailedWithData("DB is unhealty", err))
		return
	}

	data := payload.HealthCheckModel{
		Status:  "OK",
		Runtime: "go1.16",
		UpTime:  time.Since(uptime).String(),
		Version: "1.0.0",
	}

	ctx.JSON(http.StatusOK, payload.ResponseSuccess("Server is healthy", data))
	return
}
