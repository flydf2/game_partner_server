package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

// RewardOrder 悬赏订单模型
type RewardOrder struct {
	global.GVA_MODEL
	UserID        uint           `json:"userId" gorm:"index"`
	Game          string         `json:"game" gorm:"size:100;not null"`
	Content       string         `json:"content" gorm:"type:text;not null"`
	Reward        float64        `json:"reward" gorm:"type:decimal(10,2);not null"`
	PaymentMethod string         `json:"paymentMethod" gorm:"size:20;not null"`              // prepay 预付, postpay 现付
	Status        string         `json:"status" gorm:"size:20;not null;default:'available'"` // available 可抢单, ongoing 进行中, completed 已完成, draft 草稿, cancelled 已取消, expired 已过期
	TimeLeft      string         `json:"timeLeft" gorm:"size:50;not null"`
	GameRank      string         `json:"gameRank" gorm:"size:50;not null"`
	StartTime     string         `json:"startTime" gorm:"size:50;not null"`
	Duration      int            `json:"duration" gorm:"not null"`
	Location      string         `json:"location" gorm:"size:255;not null"`
	Tags          string         `json:"tags" gorm:"type:text"`         // 以逗号分隔的标签
	Requirements  string         `json:"requirements" gorm:"type:text"` // 以逗号分隔的要求
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}

// RewardOrderApplicant 抢单申请模型
type RewardOrderApplicant struct {
	global.GVA_MODEL
	OrderID        uint      `json:"orderId" gorm:"index;not null"`
	UserID         uint      `json:"userId" gorm:"index;not null"`
	Recommendation string    `json:"recommendation" gorm:"type:text"`
	VoiceUrl       string    `json:"voiceUrl" gorm:"size:255"`
	RecordUrl      string    `json:"recordUrl" gorm:"size:255"`
	Status         string    `json:"status" gorm:"size:20;not null;default:'pending'"` // pending 待处理, approved 已通过, rejected 已拒绝
	AppliedAt      time.Time `json:"appliedAt"`
}

// RewardOrderPayment 订单支付模型
type RewardOrderPayment struct {
	global.GVA_MODEL
	OrderID       uint      `json:"orderId" gorm:"index;not null"`
	Amount        float64   `json:"amount" gorm:"type:decimal(10,2);not null"`
	PaymentMethod string    `json:"paymentMethod" gorm:"size:20;not null"` // wechat, alipay, etc.
	TransactionID string    `json:"transactionId" gorm:"size:100;not null"`
	PaymentStatus string    `json:"paymentStatus" gorm:"size:20;not null;default:'pending'"` // pending 待支付, success 支付成功, failed 支付失败
	PaidAt        time.Time `json:"paidAt"`
}

// RewardOrderReview 订单评价模型
type RewardOrderReview struct {
	global.GVA_MODEL
	OrderID    uint      `json:"orderId" gorm:"index;not null"`
	Rating     int       `json:"rating" gorm:"not null"` // 1-5星
	Review     string    `json:"review" gorm:"type:text"`
	Images     string    `json:"images" gorm:"type:text"` // 以逗号分隔的图片URL
	ReviewedAt time.Time `json:"reviewedAt"`
}
