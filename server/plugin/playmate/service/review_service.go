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
func (s *ReviewService) GetReviews(page, pageSize int) ([]model.Review, int64, error) {
	var reviews []model.Review
	var total int64

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询总数
	if err := global.GVA_DB.Model(&model.Review{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 查询数据
	if err := global.GVA_DB.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&reviews).Error; err != nil {
		return nil, 0, err
	}

	return reviews, total, nil
}