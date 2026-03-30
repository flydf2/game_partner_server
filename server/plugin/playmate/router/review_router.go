package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/gin-gonic/gin"
)

type ReviewRouter struct{}

// InitReviewRouter 初始化评价路由
func (r *ReviewRouter) InitReviewRouter(router *gin.RouterGroup) {
	reviewRouter := router.Group("/reviews")
	{
		// 不需要认证的路由
		reviewRouter.GET("", api.ApiGroupApp.ReviewApi.GetReviews)

		// 需要认证的路由
		authRouter := reviewRouter.Group("/")
		authRouter.Use(middleware.CombinedAuthMiddleware())
		{
			authRouter.POST("", api.ApiGroupApp.ReviewApi.SubmitReview)
		}
	}
}
