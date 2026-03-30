package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/gin-gonic/gin"
)

type LeaderboardRouter struct{}

// InitLeaderboardRouter 初始化排行榜路由
func (r *LeaderboardRouter) InitLeaderboardRouter(router *gin.RouterGroup) {
	leaderboardRouter := router.Group("/leaderboards")
	{
		// 不需要认证的路由
		leaderboardRouter.GET("", api.ApiGroupApp.LeaderboardApi.GetLeaderboards)
		leaderboardRouter.GET("/:id", api.ApiGroupApp.LeaderboardApi.GetLeaderboardById)
		leaderboardRouter.GET("/:id/items", api.ApiGroupApp.LeaderboardApi.GetLeaderboardWithItems)
		leaderboardRouter.GET("/:id/items-only", api.ApiGroupApp.LeaderboardApi.GetLeaderboardItems)

		// 需要认证的路由
		authRouter := leaderboardRouter.Group("/")
		authRouter.Use(middleware.CombinedAuthMiddleware())
		{
			authRouter.POST("", api.ApiGroupApp.LeaderboardApi.CreateLeaderboard)
			authRouter.PUT("/:id", api.ApiGroupApp.LeaderboardApi.UpdateLeaderboard)
			authRouter.DELETE("/:id", api.ApiGroupApp.LeaderboardApi.DeleteLeaderboard)
			authRouter.POST("/:id/generate", api.ApiGroupApp.LeaderboardApi.GenerateLeaderboard)
		}
	}
}
