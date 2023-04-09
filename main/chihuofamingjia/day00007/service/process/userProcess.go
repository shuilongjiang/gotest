package process

import (
	"encoding/json"
	"fmt"
	"net"
	"test/main/chihuofamingjia/day00007/common/message"
	"test/main/chihuofamingjia/day00007/service/model"
	"test/main/chihuofamingjia/day00007/service/utils"
)

type UserProcess struct {
	Conn   net.Conn
	userId string
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("ServerProcessRegister--Json--", err)
		return
	}
	err = model.MyUserDao.Register(&registerMes.User)
	var registerResMes message.RegisterResMes

	if err != nil {
		registerResMes.Code = 500
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Error = err.Error()

		}
		fmt.Println("", err)
	} else {
		registerResMes.Code = 200
	}
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("", err)
	}
	var mesRes message.Message
	mesRes.Type = mes.Type
	mesRes.Data = string(data)

	data, err = json.Marshal(mesRes)
	if err != nil {
		fmt.Println("", err)
	}
	//先初始化结构体
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return

}
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("ServerProcessLogin--Json--", err)
		return
	}
	var loginResMes message.LoginResMes
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		loginResMes.Code = 500
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PDW_ERROR {
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Error = "服務器内部錯誤"
		}

	} else {
		loginResMes.Code = 200
		this.userId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		this.NotifyOthersOnlineUser(loginMes.UserId)
		var userId []string
		for id := range userMgr.OnlineUsers {
			userId = append(userId, id)
		}
		loginResMes.UsersId = userId
		fmt.Println(user, "登录校验通过,", userId)
	}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("", err)
	}
	var mesRes message.Message
	mesRes.Type = mes.Type
	mesRes.Data = string(data)

	data, err = json.Marshal(mesRes)
	if err != nil {
		fmt.Println("", err)
	}
	//先初始化结构体
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

func (p *UserProcess) NotifyOthersOnlineUser(userId string) {
	for id, Ou := range userMgr.OnlineUsers {
		if id == userId {
			continue
		}
		Ou.NotifyMeToOtherOnlineUser(userId)
	}
}
func (p *UserProcess) NotifyMeToOtherOnlineUser(userId string) {

	notify := message.NotifyUserStatus{
		UserId:     userId,
		UserStatus: message.UserOnline,
	}
	notifyData, err := json.Marshal(notify)
	if err != nil {
		fmt.Println("notifyData-json-err", err)
		return
	}
	message := message.Message{
		Type: message.NotifyUserStatusType,
		Data: string(notifyData),
	}

	mesData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("mesData-json-err", err)
		return
	}
	tf := &utils.Transfer{
		Conn: p.Conn,
	}
	err = tf.WritePkg(mesData)
	if err != nil {
		fmt.Println("notifyOnline-err,", err)
	}
	return
}
