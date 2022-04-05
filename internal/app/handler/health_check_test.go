package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
	"time"

	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/service"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	healthHandler HealthCheckHandler
	testServer    *echo.Echo
}

func (suite *Suite) SetupTest() {
	handlerOption := HandlerOption{
		Services: &service.Services{
			HealthCheck: NewHealthCheckServiceMock(),
		},
	}
	suite.healthHandler = HealthCheckHandler{
		HandlerOption: handlerOption,
	}
	suite.testServer = echo.New()
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (suite *Suite) TestHealthCheckHandler_ServerHealty() {
	req := httptest.NewRequest(echo.GET, "/health", nil)
	rec := httptest.NewRecorder()
	ctx := suite.testServer.NewContext(req, rec)
	uptime := time.Now()
	resBody, _ := json.Marshal(payload.ResponseSuccess("Server is healthy", payload.HealthCheckModel{
		Status:  "OK",
		Runtime: "go1.16",
		Version: "1.0.0",
	}))
	ctx.Set(commons.API_SERVER_TIME_KEY, uptime)
	suite.healthHandler.Services.HealthCheck.(*HealthCheckServiceMockTest).On("HealthCheckDbMysql").Return(nil)
	err := suite.healthHandler.HealthCheck(ctx)
	m := regexp.MustCompile("(.*?)\\\"uptime\\\":\\\"[\\wµ.]+\\\"(.*)")
	recBody := m.ReplaceAllString(rec.Body.String(), `${1}"uptime":""$2`)
	suite.NoError(err)
	suite.Equal(http.StatusOK, rec.Code)
	suite.Equal(string(resBody)+"\n", string(recBody))
}

func (suite *Suite) TestHealthCheckHandler_ServerUnhealthy() {
	req := httptest.NewRequest(echo.GET, "/health", nil)
	rec := httptest.NewRecorder()
	ctx := suite.testServer.NewContext(req, rec)
	uptime := time.Now()
	resBody, _ := json.Marshal(payload.ResponseFailedWithData("DB is unhealty", commons.ErrDBConn))
	ctx.Set(commons.API_SERVER_TIME_KEY, uptime)

	suite.healthHandler.Services.HealthCheck.(*HealthCheckServiceMockTest).On("HealthCheckDbMysql").Return(commons.ErrDBConn)

	err := suite.healthHandler.HealthCheck(ctx)
	suite.Error(err)
	suite.Equal(http.StatusInternalServerError, rec.Code)
	suite.Equal(string(resBody)+"\n", rec.Body.String())

}
