package basic

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	latestTime := c.Query("latest_time")

	fmt.Println(token)
	fmt.Println(latestTime)
}
