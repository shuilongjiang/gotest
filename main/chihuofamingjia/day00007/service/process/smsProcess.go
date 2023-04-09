package process

import (
	"encoding/json"
	"fmt"
	"net"
	"test/main/chihuofamingjia/day00007/common/message"
	"test/main/chihuofamingjia/day00007/service/utils"
)

type SmsProcess struct {
}

func (p *SmsProcess) SendGroup(mes *message.Message) (err error) {
	var smsMsg message.SmsMsg
	err = json.Unmarshal([]byte(mes.Data), &smsMsg)
	if err != nil {
		fmt.Println("ServerProcessLogin-smsMsg-Json--", err)
		return
	}
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("ServerProcessLogin-data-Json--", err)
		return
	}
	for id, up := range userMgr.OnlineUsers {
		if id == smsMsg.UserId {
			continue
		}
		p.SendEachOnline(up.Conn, data)
	}
	return
}
func (p *SmsProcess) SendEachOnline(conn net.Conn, info []byte) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(info)
	if err != nil {
		fmt.Println("SendEachOnline-转发失败-err,", err)
		return
	}
	return
}
