package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
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
// @Success  200  {object} response.Response{msg=string} "提交提现成功"
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
		response.FailWithMessage("未获取到用户ID", c)
		return
	}
	if _, err := service.ServiceGroupApp.WithdrawalService.SubmitWithdrawal(userID, req); err != nil {
		response.FailWithError(err, c)
		return
	}
	response.OkWithMessage("提交提现成功", c)
}

// GetWithdrawalRecords 获取提现记录
// @Tags     Withdrawal
// @Summary  获取提现记录
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    status     query    string  false "状态"
// @Param    method     query    string  false "提现方式"
// @Param    minAmount  query    number  false "最小金额"
// @Param    maxAmount  query    number  false "最大金额"
// @Param    startTime  query    string  false "开始时间"
// @Param    endTime    query    string  false "结束时间"
// @Param    page       query    int     false "页码"
// @Param    pageSize   query    int     false "每页数量"
// @Success  200        {object} response.Response{data=[]model.Withdrawal,pagination=map[string]int64} "获取提现记录成功"
// @Router   /playmate/withdrawals [get]
func (a *WithdrawalApi) GetWithdrawalRecords(c *gin.Context) {
	var search request.WithdrawalSearch
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

	withdrawals, total, err := service.ServiceGroupApp.WithdrawalService.GetWithdrawalRecords(userID, search)
	if err != nil {
		response.FailWithMessage("获取提现记录失败", c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": withdrawals,
		"pagination": gin.H{
			"currentPage": search.Page,
			"totalPages":  (total + int64(search.PageSize) - 1) / int64(search.PageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}
