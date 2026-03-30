package model

import (
	"time"

	"gorm.io/gorm"
)

// ExpertOrderSetting 专家订单设置模型
type ExpertOrderSetting struct {
	ID                  uint           `gorm:"primaryKey" json:"id"`
	CreatedAt           time.Time      `json:"createdAt"`
	UpdatedAt           time.Time      `json:"updatedAt"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"-"`
	ExpertID            uint           `json:"expertId" gorm:"uniqueIndex;not null;comment:专家ID"`
	AutoAccept          bool           `json:"autoAccept" gorm:"default:false;comment:是否自动接单"`
	AcceptMode          string         `json:"acceptMode" gorm:"default:'manual';comment:接单模式 manual-手动 auto-自动"`
	MaxDailyOrders      int            `json:"maxDailyOrders" gorm:"default:10;comment:每日最大接单数"`
	ServiceStartTime    string         `json:"serviceStartTime" gorm:"default:'09:00';comment:服务开始时间"`
	ServiceEndTime      string         `json:"serviceEndTime" gorm:"default:'23:00';comment:服务结束时间"`
	RestDays            string         `json:"restDays" gorm:"comment:休息日 用逗号分隔 0-6代表周日到周六"`
	MinOrderAmount      float64        `json:"minOrderAmount" gorm:"default:0;comment:最小订单金额"`
	MaxOrderAmount      float64        `json:"maxOrderAmount" gorm:"default:0;comment:最大订单金额 0表示不限"`
	AcceptableGames     string         `json:"acceptableGames" gorm:"comment:可接单游戏 用逗号分隔"`
	RejectWithoutVoice  bool           `json:"rejectWithoutVoice" gorm:"default:false;comment:是否拒绝无语音介绍的订单"`
	NotificationEnabled bool           `json:"notificationEnabled" gorm:"default:true;comment:是否开启订单通知"`
}

// TableName 设置ExpertOrderSetting表名
func (ExpertOrderSetting) TableName() string {
	return "game_partner_expert_order_settings"
}

// ExpertService 专家服务项目模型
type ExpertService struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	ExpertID    uint           `json:"expertId" gorm:"index;not null;comment:专家ID"`
	Name        string         `json:"name" gorm:"not null;comment:服务名称"`
	Description string         `json:"description" gorm:"comment:服务描述"`
	Price       float64        `json:"price" gorm:"not null;comment:服务价格"`
	Unit        string         `json:"unit" gorm:"default:'hour';comment:计价单位 hour-小时 game-局"`
	Duration    int            `json:"duration" gorm:"default:60;comment:服务时长(分钟)"`
	GameID      uint           `json:"gameId" gorm:"comment:关联游戏ID"`
	GameName    string         `json:"gameName" gorm:"comment:游戏名称"`
	IsEnabled   bool           `json:"isEnabled" gorm:"default:true;comment:是否启用"`
	SortOrder   int            `json:"sortOrder" gorm:"default:0;comment:排序"`
}

// TableName 设置ExpertService表名
func (ExpertService) TableName() string {
	return "game_partner_expert_services"
}

// TodayRecommendation 今日推荐模型
type TodayRecommendation struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	ExpertID      uint           `json:"expertId" gorm:"index;not null;comment:专家ID"`
	Title         string         `json:"title" gorm:"not null;comment:推荐标题"`
	Content       string         `json:"content" gorm:"type:text;comment:推荐内容"`
	ImageURL      string         `json:"imageUrl" gorm:"comment:推荐图片"`
	ServiceID     uint           `json:"serviceId" gorm:"comment:关联服务ID"`
	ServiceName   string         `json:"serviceName" gorm:"comment:服务名称"`
	Price         float64        `json:"price" gorm:"comment:推荐价格"`
	OriginalPrice float64        `json:"originalPrice" gorm:"comment:原价"`
	IsEnabled     bool           `json:"isEnabled" gorm:"default:true;comment:是否启用"`
	SortOrder     int            `json:"sortOrder" gorm:"default:0;comment:排序"`
	StartTime     *time.Time     `json:"startTime" gorm:"comment:开始时间"`
	EndTime       *time.Time     `json:"endTime" gorm:"comment:结束时间"`
}

// TableName 设置TodayRecommendation表名
func (TodayRecommendation) TableName() string {
	return "game_partner_today_recommendations"
}
