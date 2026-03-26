package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
)

type PlaymateRouter struct{}

// InitPlaymateRouter 初始化playmate路由
func (r *PlaymateRouter) InitPlaymateRouter(router *gin.RouterGroup) {
	playmateRouter := router.Group("/playmates")
	{
		playmateRouter.GET("", api.ApiGroupApp.PlaymateApi.GetPlaymates)
		playmateRouter.GET("/search", api.ApiGroupApp.PlaymateApi.SearchPlaymates)
		playmateRouter.GET("/suggestions", api.ApiGroupApp.PlaymateApi.GetSearchSuggestions)
		playmateRouter.GET("/:id", api.ApiGroupApp.PlaymateApi.GetPlaymateById)
		playmateRouter.POST("", api.ApiGroupApp.PlaymateApi.CreatePlaymate)
		playmateRouter.PUT("/:id", api.ApiGroupApp.PlaymateApi.UpdatePlaymate)
		playmateRouter.DELETE("/:id", api.ApiGroupApp.PlaymateApi.DeletePlaymate)
	}

	expertRouter := router.Group("/experts")
	{
		expertRouter.GET("/:id", api.ApiGroupApp.PlaymateApi.GetExpertDetail)
		expertRouter.POST("/:id/follow", api.ApiGroupApp.PlaymateApi.FollowExpert)
		expertRouter.DELETE("/:id/follow", api.ApiGroupApp.PlaymateApi.UnfollowExpert)
		expertRouter.GET("/:id/reviews", api.ApiGroupApp.ReviewApi.GetExpertReviews)
		expertRouter.GET("/:id/voice", api.ApiGroupApp.PlaymateApi.GetExpertVoice)
	}
}
