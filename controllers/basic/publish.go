package basic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"tiktok_Demo/models"
)

type VideoListResponse struct {
	models.Response
	VideoList []models.Video `json:"video_list"`
}

func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title := c.PostForm("title")

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)

	user, err := models.QueryUserInfoByToken(token)
	if err == nil {
		finalName := fmt.Sprintf("%d_%s", user.UserID, filename)
		saveFile := filepath.Join("/root/video", finalName)

		var video models.Video
		video.PlayURL = saveFile
		video.Title = title
		video.UserID = user.UserID
		models.AddVideo(video)

		if err := c.SaveUploadedFile(data, saveFile); err != nil {
			c.JSON(http.StatusOK, models.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		StatusMsg:  filename + " uploaded successfully",
	})
}

func PublishList(c *gin.Context) {

	// token := c.Query("token")
	userId := c.Query("user_id")

	UserID, _ := strconv.ParseInt(userId, 10, 64)

	videoList, err := models.QueryVideoListByUserId(UserID)

	var userInfo []models.UserInfo

	for _, video := range videoList {
		info, _ := models.QueryUserInfoByID(video.UserID)
		userInfo = append(userInfo, info)
	}

	if err == nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: models.Response{
				StatusCode: 0,
			},
			VideoList: videoList,
		})
	}

}
