package models

import "fmt"

type Users struct {
	UserID            uint   `gorm:"primaryKey"`           // 用户id
	UserName          string `gorm:"column:Name"`          // 用户名称
	UserPwd           string `gorm:"column:Pwd"`           // 用户密码
	UserFollowCount   int64  `gorm:"column:FollowCount"`   // 用户关注数
	UserFollowerCount int64  `gorm:"column:FollowerCount"` // 用户粉丝数
}

func (value Users) TableName() string {
	return "users"
}

func UserLogin(Name string, Pwd string) Users {

	var user Users

	DB.Select("ID, `Name`, Pwd").Where("`Name` = ? AND Pwd = ?", Name, Pwd).Take(&user)

	fmt.Println(user.UserID)
	fmt.Println(user.UserName)
	fmt.Println(user.UserPwd)

	return user
}
