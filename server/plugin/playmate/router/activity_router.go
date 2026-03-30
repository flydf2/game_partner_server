package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/gin-gonic/gin"
)

type ActivityRouter struct{}

// InitActivityRouter 初始化活动路由
func (r *ActivityRouter) InitActivityRouter(router *gin.RouterGroup) {
	activityRouter := router.Group("/activities")
	{
		activityRouter.GET("", api.ApiGroupApp.ActivityApi.GetActivities)
	}
}
