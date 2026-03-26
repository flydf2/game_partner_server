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
	User      User           `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	Content   string         `json:"content"`
	Images    string         `gorm:"type:varchar(4096)" json:"images"` // 用逗号分隔的图片URL
	Likes     int            `json:"likes"`
	Comments  int            `json:"comments"`
	Game      string         `json:"game"`
}

// TableName 设置CommunityPost表名
func (CommunityPost) TableName() string {
	return "game_partner_community_posts"
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

// TableName 设置Comment表名
func (Comment) TableName() string {
	return "game_partner_comments"
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

// TableName 设置Recommendation表名
func (Recommendation) TableName() string {
	return "game_partner_recommendations"
}

// UserFollow 用户关注模型
type UserFollow struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint           `json:"userId"`
	FollowID  uint           `json:"followId"` // 关注的用户ID
}

// TableName 设置UserFollow表名
func (UserFollow) TableName() string {
	return "game_partner_user_follows"
}

// UserFavorite 用户收藏模型
type UserFavorite struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"createdAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	UserID     uint           `json:"userId"`
	PlaymateID uint           `json:"playmateId"`
}

// TableName 设置UserFavorite表名
func (UserFavorite) TableName() string {
	return "game_partner_user_favorites"
}

// UserBrowseHistory 用户浏览历史模型
type UserBrowseHistory struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"createdAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	UserID     uint           `json:"userId"`
	PlaymateID uint           `json:"playmateId"`
	ViewedAt   time.Time      `json:"viewedAt"`
}

// TableName 设置UserBrowseHistory表名
func (UserBrowseHistory) TableName() string {
	return "game_partner_user_browse_histories"
}
