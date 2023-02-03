package socializing

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok_Demo/models"
	"time"
)

type ChatResponse struct {
	models.Response
	MessageList []models.Message `json:"message_list"`
}

func ChatAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")
	timeStr := time.Now().Format("2006-01-02 15:04:05")

	Count, err := models.QueryMessageListCountByTokenAndUserID(token, toUserId)
	if err == nil {
		models.InsertOwnerMessageList(token, Count, toUserId, content, timeStr)
		models.InsertToUserIDMessageList(token, Count, toUserId, content, timeStr)
	}

}

func ChatHistory(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")

	MessageList, err := models.QueryMessageListByTokenAndUserID(token, toUserId)

	if err == nil {
		c.JSON(http.StatusOK, ChatResponse{
			Response:    models.Response{StatusCode: 0},
			MessageList: MessageList,
		})
	} else {
		c.JSON(http.StatusOK, models.Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
	}

}
