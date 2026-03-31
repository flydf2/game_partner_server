package router

import (
	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
)

// RewardOrderRouter 悬赏订单路由
type RewardOrderRouter struct{}

// InitRewardOrderRouter 初始化悬赏订单路由
func (r *RewardOrderRouter) InitRewardOrderRouter(Router *gin.RouterGroup) {
	rewardOrderRouter := Router.Group("/reward")
	{
		// 不需要认证的路由
		// 获取悬赏订单列表
		rewardOrderRouter.GET("", api.ApiGroupApp.RewardOrderApi.GetRewardOrders)
		// 获取悬赏订单详情
		rewardOrderRouter.GET("/:orderId", api.ApiGroupApp.RewardOrderApi.GetRewardOrderDetail)
		// 获取抢单者列表
		rewardOrderRouter.GET("/:orderId/applicants", api.ApiGroupApp.RewardOrderApi.GetApplicants)

		// 需要认证的路由
		// 发布悬赏订单
		rewardOrderRouter.POST("", middleware.CombinedAuthMiddleware(), api.ApiGroupApp.RewardOrderApi.PublishReward)
		// 获取我的悬赏订单列表
		rewardOrderRouter.GET("/my", middleware.CombinedAuthMiddleware(), api.ApiGroupApp.RewardOrderApi.GetMyRewardOrders)
		// 选择抢单者
		rewardOrderRouter.POST("/:orderId/select-applicant", middleware.CombinedAuthMiddleware(), api.ApiGroupApp.RewardOrderApi.SelectApplicant)
		// 抢单
		rewardOrderRouter.POST("/:orderId/grab", middleware.CombinedAuthMiddleware(), api.ApiGroupApp.RewardOrderApi.GrabRewardOrder)
		// 发布订单
		rewardOrderRouter.POST("/:orderId/publish", middleware.CombinedAuthMiddleware(), api.ApiGroupApp.RewardOrderApi.PublishRewardOrder)
		// 支付订单
		rewardOrderRouter.POST("/:orderId/pay", middleware.CombinedAuthMiddleware(), api.ApiGroupApp.RewardOrderApi.PayRewardOrder)
		// 确认服务
		rewardOrderRouter.POST("/:orderId/confirm", middleware.CombinedAuthMiddleware(), api.ApiGroupApp.RewardOrderApi.ConfirmService)
		// 分享悬赏订单
		rewardOrderRouter.POST("/:orderId/share", middleware.CombinedAuthMiddleware(), api.ApiGroupApp.RewardOrderApi.ShareRewardOrder)
		// 取消订单
		rewardOrderRouter.POST("/:orderId/cancel", middleware.CombinedAuthMiddleware(), api.ApiGroupApp.RewardOrderApi.CancelRewardOrder)
	}

	// 抢单相关路由
	grabOrderRouter := Router.Group("/grab-orders")
	{
		// 需要认证的路由
		// 获取抢单详情
		grabOrderRouter.GET("/:id", middleware.CombinedAuthMiddleware(), api.ApiGroupApp.RewardOrderApi.GetGrabOrderDetail)
		// 撤回抢单申请
		grabOrderRouter.POST("/:id/withdraw", middleware.CombinedAuthMiddleware(), api.ApiGroupApp.RewardOrderApi.WithdrawGrabOrder)
	}
}
