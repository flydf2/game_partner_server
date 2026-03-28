package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/gin-gonic/gin"
)

type PlaymateRouter struct{}

// InitPlaymateRouter 初始化playmate路由
func (r *PlaymateRouter) InitPlaymateRouter(router *gin.RouterGroup) {
	playmateRouter := router.Group("/playmates")
	{
		playmateRouter.GET("", api.ApiGroupApp.PlaymateApi.GetPlaymates)
		playmateRouter.GET("/search", api.ApiGroupApp.PlaymateApi.SearchPlaymates)
		playmateRouter.GET("/suggestions", api.ApiGroupApp.PlaymateApi.GetSearchSuggestions)
		playmateRouter.GET("/leaderboard", api.ApiGroupApp.PlaymateApi.GetLeaderboard)
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
		expertRouter.GET("/:id/status", api.ApiGroupApp.PlaymateApi.GetExpertStatus)
	}

	// 技能相关路由
	skillRouter := router.Group("/skills")
	{
		skillRouter.GET("", api.ApiGroupApp.PlaymateApi.GetSkills)
		skillRouter.POST("", api.ApiGroupApp.PlaymateApi.AddSkill)
		skillRouter.PUT("/:id", api.ApiGroupApp.PlaymateApi.UpdateSkill)
		skillRouter.DELETE("/:id", api.ApiGroupApp.PlaymateApi.DeleteSkill)
	}

	// 匹配历史相关路由
	matchHistoryRouter := router.Group("/match-history")
	{
		matchHistoryRouter.GET("", api.ApiGroupApp.PlaymateApi.GetMatchHistory)
		matchHistoryRouter.GET("/matches", api.ApiGroupApp.PlaymateApi.GetMatchHistoryMatches)
		matchHistoryRouter.GET("/:id", api.ApiGroupApp.PlaymateApi.GetMatchHistoryById)
	}
}
