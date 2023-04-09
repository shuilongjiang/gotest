package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	fmt.Println("redisPool--初始化")
	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
