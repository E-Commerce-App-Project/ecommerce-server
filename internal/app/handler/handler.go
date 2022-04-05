package handler

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/service"
)

type HandlerOption struct {
	commons.Options
	*service.Services
}
