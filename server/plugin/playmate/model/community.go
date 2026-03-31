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
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	PostID       uint           `json:"postId"`
	UserID       uint           `json:"userId"`
	User         User           `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	Content      string         `json:"content"`
	Likes        int            `json:"likes"`
	ParentID     *uint          `json:"parentId"` // 被引用的评论ID
	ParentComment *Comment      `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE" json:"parentComment,omitempty"` // 被引用的评论
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

// Topic 话题模型
type Topic struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
	Title            string         `json:"title"`
	Description      string         `json:"description"`
	Cover            string         `json:"cover"`
	ParticipantCount int            `json:"participantCount"`
	PostCount        int            `json:"postCount"`
}

// TableName 设置Topic表名
func (Topic) TableName() string {
	return "game_partner_topics"
}

// UserTopicFollow 用户话题关注模型
type UserTopicFollow struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint           `json:"userId"`
	TopicID   uint           `json:"topicId"`
}

// TableName 设置UserTopicFollow表名
func (UserTopicFollow) TableName() string {
	return "game_partner_user_topic_follows"
}

// CommunityBid 社区投标模型
type CommunityBid struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	PostID    uint           `json:"postId"`
	UserID    uint           `json:"userId"`
	User      User           `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user"`
	Message   string         `json:"message"`
	Status    string         `json:"status"` // pending, accepted, rejected, cancelled
}

// TableName 设置CommunityBid表名
func (CommunityBid) TableName() string {
	return "game_partner_community_bids"
}
