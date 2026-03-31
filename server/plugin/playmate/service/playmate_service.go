package service

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
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
			return model.Playmate{}, response.NewPlaymateError(response.ErrPlaymateNotFound)
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
		return response.NewPlaymateError(response.ErrAlreadyFollowed)
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
		return response.NewPlaymateError(response.ErrNotFollowed)
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

	// 查询评价并关联用户信息
	if err := query.Preload("User").Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&reviews).Error; err != nil {
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

// GetExpertStatus 获取专家状态
func (s *PlaymateService) GetExpertStatus(expertID uint) (map[string]interface{}, error) {
	// 获取专家基本信息
	playmate, err := s.GetPlaymateById(expertID)
	if err != nil {
		return nil, err
	}

	// 构建状态响应，匹配前端期望的数据结构
	status := map[string]interface{}{
		"isFollowing": false,          // 这里可以根据实际业务逻辑判断是否关注
		"isFavorite":  false,          // 这里可以根据实际业务逻辑判断是否收藏
		"isLiked":     false,          // 这里可以根据实际业务逻辑判断是否点赞
		"likeCount":   playmate.Likes, // 使用playmate的点赞数
	}

	return status, nil
}

// GetSkills 获取技能列表（支持分页和搜索）
func (s *PlaymateService) GetSkills(search request.SkillSearch) ([]model.PlaymateSkill, int64, error) {
	var skills []model.PlaymateSkill
	var total int64

	db := global.GVA_DB
	query := db.Model(&model.PlaymateSkill{})

	// 应用过滤条件
	if search.Game != "" {
		// 这里需要根据实际业务逻辑来实现游戏过滤
		// 暂时注释掉，因为PlaymateSkill模型中没有Game字段
		// query = query.Where("game = ?", search.Game)
	}

	if search.Level != "" {
		query = query.Where("level = ?", search.Level)
	}

	if search.Keyword != "" {
		keyword := fmt.Sprintf("%%%s%%", search.Keyword)
		query = query.Where("name LIKE ? OR description LIKE ?", keyword, keyword)
	}

	// 计算总数
	query.Count(&total)

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
	if err := query.Find(&skills).Error; err != nil {
		return nil, 0, err
	}

	return skills, total, nil
}

// AddSkill 添加技能
func (s *PlaymateService) AddSkill(userID uint, req request.AddSkillRequest) (model.PlaymateSkill, error) {
	skill := model.PlaymateSkill{
		PlaymateID:  userID,
		Name:        req.Skill,
		Price:       req.Price,
		Level:       req.Level,
		Description: req.Description,
	}

	if err := global.GVA_DB.Create(&skill).Error; err != nil {
		return model.PlaymateSkill{}, err
	}

	return skill, nil
}

// UpdateSkill 更新技能
func (s *PlaymateService) UpdateSkill(skillID uint, req request.UpdateSkillRequest) (model.PlaymateSkill, error) {
	var skill model.PlaymateSkill
	if err := global.GVA_DB.First(&skill, skillID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.PlaymateSkill{}, response.NewPlaymateError(response.ErrOrderNotFound)
		}
		return model.PlaymateSkill{}, err
	}

	if req.Skill != "" {
		skill.Name = req.Skill
	}
	if req.Price > 0 {
		skill.Price = req.Price
	}
	if req.Level != "" {
		skill.Level = req.Level
	}
	if req.Description != "" {
		skill.Description = req.Description
	}

	if err := global.GVA_DB.Save(&skill).Error; err != nil {
		return model.PlaymateSkill{}, err
	}

	return skill, nil
}

// DeleteSkill 删除技能
func (s *PlaymateService) DeleteSkill(skillID uint) error {
	if err := global.GVA_DB.Delete(&model.PlaymateSkill{}, skillID).Error; err != nil {
		return err
	}
	return nil
}

// GetLeaderboard 获取排行榜
func (s *PlaymateService) GetLeaderboard(search request.LeaderboardSearch) ([]model.Playmate, int64, error) {
	var playmates []model.Playmate
	var total int64

	db := global.GVA_DB
	query := db.Model(&model.Playmate{})

	// 按游戏过滤
	if search.Game != "" {
		query = query.Where("game = ?", search.Game)
	}

	// 计算总数
	query.Count(&total)

	// 按评分降序排序
	query = query.Order("rating DESC")

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
	if err := query.Find(&playmates).Error; err != nil {
		return nil, 0, err
	}

	return playmates, total, nil
}

