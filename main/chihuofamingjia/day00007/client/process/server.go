package process

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"test/main/chihuofamingjia/day00007/common/message"
	"test/main/chihuofamingjia/day00007/service/utils"
)

func ShowMenu() {
	fmt.Println("----------恭喜XXX登录成功-----------------")
	fmt.Println("----------1、显示在线用户列表----------")
	fmt.Println("----------2、发送消息----------")
	fmt.Println("----------3、信息列表----------")
	fmt.Println("----------4、退出系统----------")
	fmt.Println("----------请选择1-4----------")
	var key int

	smsProcess := SmsProcess{}
	var content string
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outPutOnlineUser()
	case 2:

		fmt.Println("请输入要发送的内容")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroup(content)
	case 3:
		fmt.Println("")
	case 4:
		fmt.Println("----------退出系统----------")
		os.Exit(0)
	default:
		fmt.Println("----------输入有误，请重新输入----------")
	}

}
func ServerProcessMes(conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客戶端在等待服务器發送的消息")
		mes, err := tf.ReadPag()
		if err != nil {
			fmt.Println("等待读取出错，err=", err)
		}
		fmt.Println("等待读取成功，mes=", mes)
		switch mes.Type {
		case message.NotifyUserStatusType:
			var notifyMes message.NotifyUserStatus
			err = json.Unmarshal([]byte(mes.Data), &notifyMes)
			updateUsersStatus(&notifyMes)
		case message.SmsMsgType:
			OutPutGroupMsg(&mes)
		default:
			fmt.Println("服务器返回位置消息，", mes)
		}
	}
}
