package api

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
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
// @Router   /withdrawals [post]
func (a *WithdrawalApi) SubmitWithdrawal(c *gin.Context) {
	var req request.SubmitWithdrawalRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值
	if _, err := service.ServiceGroupApp.WithdrawalService.SubmitWithdrawal(userID, req); err != nil {
		response.FailWithMessage("提交提现失败", c)
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
// @Success  200  {object} response.Response{data=[]model.Withdrawal} "获取提现记录成功"
// @Router   /withdrawals [get]
func (a *WithdrawalApi) GetWithdrawalRecords(c *gin.Context) {
	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值
	withdrawals, err := service.ServiceGroupApp.WithdrawalService.GetWithdrawalRecords(userID)
	if err != nil {
		response.FailWithMessage("获取提现记录失败", c)
		return
	}
	response.OkWithData(withdrawals, c)
}
