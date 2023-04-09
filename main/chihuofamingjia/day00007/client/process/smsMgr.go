package process

import (
	"encoding/json"
	"fmt"
	"test/main/chihuofamingjia/day00007/common/message"
)

func OutPutGroupMsg(mes *message.Message) {
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(mes.Data), &smsMsg)
	if err != nil {
		fmt.Println("OutPutGroupMsg-smsMsg-Json--", err)
		return
	}
	content := smsMsg.UserId + ":对大家说：" + smsMsg.Content
	fmt.Println("------------群发消息开始---------------")
	fmt.Println(content)
	fmt.Println("------------群发消息结束---------------")
}
