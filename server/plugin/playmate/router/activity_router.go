package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
)

type ActivityRouter struct{}

// InitActivityRouter 初始化活动路由
func (r *ActivityRouter) InitActivityRouter(router *gin.RouterGroup) {
	activityRouter := router.Group("/activities")
	{
		activityRouter.GET("", api.ApiGroupApp.ActivityApi.GetActivities)
	}
}
