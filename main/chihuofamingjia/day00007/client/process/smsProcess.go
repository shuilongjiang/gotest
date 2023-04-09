package process

import (
	"encoding/json"
	"fmt"
	"test/main/chihuofamingjia/day00007/common/message"
	"test/main/chihuofamingjia/day00007/service/utils"
)

type SmsProcess struct {
}

func (p *SmsProcess) SendGroup(content string) (err error) {
	msg := message.Message{
		Type: message.SmsMsgType,
	}
	var smsMsg message.SmsMsg
	smsMsg.Content = content
	smsMsg.UserId = curUser.UserId
	smsMsgData, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("SendGroup-json-smsMsg-err,", err)
		return
	}
	msg.Data = string(smsMsgData)
	msgData, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("SendGroup-json-msgData-err,", err)
		return
	}
	tf := &utils.Transfer{
		Conn: curUser.Conn,
	}
	err = tf.WritePkg(msgData)
	if err != nil {
		fmt.Println("SendGroup-WritePkg-err,", err)
		return
	}
	return
}
