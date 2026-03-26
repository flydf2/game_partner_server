package service

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
)

// PlaymateService 陪玩服务
type PlaymateService struct{}

// GetPlaymates 获取陪玩列表
func (s *PlaymateService) GetPlaymates(search request.PlaymateSearch) ([]model.Playmate, int64, error) {
	var playmates []model.Playmate
	var total int64

	db := global.GVA_DB
	query := db.Model(&model.Playmate{})

	// 应用过滤条件
	if search.Game != "" {
		query = query.Where("game = ?", search.Game)
	}

	if search.Online != nil {
		query = query.Where("is_online = ?", *search.Online)
	}

	if search.PriceRange != "" {
		if search.PriceRange == "100+" {
			query = query.Where("price >= ?", 100)
		} else {
			rangeParts := strings.Split(search.PriceRange, "-")
			if len(rangeParts) == 2 {
				minPrice := rangeParts[0]
				maxPrice := rangeParts[1]
				query = query.Where("price >= ? AND price <= ?", minPrice, maxPrice)
			}
		}
	}

	if search.Rank != "" {
		query = query.Where("rank = ?", search.Rank)
	}

	if search.Gender != "" {
		query = query.Where("gender = ?", search.Gender)
	}

	if search.Keyword != "" {
		keyword := fmt.Sprintf("%%%s%%", search.Keyword)
		query = query.Where("nickname LIKE ? OR tags LIKE ?", keyword, keyword)
	}

	// 计算总数
	query.Count(&total)

	// 应用排序
	switch search.SortBy {
	case "rating":
		query = query.Order("rating DESC")
	case "price_asc":
		query = query.Order("price ASC")
	case "price_desc":
		query = query.Order("price DESC")
	case "newest":
		query = query.Order("created_at DESC")
	default:
		query = query.Order("created_at DESC")
	}

	// 应用分页
	page := search.Page
	if page <= 0 {
		page = 1
	}

	pageSize := search.PageSize
	if pageSize <= 0 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	// 执行查询
	if err := query.Find(&playmates).Error; err != nil {
		return nil, 0, err
	}

	return playmates, total, nil
}

// GetPlaymateById 根据ID获取陪玩详情
func (s *PlaymateService) GetPlaymateById(id uint) (model.Playmate, error) {
	var playmate model.Playmate
	if err := global.GVA_DB.First(&playmate, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Playmate{}, errors.New("陪玩不存在")
		}
		return model.Playmate{}, err
	}
	return playmate, nil
}

// CreatePlaymate 创建陪玩
func (s *PlaymateService) CreatePlaymate(playmate model.Playmate) (model.Playmate, error) {
	if err := global.GVA_DB.Create(&playmate).Error; err != nil {
		return model.Playmate{}, err
	}
	return playmate, nil
}

// UpdatePlaymate 更新陪玩信息
func (s *PlaymateService) UpdatePlaymate(playmate model.Playmate) (model.Playmate, error) {
	if err := global.GVA_DB.Save(&playmate).Error; err != nil {
		return model.Playmate{}, err
	}
	return playmate, nil
}

