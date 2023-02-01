package router_tiktok

import (
	"github.com/gin-gonic/gin"
	"tiktok_Demo/controllers/basic"
	"tiktok_Demo/controllers/socializing"
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
	apiRouter.POST("/publish/action/", basic.Publish)
	apiRouter.GET("/publish/list/", basic.PublishList)

	// 互动接口
	apiRouter.POST("/favorite/action/")
	apiRouter.GET("/favorite/list/")
	apiRouter.POST("/comment/action/")
	apiRouter.GET("/comment/list/")

	// 社交接口
	apiRouter.POST("/relation/action/")
	apiRouter.GET("/relation/follow/list/", socializing.FollowList)
	apiRouter.GET("/relation/follower/list/", socializing.FollowerList)
	apiRouter.GET("/relation/friend/list/")
	apiRouter.GET("/message/chat/")
	apiRouter.POST("/message/action/")

	return r
}
