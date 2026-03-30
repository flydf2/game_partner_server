package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/gin-gonic/gin"
)

// AppealRouter 申诉路由
type AppealRouter struct{}

// InitAppealRouter 初始化申诉路由
func (r *AppealRouter) InitAppealRouter(Router *gin.RouterGroup) {
	// 申诉相关路由 - 需要认证
	appealRouter := Router.Group("appeals")
	appealRouter.Use(middleware.CombinedAuthMiddleware())
	appealApi := api.ApiGroupApp.AppealApi
	{
		appealRouter.GET("", appealApi.GetAppeals)
		appealRouter.GET("/my", appealApi.GetMyAppeals)
		appealRouter.GET("/:id", appealApi.GetAppealDetail)
		appealRouter.POST("", appealApi.CreateAppeal)
		appealRouter.PUT("/:id", appealApi.UpdateAppeal)
		appealRouter.DELETE("/:id", appealApi.DeleteAppeal)
		appealRouter.PUT("/:id/handle", appealApi.HandleAppeal)
	}
}
