package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/gin-gonic/gin"
)

type CategoryRouter struct{}

// InitCategoryRouter 初始化分类路由
func (r *CategoryRouter) InitCategoryRouter(router *gin.RouterGroup) {
	categoryRouter := router.Group("/categories")
	{
		categoryRouter.GET("", api.ApiGroupApp.CategoryApi.GetCategories)
	}
}
