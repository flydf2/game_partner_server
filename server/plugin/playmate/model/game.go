package model

import (
	"time"

	"gorm.io/gorm"
)

// Game 游戏模型
type Game struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `json:"name"`
	Icon        string         `json:"icon"`
	Category    string         `json:"category"`
	CategoryIDs []uint         `gorm:"type:json" json:"categoryIds"`
	Status      string         `json:"status"`
	Platform    string         `json:"platform"`
	Image       string         `gorm:"type:varchar(2048)" json:"image"`
	PageStyle   string         `gorm:"type:text" json:"pageStyle"` // 页面风格配置，存储JSON格式数据
}

// TableName 设置Game表名
func (Game) TableName() string {
	return "game_partner_games"
}

// Activity 活动模型
type Activity struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Discount    float64        `json:"discount"`
	Type        string         `json:"type"` // discount, weekend
	ValidUntil  time.Time      `json:"validUntil"`
}

// TableName 设置Activity表名
func (Activity) TableName() string {
	return "game_partner_activities"
}

// Category 分类模型
type Category struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `json:"name"`
	Icon      string         `json:"icon"`
}

// TableName 设置Category表名
func (Category) TableName() string {
	return "game_partner_categories"
}

// GameCategory 游戏分类模型
type GameCategory struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Name      string         `json:"name"`
	Label     string         `json:"label"`
}

// TableName 设置GameCategory表名
func (GameCategory) TableName() string {
	return "game_partner_game_categories"
}
