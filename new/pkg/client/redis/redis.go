package redis

import (
	"github.com/go-redis/redis/v8"
)

type Client struct {
	redis *redis.Client
}

func NewRedisClient(o *Options, stopCh <-chan struct{}) (c *Client, err error) {
	rdb := redis.NewClient(&redis.Options{
		Network:            o.Network,
		Addr:               o.Addr,
		Username:           o.Username,
		Password:           o.Password,
		DB:                 o.DB,
		MaxRetries:         o.MaxRetries,
		MinRetryBackoff:    o.MinRetryBackoff,
		MaxRetryBackoff:    o.MaxRetryBackoff,
		DialTimeout:        o.DialTimeout,
		ReadTimeout:        o.ReadTimeout,
		WriteTimeout:       o.WriteTimeout,
		PoolSize:           o.PoolSize,
		MinIdleConns:       o.MinIdleConns,
		MaxConnAge:         o.MaxConnAge,
		PoolTimeout:        o.PoolTimeout,
		IdleTimeout:        o.IdleTimeout,
		IdleCheckFrequency: o.IdleCheckFrequency,
	})

	go func() {
		<-stopCh
		_ = rdb.Close()
	}()

	return &Client{redis: rdb}, rdb.Ping(context.Background()).Err()

}

func (c *Client) Get() *redis.Client {
	if c == nil {
		return nil
	}
	return c.redis
}
