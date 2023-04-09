package process

import (
	"encoding/json"
	"fmt"
	"net"
	"test/main/chihuofamingjia/day00007/client/utils"
	"test/main/chihuofamingjia/day00007/common/message"
)

type UserProcess struct {
}

func (p *UserProcess) Register(userId, userPwd, userName string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	defer conn.Close()
	if err != nil {
		fmt.Println("客户端，建立链接失败", err)
		return err
	}
	registerMes := message.RegisterMes{}
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	loginMesDate, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("loginMes-data-Json序列化失败", err)
		return err
	}

	mess := message.Message{}
	mess.Type = message.RegisterMesType
	mess.Data = string(loginMesDate)
	//fmt.Println(string(loginMesDate))

	date, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("message-Json序列化失败", err)
		return err
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(date)
	if err != nil {
		fmt.Println("客户端发送信息失败", err)
		return err
	}
	mes, err := tf.ReadPag()
	if err != nil {
		fmt.Println("客户端接受返回信息错误，", err)
		return
	}

	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功！")
		//go ServerProcessMes(conn)
		//for {
		//	ShowMenu()
		//}
	} else if registerResMes.Code == 500 {
		fmt.Println("注册失败-----", registerResMes.Error)
	}
	return
}

func (p *UserProcess) Login(userId string, userPwd string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8889")
	defer conn.Close()
	if err != nil {
		fmt.Println("客户端，建立链接失败", err)
		return err
	}

	loginMes := message.LoginMes{}
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//loginMes.UserName = userPwd
	loginMesDate, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("loginMes-data-Json序列化失败", err)
		return err
	}

	mess := message.Message{}
	mess.Type = message.LoginMesType
	mess.Data = string(loginMesDate)
	//fmt.Println(string(loginMesDate))

	date, err := json.Marshal(mess)
	if err != nil {
		fmt.Println("message-Json序列化失败", err)
		return err
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(date)
	if err != nil {
		fmt.Println("客户端发送信息失败", err)
		return err
	}
	mes, err := tf.ReadPag()
	if err != nil {
		fmt.Println("客户端接受返回信息错误，", err)
		return
	}
	var loginRes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginRes)
	if loginRes.Code == 200 {
		fmt.Println("登录成功！")
		fmt.Println("在线用户列表！")
		fmt.Println("---------------")
		for _, v := range loginRes.UsersId {
			if v == userId {
				continue
			}
			fmt.Println(v)
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println("---------------")
		go ServerProcessMes(conn)
		for {
			ShowMenu()
		}

	} else if loginRes.Code == 500 {
		fmt.Println("登录失败-----", loginRes.Error)
	}
	return
}
