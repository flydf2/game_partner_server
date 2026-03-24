package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
)

// ActivityService 活动服务
type ActivityService struct{}

// GetActivities 获取活动列表
func (s *ActivityService) GetActivities() ([]model.Activity, error) {
	var activities []model.Activity
	if err := global.GVA_DB.Find(&activities).Error; err != nil {
		return nil, err
	}

	return activities, nil
}