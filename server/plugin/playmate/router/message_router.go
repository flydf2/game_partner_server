package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
)

type MessageRouter struct{}

// InitMessageRouter 初始化消息路由
func (r *MessageRouter) InitMessageRouter(router *gin.RouterGroup) {
	messageRouter := router.Group("/messages")
	{
		messageRouter.GET("", api.ApiGroupApp.MessageApi.GetMessages)
		messageRouter.GET("/chat/:userId", api.ApiGroupApp.MessageApi.GetChatMessages)
		messageRouter.POST("/chat/:userId", api.ApiGroupApp.MessageApi.SendMessage)
	}
}
