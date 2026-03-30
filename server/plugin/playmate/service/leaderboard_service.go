package service

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
)

// LeaderboardService 排行榜服务
type LeaderboardService struct{}

// GetLeaderboards 获取排行榜列表
func (s *LeaderboardService) GetLeaderboards(search request.LeaderboardSearch) ([]model.Leaderboard, int64, error) {
	var leaderboards []model.Leaderboard
	var total int64

	db := global.GVA_DB
	query := db.Model(&model.Leaderboard{})

	// 按类型过滤
	if search.Type != "" {
		query = query.Where("type = ?", search.Type)
	}

	// 按游戏过滤
	if search.Game != "" {
		query = query.Where("game = ?", search.Game)
	}

	// 只查询启用的榜单
	query = query.Where("status = ?", 1)

	// 计算总数
	query.Count(&total)

	// 按排序顺序和时间排序
	query = query.Order("sort_order ASC, created_at DESC")

	// 应用分页
	page := search.Page
	if page <= 0 {
		page = 1
	}

	pageSize := search.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	// 执行查询
	if err := query.Find(&leaderboards).Error; err != nil {
		return nil, 0, err
	}

	return leaderboards, total, nil
}

// GetLeaderboardById 根据ID获取排行榜详情
func (s *LeaderboardService) GetLeaderboardById(id uint) (model.Leaderboard, error) {
	var leaderboard model.Leaderboard
	if err := global.GVA_DB.First(&leaderboard, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Leaderboard{}, response.NewPlaymateError(response.ErrPlaymateNotFound)
		}
		return model.Leaderboard{}, err
	}
	return leaderboard, nil
}

// GetLeaderboardWithItems 获取排行榜及其条目
func (s *LeaderboardService) GetLeaderboardWithItems(id uint, page, pageSize int) (model.Leaderboard, []model.LeaderboardItem, int64, error) {
	var leaderboard model.Leaderboard
	var items []model.LeaderboardItem
	var total int64

	// 获取排行榜信息
	if err := global.GVA_DB.First(&leaderboard, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Leaderboard{}, nil, 0, response.NewPlaymateError(response.ErrPlaymateNotFound)
		}
		return model.Leaderboard{}, nil, 0, err
	}

	// 获取排行榜条目
	query := global.GVA_DB.Model(&model.LeaderboardItem{}).Where("leaderboard_id = ?", id)

	// 计算总数
	query.Count(&total)

	// 应用分页
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// 查询条目并关联陪玩信息
	if err := query.Preload("Playmate").Order("rank ASC").Offset(offset).Limit(pageSize).Find(&items).Error; err != nil {
		return leaderboard, nil, 0, err
	}

	return leaderboard, items, total, nil
}

// CreateLeaderboard 创建排行榜
func (s *LeaderboardService) CreateLeaderboard(req request.CreateLeaderboardRequest) (model.Leaderboard, error) {
	leaderboard := model.Leaderboard{
		Name:        req.Name,
		Type:        model.LeaderboardType(req.Type),
		Game:        req.Game,
		Description: req.Description,
		Status:      req.Status,
		SortOrder:   req.SortOrder,
	}

	// 解析时间
	if req.StartTime != "" {
		startTime, err := time.Parse("2006-01-02 15:04:05", req.StartTime)
		if err == nil {
			leaderboard.StartTime = startTime
		}
	}
	if req.EndTime != "" {
		endTime, err := time.Parse("2006-01-02 15:04:05", req.EndTime)
		if err == nil {
			leaderboard.EndTime = endTime
		}
	}

	if err := global.GVA_DB.Create(&leaderboard).Error; err != nil {
		return model.Leaderboard{}, err
	}

	return leaderboard, nil
}

// UpdateLeaderboard 更新排行榜
func (s *LeaderboardService) UpdateLeaderboard(id uint, req request.UpdateLeaderboardRequest) (model.Leaderboard, error) {
	var leaderboard model.Leaderboard
	if err := global.GVA_DB.First(&leaderboard, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Leaderboard{}, response.NewPlaymateError(response.ErrPlaymateNotFound)
		}
		return model.Leaderboard{}, err
	}

	// 更新字段
	if req.Name != "" {
		leaderboard.Name = req.Name
	}
	if req.Type != "" {
		leaderboard.Type = model.LeaderboardType(req.Type)
	}
	if req.Game != "" {
		leaderboard.Game = req.Game
	}
	if req.Description != "" {
		leaderboard.Description = req.Description
	}
	if req.StartTime != "" {
		startTime, err := time.Parse("2006-01-02 15:04:05", req.StartTime)
		if err == nil {
			leaderboard.StartTime = startTime
		}
	}
	if req.EndTime != "" {
		endTime, err := time.Parse("2006-01-02 15:04:05", req.EndTime)
		if err == nil {
			leaderboard.EndTime = endTime
		}
	}
	leaderboard.Status = req.Status
	leaderboard.SortOrder = req.SortOrder

	if err := global.GVA_DB.Save(&leaderboard).Error; err != nil {
		return model.Leaderboard{}, err
	}

	return leaderboard, nil
}

