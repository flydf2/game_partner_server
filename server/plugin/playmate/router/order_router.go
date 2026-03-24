package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
)

type OrderRouter struct{}

// InitOrderRouter 初始化订单路由
func (r *OrderRouter) InitOrderRouter(router *gin.RouterGroup) {
	orderRouter := router.Group("/orders")
	{
		orderRouter.GET("", api.ApiGroupApp.OrderApi.GetOrders)
		orderRouter.GET("/:id", api.ApiGroupApp.OrderApi.GetOrderDetail)
		orderRouter.POST("", api.ApiGroupApp.OrderApi.CreateOrder)
		orderRouter.GET("/:id/confirmation", api.ApiGroupApp.OrderApi.GetOrderConfirmation)
	}
}