// GetMatchHistory 获取用户的匹配历史
func (s *PlaymateService) GetMatchHistory(userID uint, page, pageSize int) ([]model.MatchHistory, int64, error) {
	var histories []model.MatchHistory
	var total int64

	query := global.GVA_DB.Model(&model.MatchHistory{}).Where("user_id = ?", userID)

	// 计算总数
	query.Count(&total)

	// 应用分页
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	// 查询历史
	if err := query.Order("created_at DESC").Offset(offset).Limit(pageSize).Find(&histories).Error; err != nil {
		return nil, 0, err
	}

	return histories, total, nil
}

// GetMatchHistoryById 根据ID获取匹配历史详情
func (s *PlaymateService) GetMatchHistoryById(id uint) (model.MatchHistory, error) {
	var history model.MatchHistory
	if err := global.GVA_DB.First(&history, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.MatchHistory{}, response.NewPlaymateError(response.ErrOrderNotFound)
		}
		return model.MatchHistory{}, err
	}
	return history, nil
}

// CreateMatchHistory 创建匹配历史
func (s *PlaymateService) CreateMatchHistory(history model.MatchHistory) (model.MatchHistory, error) {
	if err := global.GVA_DB.Create(&history).Error; err != nil {
		return model.MatchHistory{}, err
	}
	return history, nil
}

// UpdateMatchHistory 更新匹配历史
func (s *PlaymateService) UpdateMatchHistory(history model.MatchHistory) (model.MatchHistory, error) {
	if err := global.GVA_DB.Save(&history).Error; err != nil {
		return model.MatchHistory{}, err
	}
	return history, nil
}

// UpdateExpertAutoReply 更新专家自动回复设置
func (s *PlaymateService) UpdateExpertAutoReply(expertID uint, autoReplyEnabled bool) error {
	// 查找专家订单设置
	var setting model.ExpertOrderSetting
	result := global.GVA_DB.Where("expert_id = ?", expertID).First(&setting)

	// 如果不存在，创建新的设置
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			setting = model.ExpertOrderSetting{
				ExpertID:            expertID,
				AutoAccept:          false,
				AcceptMode:          "manual",
				MaxDailyOrders:      10,
				ServiceStartTime:    "09:00",
				ServiceEndTime:      "23:00",
				MinOrderAmount:      0,
				MaxOrderAmount:      0,
				RejectWithoutVoice:  false,
				NotificationEnabled: autoReplyEnabled,
			}
			return global.GVA_DB.Create(&setting).Error
		}
		return result.Error
	}

	// 更新设置
	setting.NotificationEnabled = autoReplyEnabled
	return global.GVA_DB.Save(&setting).Error
}

// UpdateExpertOrderStatus 更新专家订单状态设置
func (s *PlaymateService) UpdateExpertOrderStatus(expertID uint, autoAccept bool, acceptMode string) error {
	// 查找专家订单设置
	var setting model.ExpertOrderSetting
	result := global.GVA_DB.Where("expert_id = ?", expertID).First(&setting)

	// 如果不存在，创建新的设置
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			setting = model.ExpertOrderSetting{
				ExpertID:            expertID,
				AutoAccept:          autoAccept,
				AcceptMode:          acceptMode,
				MaxDailyOrders:      10,
				ServiceStartTime:    "09:00",
				ServiceEndTime:      "23:00",
				MinOrderAmount:      0,
				MaxOrderAmount:      0,
				RejectWithoutVoice:  false,
				NotificationEnabled: true,
			}
			return global.GVA_DB.Create(&setting).Error
		}
		return result.Error
	}

	// 更新设置
	setting.AutoAccept = autoAccept
	setting.AcceptMode = acceptMode
	return global.GVA_DB.Save(&setting).Error
}

// UpdateExpertTimeSlots 更新专家服务时间设置
func (s *PlaymateService) UpdateExpertTimeSlots(expertID uint, startTime, endTime, restDays string) error {
	// 查找专家订单设置
	var setting model.ExpertOrderSetting
	result := global.GVA_DB.Where("expert_id = ?", expertID).First(&setting)

	// 如果不存在，创建新的设置
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			setting = model.ExpertOrderSetting{
				ExpertID:            expertID,
				AutoAccept:          false,
				AcceptMode:          "manual",
				MaxDailyOrders:      10,
				ServiceStartTime:    startTime,
				ServiceEndTime:      endTime,
				RestDays:            restDays,
				MinOrderAmount:      0,
				MaxOrderAmount:      0,
				RejectWithoutVoice:  false,
				NotificationEnabled: true,
			}
			return global.GVA_DB.Create(&setting).Error
		}
		return result.Error
	}

	// 更新设置
	setting.ServiceStartTime = startTime
	setting.ServiceEndTime = endTime
	setting.RestDays = restDays
	return global.GVA_DB.Save(&setting).Error
}

