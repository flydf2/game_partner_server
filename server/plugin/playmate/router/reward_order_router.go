package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
)

type RewardOrderRouter struct{}

// InitRewardOrderRouter 初始化奖励订单路由
func (r *RewardOrderRouter) InitRewardOrderRouter(router *gin.RouterGroup) {
	rewardOrderRouter := router.Group("/reward-orders")
	{
		rewardOrderRouter.GET("", api.ApiGroupApp.RewardOrderApi.GetRewardOrders)
		rewardOrderRouter.GET("/:id", api.ApiGroupApp.RewardOrderApi.GetRewardOrderDetail)
		rewardOrderRouter.POST("", api.ApiGroupApp.RewardOrderApi.CreateRewardOrder)
		rewardOrderRouter.PUT("/:id", api.ApiGroupApp.RewardOrderApi.UpdateRewardOrder)
		rewardOrderRouter.DELETE("/:id", api.ApiGroupApp.RewardOrderApi.DeleteRewardOrder)
		rewardOrderRouter.POST("/:id/grab", api.ApiGroupApp.RewardOrderApi.GrabRewardOrder)
	}
}
