package main

import (
	"fmt"
	"os"
	"test/main/chihuofamingjia/day00007/client/process"
)

var userId string
var userPsw string
var userName string

func main() {
	var key string
	for {
		fmt.Println("----------------欢迎登录多人聊天系统----------------")
		fmt.Println("1、登录聊天室")
		fmt.Println("2、注册用户")
		fmt.Println("3、退出用户")
		fmt.Println("请选择1-3")
		fmt.Scanf("%s\n", &key)
		userProcess := &process.UserProcess{}
		switch key {
		case "1":
			fmt.Println("请输入用户的ID号")
			fmt.Scanf("%s\n", &userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanf("%s\n", &userPsw)

			userProcess.Login(userId, userPsw)

		case "2":
			fmt.Println("请输入注册的ID号")
			fmt.Scanf("%s\n", &userId)
			fmt.Println("请输入注册的密码")
			fmt.Scanf("%s\n", &userPsw)
			fmt.Println("请输入注册的昵称")
			fmt.Scanf("%s\n", &userName)
			userProcess.Register(userId, userPsw, userName)

		case "3":
			fmt.Println("退出系统")
			os.Exit(1)
		default:
			fmt.Println("输入有误")
		}
	}
}
