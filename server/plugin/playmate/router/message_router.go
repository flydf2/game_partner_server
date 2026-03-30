package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/gin-gonic/gin"
)

type MessageRouter struct{}

// InitMessageRouter 初始化消息路由
func (r *MessageRouter) InitMessageRouter(router *gin.RouterGroup) {
	// 会话相关路由 - 需要认证
	conversationRouter := router.Group("/conversations")
	conversationRouter.Use(middleware.CombinedAuthMiddleware())
	{
		conversationRouter.GET("", api.ApiGroupApp.MessageApi.GetConversations)
		conversationRouter.PUT("/by-user/:userId/read", api.ApiGroupApp.MessageApi.MarkConversationAsRead)
		conversationRouter.PUT("/by-id/:id/read", api.ApiGroupApp.MessageApi.MarkConversationAsReadByID)
		conversationRouter.PUT("/by-id/:id/archive", api.ApiGroupApp.MessageApi.ArchiveConversation)
	}

	// 消息相关路由 - 需要认证
	messageRouter := router.Group("/messages")
	messageRouter.Use(middleware.CombinedAuthMiddleware())
	{
		messageRouter.GET("", api.ApiGroupApp.MessageApi.GetMessages)
		messageRouter.GET("/chat/:userId", api.ApiGroupApp.MessageApi.GetChatMessages)
		messageRouter.POST("/chat/:userId", api.ApiGroupApp.MessageApi.SendMessage)
		messageRouter.PUT("/by-id/:id/read", api.ApiGroupApp.MessageApi.MarkMessageAsRead)
	}
}
