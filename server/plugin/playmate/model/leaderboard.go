package model

import (
	"time"

	"gorm.io/gorm"
)

// LeaderboardType 排行榜类型
type LeaderboardType string

const (
	// LeaderboardTypeWeekly 周榜
	LeaderboardTypeWeekly LeaderboardType = "weekly"
	// LeaderboardTypeMonthly 月榜
	LeaderboardTypeMonthly LeaderboardType = "monthly"
)

// Leaderboard 排行榜模型
type Leaderboard struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   time.Time       `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt  `gorm:"index" json:"-"`
	Name        string          `gorm:"type:varchar(255);not null;comment:榜单名称（话题名称、游戏名称等）" json:"name"`
	Type        LeaderboardType `gorm:"type:varchar(20);not null;comment:榜单类型（weekly-周榜, monthly-月榜）" json:"type"`
	Game        string          `gorm:"type:varchar(100);comment:关联游戏" json:"game"`
	StartTime   time.Time       `json:"startTime"`
	EndTime     time.Time       `json:"endTime"`
	Description string          `gorm:"type:text" json:"description"`
	Status      int             `gorm:"default:1;comment:状态（1-启用，0-禁用）" json:"status"`
	SortOrder   int             `gorm:"default:0;comment:排序顺序" json:"sortOrder"`
}

// TableName 设置Leaderboard表名
func (Leaderboard) TableName() string {
	return "game_partner_leaderboards"
}

// LeaderboardItem 排行榜条目模型
type LeaderboardItem struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	LeaderboardID uint     `gorm:"not null;index" json:"leaderboardId"`
	PlaymateID   uint      `gorm:"not null;index" json:"playmateId"`
	Rank         int       `gorm:"not null" json:"rank"`
	Score        float64   `json:"score"`
	OrderCount   int       `json:"orderCount"`
	Rating       float64   `json:"rating"`
	Likes        int       `json:"likes"`
	Playmate     Playmate  `gorm:"foreignKey:PlaymateID" json:"playmate,omitempty"`
}

// TableName 设置LeaderboardItem表名
func (LeaderboardItem) TableName() string {
	return "game_partner_leaderboard_items"
}
