package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
)

type CategoryRouter struct{}

// InitCategoryRouter 初始化分类路由
func (r *CategoryRouter) InitCategoryRouter(router *gin.RouterGroup) {
	categoryRouter := router.Group("/categories")
	{
		categoryRouter.GET("", api.ApiGroupApp.CategoryApi.GetCategories)
	}
}
