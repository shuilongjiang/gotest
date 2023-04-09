package main

import (
	"fmt"
	"net"
	"test/main/chihuofamingjia/day00007/service/model"
	"time"
)

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	fmt.Println("服务器在8889端口进行监听")
	listen, err := net.Listen("tcp", "127.0.0.1:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("服务端监听失败", err)
		return
	}
	for {
		fmt.Println("服务器等待客户端来链接服务器")
		conn, errListen := listen.Accept()
		if errListen != nil {
			fmt.Println("服务端获取链接失败", errListen)
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	//读取客户端发送的信息
	defer conn.Close()
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("服务器的协程出问题了：", err)
		return
	}
}
