package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/gin-gonic/gin"
)

// AppealRouter 申诉路由
type AppealRouter struct{}

// InitAppealRouter 初始化申诉路由
func (r *AppealRouter) InitAppealRouter(Router *gin.RouterGroup) {
	appealRouter := Router.Group("appeals")
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
