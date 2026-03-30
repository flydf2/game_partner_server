package service

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
)

// TournamentService 赛事服务
type TournamentService struct{}

// GetTournamentList 获取赛事列表
func (s *TournamentService) GetTournamentList(search request.TournamentSearch) ([]model.Tournament, int64, error) {
	var tournaments []model.Tournament
	var total int64

	db := global.GVA_DB.Model(&model.Tournament{})

	// 应用搜索条件
	if search.Status != "" {
		db = db.Where("status = ?", search.Status)
	}
	if search.Game != "" {
		db = db.Where("game = ?", search.Game)
	}
	if search.GameID > 0 {
		db = db.Where("game_id = ?", search.GameID)
	}
	if search.Keyword != "" {
		db = db.Where("title LIKE ? OR description LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%")
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 设置分页默认值
	if search.Page <= 0 {
		search.Page = 1
	}
	if search.PageSize <= 0 {
		search.PageSize = 10
	}

	offset := (search.Page - 1) * search.PageSize

	// 执行查询
	if err := db.Offset(offset).Limit(search.PageSize).Order("created_at DESC").Find(&tournaments).Error; err != nil {
		return nil, 0, err
	}

	return tournaments, total, nil
}

// GetTournamentByID 根据ID获取赛事详情
func (s *TournamentService) GetTournamentByID(id uint) (model.Tournament, error) {
	var tournament model.Tournament
	if err := global.GVA_DB.First(&tournament, id).Error; err != nil {
		return model.Tournament{}, err
	}
	return tournament, nil
}

// GetTournamentTeams 获取赛事参赛队伍
func (s *TournamentService) GetTournamentTeams(tournamentID uint) ([]model.TournamentTeam, error) {
	var teams []model.TournamentTeam
	if err := global.GVA_DB.Where("tournament_id = ?", tournamentID).
		Order("CASE WHEN rank = '冠军' THEN 1 WHEN rank = '亚军' THEN 2 WHEN rank = '季军' THEN 3 ELSE 4 END, id ASC").
		Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

// GetTournamentMatches 获取赛事比赛列表
func (s *TournamentService) GetTournamentMatches(tournamentID uint) ([]model.TournamentMatch, error) {
	var matches []model.TournamentMatch
	if err := global.GVA_DB.Where("tournament_id = ?", tournamentID).
		Order("match_time ASC").
		Find(&matches).Error; err != nil {
		return nil, err
	}
	return matches, nil
}

// JoinTournament 报名参赛
func (s *TournamentService) JoinTournament(req request.JoinTournamentRequest, userID uint) error {
	// 检查赛事是否存在
	var tournament model.Tournament
	if err := global.GVA_DB.First(&tournament, req.TournamentID).Error; err != nil {
		return errors.New("赛事不存在")
	}

	// 检查赛事状态
	if tournament.Status != "upcoming" {
		return errors.New("该赛事不在报名阶段")
	}

	// 检查报名时间
	now := time.Now()
	if now.Before(tournament.RegisterStart) {
		return errors.New("报名尚未开始")
	}
	if now.After(tournament.RegisterEnd) {
		return errors.New("报名已结束")
	}

	// 检查是否已报名
	var existingRegistration model.TournamentRegistration
	if err := global.GVA_DB.Where("tournament_id = ? AND user_id = ?", req.TournamentID, userID).
		First(&existingRegistration).Error; err == nil {
		return errors.New("您已经报名过该赛事")
	}

	// 检查队伍名称是否已被使用
	var existingTeam model.TournamentTeam
	if err := global.GVA_DB.Where("tournament_id = ? AND name = ?", req.TournamentID, req.TeamName).
		First(&existingTeam).Error; err == nil {
		return errors.New("该队伍名称已被使用")
	}

	// 创建报名记录
	registration := model.TournamentRegistration{
		TournamentID: req.TournamentID,
		UserID:       userID,
		TeamName:     req.TeamName,
		ContactInfo:  req.ContactInfo,
		MembersInfo:  req.MembersInfo,
		Status:       "pending",
	}

	if err := global.GVA_DB.Create(&registration).Error; err != nil {
		return errors.New("报名失败，请稍后重试")
	}

	return nil
}

// CreateTournament 创建赛事
func (s *TournamentService) CreateTournament(tournament *model.Tournament) error {
	return global.GVA_DB.Create(tournament).Error
}

// UpdateTournament 更新赛事
func (s *TournamentService) UpdateTournament(tournament *model.Tournament) error {
	return global.GVA_DB.Save(tournament).Error
}

// DeleteTournament 删除赛事
func (s *TournamentService) DeleteTournament(id uint) error {
	return global.GVA_DB.Delete(&model.Tournament{}, id).Error
}

// CreateTournamentTeam 创建参赛队伍
func (s *TournamentService) CreateTournamentTeam(team *model.TournamentTeam) error {
	return global.GVA_DB.Create(team).Error
}

// UpdateTournamentTeam 更新参赛队伍
func (s *TournamentService) UpdateTournamentTeam(team *model.TournamentTeam) error {
	return global.GVA_DB.Save(team).Error
}

// DeleteTournamentTeam 删除参赛队伍
func (s *TournamentService) DeleteTournamentTeam(id uint) error {
	return global.GVA_DB.Delete(&model.TournamentTeam{}, id).Error
}

// CreateTournamentMatch 创建比赛
func (s *TournamentService) CreateTournamentMatch(match *model.TournamentMatch) error {
	return global.GVA_DB.Create(match).Error
}

// UpdateTournamentMatch 更新比赛
func (s *TournamentService) UpdateTournamentMatch(match *model.TournamentMatch) error {
	return global.GVA_DB.Save(match).Error
}

// DeleteTournamentMatch 删除比赛
func (s *TournamentService) DeleteTournamentMatch(id uint) error {
	return global.GVA_DB.Delete(&model.TournamentMatch{}, id).Error
}

// ApproveRegistration 审核通过报名
func (s *TournamentService) ApproveRegistration(registrationID uint) error {
	var registration model.TournamentRegistration
	if err := global.GVA_DB.First(&registration, registrationID).Error; err != nil {
		return errors.New("报名记录不存在")
	}

	// 更新报名状态
	registration.Status = "approved"
	if err := global.GVA_DB.Save(&registration).Error; err != nil {
		return err
	}

	// 创建参赛队伍
	team := model.TournamentTeam{
		TournamentID: registration.TournamentID,
		Name:         registration.TeamName,
		LeaderID:     registration.UserID,
		Status:       "approved",
	}

	// 解析成员信息
	if registration.MembersInfo != "" {
		var membersInfo map[string]interface{}
		if err := json.Unmarshal([]byte(registration.MembersInfo), &membersInfo); err == nil {
			if members, ok := membersInfo["members"].([]interface{}); ok {
				team.Members = len(members)
			}
		}
	}

	if err := global.GVA_DB.Create(&team).Error; err != nil {
		return err
	}

	// 更新赛事参赛人数
	return global.GVA_DB.Model(&model.Tournament{}).
		Where("id = ?", registration.TournamentID).
		UpdateColumn("participants", global.GVA_DB.Raw("participants + 1")).Error
}

// RejectRegistration 拒绝报名
func (s *TournamentService) RejectRegistration(registrationID uint, remark string) error {
	return global.GVA_DB.Model(&model.TournamentRegistration{}).
		Where("id = ?", registrationID).
		Updates(map[string]interface{}{
			"status":        "rejected",
			"review_remark": remark,
		}).Error
}

// GetTournamentRegistrations 获取赛事报名列表
func (s *TournamentService) GetTournamentRegistrations(tournamentID uint, status string) ([]model.TournamentRegistration, error) {
	var registrations []model.TournamentRegistration
	db := global.GVA_DB.Where("tournament_id = ?", tournamentID)
	if status != "" {
		db = db.Where("status = ?", status)
	}
	if err := db.Order("created_at DESC").Find(&registrations).Error; err != nil {
		return nil, err
	}
	return registrations, nil
}
