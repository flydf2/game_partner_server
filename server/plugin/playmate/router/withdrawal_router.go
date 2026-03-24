package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
)

type WithdrawalRouter struct{}

// InitWithdrawalRouter 初始化提现路由
func (r *WithdrawalRouter) InitWithdrawalRouter(router *gin.RouterGroup) {
	withdrawalRouter := router.Group("/withdrawals")
	{
		withdrawalRouter.POST("", api.ApiGroupApp.WithdrawalApi.SubmitWithdrawal)
		withdrawalRouter.GET("", api.ApiGroupApp.WithdrawalApi.GetWithdrawalRecords)
	}
}
