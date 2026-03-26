package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
)

// RewardOrderApi 奖励订单API
type RewardOrderApi struct{}

// GetRewardOrders 获取奖励订单列表
// @Tags     RewardOrder
// @Summary  获取奖励订单列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    game           query    string  false "游戏"
// @Param    status         query    string  false "订单状态"
// @Param    paymentMethod  query    string  false "支付方式"
// @Param    keyword        query    string  false "关键词"
// @Param    page           query    int     false "页码"
// @Param    pageSize       query    int     false "每页数量"
// @Success  200            {object} response.Response{data=[]model.RewardOrder,pagination=map[string]int64} "获取成功"
// @Router   /playmate/reward-orders [get]
func (a *RewardOrderApi) GetRewardOrders(c *gin.Context) {
	var search request.RewardOrderSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 设置默认值
	if search.Page <= 0 {
		search.Page = 1
	}
	if search.PageSize <= 0 {
		search.PageSize = 20
	}

	orders, total, err := service.ServiceGroupApp.RewardOrderService.GetRewardOrders(search)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": orders,
		"pagination": gin.H{
			"currentPage": search.Page,
			"totalPages":  (total + int64(search.PageSize) - 1) / int64(search.PageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// GetRewardOrderDetail 获取奖励订单详情
// @Tags     RewardOrder
// @Summary  获取奖励订单详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint  true "订单ID"
// @Success  200  {object} response.Response{data=model.RewardOrder} "获取成功"
// @Router   /playmate/reward-orders/{id} [get]
func (a *RewardOrderApi) GetRewardOrderDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	order, err := service.ServiceGroupApp.RewardOrderService.GetRewardOrderDetail(uint(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(order, "获取成功", c)
}

// CreateRewardOrder 创建奖励订单
// @Tags     RewardOrder
// @Summary  创建奖励订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.CreateRewardOrderRequest  true "奖励订单信息"
// @Success  200   {object} response.Response{data=model.RewardOrder} "创建成功"
// @Router   /playmate/reward-orders [post]
func (a *RewardOrderApi) CreateRewardOrder(c *gin.Context) {
	var req request.CreateRewardOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	// 这里应该从上下文获取用户ID
	// userID := c.GetUint("userID")
	userID := uint(1) // 临时值

	order, err := service.ServiceGroupApp.RewardOrderService.CreateRewardOrder(userID, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(order, "创建成功", c)
}

// UpdateRewardOrder 更新奖励订单
// @Tags     RewardOrder
// @Summary  更新奖励订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id    path      uint                              true "订单ID"
// @Param    data  body      request.UpdateRewardOrderRequest  true "奖励订单更新信息"
// @Success  200   {object} response.Response{data=model.RewardOrder} "更新成功"
// @Router   /playmate/reward-orders/{id} [put]
func (a *RewardOrderApi) UpdateRewardOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var req request.UpdateRewardOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	order, err := service.ServiceGroupApp.RewardOrderService.UpdateRewardOrder(uint(id), req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(order, "更新成功", c)
}

// DeleteRewardOrder 删除奖励订单
// @Tags     RewardOrder
// @Summary  删除奖励订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint  true "订单ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /playmate/reward-orders/{id} [delete]
func (a *RewardOrderApi) DeleteRewardOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := service.ServiceGroupApp.RewardOrderService.DeleteRewardOrder(uint(id)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// GrabRewardOrder 抢奖励订单
// @Tags     RewardOrder
// @Summary  抢奖励订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint  true "订单ID"
// @Success  200  {object} response.Response{msg=string} "抢单成功"
// @Router   /playmate/reward-orders/{id}/grab [post]
func (a *RewardOrderApi) GrabRewardOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该从上下文获取用户ID
	// userID := c.GetUint("userID")
	userID := uint(1) // 临时值

	if err := service.ServiceGroupApp.RewardOrderService.GrabRewardOrder(uint(id), userID); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("抢单成功", c)
}
