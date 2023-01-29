package models

import (
	"strings"
)

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type UserInfo struct {
	UserID            int64  `gorm:"column:ID"`            // 用户id
	UserName          string `gorm:"column:Name"`          // 用户名称
	UserPwd           string `gorm:"column:Pwd"`           // 用户密码
	UserFollowCount   int64  `gorm:"column:FollowCount"`   // 用户关注数
	UserFollowerCount int64  `gorm:"column:FollowerCount"` // 用户粉丝数
}

func UserSearch(token string) Users {

	str := strings.Split(token, "-")

	name := str[0]

	var user Users

	DB.Select("ID, FollowCount, FollowerCount").Where("`Name` = ?", name).First(&user)

	return user
}
