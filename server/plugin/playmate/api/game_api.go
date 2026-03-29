package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/gin-gonic/gin"
)

type GameApi struct{}

// GetGames 获取游戏列表
// @Tags     Game
// @Summary  获取游戏列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    categoryIds query    []uint  false "分类ID列表"
// @Param    status     query    string  false "状态"
// @Param    platform   query    string  false "平台"
// @Param    keyword    query    string  false "关键词"
// @Param    page       query    int     false "页码"
// @Param    pageSize   query    int     false "每页数量"
// @Success  200        {object} response.Response{data=[]model.Game,pagination=map[string]int64} "获取游戏列表成功"
// @Router   /playmate/games [get]
func (a *GameApi) GetGames(c *gin.Context) {
	var search request.GameSearch
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

	games, total, err := service.ServiceGroupApp.GameService.GetGames(search)
	if err != nil {
		response.FailWithMessage("获取游戏列表失败", c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": games,
		"pagination": gin.H{
			"currentPage": search.Page,
			"totalPages":  (total + int64(search.PageSize) - 1) / int64(search.PageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}