// DeleteLeaderboard 删除排行榜
func (s *LeaderboardService) DeleteLeaderboard(id uint) error {
	// 先删除关联的条目
	if err := global.GVA_DB.Where("leaderboard_id = ?", id).Delete(&model.LeaderboardItem{}).Error; err != nil {
		return err
	}

	// 删除排行榜
	if err := global.GVA_DB.Delete(&model.Leaderboard{}, id).Error; err != nil {
		return err
	}

	return nil
}

// AddLeaderboardItem 添加排行榜条目
func (s *LeaderboardService) AddLeaderboardItem(leaderboardID, playmateID uint, rank int, score float64) (model.LeaderboardItem, error) {
	item := model.LeaderboardItem{
		LeaderboardID: leaderboardID,
		PlaymateID:    playmateID,
		Rank:          rank,
		Score:         score,
	}

	// 获取陪玩的统计数据
	var playmate model.Playmate
	if err := global.GVA_DB.First(&playmate, playmateID).Error; err != nil {
		return model.LeaderboardItem{}, err
	}

	item.Rating = playmate.Rating
	item.Likes = playmate.Likes

	// 统计订单数量
	var orderCount int64
	global.GVA_DB.Model(&model.Order{}).Where("playmate_id = ?", playmateID).Count(&orderCount)
	item.OrderCount = int(orderCount)

	if err := global.GVA_DB.Create(&item).Error; err != nil {
		return model.LeaderboardItem{}, err
	}

	return item, nil
}

// RemoveLeaderboardItem 移除排行榜条目
func (s *LeaderboardService) RemoveLeaderboardItem(itemID uint) error {
	if err := global.GVA_DB.Delete(&model.LeaderboardItem{}, itemID).Error; err != nil {
		return err
	}
	return nil
}

// UpdateLeaderboardItemRank 更新排行榜条目排名
func (s *LeaderboardService) UpdateLeaderboardItemRank(itemID uint, rank int) error {
	if err := global.GVA_DB.Model(&model.LeaderboardItem{}).Where("id = ?", itemID).Update("rank", rank).Error; err != nil {
		return err
	}
	return nil
}

// GetLeaderboardItems 获取排行榜条目列表
func (s *LeaderboardService) GetLeaderboardItems(leaderboardID uint, page, pageSize int) ([]model.LeaderboardItem, int64, error) {
	var items []model.LeaderboardItem
	var total int64

	query := global.GVA_DB.Model(&model.LeaderboardItem{}).Where("leaderboard_id = ?", leaderboardID)

	// 计算总数
	query.Count(&total)

	// 应用分页
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	// 查询条目并关联陪玩信息
	if err := query.Preload("Playmate").Order("rank ASC").Offset(offset).Limit(pageSize).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

// GenerateLeaderboard 生成排行榜（根据陪玩数据自动计算排名）
func (s *LeaderboardService) GenerateLeaderboard(leaderboardID uint) error {
	// 获取排行榜信息
	var leaderboard model.Leaderboard
	if err := global.GVA_DB.First(&leaderboard, leaderboardID).Error; err != nil {
		return err
	}

	// 删除旧的条目
	if err := global.GVA_DB.Where("leaderboard_id = ?", leaderboardID).Delete(&model.LeaderboardItem{}).Error; err != nil {
		return err
	}

	// 根据游戏过滤陪玩
	var playmates []model.Playmate
	query := global.GVA_DB.Model(&model.Playmate{})
	if leaderboard.Game != "" {
		query = query.Where("game = ?", leaderboard.Game)
	}

	// 按评分和点赞数排序
	if err := query.Order("rating DESC, likes DESC").Limit(100).Find(&playmates).Error; err != nil {
		return err
	}

	// 创建新的排行榜条目
	for i, playmate := range playmates {
		item := model.LeaderboardItem{
			LeaderboardID: leaderboardID,
			PlaymateID:    playmate.ID,
			Rank:          i + 1,
			Score:         playmate.Rating,
			Rating:        playmate.Rating,
			Likes:         playmate.Likes,
		}

		// 统计订单数量
		var orderCount int64
		global.GVA_DB.Model(&model.Order{}).Where("playmate_id = ?", playmate.ID).Count(&orderCount)
		item.OrderCount = int(orderCount)

		if err := global.GVA_DB.Create(&item).Error; err != nil {
			return err
		}
	}

	return nil
}
