package model

import (
	"time"

	"gorm.io/gorm"
)

// Review 评价模型
type Review struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint           `json:"userId"`
	PlaymateID uint          `json:"playmateId"`
	Rating    int            `json:"rating"`
	Content   string         `json:"content"`
	Images    string         `gorm:"type:varchar(4096)" json:"images"` // 用逗号分隔的图片URL
	Tags      string         `json:"tags"`   // 用逗号分隔的标签
}

// TableName 设置Review表名
func (Review) TableName() string {
	return "game_partner_reviews"
}

// Withdrawal 提现模型
type Withdrawal struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	UserID      uint           `json:"userId"`
	Amount      float64        `json:"amount"`
	Fee         float64        `json:"fee"`
	ActualAmount float64       `json:"actualAmount"`
	Method      string         `json:"method"` // wechat, alipay, bank
	Status      string         `json:"status"` // pending, processing, completed, failed
	CompletedAt *time.Time     `json:"completedAt"`
	FailedReason string        `json:"failedReason"`
}

// TableName 设置Withdrawal表名
func (Withdrawal) TableName() string {
	return "game_partner_withdrawals"
}