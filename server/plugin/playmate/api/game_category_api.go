package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/gin-gonic/gin"
)

type GameCategoryApi struct{}

// GetCategories 获取游戏分类列表
// @Tags     GameCategory
// @Summary  获取游戏分类列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    status     query    string  false "状态"
// @Param    parentId   query    uint    false "父分类ID"
// @Param    keyword    query    string  false "关键词"
// @Param    page       query    int     false "页码"
// @Param    pageSize   query    int     false "每页数量"
// @Success  200        {object} response.Response{data=[]model.GameCategory,pagination=map[string]int64} "获取游戏分类列表成功"
// @Router   /playmate/game-categories [get]
func (a *GameCategoryApi) GetCategories(c *gin.Context) {
	var search request.GameCategorySearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 设置默认值
	if search.Page <= 0 {
		search.Page = 1
	}
	if search.PageSize <= 0 {
		search.PageSize = 10
	}

	categories, total, err := service.ServiceGroupApp.GameCategoryService.GetCategories(search)
	if err != nil {
		response.FailWithMessage("获取游戏分类列表失败", c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": categories,
		"pagination": gin.H{
			"currentPage": search.Page,
			"totalPages":  (total + int64(search.PageSize) - 1) / int64(search.PageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// GetGamesByCategory 根据分类获取游戏列表
// @Tags     GameCategory
// @Summary  根据分类获取游戏列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    category path     string true "分类ID"
// @Success  200  {object} response.Response{data=[]model.Game} "根据分类获取游戏列表成功"
// @Router   /playmate/game-categories/{category}/games [get]
func (a *GameCategoryApi) GetGamesByCategory(c *gin.Context) {
	category := c.Param("category")
	games, err := service.ServiceGroupApp.GameCategoryService.GetGamesByCategory(category)
	if err != nil {
		response.FailWithMessage("根据分类获取游戏列表失败", c)
		return
	}
	response.OkWithData(games, c)
}
