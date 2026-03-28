package service

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
)

// ExpertOrderSettingService 专家订单设置服务
type ExpertOrderSettingService struct{}

// GetOrderSetting 获取专家订单设置
func (s *ExpertOrderSettingService) GetOrderSetting(expertID uint) (*model.ExpertOrderSetting, error) {
	var setting model.ExpertOrderSetting
	err := global.GVA_DB.Where("expert_id = ?", expertID).First(&setting).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果没有设置，返回默认设置
			return &model.ExpertOrderSetting{
				ExpertID:            expertID,
				AutoAccept:          false,
				AcceptMode:          "manual",
				MaxDailyOrders:      10,
				ServiceStartTime:    "09:00",
				ServiceEndTime:      "23:00",
				MinOrderAmount:      0,
				MaxOrderAmount:      0,
				RejectWithoutVoice:  false,
				NotificationEnabled: true,
			}, nil
		}
		return nil, err
	}
	return &setting, nil
}

// UpdateOrderSetting 更新专家订单设置
func (s *ExpertOrderSettingService) UpdateOrderSetting(setting *model.ExpertOrderSetting) (*model.ExpertOrderSetting, error) {
	var existing model.ExpertOrderSetting
	err := global.GVA_DB.Where("expert_id = ?", setting.ExpertID).First(&existing).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建新设置
			if err := global.GVA_DB.Create(setting).Error; err != nil {
				return nil, err
			}
			return setting, nil
		}
		return nil, err
	}

	// 更新现有设置
	setting.ID = existing.ID
	setting.CreatedAt = existing.CreatedAt
	if err := global.GVA_DB.Save(setting).Error; err != nil {
		return nil, err
	}

	return setting, nil
}

// GetExpertServices 获取专家服务列表
func (s *ExpertOrderSettingService) GetExpertServices(expertID uint) ([]model.ExpertService, error) {
	var services []model.ExpertService
	err := global.GVA_DB.Where("expert_id = ?", expertID).
		Order("sort_order ASC, created_at DESC").
		Find(&services).Error
	if err != nil {
		return nil, err
	}
	return services, nil
}

// GetExpertServiceByID 根据ID获取专家服务
func (s *ExpertOrderSettingService) GetExpertServiceByID(serviceID uint) (*model.ExpertService, error) {
	var service model.ExpertService
	err := global.GVA_DB.First(&service, serviceID).Error
	if err != nil {
		return nil, err
	}
	return &service, nil
}

// CreateExpertService 创建专家服务
func (s *ExpertOrderSettingService) CreateExpertService(service *model.ExpertService) (*model.ExpertService, error) {
	if err := global.GVA_DB.Create(service).Error; err != nil {
		return nil, err
	}
	return service, nil
}

// UpdateExpertService 更新专家服务
func (s *ExpertOrderSettingService) UpdateExpertService(service *model.ExpertService) (*model.ExpertService, error) {
	if err := global.GVA_DB.Save(service).Error; err != nil {
		return nil, err
	}
	return service, nil
}

// DeleteExpertService 删除专家服务
func (s *ExpertOrderSettingService) DeleteExpertService(serviceID uint) error {
	return global.GVA_DB.Delete(&model.ExpertService{}, serviceID).Error
}

// GetTodayRecommendations 获取今日推荐列表
func (s *ExpertOrderSettingService) GetTodayRecommendations(expertID uint) ([]model.TodayRecommendation, error) {
	var recommendations []model.TodayRecommendation
	now := time.Now()

	err := global.GVA_DB.Where("expert_id = ? AND is_enabled = ?", expertID, true).
		Where("(start_time IS NULL OR start_time <= ?) AND (end_time IS NULL OR end_time >= ?)", now, now).
		Order("sort_order ASC, created_at DESC").
		Find(&recommendations).Error
	if err != nil {
		return nil, err
	}
	return recommendations, nil
}

// GetTodayRecommendationByID 根据ID获取今日推荐
func (s *ExpertOrderSettingService) GetTodayRecommendationByID(recommendationID uint) (*model.TodayRecommendation, error) {
	var recommendation model.TodayRecommendation
	err := global.GVA_DB.First(&recommendation, recommendationID).Error
	if err != nil {
		return nil, err
	}
	return &recommendation, nil
}

// CreateTodayRecommendation 创建今日推荐
func (s *ExpertOrderSettingService) CreateTodayRecommendation(recommendation *model.TodayRecommendation) (*model.TodayRecommendation, error) {
	if err := global.GVA_DB.Create(recommendation).Error; err != nil {
		return nil, err
	}
	return recommendation, nil
}

// UpdateTodayRecommendation 更新今日推荐
func (s *ExpertOrderSettingService) UpdateTodayRecommendation(recommendation *model.TodayRecommendation) (*model.TodayRecommendation, error) {
	if err := global.GVA_DB.Save(recommendation).Error; err != nil {
		return nil, err
	}
	return recommendation, nil
}

// DeleteTodayRecommendation 删除今日推荐
func (s *ExpertOrderSettingService) DeleteTodayRecommendation(recommendationID uint) error {
	return global.GVA_DB.Delete(&model.TodayRecommendation{}, recommendationID).Error
}

// ToggleServiceStatus 切换服务启用状态
func (s *ExpertOrderSettingService) ToggleServiceStatus(serviceID uint) (*model.ExpertService, error) {
	var service model.ExpertService
	if err := global.GVA_DB.First(&service, serviceID).Error; err != nil {
		return nil, err
	}

	service.IsEnabled = !service.IsEnabled
	if err := global.GVA_DB.Save(&service).Error; err != nil {
		return nil, err
	}

	return &service, nil
}

// ToggleRecommendationStatus 切换推荐启用状态
func (s *ExpertOrderSettingService) ToggleRecommendationStatus(recommendationID uint) (*model.TodayRecommendation, error) {
	var recommendation model.TodayRecommendation
	if err := global.GVA_DB.First(&recommendation, recommendationID).Error; err != nil {
		return nil, err
	}

	recommendation.IsEnabled = !recommendation.IsEnabled
	if err := global.GVA_DB.Save(&recommendation).Error; err != nil {
		return nil, err
	}

	return &recommendation, nil
}