// ApplyExpertVerification 提交专家认证申请
func (s *PlaymateService) ApplyExpertVerification(userID uint, req request.ExpertVerificationRequest) (model.ExpertVerification, error) {
	// 检查用户是否已有认证申请
	var existingVerification model.ExpertVerification
	result := global.GVA_DB.Where("user_id = ? AND game_id = ?", userID, req.GameID).First(&existingVerification)
	if result.Error == nil && existingVerification.Status == "pending" {
		return model.ExpertVerification{}, errors.New("已有待处理的认证申请")
	}

	// 准备认证申请数据
	verification := model.ExpertVerification{
		UserID:      userID,
		GameID:      req.GameID,
		GameName:    req.GameName,
		Rank:        req.Rank,
		Positions:   strings.Join(req.Positions, ","),
		Screenshots: strings.Join(req.Screenshots, ","),
		VoiceURL:    req.VoiceURL,
		Status:      "pending",
	}

	// 保存认证申请
	if err := global.GVA_DB.Create(&verification).Error; err != nil {
		return model.ExpertVerification{}, err
	}

	return verification, nil
}

// GetExpertVerificationStatus 获取专家认证状态
func (s *PlaymateService) GetExpertVerificationStatus(userID uint, game string) (model.ExpertVerification, error) {
	var verification model.ExpertVerification
	result := global.GVA_DB.Where("user_id = ? AND game_name = ?", userID, game).Order("created_at DESC").First(&verification)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return model.ExpertVerification{}, errors.New("未找到认证申请")
		}
		return model.ExpertVerification{}, result.Error
	}

	return verification, nil
}

// CancelExpertVerification 撤销专家认证申请
func (s *PlaymateService) CancelExpertVerification(userID, verificationID uint) error {
	// 查找认证申请
	var verification model.ExpertVerification
	result := global.GVA_DB.First(&verification, verificationID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("认证申请不存在")
		}
		return result.Error
	}

	// 检查是否是申请的用户
	if verification.UserID != userID {
		return errors.New("无权限撤销此认证申请")
	}

	// 检查认证状态是否为pending
	if verification.Status != "pending" {
		return errors.New("只能撤销待处理的认证申请")
	}

	// 更新认证状态为cancelled
	verification.Status = "cancelled"
	if err := global.GVA_DB.Save(&verification).Error; err != nil {
		return err
	}

	return nil
}

// HandleExpertVerification 处理专家认证申请
func (s *PlaymateService) HandleExpertVerification(verificationID uint, req request.HandleExpertVerificationRequest) (model.ExpertVerification, error) {
	var verification model.ExpertVerification
	if err := global.GVA_DB.First(&verification, verificationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ExpertVerification{}, errors.New("认证申请不存在")
		}
		return model.ExpertVerification{}, err
	}

	// 更新认证状态
	verification.Status = req.Status
	if req.Status == "rejected" {
		verification.Reason = req.Reason
	}

	if err := global.GVA_DB.Save(&verification).Error; err != nil {
		return model.ExpertVerification{}, err
	}

	// 如果认证通过，创建或更新专家信息
	if req.Status == "approved" {
		// 检查是否已存在专家信息
		var playmate model.Playmate
		result := global.GVA_DB.Where("user_id = ?", verification.UserID).First(&playmate)

		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				// 创建新的专家信息
				playmate = model.Playmate{
					UserID:      verification.UserID,
					Game:        verification.GameName,
					Rank:        verification.Rank,
					IsOnline:    false,
					Rating:      0,
					Likes:       0,
				}
				if err := global.GVA_DB.Create(&playmate).Error; err != nil {
					return verification, err
				}
			} else {
				return verification, result.Error
			}
		} else {
			// 更新现有专家信息
			playmate.Game = verification.GameName
			playmate.Rank = verification.Rank
			if err := global.GVA_DB.Save(&playmate).Error; err != nil {
				return verification, err
			}
		}
	}

	return verification, nil
}

