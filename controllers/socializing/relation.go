package socializing

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok_Demo/models"
)

type FollowListResponse struct {
	models.Response
	FollowList []models.UserInfo `json:"user_list"` // 用户信息列表
}

type FollowerListResponse struct {
	models.Response
	FollowerList []models.UserInfo `json:"user_list"` // 用户信息列表
}

func FollowList(c *gin.Context) {
	token := c.Query("token")
	// userId := c.Query("user_id")

	userFollowList, err := models.QueryFollowInfoListByToken(token)
	if err == nil {
		fmt.Println(userFollowList)
		c.JSON(http.StatusOK, FollowListResponse{
			Response:   models.Response{StatusCode: 0},
			FollowList: userFollowList,
		})
	}
}

func FollowerList(c *gin.Context) {
	token := c.Query("token")
	// userId := c.Query("user_id")

	userFollowerList, err := models.QueryFollowerInfoListByToken(token)
	if err == nil {
		fmt.Println(userFollowerList)
		c.JSON(http.StatusOK, FollowerListResponse{
			Response:     models.Response{StatusCode: 0},
			FollowerList: userFollowerList,
		})
	}
}
