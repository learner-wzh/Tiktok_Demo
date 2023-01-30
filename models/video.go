package models

import "strings"

// Video 数据结构
type Video struct {
	VideoID       int64    `gorm:"column:VideoID"json:"id,omitempty"`                       // 视频唯一标识
	UserID        int64    `gorm:"column:UserID"`                                           // 视频作者唯一标识
	Title         string   `gorm:"column:Title"json:"title,omitempty"`                      // 视频标题
	CommentCount  int      `gorm:"column:CommentCount"json:"comment_count,omitempty"`       // 视频的评论总数
	FavoriteCount int      `gorm:"column:FavoriteCount"json:"favorite_count,omitempty"`     // 视频的点赞总数
	CoverURL      string   `gorm:"column:CoverURL" json:"cover_url,omitempty"`              // 视频封面地址
	PlayURL       string   `gorm:"column:PlayURL"json:"play_url" json:"play_url,omitempty"` // 视频播放地址
	IsFavorite    bool     `gorm:"column:IsFavorite"json:"is_favorite,omitempty"`           // 是否喜欢
	Author        UserInfo `json:"author" gorm:"-"`                                         // 视频作者信息
}

func (value Video) TableName() string {
	return "Videos"
}

// AddVideo API：添加视频
func AddVideo(video Video) error {
	return DB.Create(video).Error
}

// QueryVideoByVideoId API：根据视频唯一标识查找视频
func QueryVideoByVideoId(videoId int64) Video {
	var video Video
	DB.Select("VideoID", "UserID", "Title", "CommentCount", "FavoriteCount", "CoverURL", "PlayURL").Where("VideoID=?", videoId).First(&video)
	return video
}

// QueryVideoCountByUserId API：根据作者唯一标识返回作者作品个数
func QueryVideoCountByUserId(userId int64) (int64, error) {
	var count int64
	err := DB.Model(&Video{}).Where("UserID=?", userId).Count(&count).Error
	return count, err
}

func QueryVideoListByToken(token string) []Video {

	strArray := strings.Split(token, "-")
	userId := strArray[0]

	var videoList []Video
	DB.Select("VideoID", "UserID", "Title", "CommentCount", "FavoriteCount", "CoverURL", "PlayURL").Where("UserID=?", userId).Find(&videoList)
	return videoList
}

// QueryVideoListByUserId API：根据作者唯一标识查找返回相关视频列表
func QueryVideoListByUserId(userId int64) ([]Video, error) {
	var videoList []Video
	err := DB.Select("VideoID", "UserID", "Title", "CommentCount", "FavoriteCount", "CoverURL", "PlayURL").Where("UserID=?", userId).Find(&videoList).Error
	return videoList, err
}
