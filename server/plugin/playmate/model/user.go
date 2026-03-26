package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Username    string         `json:"username"`
	Password    string         `json:"-"`
	Phone       string         `json:"phone"`
	Avatar      string         `json:"avatar"`
	Nickname    string         `json:"nickname"`
	VipLevel    int            `json:"vipLevel"`
	Balance     float64        `json:"balance"`
	CouponCount int            `json:"couponCount"`
}

// TableName 设置User表名
func (User) TableName() string {
	return "game_partner_users"
}

// UserSettings 用户设置模型
type UserSettings struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserID        uint      `json:"userId"`
	Notifications string    `json:"notifications"` // JSON格式存储通知设置
	Privacy       string    `json:"privacy"`       // JSON格式存储隐私设置
	Theme         string    `json:"theme"`
	Language      string    `json:"language"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// TableName 设置UserSettings表名
func (UserSettings) TableName() string {
	return "game_partner_user_settings"
}

// UserWallet 用户钱包模型
type UserWallet struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `json:"userId"`
	Balance      float64   `json:"balance"`
	Frozen       float64   `json:"frozen"`
	TotalIncome  float64   `json:"totalIncome"`
	TotalExpense float64   `json:"totalExpense"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// TableName 设置UserWallet表名
func (UserWallet) TableName() string {
	return "game_partner_user_wallets"
}

// Transaction 交易记录模型
type Transaction struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"userId"`
	Type        string    `json:"type"` // income 或 expense
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	Time        time.Time `json:"time"`
	CreatedAt   time.Time `json:"createdAt"`
}

// TableName 设置Transaction表名
func (Transaction) TableName() string {
	return "game_partner_transactions"
}
