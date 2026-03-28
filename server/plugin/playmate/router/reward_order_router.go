package router

import (
	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
)

// RewardOrderRouter 悬赏订单路由
type RewardOrderRouter struct{}

// InitRewardOrderRouter 初始化悬赏订单路由
func (r *RewardOrderRouter) InitRewardOrderRouter(Router *gin.RouterGroup) {
	rewardOrderRouter := Router.Group("/reward")
	{
		// 获取悬赏订单列表
		rewardOrderRouter.GET("", api.ApiGroupApp.RewardOrderApi.GetRewardOrders)
		// 获取我的悬赏订单列表
		rewardOrderRouter.GET("/my", api.ApiGroupApp.RewardOrderApi.GetMyRewardOrders)
		// 获取悬赏订单详情
		rewardOrderRouter.GET("/:orderId", api.ApiGroupApp.RewardOrderApi.GetRewardOrderDetail)
		// 获取抢单者列表
		rewardOrderRouter.GET("/:orderId/applicants", api.ApiGroupApp.RewardOrderApi.GetApplicants)
		// 选择抢单者
		rewardOrderRouter.POST("/:orderId/select-applicant", api.ApiGroupApp.RewardOrderApi.SelectApplicant)
		// 抢单
		rewardOrderRouter.POST("/:orderId/grab", api.ApiGroupApp.RewardOrderApi.GrabRewardOrder)
		// 发布悬赏订单
		rewardOrderRouter.POST("", api.ApiGroupApp.RewardOrderApi.PublishReward)
		// 发布订单
		rewardOrderRouter.POST("/:orderId/publish", api.ApiGroupApp.RewardOrderApi.PublishRewardOrder)
		// 支付订单
		rewardOrderRouter.POST("/:orderId/pay", api.ApiGroupApp.RewardOrderApi.PayRewardOrder)
		// 确认服务
		rewardOrderRouter.POST("/:orderId/confirm", api.ApiGroupApp.RewardOrderApi.ConfirmService)
		// 分享悬赏订单
		rewardOrderRouter.POST("/:orderId/share", api.ApiGroupApp.RewardOrderApi.ShareRewardOrder)
	}
}
