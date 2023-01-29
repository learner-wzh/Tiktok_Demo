package basic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok_Demo/json_response"
	"tiktok_Demo/models"
)

// 用户注册API：/douyin/user/register/
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	mes, userIdSequence := models.UserRegister(username, password)

	if mes {
		c.JSON(http.StatusOK, json_response.UserLoginResponse{
			Response: json_response.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	} else {
		c.JSON(http.StatusOK, json_response.UserLoginResponse{
			Response: json_response.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	}

}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + "-" + password

	user := models.UserLogin(username, password)

	mes := user.UserName + "-" + user.UserPwd

	if token == mes {
		c.JSON(http.StatusOK, json_response.UserLoginResponse{
			Response: json_response.Response{StatusCode: 0},
			UserId:   user.UserID,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, json_response.UserLoginResponse{
			Response: json_response.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	user := models.UserSearch(token)

	c.JSON(http.StatusOK, json_response.UserResponse{
		Response: json_response.Response{StatusCode: 0},
		User:     json_response.User{Id: user.UserID, Name: user.UserName, FollowCount: user.UserFollowCount, FollowerCount: user.UserFollowerCount, IsFollow: true},
	})
}
