package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

//定义全局变量pool
var pool *redis.Pool

func init(){
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp","192.168.18.130:6379")
		},
		MaxIdle: 5,
		MaxActive: 5,
		IdleTimeout: 240 * time.Second,
	}
}

func main() {
	conn := pool.Get()
	defer conn.Close()
	value, err := redis.String(conn.Do("get","test"))
	if err != nil {
		fmt.Println("出错", err)
	}

	fmt.Println(value)
}