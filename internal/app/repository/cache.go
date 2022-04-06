package repository

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/im7mortal/kmutex"
)

// ICacheRepository interface for cache repo
type ICacheRepository interface {
	WriteCache(key string, data interface{}, ttl time.Duration) (err error)
	WriteCacheIfEmpty(key string, data interface{}, ttl time.Duration) (err error)
	ReadCache(key string) (data interface{}, err error)
	DeleteCache(key string) (err error)
}

type cacheRepo struct {
	opt    Option
	kmutex *kmutex.Kmutex
}

// NewCacheRepository initiate cache repo
func NewCacheRepository(opt Option) ICacheRepository {
	return &cacheRepo{
		opt:    opt,
		kmutex: kmutex.New(),
	}
}

// WriteCache this will and must write the data to cache with corresponding key using locking
func (c *cacheRepo) WriteCache(key string, data interface{}, ttl time.Duration) (err error) {
	c.kmutex.Lock(key)
	defer c.kmutex.Unlock(key)

	// write data to cache
	conn := c.opt.CachePool.Get()
	defer conn.Close()

	_, err = conn.Do("SETEX", key, ttl.Seconds(), data)
	if err != nil {
		return err
	}

	return nil
}

// WriteCacheIfEmpty will try to write to cache, if the data still empty after locking
func (c *cacheRepo) WriteCacheIfEmpty(key string, data interface{}, ttl time.Duration) (err error) {
	c.kmutex.Lock(key)
	defer c.kmutex.Unlock(key)

	// check whether cache value is empty
	conn := c.opt.CachePool.Get()
	defer conn.Close()

	_, err = conn.Do("GET", key)
	if err != nil {
		if err == redis.ErrNil {
			return nil //return nil as the data already set, no need to overwrite
		}

		return err
	}

	// write data to cache
	_, err = conn.Do("SETEX", key, ttl.Seconds(), data)
	if err != nil {
		return err
	}

	return nil
}

func (c *cacheRepo) ReadCache(key string) (data interface{}, err error) {
	conn := c.opt.CachePool.Get()
	defer conn.Close()

	data, err = conn.Do("GET", key)
	if err != nil {
		if err == redis.ErrNil {
			return nil, nil
		}

		return nil, err
	}

	return data, nil
}

func (c *cacheRepo) DeleteCache(key string) (err error) {
	conn := c.opt.CachePool.Get()
	defer conn.Close()

	_, err = conn.Do("DEL", key)
	if err != nil {
		return err
	}

	return nil
}
