package router_tiktok

import (
	"github.com/gin-gonic/gin"
	"tiktok_Demo/controllers/basic"
	"tiktok_Demo/models"
)

func InitRouter() *gin.Engine {
	models.InitDB()

	r := gin.Default()

	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// 基础接口
	apiRouter.GET("/feed/", basic.Feed)
	apiRouter.GET("/user/", basic.UserInfo)
	apiRouter.POST("/user/register/", basic.Register)
	apiRouter.POST("/user/login/", basic.Login)
	apiRouter.POST("/publish/action/")
	apiRouter.GET("/publish/list/", basic.PublishList)

	// 互动接口
	apiRouter.POST("/favorite/action/")
	apiRouter.GET("/favorite/list/")
	apiRouter.POST("/comment/action/")
	apiRouter.GET("/comment/list/")

	// 社交接口
	apiRouter.POST("/relation/action/")
	apiRouter.GET("/relation/follow/list/")
	apiRouter.GET("/relation/follower/list/")
	apiRouter.GET("/relation/friend/list/")
	apiRouter.GET("/message/chat/")
	apiRouter.POST("/message/action/")

	return r
}
