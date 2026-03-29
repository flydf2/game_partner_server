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

// AppealService 申诉服务
type AppealService struct{}

// GetAppeals 获取申诉列表
func (s *AppealService) GetAppeals(search request.AppealSearch) ([]model.Appeal, int64, error) {
	var appeals []model.Appeal
	var total int64

	query := global.GVA_DB.Model(&model.Appeal{})

	// 应用过滤条件
	if search.Type != "" {
		query = query.Where("type = ?", search.Type)
	}

	if search.Status != "" {
		query = query.Where("status = ?", search.Status)
	}

	if search.Priority != "" {
		query = query.Where("priority = ?", search.Priority)
	}

	if search.UserID > 0 {
		query = query.Where("user_id = ?", search.UserID)
	}

	if search.OrderID > 0 {
		query = query.Where("order_id = ?", search.OrderID)
	}

	if search.Keyword != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%")
	}

	if search.StartTime != "" {
		query = query.Where("created_at >= ?", search.StartTime)
	}

	if search.EndTime != "" {
		query = query.Where("created_at <= ?", search.EndTime)
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 应用分页
	offset := (search.Page - 1) * search.PageSize
	if search.Page <= 0 {
		offset = 0
	}
	if search.PageSize <= 0 {
		search.PageSize = 10
	}

	// 执行查询
	if err := query.Offset(offset).Limit(search.PageSize).Order("created_at DESC").Find(&appeals).Error; err != nil {
		return nil, 0, err
	}

	return appeals, total, nil
}

// GetAppealDetail 获取申诉详情
func (s *AppealService) GetAppealDetail(appealID uint) (*model.Appeal, error) {
	var appeal model.Appeal
	if err := global.GVA_DB.First(&appeal, appealID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewPlaymateError(response.ErrAppealNotFound)
		}
		return nil, err
	}
	return &appeal, nil
}

// CreateAppeal 创建申诉
func (s *AppealService) CreateAppeal(userID uint, req request.CreateAppealRequest) (*model.Appeal, error) {
	// 设置默认优先级
	priority := req.Priority
	if priority == "" {
		priority = "normal"
	}

	appeal := model.Appeal{
		UserID:      userID,
		OrderID:     req.OrderID,
		Type:        req.Type,
		Title:       req.Title,
		Content:     req.Content,
		Images:      req.Images,
		Status:      "pending",
		ContactInfo: req.ContactInfo,
		Priority:    priority,
	}

	if err := global.GVA_DB.Create(&appeal).Error; err != nil {
		return nil, err
	}

	return &appeal, nil
}

// UpdateAppeal 更新申诉
func (s *AppealService) UpdateAppeal(appealID uint, req request.UpdateAppealRequest) (*model.Appeal, error) {
	var appeal model.Appeal
	if err := global.GVA_DB.First(&appeal, appealID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewPlaymateError(response.ErrAppealNotFound)
		}
		return nil, err
	}

	// 只允许更新pending状态的申诉
	if appeal.Status != "pending" {
		return nil, response.NewPlaymateError(response.ErrAppealStatusNotUpdatable)
	}

	// 更新字段
	updates := make(map[string]interface{})
	if req.Type != "" {
		updates["type"] = req.Type
	}
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Content != "" {
		updates["content"] = req.Content
	}
	if req.Images != "" {
		updates["images"] = req.Images
	}
	if req.ContactInfo != "" {
		updates["contact_info"] = req.ContactInfo
	}
	if req.Priority != "" {
		updates["priority"] = req.Priority
	}

	if len(updates) > 0 {
		if err := global.GVA_DB.Model(&appeal).Updates(updates).Error; err != nil {
			return nil, err
		}
	}

	return &appeal, nil
}

// DeleteAppeal 删除申诉
func (s *AppealService) DeleteAppeal(appealID uint) error {
	var appeal model.Appeal
	if err := global.GVA_DB.First(&appeal, appealID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewPlaymateError(response.ErrAppealNotFound)
		}
		return err
	}

	if err := global.GVA_DB.Delete(&appeal).Error; err != nil {
		return err
	}

	return nil
}

// HandleAppeal 处理申诉
func (s *AppealService) HandleAppeal(appealID uint, handlerID uint, req request.HandleAppealRequest) (*model.Appeal, error) {
	var appeal model.Appeal
	if err := global.GVA_DB.First(&appeal, appealID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewPlaymateError(response.ErrAppealNotFound)
		}
		return nil, err
	}

	// 只能处理pending或processing状态的申诉
	if appeal.Status != "pending" && appeal.Status != "processing" {
		return nil, response.NewPlaymateError(response.ErrAppealAlreadyProcessed)
	}

	now := time.Now()
	updates := map[string]interface{}{
		"status":     req.Status,
		"response":   req.Response,
		"handled_by": handlerID,
		"handled_at": now,
	}

	if err := global.GVA_DB.Model(&appeal).Updates(updates).Error; err != nil {
		return nil, err
	}

	return &appeal, nil
}

// GetUserAppeals 获取用户的申诉列表
func (s *AppealService) GetUserAppeals(userID uint, page, pageSize int) ([]model.Appeal, int64, error) {
	var appeals []model.Appeal
	var total int64

	query := global.GVA_DB.Model(&model.Appeal{}).Where("user_id = ?", userID)

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 应用分页
	offset := (page - 1) * pageSize
	if page <= 0 {
		offset = 0
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// 执行查询
	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&appeals).Error; err != nil {
		return nil, 0, err
	}

	return appeals, total, nil
}
