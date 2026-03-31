package router

import (
	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
)

// StatsRouter 统计分析路由
type StatsRouter struct{}

// InitStatsRouter 初始化统计分析路由
func (s *StatsRouter) InitStatsRouter(Router *gin.RouterGroup) {
	statsRouter := Router.Group("/stats")
	statsRouter.Use(middleware.CombinedAuthMiddleware())
	{
		statsRouter.GET("/dashboard", api.ApiGroupApp.StatsApi.GetDashboardStats)
		statsRouter.GET("/orders", api.ApiGroupApp.StatsApi.GetOrderStats)
		statsRouter.GET("/users", api.ApiGroupApp.StatsApi.GetUserStats)
		statsRouter.GET("/experts", api.ApiGroupApp.StatsApi.GetExpertStats)
		statsRouter.GET("/revenue", api.ApiGroupApp.StatsApi.GetRevenueStats)
		statsRouter.GET("/trend", api.ApiGroupApp.StatsApi.GetTrendStats)
	}
}
