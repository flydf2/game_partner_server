package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/gin-gonic/gin"
)

type GameRouter struct{}

// InitGameRouter 初始化游戏路由
func (r *GameRouter) InitGameRouter(router *gin.RouterGroup) {
	gameRouter := router.Group("/games")
	{
		gameRouter.GET("", api.ApiGroupApp.GameApi.GetGames)
	}
}
