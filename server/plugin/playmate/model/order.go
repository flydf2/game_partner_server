package model

import (
	"time"

	"gorm.io/gorm"
)

// Order 订单模型
type Order struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	UserID        uint           `json:"userId"`
	PlaymateID    uint           `json:"playmateId"`
	SkillID       uint           `gorm:"default:1" json:"skillId"`
	Game          string         `json:"game"`
	Skill         string         `json:"skill"`
	Status        string         `json:"status"` // pending, completed, cancelled
	ServiceTime   string         `json:"serviceTime"`
	Amount        float64        `json:"amount"`
	Quantity      int            `json:"quantity"`
	OrderNumber   string         `json:"orderNumber"`
	PaymentMethod string         `json:"paymentMethod"`
}

// TableName 设置Order表名
func (Order) TableName() string {
	return "game_partner_orders"
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

// TableName 设置OrderConfirmation表名
func (OrderConfirmation) TableName() string {
	return "game_partner_order_confirmations"
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

// TableName 设置Coupon表名
func (Coupon) TableName() string {
	return "game_partner_coupons"
}
