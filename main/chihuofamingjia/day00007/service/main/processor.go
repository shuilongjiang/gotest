package main

import (
	"fmt"
	"net"
	"test/main/chihuofamingjia/day00007/common/message"
	process2 "test/main/chihuofamingjia/day00007/service/process"
	"test/main/chihuofamingjia/day00007/service/utils"
)

type Processor struct {
	Conn net.Conn
}

func (this *Processor) ServerProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		userProcess := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = userProcess.ServerProcessLogin(mes)
	case message.RegisterMesType:
		userProcess := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = userProcess.ServerProcessRegister(mes)
	case message.SmsMsgType:
		fmt.Println(mes)
	default:
		fmt.Println("消息类型不存在，无法处理。。。")
	}
	return
}

func (this *Processor) process2() (err error) {
	//读取客户端发送的信息
	for {
		ts := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := ts.ReadPag()
		if err != nil {
			fmt.Println("服务器端推出-redpackage err=", err)
			return err
		}
		fmt.Println(mes)
		err = this.ServerProcessMes(&mes)
		if err != nil {
			fmt.Println("服务器端推出-ServerProcessMes-err=", err)
			return err
		}
	}
}
