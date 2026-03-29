package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
)

// NotificationService 通知服务
type NotificationService struct{}

// GetNotifications 获取通知列表
func (s *NotificationService) GetNotifications(userID uint, search request.NotificationSearch) ([]model.Notification, int64, error) {
	// 构建查询
	query := global.GVA_DB.Model(&model.Notification{}).Where("user_id = ?", userID)

	// 应用搜索条件
	if search.Type != "" {
		query = query.Where("type = ?", search.Type)
	}
	if search.Status != "" {
		if search.Status == "read" {
			query = query.Where("read = ?", true)
		} else if search.Status == "unread" {
			query = query.Where("read = ?", false)
		}
	}
	if search.StartTime != "" {
		query = query.Where("time >= ?", search.StartTime)
	}
	if search.EndTime != "" {
		query = query.Where("time <= ?", search.EndTime)
	}
	if search.Keyword != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%")
	}

	// 计算总数
	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页
	offset := (search.Page - 1) * search.PageSize

	// 执行查询
	var notifications []model.Notification
	if err := query.Offset(offset).Limit(search.PageSize).Order("time DESC").Find(&notifications).Error; err != nil {
		return nil, 0, err
	}

	return notifications, total, nil
}

// MarkAsRead 标记通知为已读
func (s *NotificationService) MarkAsRead(notificationID uint) error {
	var notification model.Notification
	if err := global.GVA_DB.First(&notification, notificationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewPlaymateError(response.ErrNotificationNotFound)
		}
		return err
	}

	notification.Read = true
	if err := global.GVA_DB.Save(&notification).Error; err != nil {
		return err
	}

	return nil
}

// MarkAllAsRead 标记所有通知为已读
func (s *NotificationService) MarkAllAsRead(userID uint) error {
	if err := global.GVA_DB.Model(&model.Notification{}).Where("user_id = ?", userID).Update("read", true).Error; err != nil {
		return err
	}

	return nil
}

// GetUnreadCount 获取未读通知数量
func (s *NotificationService) GetUnreadCount(userID uint) (int64, error) {
	var count int64
	if err := global.GVA_DB.Model(&model.Notification{}).Where("user_id = ? AND read = ?", userID, false).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}