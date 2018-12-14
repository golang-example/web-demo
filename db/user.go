package db

import (
	. "web-demo/log"
	. "web-demo/model/user"
	"time"
)

//根据userName查询用户相关信息
func SelectUserInfo(userName string) *User {
	user := User{}
	Mysql.Where("user_name = ?", userName).Find(&user)
	//Mysql.Table("user").Where("user_name = ?", userName).Find(&user)
	if Mysql.Error != nil {
		Log.Error(Mysql.Error.Error())
		return nil
	}
	return &user
}

//修改用户信息 使用事务
func UpdateUserInfo(user User) bool {
	tx := Mysql.Begin()

	if err := tx.Table("user").Where("user_name = ?", user.UserName).Update(&user).Error; err != nil {
		Log.Error(err.Error())
		return false
	}

	//测试事务回滚
	if (user.UserName == "liang") {
		tx.Rollback()
		return false
	}

	tx.Commit()
	return true
}

//插入用户
func InsertUser(userName, pwd string) bool {
	user := User{UserName: userName, Pwd: pwd, CreateTime:time.Now()}
	db := Mysql.Create(user)
	if db.Error != nil {
		Log.Error(db.Error.Error())
		return false
	}
	if db.RowsAffected == 0 {
		return false
	}
	return true
}
