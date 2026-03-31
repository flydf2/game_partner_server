package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
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
// @Param    id path     string true "分类ID"
// @Success  200  {object} response.Response{data=[]model.Game} "根据分类获取游戏列表成功"
// @Router   /playmate/game-categories/{id}/games [get]
func (a *GameCategoryApi) GetGamesByCategory(c *gin.Context) {
	category := c.Param("id")
	games, err := service.ServiceGroupApp.GameCategoryService.GetGamesByCategory(category)
	if err != nil {
		response.FailWithMessage("根据分类获取游戏列表失败", c)
		return
	}
	response.OkWithData(games, c)
}

// GetCategoryById 根据ID获取游戏分类
// @Tags     GameCategory
// @Summary  根据ID获取游戏分类
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "分类ID"
// @Success  200  {object} response.Response{data=model.GameCategory} "获取成功"
// @Router   /playmate/game-categories/{id} [get]
func (a *GameCategoryApi) GetCategoryById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	category, err := service.ServiceGroupApp.GameCategoryService.GetCategoryById(uint(id))
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(category, "获取成功", c)
}

// CreateCategory 创建游戏分类
// @Tags     GameCategory
// @Summary  创建游戏分类
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body      model.GameCategory  true "分类信息"
// @Success  200   {object}  response.Response{data=model.GameCategory} "创建成功"
// @Router   /playmate/game-categories [post]
func (a *GameCategoryApi) CreateCategory(c *gin.Context) {
	var category model.GameCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	createdCategory, err := service.ServiceGroupApp.GameCategoryService.CreateCategory(category)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(createdCategory, "创建成功", c)
}

// UpdateCategory 更新游戏分类
// @Tags     GameCategory
// @Summary  更新游戏分类
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id    path      uint              true "分类ID"
// @Param    data  body      model.GameCategory  true "分类信息"
// @Success  200   {object}  response.Response{data=model.GameCategory} "更新成功"
// @Router   /playmate/game-categories/{id} [put]
func (a *GameCategoryApi) UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var category model.GameCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	category.ID = uint(id)
	updatedCategory, err := service.ServiceGroupApp.GameCategoryService.UpdateCategory(category)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(updatedCategory, "更新成功", c)
}

// DeleteCategory 删除游戏分类
// @Tags     GameCategory
// @Summary  删除游戏分类
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id  path      uint    true "分类ID"
// @Success  200 {object}  response.Response{message=string} "删除成功"
// @Router   /playmate/game-categories/{id} [delete]
func (a *GameCategoryApi) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	err = service.ServiceGroupApp.GameCategoryService.DeleteCategory(uint(id))
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("删除成功", c)
}
