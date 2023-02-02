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

type FriendListResponse struct {
	models.Response
	FriendList []models.UserInfo `json:"user_list"` // 用户信息列表
}

type UserListResponse struct {
	models.Response
	UserList []models.UserInfo `json:"user_list"`
}

func FollowAction(c *gin.Context) {
	token := c.Query("token")
	toUserID := c.Query("to_user_id")
	actionType := c.Query("action_type")

	fmt.Println(token)
	fmt.Println(toUserID)

	if actionType == "1" {
		// 关注用户
		models.InsertFollowInfoListByToken(token, toUserID)
		err := models.InsertFollowerInfoListByToken(token, toUserID)
		if err == nil {
			isFollow := models.QueryUserInfoInFollowListByToken(token, toUserID)
			models.UpdateFollowerInfoListByUserID(token, toUserID, isFollow)

			isFollower := models.QueryUserInfoInFollowerListByToken(token, toUserID)
			models.UpdateFollowerInfoListByToken(token, toUserID, isFollower)
		}
	} else if actionType == "2" {
		// 取消关注
		fmt.Println("取消关注")
		models.DeleteFollowInfoListByToken(token, toUserID)
		err := models.DeleteFollowerInfoListByToken(token, toUserID)
		if err == nil {
			isFollower := models.QueryUserInfoInFollowerListByToken(token, toUserID)
			models.UpdateFollowerInfoListByToken(token, toUserID, !isFollower)
		}
	}
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

func FriendList(c *gin.Context) {
	token := c.Query("token")
	// userId := c.Query("user_id")

	fmt.Println("friendList")

	userFriendList, err := models.QueryFriendInfoListByToken(token)

	fmt.Println(userFriendList)

	if err == nil {
		c.JSON(http.StatusOK, FriendListResponse{
			Response: models.Response{
				StatusCode: 0,
			},
			FriendList: userFriendList,
		})
	}
}
