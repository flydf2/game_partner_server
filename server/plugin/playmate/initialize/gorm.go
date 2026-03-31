package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
)

// InitializeDB 初始化数据库
func InitializeDB() {
	global.GVA_DB.AutoMigrate(
		// 用户相关
		&model.User{},
		&model.UserSettings{},
		&model.UserWallet{},
		&model.Transaction{},

		// 陪玩相关
		&model.Playmate{},
		&model.PlaymateSkill{},
		&model.PlaymateVoiceIntroduction{},

		// 订单相关
		&model.Order{},
		&model.OrderConfirmation{},
		&model.RewardOrder{},
		&model.RewardOrderApplicant{},
		&model.RewardOrderPayment{},
		&model.RewardOrderReview{},
		&model.Coupon{},

		// 通知相关
		&model.Notification{},
		&model.Message{},
		&model.Conversation{},

		// 游戏相关
		&model.Game{},
		&model.Activity{},
		&model.Category{},
		&model.GameCategory{},

		// 评价相关
		&model.Review{},
		&model.Withdrawal{},

		// 社区相关
		&model.CommunityPost{},
		&model.Comment{},
		&model.Recommendation{},
		&model.UserFollow{},
		&model.UserFavorite{},
		&model.UserBrowseHistory{},

		// 分享相关
		&model.Share{},

		// 匹配历史相关
		&model.MatchHistory{},

		// 专家订单设置相关
		&model.ExpertOrderSetting{},
		&model.ExpertService{},
		&model.TodayRecommendation{},

		// 赛事相关
		&model.Tournament{},
		&model.TournamentTeam{},
		&model.TournamentMatch{},
		&model.TournamentRegistration{},

		// 排行榜相关
		&model.Leaderboard{},
		&model.LeaderboardItem{},

		// 专家认证相关
		&model.ExpertVerification{},
	)
}
