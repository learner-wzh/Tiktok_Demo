package models

import (
	"strings"
)

type UserInfoModel struct {
	UserID            int64  `gorm:"column:ID"`            // 用户id
	UserName          string `gorm:"column:Name"`          // 用户名称
	UserPwd           string `gorm:"column:Pwd"`           // 用户密码
	UserFollowCount   int64  `gorm:"column:FollowCount"`   // 用户关注数
	UserFollowerCount int64  `gorm:"column:FollowerCount"` // 用户粉丝数
}

func UserSearch(token string) UserModel {

	str := strings.Split(token, "-")

	name := str[0]

	var user UserModel

	DB.Select("ID, FollowCount, FollowerCount").Where("`Name` = ?", name).First(&user)

	return user
}

func QueryUserInfoByID(ID int64) UserInfoModel {
	var userInfo UserInfoModel

	DB.Select("ID, FollowCount, FollowerCount").Where("ID = ?", ID).First(&userInfo)

	return userInfo
}
