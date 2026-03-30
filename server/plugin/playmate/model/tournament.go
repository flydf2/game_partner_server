package model

import (
	"time"

	"gorm.io/gorm"
)

// Tournament 赛事模型
type Tournament struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	Title         string         `json:"title" gorm:"not null;comment:赛事标题"`
	Description   string         `json:"description" gorm:"type:text;comment:赛事描述"`
	Cover         string         `json:"cover" gorm:"type:varchar(2048);comment:封面图片"`
	Game          string         `json:"game" gorm:"comment:游戏名称"`
	GameID        uint           `json:"gameId" gorm:"comment:游戏ID"`
	Status        string         `json:"status" gorm:"default:'upcoming';comment:赛事状态 upcoming-报名中 ongoing-进行中 completed-已结束"`
	RegisterStart time.Time      `json:"registerStart" gorm:"comment:报名开始时间"`
	RegisterEnd   time.Time      `json:"registerEnd" gorm:"comment:报名结束时间"`
	MatchStart    time.Time      `json:"matchStart" gorm:"comment:比赛开始时间"`
	MatchEnd      time.Time      `json:"matchEnd" gorm:"comment:比赛结束时间"`
	Prize         string         `json:"prize" gorm:"comment:奖金池"`
	Participants  int            `json:"participants" gorm:"default:0;comment:参赛人数"`
	MaxTeams      int            `json:"maxTeams" gorm:"default:128;comment:最大参赛队伍数"`
	Rules         string         `json:"rules" gorm:"type:text;comment:赛事规则(JSON格式)"`
	Format        string         `json:"format" gorm:"comment:赛制 BO3/BO5等"`
	MinRank       string         `json:"minRank" gorm:"comment:最低段位要求"`
	Organizer     string         `json:"organizer" gorm:"comment:主办方"`
	ContactInfo   string         `json:"contactInfo" gorm:"comment:联系方式"`
}

// TableName 设置Tournament表名
func (Tournament) TableName() string {
	return "game_partner_tournaments"
}

// TournamentTeam 赛事参赛队伍模型
type TournamentTeam struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	TournamentID uint           `json:"tournamentId" gorm:"not null;index;comment:赛事ID"`
	Name         string         `json:"name" gorm:"not null;comment:队伍名称"`
	Avatar       string         `json:"avatar" gorm:"type:varchar(2048);comment:队伍头像"`
	Members      int            `json:"members" gorm:"default:5;comment:成员数量"`
	LeaderID     uint           `json:"leaderId" gorm:"comment:队长用户ID"`
	Rank         string         `json:"rank" gorm:"comment:最终排名"`
	Status       string         `json:"status" gorm:"default:'registered';comment:状态 registered-已报名 approved-已通过 rejected-已拒绝"`
}

// TableName 设置TournamentTeam表名
func (TournamentTeam) TableName() string {
	return "game_partner_tournament_teams"
}

// TournamentMatch 赛事比赛模型
type TournamentMatch struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	TournamentID uint           `json:"tournamentId" gorm:"not null;index;comment:赛事ID"`
	Round        int            `json:"round" gorm:"comment:轮次"`
	MatchTime    time.Time      `json:"matchTime" gorm:"comment:比赛时间"`
	Team1ID      uint           `json:"team1Id" gorm:"comment:队伍1ID"`
	Team1Name    string         `json:"team1Name" gorm:"comment:队伍1名称"`
	Team1Avatar  string         `json:"team1Avatar" gorm:"type:varchar(2048);comment:队伍1头像"`
	Team2ID      uint           `json:"team2Id" gorm:"comment:队伍2ID"`
	Team2Name    string         `json:"team2Name" gorm:"comment:队伍2名称"`
	Team2Avatar  string         `json:"team2Avatar" gorm:"type:varchar(2048);comment:队伍2头像"`
	Score1       *int           `json:"score1" gorm:"comment:队伍1比分"`
	Score2       *int           `json:"score2" gorm:"comment:队伍2比分"`
	Status       string         `json:"status" gorm:"default:'upcoming';comment:比赛状态 upcoming-未开始 ongoing-进行中 completed-已结束"`
	WinnerID     *uint          `json:"winnerId" gorm:"comment:获胜队伍ID"`
}

// TableName 设置TournamentMatch表名
func (TournamentMatch) TableName() string {
	return "game_partner_tournament_matches"
}

// TournamentRegistration 赛事报名记录模型
type TournamentRegistration struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	TournamentID uint           `json:"tournamentId" gorm:"not null;index;comment:赛事ID"`
	UserID       uint           `json:"userId" gorm:"not null;index;comment:用户ID"`
	TeamName     string         `json:"teamName" gorm:"not null;comment:队伍名称"`
	ContactInfo  string         `json:"contactInfo" gorm:"comment:联系信息"`
	MembersInfo  string         `json:"membersInfo" gorm:"type:text;comment:成员信息(JSON格式)"`
	Status       string         `json:"status" gorm:"default:'pending';comment:报名状态 pending-待审核 approved-已通过 rejected-已拒绝"`
	ReviewRemark string         `json:"reviewRemark" gorm:"comment:审核备注"`
}

// TableName 设置TournamentRegistration表名
func (TournamentRegistration) TableName() string {
	return "game_partner_tournament_registrations"
}
