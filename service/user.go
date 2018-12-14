package service

import (
	"web-demo/db"
	"web-demo/redis"
	."web-demo/model/user"
)

//从redis取用户信息, 不存在则从mysql数据库取,然后保存到redis
func GetUserInfo(userName string) *User {
	user := redis.GetUserInfoInRedis(userName)
	if user == nil {
		user = db.SelectUserInfo(userName)
		if user != nil && user.UserName != "" {
			redis.AddUserInfoToRedis(user)
		}
	}
	return user
}

//修改数据库用户信息,并删除redis缓存
func UpdateUserInfo(user User) bool {
	succ := db.UpdateUserInfo(user)
	if  succ {
		//redis 删除信息
		redis.DelUserInfoInRedis(user.UserName)
		return true
	}
	return false
}

//添加用户
func AddUserInfo(userName, pwd string) bool {
	succ := db.InsertUser(userName, pwd)
	if  succ {
		return true
	}
	return false
}
