package model

import (
	"time"

	"gorm.io/gorm"
)

// Order 订单模型
type Order struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	UserID       uint           `json:"userId"`
	PlaymateID   uint           `json:"playmateId"`
	Game         string         `json:"game"`
	Skill        string         `json:"skill"`
	Status       string         `json:"status"` // pending, completed, cancelled
	ServiceTime  string         `json:"serviceTime"`
	Amount       float64        `json:"amount"`
	OrderNumber  string         `json:"orderNumber"`
	PaymentMethod string        `json:"paymentMethod"`
}

// OrderConfirmation 订单确认模型
type OrderConfirmation struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	OrderID      uint      `json:"orderId"`
	PricePerHour float64   `json:"pricePerHour"`
	Duration     int       `json:"duration"`
	ServiceFee   float64   `json:"serviceFee"`
	CouponID     *uint     `json:"couponId"`
	TotalAmount  float64   `json:"totalAmount"`
	CreatedAt    time.Time `json:"createdAt"`
}

// RewardOrder 奖励订单模型
type RewardOrder struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	UserID        uint           `json:"userId"`
	Game          string         `json:"game"`
	Content       string         `json:"content"`
	Reward        float64        `json:"reward"`
	PaymentMethod string         `json:"paymentMethod"` // prepay, postpay
	Status        string         `json:"status"`       // available, grabbed, completed, cancelled
	Tags          string         `json:"tags"`         // 用逗号分隔的标签
}

// Coupon 优惠券模型
type Coupon struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Discount    float64   `json:"discount"`
	Description string    `json:"description"`
	ValidUntil  time.Time `json:"validUntil"`
	CreatedAt   time.Time `json:"createdAt"`
}