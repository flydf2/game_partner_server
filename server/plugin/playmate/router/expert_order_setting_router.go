package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/gin-gonic/gin"
)

type ExpertOrderSettingRouter struct{}

// InitExpertOrderSettingRouter 初始化专家订单设置路由
func (r *ExpertOrderSettingRouter) InitExpertOrderSettingRouter(router *gin.RouterGroup) {
	orderSettingRouter := router.Group("/expert/order-settings")
	{
		// 订单设置
		orderSettingRouter.GET("", api.ApiGroupApp.ExpertOrderSettingApi.GetOrderSetting)
		orderSettingRouter.PUT("", api.ApiGroupApp.ExpertOrderSettingApi.UpdateOrderSetting)

		// 服务管理
		orderSettingRouter.GET("/services", api.ApiGroupApp.ExpertOrderSettingApi.GetExpertServices)
		orderSettingRouter.POST("/services", api.ApiGroupApp.ExpertOrderSettingApi.CreateExpertService)
		orderSettingRouter.PUT("/services/:id", api.ApiGroupApp.ExpertOrderSettingApi.UpdateExpertService)
		orderSettingRouter.DELETE("/services/:id", api.ApiGroupApp.ExpertOrderSettingApi.DeleteExpertService)

		// 今日推荐
		orderSettingRouter.GET("/today-recommendations", api.ApiGroupApp.ExpertOrderSettingApi.GetTodayRecommendations)
		orderSettingRouter.POST("/today-recommendations", api.ApiGroupApp.ExpertOrderSettingApi.CreateTodayRecommendation)
		orderSettingRouter.PUT("/today-recommendations/:id", api.ApiGroupApp.ExpertOrderSettingApi.UpdateTodayRecommendation)
		orderSettingRouter.DELETE("/today-recommendations/:id", api.ApiGroupApp.ExpertOrderSettingApi.DeleteTodayRecommendation)
	}
}
