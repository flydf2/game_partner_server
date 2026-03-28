package model

import (
	"time"

	"gorm.io/gorm"
)

// MatchHistory 匹配历史模型
type MatchHistory struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint           `json:"userId"`
	PlaymateID uint          `json:"playmateId"`
	Playmate  Playmate       `gorm:"foreignKey:PlaymateID" json:"playmate"`
	Status    string         `json:"status"` // pending, matched, completed, cancelled
	StartTime *time.Time     `json:"startTime"`
	EndTime   *time.Time     `json:"endTime"`
	Rating    *float64       `json:"rating"`
	Review    string         `json:"review"`
}

// TableName 设置MatchHistory表名
func (MatchHistory) TableName() string {
	return "game_partner_match_histories"
}
