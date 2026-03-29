package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
)

// MessageApi 消息API
type MessageApi struct{}

// GetConversations 获取会话列表
// @Tags     Message
// @Summary  获取会话列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200        {object} response.Response{data=[]map[string]interface{}} "获取成功"
// @Router   /playmate/conversations [get]
func (a *MessageApi) GetConversations(c *gin.Context) {
	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	conversations, err := service.ServiceGroupApp.MessageService.GetConversations(userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(conversations, "获取成功", c)
}

// GetMessages 获取消息列表
// @Tags     Message
// @Summary  获取消息列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    type       query    string  false "类型"
// @Param    status     query    string  false "状态"
// @Param    senderId   query    uint    false "发送者ID"
// @Param    receiverId query    uint    false "接收者ID"
// @Param    startTime  query    string  false "开始时间"
// @Param    endTime    query    string  false "结束时间"
// @Param    keyword    query    string  false "关键词"
// @Param    page       query    int     false "页码"
// @Param    pageSize   query    int     false "每页数量"
// @Success  200        {object} response.Response{data=[]map[string]interface{},pagination=map[string]int64} "获取成功"
// @Router   /playmate/messages [get]
func (a *MessageApi) GetMessages(c *gin.Context) {
	var search request.MessageSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 设置默认值
	if search.Page <= 0 {
		search.Page = 1
	}
	if search.PageSize <= 0 {
		search.PageSize = 10
	}

	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	messages, total, err := service.ServiceGroupApp.MessageService.GetMessages(userID, search)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": messages,
		"pagination": gin.H{
			"currentPage": search.Page,
			"totalPages":  (total + int64(search.PageSize) - 1) / int64(search.PageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// GetChatMessages 获取聊天消息
// @Tags     Message
// @Summary  获取聊天消息
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    userId  path      uint    true "用户ID"
// @Success  200     {object} response.Response{data=[]model.ChatMessage} "获取成功"
// @Router   /playmate/messages/chat/{userId} [get]
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
// @Router   /playmate/messages/chat/{userId} [post]
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

	// 设置默认消息类型
	if req.Type == "" {
		req.Type = "text"
	}

	message, err := service.ServiceGroupApp.MessageService.SendMessage(currentUserID, uint(userId), req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(message, "发送成功", c)
}

// MarkMessageAsRead 标记消息为已读
// @Tags     Message
// @Summary  标记消息为已读
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "消息ID"
// @Success  200  {object} response.Response{msg=string} "标记成功"
// @Router   /playmate/messages/{id}/read [put]
func (a *MessageApi) MarkMessageAsRead(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := service.ServiceGroupApp.MessageService.MarkMessageAsRead(uint(id)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("标记成功", c)
}

// MarkConversationAsRead 标记会话为已读
// @Tags     Message
// @Summary  标记会话为已读
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    userId  path      uint    true "用户ID"
// @Success  200     {object} response.Response{msg=string} "标记成功"
// @Router   /playmate/conversations/read/{userId} [put]
func (a *MessageApi) MarkConversationAsRead(c *gin.Context) {
	userIdStr := c.Param("userId")
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该从上下文获取当前用户ID
	currentUserID := uint(1) // 临时值

	if err := service.ServiceGroupApp.MessageService.MarkConversationAsRead(currentUserID, uint(userId)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("标记成功", c)
}

// MarkConversationAsReadByID 通过会话ID标记会话为已读
// @Tags     Message
// @Summary  通过会话ID标记会话为已读
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "会话ID"
// @Success  200  {object} response.Response{msg=string} "标记成功"
// @Router   /playmate/conversations/{id}/read [put]
func (a *MessageApi) MarkConversationAsReadByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := service.ServiceGroupApp.MessageService.MarkConversationAsReadByID(uint(id)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("标记成功", c)
}

// ArchiveConversation 归档会话
// @Tags     Message
// @Summary  归档会话
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "会话ID"
// @Success  200  {object} response.Response{msg=string} "归档成功"
// @Router   /playmate/conversations/archive/{id} [put]
func (a *MessageApi) ArchiveConversation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := service.ServiceGroupApp.MessageService.ArchiveConversation(uint(id)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("归档成功", c)
}
