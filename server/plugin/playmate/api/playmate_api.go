package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
)

// PlaymateApi 陪玩API
type PlaymateApi struct{}

// GetPlaymates 获取陪玩列表
// @Tags     Playmate
// @Summary  获取陪玩列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    game       query    string  false "游戏"
// @Param    online     query    bool    false "是否在线"
// @Param    priceRange query    string  false "价格范围"
// @Param    rank       query    string  false "段位"
// @Param    gender     query    string  false "性别"
// @Param    keyword    query    string  false "关键词"
// @Param    sortBy     query    string  false "排序方式"
// @Param    page       query    int     false "页码"
// @Param    pageSize   query    int     false "每页数量"
// @Success  200        {object} response.Response{data=[]model.Playmate, pagination=map[string]int64} "获取成功"
// @Router   /playmate/playmates [get]
func (a *PlaymateApi) GetPlaymates(c *gin.Context) {
	var search request.PlaymateSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	playmates, total, err := service.ServiceGroupApp.PlaymateService.GetPlaymates(search)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 处理pageSize为0的情况，避免除以零错误
	pageSize := search.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	response.OkWithDetailed(gin.H{
		"data": playmates,
		"pagination": gin.H{
			"currentPage": search.Page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// GetExpertDetail 获取专家详情
// @Tags     Playmate
// @Summary  获取专家详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "专家ID"
// @Success  200  {object} response.Response{data=map[string]interface{}} "获取成功"
// @Router   /playmate/experts/{id} [get]
func (a *PlaymateApi) GetExpertDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	detail, err := service.ServiceGroupApp.PlaymateService.GetExpertDetail(uint(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(detail, "获取成功", c)
}

// FollowExpert 关注专家
// @Tags     Playmate
// @Summary  关注专家
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "专家ID"
// @Success  200  {object} response.Response{message=string} "关注成功"
// @Router   /playmate/experts/{id}/follow [post]
func (a *PlaymateApi) FollowExpert(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	err = service.ServiceGroupApp.PlaymateService.FollowExpert(userID, uint(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("关注成功", c)
}

// UnfollowExpert 取消关注专家
// @Tags     Playmate
// @Summary  取消关注专家
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "专家ID"
// @Success  200  {object} response.Response{message=string} "取消关注成功"
// @Router   /playmate/experts/{id}/follow [delete]
func (a *PlaymateApi) UnfollowExpert(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	err = service.ServiceGroupApp.PlaymateService.UnfollowExpert(userID, uint(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("取消关注成功", c)
}

// GetExpertReviews 获取专家评价
// @Tags     Playmate
// @Summary  获取专家评价
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id       path      uint    true "专家ID"
// @Param    page     query     int     false "页码"
// @Param    pageSize query     int     false "每页数量"
// @Success  200      {object} response.Response{data=[]model.Review, pagination=map[string]int64} "获取成功"
// @Router   /playmate/experts/{id}/reviews [get]
func (a *PlaymateApi) GetExpertReviews(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	reviews, total, err := service.ServiceGroupApp.PlaymateService.GetExpertReviews(uint(id), page, pageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"reviews": reviews,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// SearchPlaymates 搜索陪玩
// @Tags     Playmate
// @Summary  搜索陪玩
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    keyword  query    string  true "搜索关键词"
// @Param    page     query    int     false "页码"
// @Param    pageSize query    int     false "每页数量"
// @Success  200      {object} response.Response{data=[]model.Playmate, pagination=map[string]int64} "搜索成功"
// @Router   /playmate/playmates/search [get]
func (a *PlaymateApi) SearchPlaymates(c *gin.Context) {
	keyword := c.Query("keyword")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	search := request.PlaymateSearch{
		Keyword:  keyword,
		Page:     page,
		PageSize: pageSize,
	}

	playmates, total, err := service.ServiceGroupApp.PlaymateService.GetPlaymates(search)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": playmates,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "搜索成功", c)
}

// GetSearchSuggestions 获取搜索建议
// @Tags     Playmate
// @Summary  获取搜索建议
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    keyword  query    string  true "搜索关键词"
// @Success  200      {object} response.Response{data=[]string} "获取成功"
// @Router   /playmate/playmates/suggestions [get]
func (a *PlaymateApi) GetSearchSuggestions(c *gin.Context) {
	keyword := c.Query("keyword")

	suggestions, err := service.ServiceGroupApp.PlaymateService.GetSearchSuggestions(keyword)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(suggestions, "获取成功", c)
}

// GetPlaymateById 根据ID获取陪玩信息
// @Tags     Playmate
// @Summary  根据ID获取陪玩信息
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "陪玩ID"
// @Success  200  {object} response.Response{data=model.Playmate} "获取成功"
// @Router   /playmate/playmates/{id} [get]
func (a *PlaymateApi) GetPlaymateById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	playmate, err := service.ServiceGroupApp.PlaymateService.GetPlaymateById(uint(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(playmate, "获取成功", c)
}

// CreatePlaymate 创建陪玩
// @Tags     Playmate
// @Summary  创建陪玩
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body      model.Playmate  true "陪玩信息"
// @Success  200   {object}  response.Response{data=model.Playmate} "创建成功"
// @Router   /playmate/playmates [post]
func (a *PlaymateApi) CreatePlaymate(c *gin.Context) {
	var playmate model.Playmate
	if err := c.ShouldBindJSON(&playmate); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	createdPlaymate, err := service.ServiceGroupApp.PlaymateService.CreatePlaymate(playmate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(createdPlaymate, "创建成功", c)
}

// UpdatePlaymate 更新陪玩信息
// @Tags     Playmate
// @Summary  更新陪玩信息
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id    path      uint            true "陪玩ID"
// @Param    data  body      model.Playmate  true "陪玩信息"
// @Success  200   {object}  response.Response{data=model.Playmate} "更新成功"
// @Router   /playmate/playmates/{id} [put]
func (a *PlaymateApi) UpdatePlaymate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var playmate model.Playmate
	if err := c.ShouldBindJSON(&playmate); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	playmate.ID = uint(id)
	updatedPlaymate, err := service.ServiceGroupApp.PlaymateService.UpdatePlaymate(playmate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(updatedPlaymate, "更新成功", c)
}

// DeletePlaymate 删除陪玩
// @Tags     Playmate
// @Summary  删除陪玩
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id  path      uint    true "陪玩ID"
// @Success  200 {object}  response.Response{message=string} "删除成功"
// @Router   /playmate/playmates/{id} [delete]
func (a *PlaymateApi) DeletePlaymate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	err = service.ServiceGroupApp.PlaymateService.DeletePlaymate(uint(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// GetExpertVoice 获取专家语音
// @Tags     Playmate
// @Summary  获取专家语音
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "专家ID"
// @Success  200  {object} response.Response{data=map[string]string} "获取成功"
// @Router   /playmate/experts/{id}/voice [get]
func (a *PlaymateApi) GetExpertVoice(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	voice, err := service.ServiceGroupApp.PlaymateService.GetExpertVoice(uint(id))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(voice, "获取成功", c)
}

// GetSkills 获取技能列表
// @Tags     Playmate
// @Summary  获取技能列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]model.PlaymateSkill} "获取成功"
// @Router   /playmate/skills [get]
func (a *PlaymateApi) GetSkills(c *gin.Context) {
	skills, err := service.ServiceGroupApp.PlaymateService.GetSkills()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(skills, "获取成功", c)
}
