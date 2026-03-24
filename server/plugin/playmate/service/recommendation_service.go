package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
)

// RecommendationService 推荐服务
type RecommendationService struct{}

// GetRecommendations 获取推荐列表
func (s *RecommendationService) GetRecommendations() ([]model.Recommendation, error) {
	var recommendations []model.Recommendation
	if err := global.GVA_DB.Find(&recommendations).Error; err != nil {
		return nil, err
	}

	return recommendations, nil
}