package models

import (
	"errors"
	"log"
	"strconv"
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

	err := DB.Select("UserID, Name, FollowCount, FollowerCount, IsFollow").Where("UserID = ?", UserID).First(&userInfo).Error

	return userInfo, err
}

func QueryUserInfoByID(ID int64) (UserInfo, error) {
	var userInfo UserInfo

	err := DB.Select("UserID, Name, FollowCount, FollowerCount, IsFollow").Where("UserID = ?", ID).First(&userInfo).Error

	return userInfo, err
}

func QueryUserInfoNameByID(UserID string) (string, error) {
	var userInfo UserInfo

	ID, _ := strconv.ParseInt(UserID, 10, 64)

	err := DB.Select("UserID,Name").Where("UserID = ?", ID).First(&userInfo).Error

	userName := userInfo.UserName

	return userName, err

}

func QueryUserInfoInFollowListByToken(token string, toUserID string) bool {

	followName, _ := QueryUserInfoNameByID(toUserID)

	tableName := toUserID + "-" + followName + "-follow"

	strArray := strings.Split(token, "-")
	ID := strArray[0]
	followID, _ := strconv.ParseInt(ID, 10, 64)

	var followUser UserInfo
	err := DB.Table(tableName).Select("UserID").Where("UserID = ?", followID).Find(&followUser).Error

	if err == nil {
		if followUser.UserID == 0 {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

func QueryUserInfoInFollowerListByToken(token string, toUserID string) bool {
	tableName := token + "-follower"
	followerID, _ := strconv.ParseInt(toUserID, 10, 64)

	var followerUser UserInfo
	err := DB.Table(tableName).Select("UserID").Where("UserID = ?", followerID).Find(&followerUser).Error

	if err == nil {
		if followerUser.UserID == 0 {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

func QueryFollowInfoListByToken(token string) ([]UserInfo, error) {

	tableName := token + "-follow"

	var FollowInfoList []UserInfo

	err := DB.Table(tableName).Select("UserID, Name, FollowCount, FollowerCount, IsFollow").Find(&FollowInfoList).Error

	return FollowInfoList, err
}

func QueryFollowerInfoListByToken(token string) ([]UserInfo, error) {

	tableName := token + "-follower"

	var FollowerInfoList []UserInfo

	err := DB.Table(tableName).Select("UserID, Name, FollowCount, FollowerCount, IsFollow").Find(&FollowerInfoList).Error

	return FollowerInfoList, err
}

func QueryFriendInfoListByToken(token string) ([]UserInfo, error) {
	tableName := token + "-follower"

	var FriendInfoList []UserInfo

	err := DB.Table(tableName).Select("UserID, Name, FollowCount, FollowerCount, IsFollow").Where("IsFollow = 1").Find(&FriendInfoList).Error

	return FriendInfoList, err
}

func InsertFollowerInfoListByToken(token string, toUserID string) error {
	followerInfo, err := QueryUserInfoByToken(token)
	followerInfo.IsFollow = false

	if err == nil {
		followName, _ := QueryUserInfoNameByID(toUserID)

		tableName := toUserID + "-" + followName + "-follower"

		return DB.Table(tableName).Create(followerInfo).Error
	} else {
		return errors.New("插入人员信息查询失败")
	}
}

func InsertFollowInfoListByToken(token string, toUserID string) {
	followID, _ := strconv.ParseInt(toUserID, 10, 64)

	followInfo, err := QueryUserInfoByID(followID)
	followInfo.IsFollow = true

	if err == nil {
		tableName := token + "-follow"

		log.Println(DB.Table(tableName).Create(followInfo).Error)
	} else {
		log.Println(errors.New("插入人员信息查询失败"))
	}
}

func UpdateFollowerInfoListByUserID(token string, toUserID string, isFollow bool) {
	strArray := strings.Split(token, "-")
	ID := strArray[0]
	followID, _ := strconv.ParseInt(ID, 10, 64)

	followName, _ := QueryUserInfoNameByID(toUserID)

	tableName := toUserID + "-" + followName + "-follower"

	var err error

	if isFollow {
		err = DB.Table(tableName).Where("UserID = ?", followID).Update("IsFollow", true).Error
	} else {
		err = DB.Table(tableName).Where("UserID = ?", followID).Update("IsFollow", false).Error
	}

	log.Println(err)
}

func UpdateFollowerInfoListByToken(token string, toUserID string, isFollower bool) {
	followID, _ := strconv.ParseInt(toUserID, 10, 64)
	tableName := token + "-follower"

	var err error

	if isFollower {
		err = DB.Table(tableName).Where("UserID = ?", followID).Update("IsFollow", true).Error
	} else {
		err = DB.Table(tableName).Where("UserID = ?", followID).Update("IsFollow", false).Error
	}

	log.Println(err)
}

func DeleteFollowInfoListByToken(token string, toUserID string) {
	followID, _ := strconv.ParseInt(toUserID, 10, 64)
	tableName := token + "-follow"

	err := DB.Table(tableName).Delete(&UserInfo{}, followID).Error

	log.Println(err)
}

func DeleteFollowerInfoListByToken(token string, toUserID string) error {
	strArray := strings.Split(token, "-")
	ID := strArray[0]
	followerID, _ := strconv.ParseInt(ID, 10, 64)

	followName, _ := QueryUserInfoNameByID(toUserID)

	tableName := toUserID + "-" + followName + "-follower"

	return DB.Table(tableName).Delete(&UserInfo{}, followerID).Error
}
