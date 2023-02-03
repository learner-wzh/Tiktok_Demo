package basic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok_Demo/models"
	"time"
)

type FeedResponse struct {
	models.Response
	VideoList []models.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {
	//token := c.Query("token")
	//latestTime := c.Query("latest_time")

	video, err := models.ReturnVideoInRand()
	var videoList []models.Video

	videoList = append(videoList, video)

	if err == nil {
		fmt.Println(video)
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  models.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	})
}
