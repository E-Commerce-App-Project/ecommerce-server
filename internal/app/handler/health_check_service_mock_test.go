package handler

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/service"
	"github.com/stretchr/testify/mock"
)

type HealthCheckServiceMockTest struct {
	mock.Mock
}

func NewHealthCheckServiceMock() service.IHealthCheck {
	return &HealthCheckServiceMockTest{}
}

func (hm *HealthCheckServiceMockTest) HealthCheckDbMysql() error {
	args := hm.Called()
	return args.Error(0)
}

func (hm *HealthCheckServiceMockTest) HealthCheckDbCache() error {
	args := hm.Called()
	return args.Error(0)
}
