package service

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
)

// ReviewService 评价服务
type ReviewService struct{}

// SubmitReview 提交评价
func (s *ReviewService) SubmitReview(userID uint, req request.SubmitReviewRequest) (model.Review, error) {
	// 检查陪玩是否存在
	var playmate model.Playmate
	if err := global.GVA_DB.First(&playmate, req.PlaymateID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Review{}, errors.New("陪玩不存在")
		}
		return model.Review{}, err
	}

	// 检查是否已经评价过
	var existingReview model.Review
	if err := global.GVA_DB.Where("user_id = ? AND playmate_id = ?", userID, req.PlaymateID).First(&existingReview).Error; err == nil {
		return model.Review{}, errors.New("已经评价过该陪玩")
	}

	// 构建评价
	review := model.Review{
		UserID:     userID,
		PlaymateID: req.PlaymateID,
		Rating:     req.Rating,
		Content:    req.Content,
		Images:     strings.Join(req.Images, ","),
		Tags:       strings.Join(req.Tags, ","),
	}

	if err := global.GVA_DB.Create(&review).Error; err != nil {
		return model.Review{}, err
	}

	// 更新陪玩的评分
	var totalRating float64
	var count int64
	global.GVA_DB.Model(&model.Review{}).Where("playmate_id = ?", req.PlaymateID).Select("COALESCE(SUM(rating), 0) as total, COUNT(*) as count").Scan(&struct {
		Total float64
		Count int64
	}{totalRating, count})

	if count > 0 {
		newRating := totalRating / float64(count)
		playmate.Rating = newRating
		global.GVA_DB.Save(&playmate)
	}

	return review, nil
}

// GetReviews 获取评价列表
func (s *ReviewService) GetReviews(search request.ReviewSearch) ([]model.Review, int64, error) {
	var reviews []model.Review
	var total int64

	// 构建查询
	query := global.GVA_DB.Model(&model.Review{})

	// 应用搜索条件
	if search.PlaymateID > 0 {
		query = query.Where("playmate_id = ?", search.PlaymateID)
	}
	if search.MinRating > 0 {
		query = query.Where("rating >= ?", search.MinRating)
	}
	if search.MaxRating > 0 {
		query = query.Where("rating <= ?", search.MaxRating)
	}
	if search.Game != "" {
		query = query.Where("game = ?", search.Game)
	}
	if search.StartTime != "" {
		query = query.Where("created_at >= ?", search.StartTime)
	}
	if search.EndTime != "" {
		query = query.Where("created_at <= ?", search.EndTime)
	}
	if search.Keyword != "" {
		query = query.Where("content LIKE ? OR tags LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%")
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页
	offset := (search.Page - 1) * search.PageSize

	// 执行查询
	if err := query.Offset(offset).Limit(search.PageSize).Order("created_at DESC").Find(&reviews).Error; err != nil {
		return nil, 0, err
	}

	return reviews, total, nil
}