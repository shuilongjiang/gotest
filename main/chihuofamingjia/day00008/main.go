package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	defer func() {
		conn.Close()
		fmt.Println("over")
	}()
	go handle(conn)
	fmt.Println("--2222-")
}
func handle(c redis.Conn) {
	fmt.Println("---")
	c.Do("set", "name1", "vds的城市")
	r, err := redis.String(c.Do("get", "name1"))
	fmt.Println("test1--", r)
	fmt.Println("test2--", err)
}
