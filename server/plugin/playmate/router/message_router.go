package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
)

type MessageRouter struct{}

// InitMessageRouter 初始化消息路由
func (r *MessageRouter) InitMessageRouter(router *gin.RouterGroup) {
	// 会话相关路由
	conversationRouter := router.Group("/conversations")
	{
		conversationRouter.GET("", api.ApiGroupApp.MessageApi.GetConversations)
		conversationRouter.PUT("/read/:userId", api.ApiGroupApp.MessageApi.MarkConversationAsRead)
		conversationRouter.PUT("/:id/read", api.ApiGroupApp.MessageApi.MarkConversationAsReadByID)
		conversationRouter.PUT("/archive/:id", api.ApiGroupApp.MessageApi.ArchiveConversation)
	}

	// 消息相关路由
	messageRouter := router.Group("/messages")
	{
		messageRouter.GET("", api.ApiGroupApp.MessageApi.GetMessages)
		messageRouter.GET("/chat/:userId", api.ApiGroupApp.MessageApi.GetChatMessages)
		messageRouter.POST("/chat/:userId", api.ApiGroupApp.MessageApi.SendMessage)
		messageRouter.PUT("/:id/read", api.ApiGroupApp.MessageApi.MarkMessageAsRead)
	}
}
