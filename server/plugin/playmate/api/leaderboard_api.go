package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
)

// LeaderboardApi 排行榜API
type LeaderboardApi struct{}

// GetLeaderboards 获取排行榜列表
// @Tags     Leaderboard
// @Summary  获取排行榜列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    type     query    string  false "榜单类型（weekly-周榜, monthly-月榜）"
// @Param    game     query    string  false "游戏"
// @Param    page     query    int     false "页码"
// @Param    pageSize query    int     false "每页数量"
// @Success  200      {object} response.Response{data=[]model.Leaderboard, pagination=map[string]int64} "获取成功"
// @Router   /playmate/leaderboards [get]
func (a *LeaderboardApi) GetLeaderboards(c *gin.Context) {
	var search request.LeaderboardSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	leaderboards, total, err := service.ServiceGroupApp.LeaderboardService.GetLeaderboards(search)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	// 处理pageSize为0的情况，避免除以零错误
	pageSize := search.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	response.OkWithDetailed(gin.H{
		"data": leaderboards,
		"pagination": gin.H{
			"currentPage": search.Page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// GetLeaderboardById 根据ID获取排行榜详情
// @Tags     Leaderboard
// @Summary  根据ID获取排行榜详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "排行榜ID"
// @Success  200  {object} response.Response{data=model.Leaderboard} "获取成功"
// @Router   /playmate/leaderboards/{id} [get]
func (a *LeaderboardApi) GetLeaderboardById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	leaderboard, err := service.ServiceGroupApp.LeaderboardService.GetLeaderboardById(uint(id))
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(leaderboard, "获取成功", c)
}

// GetLeaderboardWithItems 获取排行榜及其条目
// @Tags     Leaderboard
// @Summary  获取排行榜及其条目
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id       path      uint    true "排行榜ID"
// @Param    page     query     int     false "页码"
// @Param    pageSize query     int     false "每页数量"
// @Success  200      {object} response.Response{data=map[string]interface{}, pagination=map[string]int64} "获取成功"
// @Router   /playmate/leaderboards/{id}/items [get]
func (a *LeaderboardApi) GetLeaderboardWithItems(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	leaderboard, items, total, err := service.ServiceGroupApp.LeaderboardService.GetLeaderboardWithItems(uint(id), page, pageSize)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"leaderboard": leaderboard,
		"items":       items,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// CreateLeaderboard 创建排行榜
// @Tags     Leaderboard
// @Summary  创建排行榜
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.CreateLeaderboardRequest  true "排行榜信息"
// @Success  200   {object}  response.Response{data=model.Leaderboard} "创建成功"
// @Router   /playmate/leaderboards [post]
func (a *LeaderboardApi) CreateLeaderboard(c *gin.Context) {
	var req request.CreateLeaderboardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	leaderboard, err := service.ServiceGroupApp.LeaderboardService.CreateLeaderboard(req)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(leaderboard, "创建成功", c)
}

// UpdateLeaderboard 更新排行榜
// @Tags     Leaderboard
// @Summary  更新排行榜
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id    path      uint                        true "排行榜ID"
// @Param    data  body      request.UpdateLeaderboardRequest  true "排行榜信息"
// @Success  200   {object}  response.Response{data=model.Leaderboard} "更新成功"
// @Router   /playmate/leaderboards/{id} [put]
func (a *LeaderboardApi) UpdateLeaderboard(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var req request.UpdateLeaderboardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	leaderboard, err := service.ServiceGroupApp.LeaderboardService.UpdateLeaderboard(uint(id), req)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(leaderboard, "更新成功", c)
}

// DeleteLeaderboard 删除排行榜
// @Tags     Leaderboard
// @Summary  删除排行榜
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id  path      uint    true "排行榜ID"
// @Success  200 {object}  response.Response{msg=string} "删除成功"
// @Router   /playmate/leaderboards/{id} [delete]
func (a *LeaderboardApi) DeleteLeaderboard(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	err = service.ServiceGroupApp.LeaderboardService.DeleteLeaderboard(uint(id))
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// GenerateLeaderboard 生成排行榜
// @Tags     Leaderboard
// @Summary  生成排行榜（根据陪玩数据自动计算排名）
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id  path      uint    true "排行榜ID"
// @Success  200 {object}  response.Response{msg=string} "生成成功"
// @Router   /playmate/leaderboards/{id}/generate [post]
func (a *LeaderboardApi) GenerateLeaderboard(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	err = service.ServiceGroupApp.LeaderboardService.GenerateLeaderboard(uint(id))
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("生成成功", c)
}

// GetLeaderboardItems 获取排行榜条目列表
// @Tags     Leaderboard
// @Summary  获取排行榜条目列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id       path      uint    true "排行榜ID"
// @Param    page     query     int     false "页码"
// @Param    pageSize query     int     false "每页数量"
// @Success  200      {object} response.Response{data=[]model.LeaderboardItem, pagination=map[string]int64} "获取成功"
// @Router   /playmate/leaderboards/{id}/items-only [get]
func (a *LeaderboardApi) GetLeaderboardItems(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	items, total, err := service.ServiceGroupApp.LeaderboardService.GetLeaderboardItems(uint(id), page, pageSize)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": items,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}
