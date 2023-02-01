package models

import (
	"strings"
)

type UserInfo struct {
	UserID            int64  `gorm:"column:UserID" json:"id,omitempty"`                   // 用户id
	UserName          string `gorm:"column:Name"json:"name,omitempty"`                    // 用户名称
	UserFollowCount   int    `gorm:"column:FollowCount"json:"follow_count,omitempty"`     // 用户关注数
	UserFollowerCount int    `gorm:"column:FollowerCount"json:"follower_count,omitempty"` // 用户粉丝数
	IsFollow          bool   `gorm:"column:IsFollow"json:"is_follow,omitempty"`           // 是否关注
}

func (value UserInfo) TableName() string {
	return "UsersInfo"
}

func QueryUserInfoByToken(token string) (UserInfo, error) {

	str := strings.Split(token, "-")

	UserID := str[0]

	var userInfo UserInfo

	err := DB.Select("UserID, FollowCount, FollowerCount, IsFollow").Where("UserID = ?", UserID).First(&userInfo).Error

	return userInfo, err
}

func QueryUserInfoByID(ID int64) (UserInfo, error) {
	var userInfo UserInfo

	err := DB.Select("UserID, FollowCount, FollowerCount, IsFollow").Where("UserID = ?", ID).First(&userInfo).Error

	return userInfo, err
}

func QueryFollowInfoListByToken(token string) ([]UserInfo, error) {

	tableName := token + "-follow"

	var FollowInfoList []UserInfo

	err := DB.Table(tableName).Select("UserID, FollowCount, FollowerCount, IsFollow").Find(&FollowInfoList).Error

	return FollowInfoList, err
}

func QueryFollowerInfoListByToken(token string) ([]UserInfo, error) {

	tableName := token + "-follower"

	var FollowerInfoList []UserInfo

	err := DB.Table(tableName).Select("UserID, FollowCount, FollowerCount, IsFollow").Find(&FollowerInfoList).Error

	return FollowerInfoList, err
}
