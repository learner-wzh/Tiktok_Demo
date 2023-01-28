package basic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok_Demo/models"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	fmt.Println(token)
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	user := models.UserLogin(username, password)

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   int64(user.UserID),
		Token:    token,
	})

	fmt.Println(token)
}
