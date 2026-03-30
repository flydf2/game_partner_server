package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/gin-gonic/gin"
)

type WithdrawalApi struct{}

// SubmitWithdrawal 提交提现
// @Tags     Withdrawal
// @Summary  提交提现
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.SubmitWithdrawalRequest true "提现信息"
// @Success  200  {object} response.Response{data=map[string]interface{}} "提交提现成功"
// @Router   /playmate/withdrawals [post]
func (a *WithdrawalApi) SubmitWithdrawal(c *gin.Context) {
	var req request.SubmitWithdrawalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	// 从上下文获取用户ID
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}
	result, err := service.ServiceGroupApp.UserService.Withdraw(userID, req)
	if err != nil {
		response.FailWithError(err, c)
		return
	}
	response.OkWithDetailed(result, "提交提现成功", c)
}

// GetWithdrawalRecords 获取提现记录
// @Tags     Withdrawal
// @Summary  获取提现记录
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    page     query    int  false "页码"
// @Param    pageSize query    int  false "每页数量"
// @Success  200      {object} response.Response{data=[]model.Withdrawal,pagination=map[string]int64} "获取提现记录成功"
// @Router   /playmate/withdrawals [get]
func (a *WithdrawalApi) GetWithdrawalRecords(c *gin.Context) {
	// 从上下文获取用户ID
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	withdrawals, total, err := service.ServiceGroupApp.UserService.GetWithdrawalList(userID, page, pageSize)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": withdrawals,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// ProcessWithdrawal 处理提现（管理员操作）
// @Tags     Withdrawal
// @Summary  处理提现
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    withdrawalId path uint true "提现记录ID"
// @Param    status body string true "处理状态（approved 或 rejected）"
// @Success  200  {object} response.Response{message=string} "处理成功"
// @Router   /playmate/withdrawals/{withdrawalId}/process [post]
func (a *WithdrawalApi) ProcessWithdrawal(c *gin.Context) {
	// 从路径参数获取提现记录ID
	withdrawalID, err := strconv.ParseUint(c.Param("withdrawalId"), 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 从请求体获取处理状态
	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 调用服务处理提现
	err = service.ServiceGroupApp.UserService.ProcessWithdrawal(uint(withdrawalID), req.Status)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("处理成功", c)
}
