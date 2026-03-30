package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/gin-gonic/gin"
)

type NotificationRouter struct{}

// InitNotificationRouter 初始化通知路由
func (r *NotificationRouter) InitNotificationRouter(router *gin.RouterGroup) {
	// 通知相关路由 - 需要认证
	notificationRouter := router.Group("/notifications")
	notificationRouter.Use(middleware.CombinedAuthMiddleware())
	{
		notificationRouter.GET("", api.ApiGroupApp.NotificationApi.GetNotifications)
		notificationRouter.GET("/unread-count", api.ApiGroupApp.NotificationApi.GetUnreadCount)
		notificationRouter.PUT("/:id/read", api.ApiGroupApp.NotificationApi.MarkAsRead)
		notificationRouter.PUT("/read-all", api.ApiGroupApp.NotificationApi.MarkAllAsRead)
	}
}
