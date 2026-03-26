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

// TableName 设置Notification表名
func (Notification) TableName() string {
	return "game_partner_notifications"
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
	Type        string         `json:"type" gorm:"default:'text'"` // text, image, voice, system
	Status      string         `json:"status" gorm:"default:'sent'"` // sent, delivered, read
	ConversationID string      `json:"conversationId" gorm:"index"` // 会话ID，用于分组聊天记录
}

// TableName 设置Message表名
func (Message) TableName() string {
	return "game_partner_messages"
}

// ChatMessage 聊天消息模型（用于前端展示）
type ChatMessage struct {
	From    string    `json:"from"` // self 或 other
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
	Type    string    `json:"type"` // text, image, voice
	Status  string    `json:"status"` // sent, delivered, read
}

// Conversation 会话模型
type Conversation struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	UserID      uint           `json:"userId" gorm:"index"` // 当前用户ID
	OtherUserID uint           `json:"otherUserId" gorm:"index"` // 对方用户ID
	LastMessage string         `json:"lastMessage"` // 最后一条消息内容
	LastTime    time.Time      `json:"lastTime"` // 最后一条消息时间
	UnreadCount int            `json:"unreadCount"` // 未读消息数
	Status      string         `json:"status"` // active, archived
}

// TableName 设置Conversation表名
func (Conversation) TableName() string {
	return "game_partner_conversations"
}