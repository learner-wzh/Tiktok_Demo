package socializing

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ChatAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	actionType := c.Query("action_type")
	content := c.Query("content")

	fmt.Println(token)
	fmt.Println(toUserId)
	fmt.Println(actionType)
	fmt.Println(content)
}

func ChatHistory(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")

	fmt.Println(token)
	fmt.Println(toUserId)
}
