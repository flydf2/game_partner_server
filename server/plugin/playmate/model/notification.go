package model

import (
	"time"

	"gorm.io/gorm"
)

// Notification 通知模型
type Notification struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint           `json:"userId"`
	Type      string         `json:"type"` // order, system, promotion, message
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	Time      time.Time      `json:"time"`
	Read      bool           `json:"read"`
	OrderID   *uint          `json:"orderId"`
}

// Message 消息模型
type Message struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	FromUserID  uint           `json:"fromUserId"`
	ToUserID    uint           `json:"toUserId"`
	Content     string         `json:"content"`
	Time        time.Time      `json:"time"`
	Read        bool           `json:"read"`
}

// ChatMessage 聊天消息模型
type ChatMessage struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	From      string         `json:"from"` // self 或 other
	Content   string         `json:"content"`
	Time      time.Time      `json:"time"`
}