// GetExpertVerificationList 获取专家认证列表
func (s *PlaymateService) GetExpertVerificationList(userID uint, search request.ExpertVerificationSearch) ([]model.ExpertVerification, int64, error) {
	var verifications []model.ExpertVerification
	var total int64

	db := global.GVA_DB
	query := db.Model(&model.ExpertVerification{})

	// 应用过滤条件
	if search.Status != "" {
		query = query.Where("status = ?", search.Status)
	}

	if search.GameID > 0 {
		query = query.Where("game_id = ?", search.GameID)
	}

	// 如果提供了用户ID，则只返回该用户的认证申请
	if userID > 0 {
		query = query.Where("user_id = ?", userID)
	} else if search.UserID > 0 {
		query = query.Where("user_id = ?", search.UserID)
	}

	if search.StartTime != "" {
		query = query.Where("created_at >= ?", search.StartTime)
	}

	if search.EndTime != "" {
		query = query.Where("created_at <= ?", search.EndTime)
	}

	// 计算总数
	query.Count(&total)

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
	query = query.Offset(offset).Limit(pageSize).Order("created_at DESC")

	// 执行查询
	if err := query.Find(&verifications).Error; err != nil {
		return nil, 0, err
	}

	return verifications, total, nil
}

// GetExpertVerificationById 根据ID获取专家认证详情
func (s *PlaymateService) GetExpertVerificationById(id uint) (model.ExpertVerification, error) {
	var verification model.ExpertVerification
	if err := global.GVA_DB.First(&verification, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ExpertVerification{}, errors.New("认证申请不存在")
		}
		return model.ExpertVerification{}, err
	}
	return verification, nil
}

// BatchHandleExpertVerification 批量处理专家认证申请
func (s *PlaymateService) BatchHandleExpertVerification(req request.BatchHandleExpertVerificationRequest) error {
	// 开始事务
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, id := range req.IDs {
		var verification model.ExpertVerification
		if err := tx.First(&verification, id).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 更新认证状态
		verification.Status = req.Status
		if req.Status == "rejected" {
			verification.Reason = req.Reason
		}

		if err := tx.Save(&verification).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 如果认证通过，创建或更新专家信息
		if req.Status == "approved" {
			// 检查是否已存在专家信息
			var playmate model.Playmate
			result := tx.Where("user_id = ?", verification.UserID).First(&playmate)

			if result.Error != nil {
				if errors.Is(result.Error, gorm.ErrRecordNotFound) {
					// 创建新的专家信息
					playmate = model.Playmate{
						UserID:      verification.UserID,
						Game:        verification.GameName,
						Rank:        verification.Rank,
						IsOnline:    false,
						Rating:      0,
						Likes:       0,
					}
					if err := tx.Create(&playmate).Error; err != nil {
						tx.Rollback()
						return err
					}
				} else {
					tx.Rollback()
					return result.Error
				}
			} else {
				// 更新现有专家信息
				playmate.Game = verification.GameName
				playmate.Rank = verification.Rank
				if err := tx.Save(&playmate).Error; err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}

	return tx.Commit().Error
}

// ExportExpertVerification 导出专家认证列表
func (s *PlaymateService) ExportExpertVerification(search request.ExpertVerificationSearch) ([]byte, error) {
	// 这里可以实现导出Excel的逻辑
	// 暂时返回空数据，实际项目中可以使用excelize等库实现
	return []byte{}, nil
}

// GetExpertVerificationStats 获取专家认证统计数据
func (s *PlaymateService) GetExpertVerificationStats(startTime, endTime string) (map[string]interface{}, error) {
	var totalCount int64
	var pendingCount int64
	var approvedCount int64
	var rejectedCount int64

	db := global.GVA_DB
	query := db.Model(&model.ExpertVerification{})

	// 应用时间过滤
	if startTime != "" {
		query = query.Where("created_at >= ?", startTime)
	}
	if endTime != "" {
		query = query.Where("created_at <= ?", endTime)
	}

	// 计算总数
	query.Count(&totalCount)

	// 计算各状态数量
	query.Where("status = ?", "pending").Count(&pendingCount)
	query.Where("status = ?", "approved").Count(&approvedCount)
	query.Where("status = ?", "rejected").Count(&rejectedCount)

	// 构建统计数据
	stats := map[string]interface{}{
		"totalCount":    totalCount,
		"pendingCount":  pendingCount,
		"approvedCount": approvedCount,
		"rejectedCount": rejectedCount,
		"approvalRate":  0.0,
	}

	// 计算通过率
	if totalCount > 0 {
		stats["approvalRate"] = float64(approvedCount) / float64(totalCount)
	}

	return stats, nil
}
