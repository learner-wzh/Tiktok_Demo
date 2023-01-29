package basic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok_Demo/json_response"
	"tiktok_Demo/models"
	"time"
)

var DemoVideos = []json_response.Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoUser = json_response.User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {

	id := (int64(2))

	video := models.QueryVideoByVideoId(id)

	Author := models.QueryUserInfoByID(id)

	var AuthorUser = json_response.User{
		Id:            Author.UserID,
		Name:          Author.UserName,
		FollowCount:   Author.UserFollowCount,
		FollowerCount: Author.UserFollowerCount,
		IsFollow:      false,
	}

	var testVideos = []json_response.Video{
		{
			Id:            video.ID,
			Author:        AuthorUser,
			PlayUrl:       video.PlayURL,
			CoverUrl:      video.CoverURL,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
		},
	}

	c.JSON(http.StatusOK, json_response.FeedResponse{
		Response:  json_response.Response{StatusCode: 0},
		VideoList: testVideos,
		NextTime:  time.Now().Unix(),
	})
}
