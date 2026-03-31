package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/gin-gonic/gin"
)

type PlaymateRouter struct{}

// InitPlaymateRouter 初始化playmate路由
func (r *PlaymateRouter) InitPlaymateRouter(router *gin.RouterGroup) {
	playmateRouter := router.Group("/playmates")
	{
		// 不需要认证的路由
		playmateRouter.GET("", api.ApiGroupApp.PlaymateApi.GetPlaymates)
		playmateRouter.GET("/search", api.ApiGroupApp.PlaymateApi.SearchPlaymates)
		playmateRouter.GET("/suggestions", api.ApiGroupApp.PlaymateApi.GetSearchSuggestions)
		playmateRouter.GET("/leaderboard", api.ApiGroupApp.PlaymateApi.GetLeaderboard)
		playmateRouter.GET("/:id", api.ApiGroupApp.PlaymateApi.GetPlaymateById)

		// 需要认证的路由
		authRouter := playmateRouter.Group("/")
		authRouter.Use(middleware.CombinedAuthMiddleware())
		{
			authRouter.POST("", api.ApiGroupApp.PlaymateApi.CreatePlaymate)
			authRouter.PUT("/:id", api.ApiGroupApp.PlaymateApi.UpdatePlaymate)
			authRouter.DELETE("/:id", api.ApiGroupApp.PlaymateApi.DeletePlaymate)
		}
	}

	expertRouter := router.Group("/experts")
	{
		// 不需要认证的路由
		expertRouter.GET("/:id", api.ApiGroupApp.PlaymateApi.GetExpertDetail)
		expertRouter.GET("/:id/reviews", api.ApiGroupApp.ReviewApi.GetExpertReviews)
		expertRouter.GET("/:id/voice", api.ApiGroupApp.PlaymateApi.GetExpertVoice)
		expertRouter.GET("/:id/status", api.ApiGroupApp.PlaymateApi.GetExpertStatus)

		// 需要认证的路由
		authRouter := expertRouter.Group("/")
		authRouter.Use(middleware.CombinedAuthMiddleware())
		{
			authRouter.POST("/:id/follow", api.ApiGroupApp.PlaymateApi.FollowExpert)
			authRouter.DELETE("/:id/follow", api.ApiGroupApp.PlaymateApi.UnfollowExpert)
		}
	}

	// 技能相关路由
	skillRouter := router.Group("/skills")
	{
		// 不需要认证的路由
		skillRouter.GET("", api.ApiGroupApp.PlaymateApi.GetSkills)

		// 需要认证的路由
		authRouter := skillRouter.Group("/")
		authRouter.Use(middleware.CombinedAuthMiddleware())
		{
			authRouter.POST("", api.ApiGroupApp.PlaymateApi.AddSkill)
			authRouter.PUT("/:id", api.ApiGroupApp.PlaymateApi.UpdateSkill)
			authRouter.DELETE("/:id", api.ApiGroupApp.PlaymateApi.DeleteSkill)
		}
	}

	// 匹配历史相关路由 - 需要认证
	matchHistoryRouter := router.Group("/match-history")
	matchHistoryRouter.Use(middleware.CombinedAuthMiddleware())
	{
		matchHistoryRouter.GET("", api.ApiGroupApp.PlaymateApi.GetMatchHistory)
		matchHistoryRouter.GET("/matches", api.ApiGroupApp.PlaymateApi.GetMatchHistoryMatches)
		matchHistoryRouter.GET("/:id", api.ApiGroupApp.PlaymateApi.GetMatchHistoryById)
	}

	// 专家认证相关路由
	verificationRouter := router.Group("/expert-verification")
	{
		// 需要认证的路由
		authRouter := verificationRouter.Group("/")
		authRouter.Use(middleware.CombinedAuthMiddleware())
		{
			authRouter.POST("/apply", api.ApiGroupApp.PlaymateApi.ApplyExpertVerification)
			authRouter.GET("/status", api.ApiGroupApp.PlaymateApi.GetExpertVerificationStatus)
			authRouter.PUT("/:id/handle", api.ApiGroupApp.PlaymateApi.HandleExpertVerification)
			authRouter.POST("/:id/cancel", api.ApiGroupApp.PlaymateApi.CancelExpertVerification)
			authRouter.GET("/my", api.ApiGroupApp.PlaymateApi.GetMyExpertVerification)
			authRouter.GET("", api.ApiGroupApp.PlaymateApi.GetExpertVerificationList)
		}
	}
}