// DeletePlaymate 删除陪玩
func (s *PlaymateService) DeletePlaymate(id uint) error {
	if err := global.GVA_DB.Delete(&model.Playmate{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetExpertDetail 获取专家详情
func (s *PlaymateService) GetExpertDetail(id uint) (map[string]interface{}, error) {
	playmate, err := s.GetPlaymateById(id)
	if err != nil {
		return nil, err
	}

	// 获取技能列表
	var skills []model.PlaymateSkill
	if err := global.GVA_DB.Where("playmate_id = ?", id).Find(&skills).Error; err != nil {
		return nil, err
	}

	// 获取语音介绍
	var voiceIntroduction model.PlaymateVoiceIntroduction
	if err := global.GVA_DB.Where("playmate_id = ?", id).First(&voiceIntroduction).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	// 获取评价列表
	var reviews []model.Review
	if err := global.GVA_DB.Where("playmate_id = ?", id).Order("created_at DESC").Limit(10).Find(&reviews).Error; err != nil {
		return nil, err
	}

	// 构建响应
	detail := map[string]interface{}{
		"id":          playmate.ID,
		"nickname":    playmate.Nickname,
		"avatar":      playmate.Avatar,
		"rating":      playmate.Rating,
		"price":       playmate.Price,
		"likes":       playmate.Likes,
		"tags":        strings.Split(playmate.Tags, ","),
		"isOnline":    playmate.IsOnline,
		"game":        playmate.Game,
		"rank":        playmate.Rank,
		"gender":      playmate.Gender,
		"description": playmate.Description,
		"level":       playmate.Level,
		"title":       playmate.Title,
		"stats": map[string]interface{}{
			"winRate":   85,
			"followers": 12000,
			"rating":    playmate.Rating,
		},
		"skills": skills,
		"voiceIntroduction": map[string]interface{}{
			"url":      voiceIntroduction.URL,
			"duration": voiceIntroduction.Duration,
		},
		"reviews": reviews,
	}

	return detail, nil
}

// FollowExpert 关注专家
func (s *PlaymateService) FollowExpert(userID, expertID uint) error {
	// 检查是否已经关注
	var follow model.UserFollow
	result := global.GVA_DB.Where("user_id = ? AND follow_id = ?", userID, expertID).First(&follow)
	if result.Error == nil {
		return errors.New("已经关注过该专家")
	}

	// 创建关注记录
	follow = model.UserFollow{
		UserID:   userID,
		FollowID: expertID,
	}

	if err := global.GVA_DB.Create(&follow).Error; err != nil {
		return err
	}

	return nil
}

// UnfollowExpert 取消关注专家
func (s *PlaymateService) UnfollowExpert(userID, expertID uint) error {
	result := global.GVA_DB.Where("user_id = ? AND follow_id = ?", userID, expertID).Delete(&model.UserFollow{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("未关注该专家")
	}

	return nil
}

// GetExpertReviews 获取专家评价
func (s *PlaymateService) GetExpertReviews(expertID uint, page, pageSize int) ([]model.Review, int64, error) {
	var reviews []model.Review
	var total int64

	query := global.GVA_DB.Model(&model.Review{}).Where("playmate_id = ?", expertID)

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

	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&reviews).Error; err != nil {
		return nil, 0, err
	}

	return reviews, total, nil
}

// GetSearchSuggestions 获取搜索建议
func (s *PlaymateService) GetSearchSuggestions(keyword string) ([]string, error) {
	// 这里可以实现更复杂的搜索建议逻辑
	// 暂时返回固定的建议
	suggestions := []string{
		"王者荣耀陪玩",
		"LOL上分",
		"萌妹陪玩",
		"技术流大神",
		"绝地求生带飞",
		"原神探索",
		"温柔语聊",
		"国服第一",
	}

	// 过滤包含关键词的建议
	var filteredSuggestions []string
	for _, suggestion := range suggestions {
		if strings.Contains(suggestion, keyword) {
			filteredSuggestions = append(filteredSuggestions, suggestion)
		}
	}

	return filteredSuggestions, nil
}

// GetExpertVoice 获取专家语音
func (s *PlaymateService) GetExpertVoice(expertID uint) (map[string]string, error) {
	var voiceIntroduction model.PlaymateVoiceIntroduction
	if err := global.GVA_DB.Where("playmate_id = ?", expertID).First(&voiceIntroduction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return map[string]string{
				"url":      "",
				"duration": "0",
			}, nil
		}
		return nil, err
	}

	return map[string]string{
		"url":      voiceIntroduction.URL,
		"duration": voiceIntroduction.Duration,
	}, nil
}

// GetSkills 获取所有技能列表
func (s *PlaymateService) GetSkills() ([]model.PlaymateSkill, error) {
	var skills []model.PlaymateSkill
	if err := global.GVA_DB.Find(&skills).Error; err != nil {
		return nil, err
	}
	return skills, nil
}
