package model

import (
	"time"

	"gorm.io/gorm"
)

// Playmate 陪玩专家模型
type Playmate struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	UserID      uint           `json:"userId"`
	Nickname    string         `json:"nickname"`
	Avatar      string         `gorm:"type:varchar(2048)" json:"avatar"`
	Rating      float64        `json:"rating"`
	Price       float64        `json:"price"`
	Likes       int            `json:"likes"`
	Tags        string         `json:"tags"` // 用逗号分隔的标签
	IsOnline    bool           `json:"isOnline"`
	Game        string         `json:"game"`
	Rank        string         `json:"rank"`
	Gender      string         `json:"gender"`
	Description string         `json:"description"`
	Level       int            `json:"level"`
	Title       string         `json:"title"`
}

// TableName 设置Playmate表名
func (Playmate) TableName() string {
	return "game_partner_playmates"
}

// PlaymateSkill 陪玩技能模型
type PlaymateSkill struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	PlaymateID  uint      `json:"playmateId"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Level       string    `json:"level"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// TableName 设置PlaymateSkill表名
func (PlaymateSkill) TableName() string {
	return "game_partner_playmate_skills"
}

// PlaymateVoiceIntroduction 语音介绍模型
type PlaymateVoiceIntroduction struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	PlaymateID uint      `json:"playmateId"`
	URL        string    `json:"url"`
	Duration   string    `json:"duration"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// TableName 设置PlaymateVoiceIntroduction表名
func (PlaymateVoiceIntroduction) TableName() string {
	return "game_partner_playmate_voice_introductions"
}

// ExpertVerification 专家认证模型
type ExpertVerification struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"userId"`
	GameID      uint      `json:"gameId"`
	GameName    string    `json:"gameName"`
	Rank        string    `json:"rank"`
	Positions   string    `json:"positions"` // 用逗号分隔的位置列表
	Screenshots string    `json:"screenshots"` // 用逗号分隔的截图URL列表
	VoiceURL    string    `json:"voiceUrl"`
	Status      string    `json:"status"` // pending, approved, rejected
	Reason      string    `json:"reason"` // 拒绝原因
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// TableName 设置ExpertVerification表名
func (ExpertVerification) TableName() string {
	return "game_partner_expert_verifications"
}
