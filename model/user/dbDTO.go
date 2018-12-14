package user

import "time"

//db user表
type User struct {
	Id           int		`gorm:"column:id"`
	UserName     string		`gorm:"column:user_name"`
	Pwd          string		`gorm:"column:pwd"`
	CreateTime   time.Time	`gorm:"column:create_time"`
}

func (user User) TableName() string {
	//表名
	return "user"
}
