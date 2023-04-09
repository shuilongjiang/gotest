package message

const (
	LoginMesType         = "LoginMes"
	LoginResMesType      = "LoginResMes"
	RegisterMesType      = "RegisterMes"
	RegisterResMesType   = "RegisterResMes"
	NotifyUserStatusType = "NotifyUserStatus"
)
const (
	UserOnline  = 0
	UserOffline = 1
	UserBusy    = 2
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //
}
type LoginMes struct {
	UserId   string `json:"userId"`   //用户ID
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}
type LoginResMes struct {
	Code    int      `json:"code"` //500 用户未注册，2000登录成功
	UsersId []string `json:"usersId"`
	Error   string   `json:"error"` //错误信息

}

type RegisterMes struct {
	//Code  int    `json:"code"`  //500 用户未注册，2000登录成功，
	//Error string `json:"error"` //错误信息
	User User `json:"user"`
}
type RegisterResMes struct {
	Code  int    `json:"code"`  //500 注册失败，200注册成功，
	Error string `json:"error"` //错误信息
}

type NotifyUserStatus struct {
	UserId     string `json:"userId"`
	UserStatus int    `json:"userStatus"`
}
