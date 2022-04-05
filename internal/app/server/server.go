package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/handler"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/service"
	"github.com/labstack/echo/v4"
)

// IServer interface for server
type IServer interface {
	StartApp()
}

type server struct {
	opt      commons.Options
	services *service.Services
}

func NewServer(opt commons.Options, services *service.Services) IServer {
	return &server{
		opt:      opt,
		services: services,
	}
}

func (s *server) StartApp() {
	e := echo.New()
	appConfig := s.opt.Config.GetAppConfig()
	idleConnectionClosed := make(chan struct{})
	var err error
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		s.opt.Logger.Info("[API] Server is shutting down")

		// We received an interrupt signal, shut down.
		if err := e.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			s.opt.Logger.Infof("[API] Fail to shutting down: %v", err)
		}
		close(idleConnectionClosed)
	}()
	addr := fmt.Sprintf(":%d", appConfig.Port)
	hOpt := handler.HandlerOption{
		Options:  s.opt,
		Services: s.services,
	}

	intiRouter(e, hOpt)
	lock := make(chan error)
	go func(lock chan error) { lock <- e.Start(addr) }(lock)

	err = <-lock
	if err != nil {
		switch err := err.(type) {
		case net.Error:
			s.opt.Logger.Error("Server failed to Start %v", err)
		}
	}

	<-idleConnectionClosed
	s.opt.Logger.Info("App closed")
}
