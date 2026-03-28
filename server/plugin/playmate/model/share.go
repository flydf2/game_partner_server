package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// Share 分享模型
type Share struct {
	global.GVA_MODEL
	UserID        uint           `json:"userId" gorm:"index"`
	OrderID       *uint          `json:"orderId" gorm:"index"` // 普通订单ID
	RewardOrderID *uint          `json:"rewardOrderId" gorm:"index"` // 悬赏订单ID
	ShareType     string         `json:"shareType" gorm:"size:20;not null"` // order 普通订单, reward 悬赏订单
	SharePlatform string         `json:"sharePlatform" gorm:"size:20;not null"` // wechat, qq, weibo, etc.
	ShareURL      string         `json:"shareUrl" gorm:"size:255;not null"`
	ShareCode     string         `json:"shareCode" gorm:"size:50;not null;uniqueIndex"`
	ClickCount    int            `json:"clickCount" gorm:"default:0"`
	Status        string         `json:"status" gorm:"size:20;not null;default:'active'"` // active, expired
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

// TableName 设置Share表名
func (Share) TableName() string {
	return "game_partner_shares"
}
