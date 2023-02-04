package interact

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok_Demo/models"
)

type VideoListResponse struct {
	models.Response
	VideoList []models.Video `json:"video_list"`
}

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoID := c.Query("video_id")
	actionType := c.Query("action_type")

	if actionType == "1" { // 点赞操作
		err := models.PlusOneFavorByVideoId(videoID)
		if err == nil {
			if err := models.InsertVideoInToFavorListByTokenAndVideoID(token, videoID); err == nil {
				c.JSON(http.StatusOK, models.Response{StatusCode: 0})
			} else {
				c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "添加喜欢列表操作失败"})
			}
		}
	} else { // 取消点赞操作
		err := models.MinusOneFavorByVideoId(videoID)
		if err == nil {
			if err := models.DeleteVideoFormFavorListByTokenAndVideoID(token, videoID); err == nil {
				c.JSON(http.StatusOK, models.Response{StatusCode: 0})
			} else {
				c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "删除喜欢列表操作失败"})
			}
		}
	}

}

func FavoriteList(c *gin.Context) {
	token := c.Query("token")

	videoList, err := models.QueryFavorListByToken(token)

	if err == nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: models.Response{
				StatusCode: 0,
			},
			VideoList: videoList,
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: models.Response{
				StatusCode: 1,
				StatusMsg:  "喜欢列表返回失败",
			},
			VideoList: nil,
		})
	}
}
