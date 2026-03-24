package api

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
)

type GameApi struct{}

// GetGames 获取游戏列表
// @Tags     Game
// @Summary  获取游戏列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]model.Game} "获取游戏列表成功"
// @Router   /games [get]
func (a *GameApi) GetGames(c *gin.Context) {
	games, err := service.ServiceGroupApp.GameService.GetGames()
	if err != nil {
		response.FailWithMessage("获取游戏列表失败", c)
		return
	}
	response.OkWithData(games, c)
}
