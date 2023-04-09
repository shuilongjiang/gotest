package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"test/main/chihuofamingjia/day00007/common/message"
)

var (
	MyUserDao *UserDao
)

type UserDao struct {
	Pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		Pool: pool,
	}
	return
}

func (p *UserDao) getUserById(conn redis.Conn, userId string) (user *message.User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", userId))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("JSON反序列化出错，err=", err)
		return
	}
	return
}
func (p *UserDao) Login(userId string, userPwd string) (user *message.User, err error) {
	conn := p.Pool.Get()
	defer conn.Close()
	user, err = p.getUserById(conn, userId)
	if err != nil {
		return
	}
	if user.UserPwd != userPwd {
		err = ERROR_USER_PDW_ERROR
		return
	}
	return
}

func (p *UserDao) Register(user *message.User) (err error) {
	conn := p.Pool.Get()
	defer conn.Close()
	_, err = p.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	err = p.insertUser(conn, user.UserId, string(data))
	if err != nil {
		fmt.Println("redis--注册用户出错")
		return
	}
	return
}

func (p *UserDao) insertUser(conn redis.Conn, userId, users string) (err error) {
	_, err = conn.Do("HSet", "users", userId, users)
	if err != nil {
		fmt.Println("redis--注册用户出错")
		return
	}
	return
}
