package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
)

// NotificationApi 通知API
type NotificationApi struct{}

// GetNotifications 获取通知列表
// @Tags     Notification
// @Summary  获取通知列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]model.Notification} "获取成功"
// @Router   /notifications [get]
func (a *NotificationApi) GetNotifications(c *gin.Context) {
	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	notifications, err := service.ServiceGroupApp.NotificationService.GetNotifications(userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(notifications, "获取成功", c)
}

// MarkAsRead 标记通知为已读
// @Tags     Notification
// @Summary  标记通知为已读
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "通知ID"
// @Success  200  {object} response.Response{message=string} "标记成功"
// @Router   /notifications/{id}/read [put]
func (a *NotificationApi) MarkAsRead(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	err = service.ServiceGroupApp.NotificationService.MarkAsRead(uint(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("标记成功", c)
}

// MarkAllAsRead 标记所有通知为已读
// @Tags     Notification
// @Summary  标记所有通知为已读
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{message=string} "全部标记为已读"
// @Router   /notifications/read-all [put]
func (a *NotificationApi) MarkAllAsRead(c *gin.Context) {
	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	err := service.ServiceGroupApp.NotificationService.MarkAllAsRead(userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("全部标记为已读", c)
}