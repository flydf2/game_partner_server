package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
)

// RewardOrderApi 悬赏订单API
type RewardOrderApi struct{}

// GetRewardOrders 获取悬赏订单列表
// @Tags     RewardOrder
// @Summary  获取悬赏订单列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    page        query    int     false "页码"
// @Param    pageSize    query    int     false "每页数量"
// @Param    status      query    string  false "订单状态"
// @Param    game        query    string  false "游戏"
// @Success  200         {object} response.Response{data=[]model.RewardOrder,pagination=map[string]int64} "获取成功"
// @Router   /playmate/reward [get]
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

// GetMyRewardOrders 获取我的悬赏订单列表
// @Tags     RewardOrder
// @Summary  获取我的悬赏订单列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    page        query    int     false "页码"
// @Param    pageSize    query    int     false "每页数量"
// @Success  200         {object} response.Response{data=[]model.RewardOrder,pagination=map[string]int64} "获取成功"
// @Router   /playmate/api/reward/my [get]
func (a *RewardOrderApi) GetMyRewardOrders(c *gin.Context) {
	// 这里应该从上下文获取用户ID
	// userID := c.GetUint("userID")
	userID := uint(1) // 临时值

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	orders, total, err := service.ServiceGroupApp.RewardOrderService.GetMyRewardOrders(userID, page, pageSize)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": orders,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// GetRewardOrderDetail 获取悬赏订单详情
// @Tags     RewardOrder
// @Summary  获取悬赏订单详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    orderId   path      uint  true "订单ID"
// @Success  200       {object} response.Response{data=model.RewardOrder} "获取成功"
// @Router   /playmate/reward/{orderId} [get]
func (a *RewardOrderApi) GetRewardOrderDetail(c *gin.Context) {
	orderIdStr := c.Param("orderId")
	orderId, err := strconv.ParseUint(orderIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	order, err := service.ServiceGroupApp.RewardOrderService.GetRewardOrderDetail(uint(orderId))
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	// 构建返回数据，包含联系信息
	responseData := gin.H{
		"id":            order.ID,
		"userId":        order.UserID,
		"userAvatar":    "https://randomuser.me/api/portraits/men/32.jpg", // 模拟数据
		"userName":      "游戏达人",                                           // 模拟数据
		"userLevel":     24,                                               // 模拟数据
		"userSpecialty": "顶级辅助",                                           // 模拟数据
		"game":          order.Game,
		"content":       order.Content,
		"reward":        order.Reward,
		"paymentMethod": order.PaymentMethod,
		"status":        order.Status,
		"timeLeft":      order.TimeLeft,
		"tags":          order.Tags,
		"requirements":  order.Requirements,
		"createdAt":     order.CreatedAt,
		"contactInfo": gin.H{
			"phone":  "138****8888",
			"wechat": "game****1234",
		},
	}

	response.OkWithDetailed(responseData, "获取成功", c)
}

// GetApplicants 获取抢单者列表
// @Tags     RewardOrder
// @Summary  获取抢单者列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    orderId   path      uint  true "订单ID"
// @Success  200       {object} response.Response{data=[]map[string]interface{}} "获取成功"
// @Router   /playmate/reward/{orderId}/applicants [get]
func (a *RewardOrderApi) GetApplicants(c *gin.Context) {
	orderIdStr := c.Param("orderId")
	orderId, err := strconv.ParseUint(orderIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	applicants, err := service.ServiceGroupApp.RewardOrderService.GetApplicants(uint(orderId))
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(applicants, "获取成功", c)
}

// SelectApplicant 选择抢单者
// @Tags     RewardOrder
// @Summary  选择抢单者
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    orderId   path      uint                      true "订单ID"
// @Param    data      body      request.SelectApplicantRequest true "选择抢单者信息"
// @Success  200       {object} response.Response{msg=string} "选择成功"
// @Router   /playmate/reward/{orderId}/select-applicant [post]
func (a *RewardOrderApi) SelectApplicant(c *gin.Context) {
	orderIdStr := c.Param("orderId")
	orderId, err := strconv.ParseUint(orderIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var req request.SelectApplicantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	if err := service.ServiceGroupApp.RewardOrderService.SelectApplicant(uint(orderId), req.ApplicantID); err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("选择成功", c)
}

// GrabRewardOrder 抢单
// @Tags     RewardOrder
// @Summary  抢单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    orderId   path      uint                      true "订单ID"
// @Param    data      body      request.GrabRewardOrderRequest false "抢单信息"
// @Success  200       {object} response.Response{data=map[string]interface{},msg=string} "抢单成功"
// @Router   /playmate/reward/{orderId}/grab [post]
func (a *RewardOrderApi) GrabRewardOrder(c *gin.Context) {
	orderIdStr := c.Param("orderId")
	orderId, err := strconv.ParseUint(orderIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该从上下文获取用户ID
	// userID := c.GetUint("userID")
	userID := uint(1) // 临时值

	var req request.GrabRewardOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 允许空请求体
		req = request.GrabRewardOrderRequest{}
	}

	if err := service.ServiceGroupApp.RewardOrderService.GrabRewardOrder(uint(orderId), userID, req); err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"orderId": orderId,
		"status":  "grabbed",
		"message": "抢单成功",
	}, "抢单成功", c)
}

// PublishReward 发布悬赏订单
// @Tags     RewardOrder
// @Summary  发布悬赏订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.CreateRewardOrderRequest  true "发布悬赏订单信息"
// @Success  200   {object} response.Response{data=map[string]uint,msg=string} "发布成功"
// @Router   /playmate/reward [post]
func (a *RewardOrderApi) PublishReward(c *gin.Context) {
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
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"orderId": order.ID,
	}, "发布成功", c)
}

