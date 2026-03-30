package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

type TournamentApi struct{}

// GetTournamentList 获取赛事列表
// @Tags     Tournament
// @Summary  获取赛事列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    status   query    string  false "赛事状态 upcoming-报名中 ongoing-进行中 completed-已结束"
// @Param    game     query    string  false "游戏名称"
// @Param    gameId   query    uint    false "游戏ID"
// @Param    keyword  query    string  false "关键词"
// @Param    page     query    int     false "页码"
// @Param    pageSize query    int     false "每页数量"
// @Success  200      {object} response.Response{data=[]model.Tournament,total=int64} "获取成功"
// @Router   /playmate/tournaments [get]
func (a *TournamentApi) GetTournamentList(c *gin.Context) {
	var search request.TournamentSearch
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

	tournaments, total, err := service.ServiceGroupApp.TournamentService.GetTournamentList(search)
	if err != nil {
		response.FailWithMessage("获取赛事列表失败", c)
		return
	}

	response.OkWithDetailed(gin.H{
		"list":  tournaments,
		"total": total,
	}, "获取成功", c)
}

// GetTournamentDetail 获取赛事详情
// @Tags     Tournament
// @Summary  获取赛事详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path     uint true "赛事ID"
// @Success  200  {object} response.Response{data=model.Tournament} "获取成功"
// @Router   /playmate/tournaments/{id} [get]
func (a *TournamentApi) GetTournamentDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("无效的赛事ID", c)
		return
	}

	tournament, err := service.ServiceGroupApp.TournamentService.GetTournamentByID(uint(id))
	if err != nil {
		response.FailWithMessage("赛事不存在", c)
		return
	}

	response.OkWithDetailed(tournament, "获取成功", c)
}

// GetTournamentTeams 获取赛事参赛队伍
// @Tags     Tournament
// @Summary  获取赛事参赛队伍
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path     uint true "赛事ID"
// @Success  200  {object} response.Response{data=[]model.TournamentTeam} "获取成功"
// @Router   /playmate/tournaments/{id}/teams [get]
func (a *TournamentApi) GetTournamentTeams(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("无效的赛事ID", c)
		return
	}

	teams, err := service.ServiceGroupApp.TournamentService.GetTournamentTeams(uint(id))
	if err != nil {
		response.FailWithMessage("获取参赛队伍失败", c)
		return
	}

	response.OkWithDetailed(teams, "获取成功", c)
}

// GetTournamentMatches 获取赛事比赛列表
// @Tags     Tournament
// @Summary  获取赛事比赛列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path     uint true "赛事ID"
// @Success  200  {object} response.Response{data=[]model.TournamentMatch} "获取成功"
// @Router   /playmate/tournaments/{id}/matches [get]
func (a *TournamentApi) GetTournamentMatches(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("无效的赛事ID", c)
		return
	}

	matches, err := service.ServiceGroupApp.TournamentService.GetTournamentMatches(uint(id))
	if err != nil {
		response.FailWithMessage("获取比赛列表失败", c)
		return
	}

	response.OkWithDetailed(matches, "获取成功", c)
}

