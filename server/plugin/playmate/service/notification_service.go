package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
)

// NotificationService 通知服务
type NotificationService struct{}

// GetNotifications 获取通知列表
func (s *NotificationService) GetNotifications(userID uint) ([]model.Notification, error) {
	var notifications []model.Notification
	if err := global.GVA_DB.Where("user_id = ?", userID).Order("time DESC").Find(&notifications).Error; err != nil {
		return nil, err
	}

	return notifications, nil
}

// MarkAsRead 标记通知为已读
func (s *NotificationService) MarkAsRead(notificationID uint) error {
	var notification model.Notification
	if err := global.GVA_DB.First(&notification, notificationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("通知不存在")
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