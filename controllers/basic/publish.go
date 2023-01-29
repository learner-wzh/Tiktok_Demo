package basic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok_Demo/json_response"
)

func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, json_response.VideoListResponse{
		Response: json_response.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