// JoinTournament 报名参赛
// @Tags     Tournament
// @Summary  报名参赛
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.JoinTournamentRequest true "报名信息"
// @Success  200  {object} response.Response{msg=string} "报名成功"
// @Router   /playmate/tournaments/join [post]
func (a *TournamentApi) JoinTournament(c *gin.Context) {
	var req request.JoinTournamentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	// 获取当前用户ID
	userID := utils.GetUserID(c)
	if userID == 0 {
		response.FailWithMessage("请先登录", c)
		return
	}

	err := service.ServiceGroupApp.TournamentService.JoinTournament(req, userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("报名成功", c)
}

// CreateTournament 创建赛事
// @Tags     Tournament
// @Summary  创建赛事
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     model.Tournament true "赛事信息"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /playmate/tournaments [post]
func (a *TournamentApi) CreateTournament(c *gin.Context) {
	var tournament model.Tournament
	if err := c.ShouldBindJSON(&tournament); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	if err := service.ServiceGroupApp.TournamentService.CreateTournament(&tournament); err != nil {
		response.FailWithMessage("创建赛事失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

// UpdateTournament 更新赛事
// @Tags     Tournament
// @Summary  更新赛事
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path     uint             true "赛事ID"
// @Param    data body     model.Tournament true "赛事信息"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /playmate/tournaments/{id} [put]
func (a *TournamentApi) UpdateTournament(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("无效的赛事ID", c)
		return
	}

	var tournament model.Tournament
	if err := c.ShouldBindJSON(&tournament); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	tournament.ID = uint(id)
	if err := service.ServiceGroupApp.TournamentService.UpdateTournament(&tournament); err != nil {
		response.FailWithMessage("更新赛事失败", c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// DeleteTournament 删除赛事
// @Tags     Tournament
// @Summary  删除赛事
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path     uint true "赛事ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /playmate/tournaments/{id} [delete]
func (a *TournamentApi) DeleteTournament(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("无效的赛事ID", c)
		return
	}

	if err := service.ServiceGroupApp.TournamentService.DeleteTournament(uint(id)); err != nil {
		response.FailWithMessage("删除赛事失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// CreateTournamentTeam 创建参赛队伍
// @Tags     Tournament
// @Summary  创建参赛队伍
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     model.TournamentTeam true "队伍信息"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /playmate/tournaments/teams [post]
func (a *TournamentApi) CreateTournamentTeam(c *gin.Context) {
	var team model.TournamentTeam
	if err := c.ShouldBindJSON(&team); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	if err := service.ServiceGroupApp.TournamentService.CreateTournamentTeam(&team); err != nil {
		response.FailWithMessage("创建队伍失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

// UpdateTournamentTeam 更新参赛队伍
// @Tags     Tournament
// @Summary  更新参赛队伍
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path     uint                 true "队伍ID"
// @Param    data body     model.TournamentTeam true "队伍信息"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /playmate/tournaments/teams/{id} [put]
func (a *TournamentApi) UpdateTournamentTeam(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("无效的队伍ID", c)
		return
	}

	var team model.TournamentTeam
	if err := c.ShouldBindJSON(&team); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	team.ID = uint(id)
	if err := service.ServiceGroupApp.TournamentService.UpdateTournamentTeam(&team); err != nil {
		response.FailWithMessage("更新队伍失败", c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// DeleteTournamentTeam 删除参赛队伍
// @Tags     Tournament
// @Summary  删除参赛队伍
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path     uint true "队伍ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /playmate/tournaments/teams/{id} [delete]
func (a *TournamentApi) DeleteTournamentTeam(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("无效的队伍ID", c)
		return
	}

	if err := service.ServiceGroupApp.TournamentService.DeleteTournamentTeam(uint(id)); err != nil {
		response.FailWithMessage("删除队伍失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// CreateTournamentMatch 创建比赛
// @Tags     Tournament
// @Summary  创建比赛
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     model.TournamentMatch true "比赛信息"
// @Success  200  {object} response.Response{msg=string} "创建成功"
// @Router   /playmate/tournaments/matches [post]
func (a *TournamentApi) CreateTournamentMatch(c *gin.Context) {
	var match model.TournamentMatch
	if err := c.ShouldBindJSON(&match); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	if err := service.ServiceGroupApp.TournamentService.CreateTournamentMatch(&match); err != nil {
		response.FailWithMessage("创建比赛失败", c)
		return
	}

	response.OkWithMessage("创建成功", c)
}

// UpdateTournamentMatch 更新比赛
// @Tags     Tournament
// @Summary  更新比赛
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id   path     uint                  true "比赛ID"
// @Param    data body     model.TournamentMatch true "比赛信息"
// @Success  200  {object} response.Response{msg=string} "更新成功"
// @Router   /playmate/tournaments/matches/{id} [put]
func (a *TournamentApi) UpdateTournamentMatch(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("无效的比赛ID", c)
		return
	}

	var match model.TournamentMatch
	if err := c.ShouldBindJSON(&match); err != nil {
		response.FailWithMessage("参数错误: "+err.Error(), c)
		return
	}

	match.ID = uint(id)
	if err := service.ServiceGroupApp.TournamentService.UpdateTournamentMatch(&match); err != nil {
		response.FailWithMessage("更新比赛失败", c)
		return
	}

	response.OkWithMessage("更新成功", c)
}

// DeleteTournamentMatch 删除比赛
// @Tags     Tournament
// @Summary  删除比赛
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path     uint true "比赛ID"
// @Success  200  {object} response.Response{msg=string} "删除成功"
// @Router   /playmate/tournaments/matches/{id} [delete]
func (a *TournamentApi) DeleteTournamentMatch(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.FailWithMessage("无效的比赛ID", c)
		return
	}

	if err := service.ServiceGroupApp.TournamentService.DeleteTournamentMatch(uint(id)); err != nil {
		response.FailWithMessage("删除比赛失败", c)
		return
	}

	response.OkWithMessage("删除成功", c)
}
