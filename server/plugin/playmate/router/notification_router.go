package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
)

type NotificationRouter struct{}

// InitNotificationRouter 初始化通知路由
func (r *NotificationRouter) InitNotificationRouter(router *gin.RouterGroup) {
	notificationRouter := router.Group("/notifications")
	{
		notificationRouter.GET("", api.ApiGroupApp.NotificationApi.GetNotifications)
		notificationRouter.GET("/unread-count", api.ApiGroupApp.NotificationApi.GetUnreadCount)
		notificationRouter.PUT("/:id/read", api.ApiGroupApp.NotificationApi.MarkAsRead)
		notificationRouter.PUT("/read-all", api.ApiGroupApp.NotificationApi.MarkAllAsRead)
	}
}
