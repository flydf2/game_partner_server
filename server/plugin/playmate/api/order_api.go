package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
)

// OrderApi 订单API
type OrderApi struct{}

// GetOrders 获取订单列表
// @Tags     Order
// @Summary  获取订单列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    status     query    string  false "订单状态"
// @Param    game       query    string  false "游戏"
// @Param    startTime  query    string  false "开始时间"
// @Param    endTime    query    string  false "结束时间"
// @Param    minAmount  query    number  false "最小金额"
// @Param    maxAmount  query    number  false "最大金额"
// @Param    keyword    query    string  false "关键词"
// @Param    page       query    int     false "页码"
// @Param    pageSize   query    int     false "每页数量"
// @Success  200        {object} response.Response{data=[]model.Order,pagination=map[string]int64} "获取成功"
// @Router   /playmate/orders [get]
func (a *OrderApi) GetOrders(c *gin.Context) {
	var search request.OrderSearch
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

	userID := uint(1)

	orders, total, err := service.ServiceGroupApp.OrderService.GetOrders(userID, search)
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

// GetOrderDetail 获取订单详情
// @Tags     Order
// @Summary  获取订单详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "订单ID"
// @Success  200  {object} response.Response{data=model.Order} "获取成功"
// @Router   /playmate/orders/{id} [get]
func (a *OrderApi) GetOrderDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	order, err := service.ServiceGroupApp.OrderService.GetOrderDetail(uint(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(order, "获取成功", c)
}

// CreateOrder 创建订单
// @Tags     Order
// @Summary  创建订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.CreateOrderRequest  true "订单信息"
// @Success  200   {object} response.Response{data=map[string]interface{}} "创建成功"
// @Router   /playmate/orders [post]
func (a *OrderApi) CreateOrder(c *gin.Context) {
	var req request.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	order, err := service.ServiceGroupApp.OrderService.CreateOrder(userID, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"orderId":     order.ID,
		"orderNumber": order.OrderNumber,
		"status":      order.Status,
		"message":     "订单创建成功",
	}, "创建成功", c)
}

// GetOrderConfirmation 获取订单确认信息
// @Tags     Order
// @Summary  获取订单确认信息
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "订单ID"
// @Success  200  {object} response.Response{data=map[string]interface{}} "获取成功"
// @Router   /playmate/orders/{id}/confirmation [get]
func (a *OrderApi) GetOrderConfirmation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	confirmation, err := service.ServiceGroupApp.OrderService.GetOrderConfirmation(uint(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(confirmation, "获取成功", c)
}

// CancelOrder 取消订单
// @Tags     Order
// @Summary  取消订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "订单ID"
// @Success  200  {object} response.Response{msg=string} "取消成功"
// @Router   /playmate/orders/{id}/cancel [post]
func (a *OrderApi) CancelOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	err = service.ServiceGroupApp.OrderService.CancelOrder(uint(id), userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("订单取消成功", c)
}
