package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
)

// AppealApi 申诉API
type AppealApi struct{}

// GetAppeals 获取申诉列表
// @Tags     Appeal
// @Summary  获取申诉列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    type      query    string  false "申诉类型"
// @Param    status    query    string  false "状态"
// @Param    priority  query    string  false "优先级"
// @Param    userId    query    uint    false "用户ID"
// @Param    orderId   query    uint    false "订单ID"
// @Param    keyword   query    string  false "关键词"
// @Param    startTime query    string  false "开始时间"
// @Param    endTime   query    string  false "结束时间"
// @Param    page      query    int     false "页码"
// @Param    pageSize  query    int     false "每页数量"
// @Success  200       {object} response.Response{data=[]model.Appeal, pagination=map[string]int64} "获取成功"
// @Router   /playmate/appeals [get]
func (a *AppealApi) GetAppeals(c *gin.Context) {
	var search request.AppealSearch
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

	appeals, total, err := service.ServiceGroupApp.AppealService.GetAppeals(search)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": appeals,
		"pagination": gin.H{
			"currentPage": search.Page,
			"totalPages":  (total + int64(search.PageSize) - 1) / int64(search.PageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// GetAppealDetail 获取申诉详情
// @Tags     Appeal
// @Summary  获取申诉详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "申诉ID"
// @Success  200  {object} response.Response{data=model.Appeal} "获取成功"
// @Router   /playmate/appeals/{id} [get]
func (a *AppealApi) GetAppealDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	appeal, err := service.ServiceGroupApp.AppealService.GetAppealDetail(uint(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(appeal, "获取成功", c)
}

// CreateAppeal 创建申诉
// @Tags     Appeal
// @Summary  创建申诉
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.CreateAppealRequest  true "申诉信息"
// @Success  200   {object}  response.Response{data=model.Appeal} "创建成功"
// @Router   /playmate/appeals [post]
func (a *AppealApi) CreateAppeal(c *gin.Context) {
	var req request.CreateAppealRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	appeal, err := service.ServiceGroupApp.AppealService.CreateAppeal(userID, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(appeal, "创建成功", c)
}

// UpdateAppeal 更新申诉
// @Tags     Appeal
// @Summary  更新申诉
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id    path      uint                       true "申诉ID"
// @Param    data  body      request.UpdateAppealRequest  true "申诉信息"
// @Success  200   {object}  response.Response{data=model.Appeal} "更新成功"
// @Router   /playmate/appeals/{id} [put]
func (a *AppealApi) UpdateAppeal(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var req request.UpdateAppealRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	appeal, err := service.ServiceGroupApp.AppealService.UpdateAppeal(uint(id), req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(appeal, "更新成功", c)
}

// DeleteAppeal 删除申诉
// @Tags     Appeal
// @Summary  删除申诉
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id  path      uint    true "申诉ID"
// @Success  200 {object}  response.Response{message=string} "删除成功"
// @Router   /playmate/appeals/{id} [delete]
func (a *AppealApi) DeleteAppeal(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := service.ServiceGroupApp.AppealService.DeleteAppeal(uint(id)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// HandleAppeal 处理申诉
// @Tags     Appeal
// @Summary  处理申诉
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id    path      uint                     true "申诉ID"
// @Param    data  body      request.HandleAppealRequest  true "处理信息"
// @Success  200   {object}  response.Response{data=model.Appeal} "处理成功"
// @Router   /playmate/appeals/{id}/handle [put]
func (a *AppealApi) HandleAppeal(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var req request.HandleAppealRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该从上下文获取处理人ID
	handlerID := uint(1) // 临时值

	appeal, err := service.ServiceGroupApp.AppealService.HandleAppeal(uint(id), handlerID, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(appeal, "处理成功", c)
}

// GetMyAppeals 获取我的申诉列表
// @Tags     Appeal
// @Summary  获取我的申诉列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    page     query    int     false "页码"
// @Param    pageSize query    int     false "每页数量"
// @Success  200      {object} response.Response{data=[]model.Appeal, pagination=map[string]int64} "获取成功"
// @Router   /playmate/appeals/my [get]
func (a *AppealApi) GetMyAppeals(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	appeals, total, err := service.ServiceGroupApp.AppealService.GetUserAppeals(userID, page, pageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": appeals,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}
