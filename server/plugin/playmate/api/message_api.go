package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
)

// MessageApi 消息API
type MessageApi struct{}

// GetMessages 获取消息列表
// @Tags     Message
// @Summary  获取消息列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]map[string]interface{}} "获取成功"
// @Router   /messages [get]
func (a *MessageApi) GetMessages(c *gin.Context) {
	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	messages, err := service.ServiceGroupApp.MessageService.GetMessages(userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(messages, "获取成功", c)
}

// GetChatMessages 获取聊天消息
// @Tags     Message
// @Summary  获取聊天消息
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    userId  path      uint    true "用户ID"
// @Success  200     {object} response.Response{data=[]model.ChatMessage} "获取成功"
// @Router   /messages/chat/{userId} [get]
func (a *MessageApi) GetChatMessages(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该从上下文获取当前用户ID
	currentUserID := uint(1) // 临时值

	chatMessages, err := service.ServiceGroupApp.MessageService.GetChatMessages(currentUserID, uint(userId))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(chatMessages, "获取成功", c)
}

// SendMessage 发送消息
// @Tags     Message
// @Summary  发送消息
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    userId  path      uint                   true "用户ID"
// @Param    data    body      request.SendMessageRequest  true "消息内容"
// @Success  200     {object} response.Response{data=model.Message} "发送成功"
// @Router   /messages/chat/{userId} [post]
func (a *MessageApi) SendMessage(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var req request.SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该从上下文获取当前用户ID
	currentUserID := uint(1) // 临时值

	message, err := service.ServiceGroupApp.MessageService.SendMessage(currentUserID, uint(userId), req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(message, "发送成功", c)
}