// PublishRewardOrder 发布订单
// @Tags     RewardOrder
// @Summary  发布订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    orderId   path      uint  true "订单ID"
// @Success  200       {object} response.Response{msg=string} "发布成功"
// @Router   /playmate/reward/{orderId}/publish [post]
func (a *RewardOrderApi) PublishRewardOrder(c *gin.Context) {
	orderIdStr := c.Param("orderId")
	orderId, err := strconv.ParseUint(orderIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := service.ServiceGroupApp.RewardOrderService.PublishRewardOrder(uint(orderId)); err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("发布成功", c)
}

// PayRewardOrder 支付订单
// @Tags     RewardOrder
// @Summary  支付订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    orderId   path      uint                      true "订单ID"
// @Param    data      body      request.PayRewardOrderRequest true "支付信息"
// @Success  200       {object} response.Response{data=map[string]string,msg=string} "支付成功"
// @Router   /playmate/reward/{orderId}/pay [post]
func (a *RewardOrderApi) PayRewardOrder(c *gin.Context) {
	orderIdStr := c.Param("orderId")
	orderId, err := strconv.ParseUint(orderIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var req request.PayRewardOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	transactionId, err := service.ServiceGroupApp.RewardOrderService.PayRewardOrder(uint(orderId), req)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"transactionId": transactionId,
	}, "支付成功", c)
}

// ConfirmService 确认服务
// @Tags     RewardOrder
// @Summary  确认服务
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    orderId   path      uint                      true "订单ID"
// @Param    data      body      request.ConfirmServiceRequest true "确认服务信息"
// @Success  200       {object} response.Response{data=map[string]float64,msg=string} "服务确认成功"
// @Router   /playmate/reward/{orderId}/confirm [post]
func (a *RewardOrderApi) ConfirmService(c *gin.Context) {
	orderIdStr := c.Param("orderId")
	orderId, err := strconv.ParseUint(orderIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var req request.ConfirmServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	settlementAmount, err := service.ServiceGroupApp.RewardOrderService.ConfirmService(uint(orderId), req)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"settlementAmount": settlementAmount,
	}, "服务确认成功", c)
}

// ShareRewardOrder 分享悬赏订单
// @Tags     RewardOrder
// @Summary  分享悬赏订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    orderId   path      uint    true "订单ID"
// @Param    data      body      map[string]string  true "分享信息，包含platform字段"
// @Success  200       {object} response.Response{data=map[string]interface{}} "分享成功"
// @Router   /playmate/reward/{orderId}/share [post]
func (a *RewardOrderApi) ShareRewardOrder(c *gin.Context) {
	orderIdStr := c.Param("orderId")
	orderId, err := strconv.ParseUint(orderIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	// 绑定请求数据
	var req map[string]string
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 获取分享平台
	platform := req["platform"]
	if platform == "" {
		platform = "unknown"
	}

	// 调用服务层分享方法
	shareData, err := service.ServiceGroupApp.RewardOrderService.ShareRewardOrder(uint(orderId), userID, platform)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(shareData, "分享成功", c)
}
