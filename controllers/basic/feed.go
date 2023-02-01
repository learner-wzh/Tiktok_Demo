package basic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"tiktok_Demo/models"
)

type FeedResponse struct {
	models.Response
	VideoList []models.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	latestTime := c.Query("latest_time")

	fmt.Println(token)
	fmt.Println(latestTime)

	//c.JSON(http.StatusOK, FeedResponse{
	//	Response:  models.Response{StatusCode: 0},
	//	VideoList: ,
	//	NextTime:  time.Now().Unix(),
	//})
}
