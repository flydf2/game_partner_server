package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
)

type GameCategoryRouter struct{}

// InitGameCategoryRouter 初始化游戏分类路由
func (r *GameCategoryRouter) InitGameCategoryRouter(router *gin.RouterGroup) {
	gameCategoryRouter := router.Group("/game-categories")
	{
		gameCategoryRouter.GET("", api.ApiGroupApp.GameCategoryApi.GetCategories)
		gameCategoryRouter.GET("/:category/games", api.ApiGroupApp.GameCategoryApi.GetGamesByCategory)
	}
}
