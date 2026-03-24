package model

import (
	"time"

	"gorm.io/gorm"
)

// CommunityPost 社区帖子模型
type CommunityPost struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint           `json:"userId"`
	Content   string         `json:"content"`
	Images    string         `json:"images"` // 用逗号分隔的图片URL
	Likes     int            `json:"likes"`
	Comments  int            `json:"comments"`
	Game      string         `json:"game"`
}

// Comment 评论模型
type Comment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	PostID    uint           `json:"postId"`
	UserID    uint           `json:"userId"`
	Content   string         `json:"content"`
	Likes     int            `json:"likes"`
}

// Recommendation 推荐模型
type Recommendation struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Type      string         `json:"type"` // playmate, activity
	Data      string         `json:"data"` // JSON格式存储推荐数据
}

// UserFollow 用户关注模型
type UserFollow struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint           `json:"userId"`
	FollowID  uint           `json:"followId"` // 关注的用户ID
}

// UserFavorite 用户收藏模型
type UserFavorite struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint           `json:"userId"`
	PlaymateID uint          `json:"playmateId"`
}

// UserBrowseHistory 用户浏览历史模型
type UserBrowseHistory struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint           `json:"userId"`
	PlaymateID uint          `json:"playmateId"`
	ViewedAt  time.Time      `json:"viewedAt"`
}