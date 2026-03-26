package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
)

// NotificationApi 通知API
type NotificationApi struct{}

// GetNotifications 获取通知列表
// @Tags     Notification
// @Summary  获取通知列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    type       query    string  false "类型"
// @Param    status     query    string  false "状态"
// @Param    startTime  query    string  false "开始时间"
// @Param    endTime    query    string  false "结束时间"
// @Param    keyword    query    string  false "关键词"
// @Param    page       query    int     false "页码"
// @Param    pageSize   query    int     false "每页数量"
// @Success  200        {object} response.Response{data=[]model.Notification,pagination=map[string]int64} "获取成功"
// @Router   /playmate/notifications [get]
func (a *NotificationApi) GetNotifications(c *gin.Context) {
	var search request.NotificationSearch
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

	notifications, total, err := service.ServiceGroupApp.NotificationService.GetNotifications(userID, search)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": notifications,
		"pagination": gin.H{
			"currentPage": search.Page,
			"totalPages":  (total + int64(search.PageSize) - 1) / int64(search.PageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// MarkAsRead 标记通知为已读
// @Tags     Notification
// @Summary  标记通知为已读
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "通知ID"
// @Success  200  {object} response.Response{message=string} "标记成功"
// @Router   /playmate/notifications/{id}/read [put]
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
// @Router   /playmate/notifications/read-all [put]
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
