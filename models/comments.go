package models

import (
	"strconv"
	"time"
)

type Comment struct {
	CommentID  int64    `gorm:"column:CommentID" json:"id"`            // 评论id
	VideoID    int64    `gorm:"column:VideoID"`                        // 视频唯一标识
	UserID     int64    `gorm:"column:UserID"`                         // 评论人唯一标识
	CommentStr string   `gorm:"column:CommentStr" json:"content"`      // 评论内容
	CreateDate string   `gorm:"column:CommentDate" json:"create_date"` // 评论发布日期，格式 mm-dd
	User       UserInfo `json:"user" gorm:"-"`                         // 评论用户信息
}

func (value Comment) TableName() string {
	return "Comments"
}

func QueryCommentListCountByVideoID(videoID string) int64 {
	VideoID, _ := strconv.ParseInt(videoID, 10, 64)
	var Count int64
	DB.Table("Comments").Select("COUNT(*)").Where("VideoID = ?", VideoID).Find(&Count)
	return Count
}

func InsertCommentListByVideoID(videoID string, userID string, commentStr string) (Comment, error) {

	Count := QueryCommentListCountByVideoID(videoID)

	VideoID, _ := strconv.ParseInt(videoID, 10, 64)
	UserID, _ := strconv.ParseInt(userID, 10, 64)

	comment := Comment{
		CommentID:  Count + 1,
		VideoID:    VideoID,
		UserID:     UserID,
		CommentStr: commentStr,
		CreateDate: time.Now().Format("01-02"),
	}

	return comment, DB.Create(&comment).Error

}

func QueryCommentListByVideoID(videoID string) ([]Comment, error) {
	VideoID, _ := strconv.ParseInt(videoID, 10, 64)

	var commentList []Comment
	err := DB.Select("CommentID, UserID, CommentStr, CommentDate").Where("VideoID = ?", VideoID).Order("CommentID DESC").Find(&commentList).Error

	for _, video := range commentList {
		userID := video.UserID
		video.User, _ = QueryUserInfoByID(userID)
	}

	return commentList, err
}
