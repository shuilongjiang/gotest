package message

type User struct {
	UserId     string `json:"userId"`
	UserPwd    string `json:"userPwd"`
	UserName   string `json:"userName"`
	UserStatus int    `json:"userStatus"` //用户状态
}
