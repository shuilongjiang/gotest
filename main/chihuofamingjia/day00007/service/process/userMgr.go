package process

import "fmt"

var (
	userMgr *UserMgr
)

type UserMgr struct {
	OnlineUsers map[string]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		OnlineUsers: make(map[string]*UserProcess, 1024),
	}
}
func (p *UserMgr) AddOnlineUser(up *UserProcess) {
	p.OnlineUsers[up.userId] = up
}

func (p *UserMgr) DeleteOnlineUser(userId string) {
	delete(p.OnlineUsers, userId)
}

func (p *UserMgr) GetAllOnlineUsers() map[string]*UserProcess {
	return p.OnlineUsers
}

func (p *UserMgr) GetOnlineUserById(userId string) (up *UserProcess, err error) {
	up, ok := p.OnlineUsers[userId]
	if !ok {
		err = fmt.Errorf("用户%d,不在线", userId)
		return
	}
	return
}
