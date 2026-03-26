package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
)

// ActivityService 活动服务
type ActivityService struct{}

// GetActivities 获取活动列表
func (s *ActivityService) GetActivities(search request.ActivitySearch) ([]model.Activity, int64, error) {
	var activities []model.Activity
	var total int64

	db := global.GVA_DB.Model(&model.Activity{})

	// 应用搜索条件
	if search.Status != "" {
		db = db.Where("status = ?", search.Status)
	}
	if search.Type != "" {
		db = db.Where("type = ?", search.Type)
	}
	if search.StartTime != "" {
		db = db.Where("start_time >= ?", search.StartTime)
	}
	if search.EndTime != "" {
		db = db.Where("end_time <= ?", search.EndTime)
	}
	if search.Keyword != "" {
		db = db.Where("title LIKE ? OR description LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%")
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页
	offset := (search.Page - 1) * search.PageSize

	// 执行查询
	if err := db.Offset(offset).Limit(search.PageSize).Order("created_at DESC").Find(&activities).Error; err != nil {
		return nil, 0, err
	}

	return activities, total, nil
}