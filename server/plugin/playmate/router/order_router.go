package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/gin-gonic/gin"
)

type OrderRouter struct{}

// InitOrderRouter 初始化订单路由
func (r *OrderRouter) InitOrderRouter(router *gin.RouterGroup) {
	// 订单相关路由 - 需要认证
	orderRouter := router.Group("/orders")
	orderRouter.Use(middleware.CombinedAuthMiddleware())
	{
		orderRouter.GET("", api.ApiGroupApp.OrderApi.GetOrders)
		orderRouter.GET("/:id", api.ApiGroupApp.OrderApi.GetOrderDetail)
		orderRouter.POST("", api.ApiGroupApp.OrderApi.CreateOrder)
		orderRouter.GET("/:id/confirmation", api.ApiGroupApp.OrderApi.GetOrderConfirmation)
		orderRouter.POST("/:id/cancel", api.ApiGroupApp.OrderApi.CancelOrder)
		orderRouter.POST("/:id/confirm", api.ApiGroupApp.OrderApi.ConfirmOrder)
		orderRouter.POST("/:id/accept", api.ApiGroupApp.OrderApi.AcceptOrder)
		orderRouter.POST("/:id/reject", api.ApiGroupApp.OrderApi.RejectOrder)
		orderRouter.POST("/:id/share", api.ApiGroupApp.OrderApi.ShareOrder)
	}
}
