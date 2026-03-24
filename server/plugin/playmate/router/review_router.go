package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
)

type ReviewRouter struct{}

// InitReviewRouter 初始化评价路由
func (r *ReviewRouter) InitReviewRouter(router *gin.RouterGroup) {
	reviewRouter := router.Group("/reviews")
	{
		reviewRouter.POST("", api.ApiGroupApp.ReviewApi.SubmitReview)
		reviewRouter.GET("", api.ApiGroupApp.ReviewApi.GetReviews)
	}
}
