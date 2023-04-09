package model

import (
	"net"
	"test/main/chihuofamingjia/day00007/common/message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
