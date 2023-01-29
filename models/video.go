package models

// Video 数据结构
type Video struct {
	ID            int64  `gorm:"column:ID"`            // 视频唯一标识
	Title         string `gorm:"column:Title"`         // 视频标题
	CommentCount  int64  `gorm:"column:CommentCount"`  // 视频的评论总数
	FavoriteCount int64  `gorm:"column:FavoriteCount"` // 视频的点赞总数
	CoverURL      string `gorm:"column:CoverURL"`      // 视频封面地址
	PlayURL       string `gorm:"column:PlayURL"`       // 视频播放地址
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
	DB.Select("ID", "UserID", "Title", "CommentCount", "FavoriteCount", "CoverURL", "PlayURL").Where("ID=?", videoId).First(&video)
	return video
}

// QueryVideoCountByUserId API：根据作者唯一标识和查找相关视频个数
func QueryVideoCountByUserId(userId int64, count *int64) error {
	return DB.Model(&Video{}).Where("UserID=?", userId).Count(count).Error
}

// QueryVideoListByUserId API：根据作者唯一标识查找返回相关视频列表
func QueryVideoListByUserId(userId int64) []Video {
	var videoList []Video
	DB.Where("UserID=?", userId).
		Select([]string{"ID", "UserID", "Title", "CommentCount", "FavoriteCount", "CoverURL", "PlayURL"}).
		Find(&videoList)
	return videoList
}
