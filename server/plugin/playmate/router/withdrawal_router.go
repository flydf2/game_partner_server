package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
)

type WithdrawalRouter struct{}

// InitWithdrawalRouter 初始化提现路由
func (r *WithdrawalRouter) InitWithdrawalRouter(router *gin.RouterGroup) {
	// 提现相关路由 - 需要认证
	withdrawalRouter := router.Group("/withdrawals")
	withdrawalRouter.Use(middleware.CombinedAuthMiddleware())
	{
		withdrawalRouter.POST("", api.ApiGroupApp.WithdrawalApi.SubmitWithdrawal)
		withdrawalRouter.GET("", api.ApiGroupApp.WithdrawalApi.GetWithdrawalRecords)
	}
}
