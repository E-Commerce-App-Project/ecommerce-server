package service

import "github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"

type IHealthCheck interface {
	HealthCheckDbMysql() error
	HealthCheckDbCache() error
}

type healthCheck struct {
	opt Option
}

func NewHealthCheck(opt Option) IHealthCheck {
	return &healthCheck{
		opt: opt,
	}
}

func (h *healthCheck) HealthCheckDbMysql() (err error) {
	d, err := h.opt.DbMysql.DB()
	if err = d.Ping(); err != nil {
		defer d.Close()
		err = commons.ErrDBConn
	}
	return
}

func (h *healthCheck) HealthCheckDbCache() (err error) {
	cacheConn := h.opt.CachePool.Get()
	_, err = cacheConn.Do("PING")
	if err != nil {
		h.opt.Logger.Panicf("Failed to ping cache | %v", err)
		err = commons.ErrCacheConn

		return
	}

	defer cacheConn.Close()

	return
}
