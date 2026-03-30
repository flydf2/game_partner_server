package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
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
		response.FailWithError(err, c)
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
		response.FailWithError(err, c)
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

	// 从上下文获取用户ID
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.FailWithMessage("未获取到用户ID", c)
		return
	}

	err = service.ServiceGroupApp.PlaymateService.FollowExpert(userID, uint(id))
	if err != nil {
		response.FailWithError(err, c)
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

	// 从上下文获取用户ID
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.FailWithMessage("未获取到用户ID", c)
		return
	}

	err = service.ServiceGroupApp.PlaymateService.UnfollowExpert(userID, uint(id))
	if err != nil {
		response.FailWithError(err, c)
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
		response.FailWithError(err, c)
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
		response.FailWithError(err, c)
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
		response.FailWithError(err, c)
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
		response.FailWithError(err, c)
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
// @Param    data  body      request.CreatePlaymateRequest  true "陪玩信息"
// @Success  200   {object}  response.Response{data=model.Playmate} "创建成功"
// @Router   /playmate/playmates [post]
func (a *PlaymateApi) CreatePlaymate(c *gin.Context) {
	var req request.CreatePlaymateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	playmate := model.Playmate{
		UserID:      req.UserID,
		Nickname:    req.Nickname,
		Avatar:      req.Avatar,
		Price:       req.Price,
		Tags:        req.Tags,
		Game:        req.Game,
		Rank:        req.Rank,
		Gender:      req.Gender,
		Description: req.Description,
		Level:       req.Level,
		Title:       req.Title,
		IsOnline:    false,
		Rating:      0,
		Likes:       0,
	}

	createdPlaymate, err := service.ServiceGroupApp.PlaymateService.CreatePlaymate(playmate)
	if err != nil {
		response.FailWithError(err, c)
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
// @Param    id    path      uint                     true "陪玩ID"
// @Param    data  body      request.UpdatePlaymateRequest  true "陪玩信息"
// @Success  200   {object}  response.Response{data=model.Playmate} "更新成功"
// @Router   /playmate/playmates/{id} [put]
func (a *PlaymateApi) UpdatePlaymate(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var req request.UpdatePlaymateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 先获取现有陪玩信息
	existingPlaymate, err := service.ServiceGroupApp.PlaymateService.GetPlaymateById(uint(id))
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	// 更新字段
	if req.Nickname != "" {
		existingPlaymate.Nickname = req.Nickname
	}
	if req.Avatar != "" {
		existingPlaymate.Avatar = req.Avatar
	}
	if req.Price > 0 {
		existingPlaymate.Price = req.Price
	}
	if req.Tags != "" {
		existingPlaymate.Tags = req.Tags
	}
	existingPlaymate.IsOnline = req.IsOnline
	if req.Game != "" {
		existingPlaymate.Game = req.Game
	}
	if req.Rank != "" {
		existingPlaymate.Rank = req.Rank
	}
	if req.Gender != "" {
		existingPlaymate.Gender = req.Gender
	}
	if req.Description != "" {
		existingPlaymate.Description = req.Description
	}
	if req.Level > 0 {
		existingPlaymate.Level = req.Level
	}
	if req.Title != "" {
		existingPlaymate.Title = req.Title
	}

	updatedPlaymate, err := service.ServiceGroupApp.PlaymateService.UpdatePlaymate(existingPlaymate)
	if err != nil {
		response.FailWithError(err, c)
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
		response.FailWithError(err, c)
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
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(voice, "获取成功", c)
}

// GetExpertStatus 获取专家状态
// @Tags     Playmate
// @Summary  获取专家状态
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "专家ID"
// @Success  200  {object} response.Response{data=map[string]interface{}} "获取成功"
// @Router   /playmate/experts/{id}/status [get]
func (a *PlaymateApi) GetExpertStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	status, err := service.ServiceGroupApp.PlaymateService.GetExpertStatus(uint(id))
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(status, "获取成功", c)
}

// GetSkills 获取技能列表
// @Tags     Playmate
// @Summary  获取技能列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    game     query    string  false "游戏"
// @Param    level    query    string  false "等级"
// @Param    keyword  query    string  false "关键词"
// @Param    page     query    int     false "页码"
// @Param    pageSize query    int     false "每页数量"
// @Success  200      {object} response.Response{data=[]model.PlaymateSkill, pagination=map[string]int64} "获取成功"
// @Router   /playmate/skills [get]
func (a *PlaymateApi) GetSkills(c *gin.Context) {
	var search request.SkillSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	skills, total, err := service.ServiceGroupApp.PlaymateService.GetSkills(search)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	// 处理pageSize为0的情况，避免除以零错误
	pageSize := search.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}

	response.OkWithDetailed(gin.H{
		"data": skills,
		"pagination": gin.H{
			"currentPage": search.Page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// AddSkill 添加技能
// @Tags     Playmate
// @Summary  添加技能
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.AddSkillRequest  true "技能信息"
// @Success  200   {object}  response.Response{data=model.PlaymateSkill} "添加成功"
// @Router   /playmate/skills [post]
func (a *PlaymateApi) AddSkill(c *gin.Context) {
	var req request.AddSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 从上下文获取用户ID
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.FailWithMessage("未获取到用户ID", c)
		return
	}

	skill, err := service.ServiceGroupApp.PlaymateService.AddSkill(userID, req)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(skill, "添加成功", c)
}

// UpdateSkill 更新技能
// @Tags     Playmate
// @Summary  更新技能
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id    path      uint                        true "技能ID"
// @Param    data  body      request.UpdateSkillRequest  true "技能信息"
// @Success  200   {object}  response.Response{data=model.PlaymateSkill} "更新成功"
// @Router   /playmate/skills/{id} [put]
func (a *PlaymateApi) UpdateSkill(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var req request.UpdateSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	skill, err := service.ServiceGroupApp.PlaymateService.UpdateSkill(uint(id), req)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(skill, "更新成功", c)
}

// DeleteSkill 删除技能
// @Tags     Playmate
// @Summary  删除技能
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id  path      uint    true "技能ID"
// @Success  200 {object}  response.Response{msg=string} "删除成功"
// @Router   /playmate/skills/{id} [delete]
func (a *PlaymateApi) DeleteSkill(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	err = service.ServiceGroupApp.PlaymateService.DeleteSkill(uint(id))
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// GetLeaderboard 获取排行榜
// @Tags     Playmate
// @Summary  获取陪玩排行榜
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    game     query    string  false "游戏"
// @Param    page     query    int     false "页码"
// @Param    pageSize query    int     false "每页数量"
// @Success  200      {object} response.Response{data=[]model.Playmate, pagination=map[string]int64} "获取成功"
// @Router   /playmate/playmates/leaderboard [get]
func (a *PlaymateApi) GetLeaderboard(c *gin.Context) {
	var search request.LeaderboardSearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	playmates, total, err := service.ServiceGroupApp.PlaymateService.GetLeaderboard(search)
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
		"data": playmates,
		"pagination": gin.H{
			"currentPage": search.Page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// GetMatchHistory 获取匹配历史列表
// @Tags     Playmate
// @Summary  获取用户的匹配历史
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    page     query    int     false "页码"
// @Param    pageSize query    int     false "每页数量"
// @Success  200      {object} response.Response{data=[]model.MatchHistory, pagination=map[string]int64} "获取成功"
// @Router   /playmate/match-history [get]
func (a *PlaymateApi) GetMatchHistory(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	// 从上下文获取用户ID
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.FailWithMessage("未获取到用户ID", c)
		return
	}

	histories, total, err := service.ServiceGroupApp.PlaymateService.GetMatchHistory(userID, page, pageSize)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": histories,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// GetMatchHistoryMatches 获取匹配历史匹配记录
// @Tags     Playmate
// @Summary  获取用户的匹配历史匹配记录
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    page     query    int     false "页码"
// @Param    pageSize query    int     false "每页数量"
// @Success  200      {object} response.Response{data=[]model.MatchHistory, pagination=map[string]int64} "获取成功"
// @Router   /playmate/match-history/matches [get]
func (a *PlaymateApi) GetMatchHistoryMatches(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))

	// 从上下文获取用户ID
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.FailWithMessage("未获取到用户ID", c)
		return
	}

	histories, total, err := service.ServiceGroupApp.PlaymateService.GetMatchHistory(userID, page, pageSize)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": histories,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// GetMatchHistoryById 获取匹配历史详情
// @Tags     Playmate
// @Summary  根据ID获取匹配历史详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path      uint    true "匹配历史ID"
// @Success  200  {object} response.Response{data=model.MatchHistory} "获取成功"
// @Router   /playmate/match-history/{id} [get]
func (a *PlaymateApi) GetMatchHistoryById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	history, err := service.ServiceGroupApp.PlaymateService.GetMatchHistoryById(uint(id))
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(history, "获取成功", c)
}
