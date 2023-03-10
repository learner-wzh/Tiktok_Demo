package basic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok_Demo/models"
)

// 用户注册API：/douyin/user/register/
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	mes, userIdSequence := models.UserRegister(username, password)

	if mes {
		c.JSON(http.StatusOK, models.UserLoginResponse{
			Response: models.Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	} else {
		c.JSON(http.StatusOK, models.UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "User already exist"},
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
		token = strconv.FormatInt(user.UserID, 10) + "-" + user.UserName
		c.JSON(http.StatusOK, models.UserLoginResponse{
			Response: models.Response{StatusCode: 0},
			UserId:   user.UserID,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, models.UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	// token := c.Query("token")
	userId := c.Query("user_id")

	UserID, _ := strconv.ParseInt(userId, 10, 64)

	userInfo, err := models.QueryUserInfoByID(UserID)

	if err == nil {
		c.JSON(http.StatusOK, models.UserResponse{
			Response: models.Response{StatusCode: 0},
			User:     userInfo,
		})
	} else {
		c.JSON(http.StatusOK, models.UserResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}

}
