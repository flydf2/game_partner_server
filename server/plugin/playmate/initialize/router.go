package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/router"
	"github.com/gin-gonic/gin"
)

// InitializeRouter 初始化路由
func InitializeRouter(routerGroup *gin.RouterGroup) {
	// 初始化API
	api.ApiGroupApp = &api.ApiGroup{}

	routerGroup = routerGroup.Group("/playmate")

	// 初始化RouterGroupApp
	router.RouterGroupApp = &router.RouterGroup{}

	// 初始化playmate路由
	router.RouterGroupApp.InitPlaymateRouter(routerGroup)

	// 初始化用户路由
	router.RouterGroupApp.InitUserRouter(routerGroup)

	// 初始化订单路由
	router.RouterGroupApp.InitOrderRouter(routerGroup)

	// 初始化通知路由
	router.RouterGroupApp.InitNotificationRouter(routerGroup)

	// 初始化消息路由
	router.RouterGroupApp.InitMessageRouter(routerGroup)

	// 初始化游戏路由
	router.RouterGroupApp.InitGameRouter(routerGroup)

	// 初始化活动路由
	router.RouterGroupApp.InitActivityRouter(routerGroup)

	// 初始化评价路由
	router.RouterGroupApp.InitReviewRouter(routerGroup)

	// 初始化提现路由
	router.RouterGroupApp.InitWithdrawalRouter(routerGroup)

	// 初始化社区路由
	router.RouterGroupApp.InitCommunityRouter(routerGroup)

	// 初始化分类路由
	router.RouterGroupApp.InitCategoryRouter(routerGroup)

	// 初始化游戏分类路由
	router.RouterGroupApp.InitGameCategoryRouter(routerGroup)

	// 初始化奖励订单路由
	router.RouterGroupApp.InitRewardOrderRouter(routerGroup)

	// 初始化申诉路由
	router.RouterGroupApp.InitAppealRouter(routerGroup)

	// 初始化文件上传路由
	router.RouterGroupApp.InitUploadRouter(routerGroup)

	// 初始化测试工具路由（仅用于自动化测试）
	router.RouterGroupApp.InitTestToolRouter(routerGroup)

	// 初始化专家订单设置路由
	router.RouterGroupApp.InitExpertOrderSettingRouter(routerGroup)

	// 初始化赛事路由
	router.RouterGroupApp.InitTournamentRouter(routerGroup)

	// 初始化排行榜路由
	router.RouterGroupApp.InitLeaderboardRouter(routerGroup)
}
