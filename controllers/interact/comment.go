package interact

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"tiktok_Demo/models"
)

type CommentListResponse struct {
	models.Response
	CommentList []models.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	models.Response
	Comment models.Comment `json:"comment,omitempty"`
}

func CommentAction(c *gin.Context) {
	token := c.Query("token")
	videoID := c.Query("video_id")
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	commentID := c.Query("comment_id")

	strArray := strings.Split(token, "-")
	userID := strArray[0]
	UserID, _ := strconv.ParseInt(userID, 10, 64)

	if actionType == "1" { //发布评论
		Comment, err := models.InsertCommentListByVideoID(videoID, userID, commentText)
		if err == nil {
			userInfo, _ := models.QueryUserInfoByID(UserID)
			Comment.User = userInfo
			c.JSON(http.StatusOK, CommentActionResponse{Response: models.Response{StatusCode: 0},
				Comment: Comment})
		}
	} else { // 删除评论
		fmt.Println(commentID)
	}
}

func CommentList(c *gin.Context) {
	//token := c.Query("token")
	videoID := c.Query("video_id")

	commentLit, err := models.QueryCommentListByVideoID(videoID)

	if err == nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    models.Response{StatusCode: 0},
			CommentList: commentLit,
		})
	} else {
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    models.Response{StatusCode: 1, StatusMsg: "评论列表获取失败"},
			CommentList: nil,
		})
	}
}
