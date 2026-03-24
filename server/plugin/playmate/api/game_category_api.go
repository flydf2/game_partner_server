package api

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
)

type GameCategoryApi struct{}

// GetCategories 获取游戏分类列表
// @Tags     GameCategory
// @Summary  获取游戏分类列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]model.GameCategory} "获取游戏分类列表成功"
// @Router   /game-categories [get]
func (a *GameCategoryApi) GetCategories(c *gin.Context) {
	categories, err := service.ServiceGroupApp.GameCategoryService.GetCategories()
	if err != nil {
		response.FailWithMessage("获取游戏分类列表失败", c)
		return
	}
	response.OkWithData(categories, c)
}

// GetGamesByCategory 根据分类获取游戏列表
// @Tags     GameCategory
// @Summary  根据分类获取游戏列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    category path     string true "分类ID"
// @Success  200  {object} response.Response{data=[]model.Game} "根据分类获取游戏列表成功"
// @Router   /game-categories/{category}/games [get]
func (a *GameCategoryApi) GetGamesByCategory(c *gin.Context) {
	category := c.Param("category")
	games, err := service.ServiceGroupApp.GameCategoryService.GetGamesByCategory(category)
	if err != nil {
		response.FailWithMessage("根据分类获取游戏列表失败", c)
		return
	}
	response.OkWithData(games, c)
}
