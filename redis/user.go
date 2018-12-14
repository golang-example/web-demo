package redis

import (
	. "web-demo/log"
	"encoding/json"
	. "web-demo/model/user"
)

//从redis获取用户信息
func GetUserInfoInRedis(userName string) *User {
	if userName == "" {
		return nil
	}
	key := "userinfo_" + userName
	userInfoStr, err := kGet(key)
	if err != nil {
		Log.Error(err.Error())
		return nil
	}
	if userInfoStr == "" {
		return nil
	}
	var user User
	if err := json.Unmarshal([]byte(userInfoStr), &user); err != nil {
		Log.Error(err.Error())
		return nil
	}
	return &user
}

//用户信息保存到redis
func AddUserInfoToRedis(user *User) {
	userJson, err := json.Marshal(user)
	if err != nil {
		Log.Error(err.Error())
		return
	}
	key := "userinfo_" + user.UserName
	//有效期10分钟
	kSetex(key, 600, string(userJson))
}

//删除 redis里用户信息
func DelUserInfoInRedis(userName string) {
	if userName == "" {
		return
	}
	key := "userinfo_" + userName
	if _, err := kDel(key); err != nil {
		Log.Error(err.Error())
	}
}
