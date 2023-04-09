package process

import (
	"fmt"
	"test/main/chihuofamingjia/day00007/common/message"
)

var onlineUsers map[string]*message.User = make(map[string]*message.User)

func updateUsersStatus(status *message.NotifyUserStatus) {
	user, ok := onlineUsers[status.UserId]
	if !ok {
		user = &message.User{
			UserId: status.UserId,
		}
	}
	user.UserStatus = status.UserStatus
	onlineUsers[status.UserId] = user
	outPutOnlineUser()
}
func outPutOnlineUser() {
	fmt.Println("当前在线用户列表如下")
	for id, user := range onlineUsers {
		fmt.Println("用户", id, "上线了，状态，", user)
	}

}
