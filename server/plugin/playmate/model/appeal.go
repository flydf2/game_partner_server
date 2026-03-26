package model

import (
	"time"

	"gorm.io/gorm"
)

// Appeal 申诉模型
type Appeal struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	UserID      uint           `json:"userId" gorm:"index"` // 申诉用户ID
	OrderID     uint           `json:"orderId" gorm:"index"` // 关联订单ID
	Type        string         `json:"type"`                 // 申诉类型：order, payment, service, other
	Title       string         `json:"title"`                // 申诉标题
	Content     string         `json:"content"`              // 申诉内容
	Images      string         `json:"images"`               // 图片证据，逗号分隔
	Status      string         `json:"status"`               // 状态：pending, processing, resolved, rejected
	Response    string         `json:"response"`             // 处理回复
	HandledBy   *uint          `json:"handledBy"`            // 处理人ID
	HandledAt   *time.Time     `json:"handledAt"`            // 处理时间
	ContactInfo string         `json:"contactInfo"`          // 联系方式
	Priority    string         `json:"priority"`             // 优先级：low, normal, high, urgent
}

// TableName 设置Appeal表名
func (Appeal) TableName() string {
	return "appeals"
}
