package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
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

	// 从上下文获取用户ID
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.FailWithMessage("未获取到用户ID", c)
		return
	}

	orders, total, err := service.ServiceGroupApp.OrderService.GetOrders(userID, search)
	if err != nil {
		response.FailWithError(err, c)
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
		response.FailWithError(err, c)
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

	// 从上下文获取用户ID
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.FailWithMessage("未获取到用户ID", c)
		return
	}

	order, err := service.ServiceGroupApp.OrderService.CreateOrder(userID, req)
	if err != nil {
		response.FailWithError(err, c)
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
		response.FailWithError(err, c)
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

	// 从上下文获取用户ID
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.FailWithMessage("未获取到用户ID", c)
		return
	}

	err = service.ServiceGroupApp.OrderService.CancelOrder(uint(id), userID)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("订单取消成功", c)
}

// ConfirmOrder 确认订单
// @Tags     Order
// @Summary  确认订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "订单ID"
// @Success  200  {object} response.Response{msg=string} "确认成功"
// @Router   /playmate/orders/{id}/confirm [post]
func (a *OrderApi) ConfirmOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 从上下文获取用户ID
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.FailWithMessage("未获取到用户ID", c)
		return
	}

	err = service.ServiceGroupApp.OrderService.ConfirmOrder(uint(id), userID)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("订单确认成功", c)
}

// AcceptOrder 接受订单
// @Tags     Order
// @Summary  接受订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "订单ID"
// @Success  200  {object} response.Response{msg=string} "接受成功"
// @Router   /playmate/orders/{id}/accept [post]
func (a *OrderApi) AcceptOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 从上下文获取用户ID
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.FailWithMessage("未获取到用户ID", c)
		return
	}

	err = service.ServiceGroupApp.OrderService.AcceptOrder(uint(id), userID)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("订单接受成功", c)
}

// RejectOrder 拒绝订单
// @Tags     Order
// @Summary  拒绝订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "订单ID"
// @Success  200  {object} response.Response{msg=string} "拒绝成功"
// @Router   /playmate/orders/{id}/reject [post]
func (a *OrderApi) RejectOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 从上下文获取用户ID
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.FailWithMessage("未获取到用户ID", c)
		return
	}

	err = service.ServiceGroupApp.OrderService.RejectOrder(uint(id), userID)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("订单拒绝成功", c)
}

// ShareOrder 分享订单
// @Tags     Order
// @Summary  分享订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint                     true "订单ID"
// @Param    data  body      request.ShareOrderRequest  true "分享信息，包含platform字段"
// @Success  200  {object} response.Response{data=map[string]interface{}} "分享成功"
// @Router   /playmate/orders/{id}/share [post]
func (a *OrderApi) ShareOrder(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 从上下文获取用户ID
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.FailWithMessage("未获取到用户ID", c)
		return
	}

	// 绑定请求数据
	var req request.ShareOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 调用服务层分享方法
	shareData, err := service.ServiceGroupApp.OrderService.ShareOrder(uint(id), userID, req.Platform)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(shareData, "分享成功", c)
}

// GetAllOrders 获取所有订单列表（管理员）
// @Tags     Order
// @Summary  获取所有订单列表（管理员）
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    status     query    string  false "订单状态"
// @Param    game       query    string  false "游戏"
// @Param    startTime  query    string  false "开始时间"
// @Param    endTime    query    string  false "结束时间"
// @Param    minAmount  query    number  false "最小金额"
// @Param    maxAmount  query    number  false "最大金额"
// @Param    userId     query    uint    false "用户ID"
// @Param    playmateId query    uint    false "陪玩ID"
// @Param    keyword    query    string  false "关键词"
// @Param    page       query    int     false "页码"
// @Param    pageSize   query    int     false "每页数量"
// @Success  200        {object} response.Response{data=[]model.Order,pagination=map[string]int64} "获取成功"
// @Router   /playmate/orders/all [get]
func (a *OrderApi) GetAllOrders(c *gin.Context) {
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

	orders, total, err := service.ServiceGroupApp.OrderService.GetAllOrders(search)
	if err != nil {
		response.FailWithError(err, c)
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

// BatchHandleOrders 批量处理订单
// @Tags     Order
// @Summary  批量处理订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.BatchHandleOrdersRequest  true "批量处理信息"
// @Success  200   {object}  response.Response{message=string} "处理成功"
// @Router   /playmate/orders/batch-handle [post]
func (a *OrderApi) BatchHandleOrders(c *gin.Context) {
	var req request.BatchHandleOrdersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	err := service.ServiceGroupApp.OrderService.BatchHandleOrders(req)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("处理成功", c)
}

// GetOrderStats 获取订单统计数据
// @Tags     Order
// @Summary  获取订单统计数据
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    startTime query    string  false "开始时间"
// @Param    endTime   query    string  false "结束时间"
// @Param    game      query    string  false "游戏"
// @Success  200       {object}  response.Response{data=map[string]interface{}} "获取成功"
// @Router   /playmate/orders/stats [get]
func (a *OrderApi) GetOrderStats(c *gin.Context) {
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	game := c.Query("game")

	stats, err := service.ServiceGroupApp.OrderService.GetOrderStats(startTime, endTime, game)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(stats, "获取成功", c)
}

// ExportOrders 导出订单列表
// @Tags     Order
// @Summary  导出订单列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/vnd.ms-excel
// @Param    status     query    string  false "订单状态"
// @Param    game       query    string  false "游戏"
// @Param    startTime  query    string  false "开始时间"
// @Param    endTime    query    string  false "结束时间"
// @Param    minAmount  query    number  false "最小金额"
// @Param    maxAmount  query    number  false "最大金额"
// @Param    userId     query    uint    false "用户ID"
// @Param    playmateId query    uint    false "陪玩ID"
// @Success  200        {file}    file    "导出成功"
// @Router   /playmate/orders/export [get]
func (a *OrderApi) ExportOrders(c *gin.Context) {
	var search request.OrderSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	excelData, err := service.ServiceGroupApp.OrderService.ExportOrders(search)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	c.Header("Content-Type", "application/vnd.ms-excel")
	c.Header("Content-Disposition", "attachment; filename=orders.xlsx")
	c.Data(200, "application/vnd.ms-excel", excelData)
}
