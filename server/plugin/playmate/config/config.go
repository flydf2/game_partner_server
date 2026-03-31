package config

import (
	model "github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// gp 辅助函数，将bool值转换为*bool类型
func gp(b bool) *bool {
	return &b
}

// GetApis 获取API配置
func GetApis() []model.SysApi {
	return []model.SysApi{
		// 陪玩相关API
		{Path: "/playmates", Method: "GET", Description: "获取陪玩列表", ApiGroup: "playmate"},
		{Path: "/playmates/search", Method: "GET", Description: "搜索陪玩", ApiGroup: "playmate"},
		{Path: "/playmates/suggestions", Method: "GET", Description: "获取搜索建议", ApiGroup: "playmate"},
		{Path: "/playmates/:id", Method: "GET", Description: "获取陪玩详情", ApiGroup: "playmate"},
		{Path: "/playmates", Method: "POST", Description: "创建陪玩", ApiGroup: "playmate"},
		{Path: "/playmates/:id", Method: "PUT", Description: "更新陪玩", ApiGroup: "playmate"},
		{Path: "/playmates/:id", Method: "DELETE", Description: "删除陪玩", ApiGroup: "playmate"},

		// 专家相关API
		{Path: "/experts/:id", Method: "GET", Description: "获取专家详情", ApiGroup: "playmate"},
		{Path: "/experts/:id/follow", Method: "POST", Description: "关注专家", ApiGroup: "playmate"},
		{Path: "/experts/:id/follow", Method: "DELETE", Description: "取消关注专家", ApiGroup: "playmate"},
		{Path: "/experts/:id/reviews", Method: "GET", Description: "获取专家评价", ApiGroup: "playmate"},
		{Path: "/experts/verifications", Method: "GET", Description: "获取专家认证列表", ApiGroup: "playmate"},
		{Path: "/experts/verifications/:id", Method: "GET", Description: "获取专家认证详情", ApiGroup: "playmate"},
		{Path: "/experts/verifications/batch", Method: "POST", Description: "批量处理专家认证", ApiGroup: "playmate"},
		{Path: "/experts/verifications/export", Method: "GET", Description: "导出专家认证数据", ApiGroup: "playmate"},
		{Path: "/experts/verifications/stats", Method: "GET", Description: "获取专家认证统计", ApiGroup: "playmate"},

		// 用户相关API
		{Path: "/user/info", Method: "GET", Description: "获取用户信息", ApiGroup: "playmate"},
		{Path: "/user/profile", Method: "PUT", Description: "更新用户资料", ApiGroup: "playmate"},
		{Path: "/user/settings", Method: "GET", Description: "获取用户设置", ApiGroup: "playmate"},
		{Path: "/user/settings", Method: "PUT", Description: "更新用户设置", ApiGroup: "playmate"},
		{Path: "/user/following", Method: "GET", Description: "获取关注列表", ApiGroup: "playmate"},
		{Path: "/user/favorites", Method: "GET", Description: "获取收藏列表", ApiGroup: "playmate"},
		{Path: "/user/history", Method: "GET", Description: "获取浏览历史", ApiGroup: "playmate"},
		{Path: "/user/wallet", Method: "GET", Description: "获取钱包信息", ApiGroup: "playmate"},
		{Path: "/auth/login", Method: "POST", Description: "用户登录", ApiGroup: "playmate"},
		{Path: "/auth/register", Method: "POST", Description: "用户注册", ApiGroup: "playmate"},
		{Path: "/auth/logout", Method: "POST", Description: "用户登出", ApiGroup: "playmate"},
		{Path: "/auth/refresh", Method: "POST", Description: "刷新令牌", ApiGroup: "playmate"},
		{Path: "/users", Method: "GET", Description: "获取用户列表", ApiGroup: "playmate"},
		{Path: "/users/:id", Method: "GET", Description: "获取用户详情", ApiGroup: "playmate"},
		{Path: "/users/:id", Method: "PUT", Description: "更新用户", ApiGroup: "playmate"},
		{Path: "/users/:id/disable", Method: "POST", Description: "禁用用户", ApiGroup: "playmate"},
		{Path: "/users/:id/enable", Method: "POST", Description: "启用用户", ApiGroup: "playmate"},
		{Path: "/users/:id/reset-password", Method: "POST", Description: "重置用户密码", ApiGroup: "playmate"},
		{Path: "/users/stats", Method: "GET", Description: "获取用户统计", ApiGroup: "playmate"},
		{Path: "/users/export", Method: "GET", Description: "导出用户数据", ApiGroup: "playmate"},

		// 订单相关API
		{Path: "/orders", Method: "GET", Description: "获取订单列表", ApiGroup: "playmate"},
		{Path: "/orders/:id", Method: "GET", Description: "获取订单详情", ApiGroup: "playmate"},
		{Path: "/orders", Method: "POST", Description: "创建订单", ApiGroup: "playmate"},
		{Path: "/orders/:id/confirmation", Method: "GET", Description: "获取订单确认", ApiGroup: "playmate"},
		{Path: "/orders/all", Method: "GET", Description: "获取所有订单列表", ApiGroup: "playmate"},
		{Path: "/orders/batch", Method: "POST", Description: "批量处理订单", ApiGroup: "playmate"},
		{Path: "/orders/stats", Method: "GET", Description: "获取订单统计", ApiGroup: "playmate"},
		{Path: "/orders/export", Method: "GET", Description: "导出订单数据", ApiGroup: "playmate"},

		// 通知相关API
		{Path: "/notifications", Method: "GET", Description: "获取通知列表", ApiGroup: "playmate"},
		{Path: "/notifications/:id/read", Method: "PUT", Description: "标记通知已读", ApiGroup: "playmate"},
		{Path: "/notifications/read-all", Method: "PUT", Description: "标记所有通知已读", ApiGroup: "playmate"},

		// 消息相关API
		{Path: "/messages", Method: "GET", Description: "获取消息列表", ApiGroup: "playmate"},
		{Path: "/messages/chat/:userId", Method: "GET", Description: "获取聊天消息", ApiGroup: "playmate"},
		{Path: "/messages/chat/:userId", Method: "POST", Description: "发送消息", ApiGroup: "playmate"},

		// 游戏相关API
		{Path: "/games", Method: "GET", Description: "获取游戏列表", ApiGroup: "playmate"},
		{Path: "/activities", Method: "GET", Description: "获取活动列表", ApiGroup: "playmate"},

		// 评价相关API
		{Path: "/reviews", Method: "POST", Description: "提交评价", ApiGroup: "playmate"},
		{Path: "/reviews", Method: "GET", Description: "获取评价列表", ApiGroup: "playmate"},

		// 提现相关API
		{Path: "/withdrawals", Method: "POST", Description: "提交提现申请", ApiGroup: "playmate"},
		{Path: "/withdrawals", Method: "GET", Description: "获取提现记录", ApiGroup: "playmate"},

		// 社区相关API
		{Path: "/community/posts", Method: "GET", Description: "获取社区帖子", ApiGroup: "playmate"},
		{Path: "/community/posts/:postId", Method: "GET", Description: "获取帖子详情", ApiGroup: "playmate"},
		{Path: "/community/posts/:postId/like", Method: "POST", Description: "点赞帖子", ApiGroup: "playmate"},
		{Path: "/community/posts/:postId/comments", Method: "POST", Description: "评论帖子", ApiGroup: "playmate"},

		// 分类相关API
		{Path: "/categories", Method: "GET", Description: "获取分类列表", ApiGroup: "playmate"},
		{Path: "/game-categories", Method: "GET", Description: "获取游戏分类", ApiGroup: "playmate"},
		{Path: "/game-categories/:id", Method: "GET", Description: "获取游戏分类详情", ApiGroup: "playmate"},
		{Path: "/game-categories", Method: "POST", Description: "创建游戏分类", ApiGroup: "playmate"},
		{Path: "/game-categories/:id", Method: "PUT", Description: "更新游戏分类", ApiGroup: "playmate"},
		{Path: "/game-categories/:id", Method: "DELETE", Description: "删除游戏分类", ApiGroup: "playmate"},
		{Path: "/game-categories/:category/games", Method: "GET", Description: "获取分类游戏", ApiGroup: "playmate"},

		// 统计分析API
		{Path: "/stats/dashboard", Method: "GET", Description: "获取仪表盘统计数据", ApiGroup: "playmate"},
		{Path: "/stats/orders", Method: "GET", Description: "获取订单统计数据", ApiGroup: "playmate"},
		{Path: "/stats/users", Method: "GET", Description: "获取用户统计数据", ApiGroup: "playmate"},
		{Path: "/stats/experts", Method: "GET", Description: "获取专家统计数据", ApiGroup: "playmate"},
		{Path: "/stats/revenue", Method: "GET", Description: "获取收入统计数据", ApiGroup: "playmate"},
		{Path: "/stats/trend", Method: "GET", Description: "获取趋势统计数据", ApiGroup: "playmate"},
	}
}

// GetDictionaries 获取字典配置
func GetDictionaries() []model.SysDictionary {
	return []model.SysDictionary{
		{
			Name:   "陪玩状态",
			Type:   "playmate_status",
			Desc:   "陪玩的在线状态",
			Status: gp(true),
		},
		{
			Name:   "订单状态",
			Type:   "order_status",
			Desc:   "订单的状态",
			Status: gp(true),
		},
		{
			Name:   "用户角色",
			Type:   "user_role",
			Desc:   "用户的角色类型",
			Status: gp(true),
		},
		{
			Name:   "游戏分类",
			Type:   "game_category",
			Desc:   "游戏的分类",
			Status: gp(true),
		},
		{
			Name:   "评价等级",
			Type:   "review_level",
			Desc:   "评价的等级",
			Status: gp(true),
		},
	}
}

// GetDictionaryDetails 获取字典详情配置
func GetDictionaryDetails() map[string][]model.SysDictionaryDetail {
	return map[string][]model.SysDictionaryDetail{
		"playmate_status": {
			{Label: "在线", Value: "online", Status: gp(true), Sort: 1},
			{Label: "离线", Value: "offline", Status: gp(true), Sort: 2},
			{Label: "忙碌", Value: "busy", Status: gp(true), Sort: 3},
		},
		"order_status": {
			{Label: "待确认", Value: "pending", Status: gp(true), Sort: 1},
			{Label: "已确认", Value: "confirmed", Status: gp(true), Sort: 2},
			{Label: "进行中", Value: "ongoing", Status: gp(true), Sort: 3},
			{Label: "已完成", Value: "completed", Status: gp(true), Sort: 4},
			{Label: "已取消", Value: "cancelled", Status: gp(true), Sort: 5},
		},
		"user_role": {
			{Label: "普通用户", Value: "user", Status: gp(true), Sort: 1},
			{Label: "陪玩", Value: "playmate", Status: gp(true), Sort: 2},
			{Label: "管理员", Value: "admin", Status: gp(true), Sort: 3},
		},
		"game_category": {
			{Label: "MOBA", Value: "moba", Status: gp(true), Sort: 1},
			{Label: "FPS", Value: "fps", Status: gp(true), Sort: 2},
			{Label: "RPG", Value: "rpg", Status: gp(true), Sort: 3},
			{Label: "策略", Value: "strategy", Status: gp(true), Sort: 4},
			{Label: "休闲", Value: "casual", Status: gp(true), Sort: 5},
		},
		"review_level": {
			{Label: "非常满意", Value: "5", Status: gp(true), Sort: 1},
			{Label: "满意", Value: "4", Status: gp(true), Sort: 2},
			{Label: "一般", Value: "3", Status: gp(true), Sort: 3},
			{Label: "不满意", Value: "2", Status: gp(true), Sort: 4},
			{Label: "非常不满意", Value: "1", Status: gp(true), Sort: 5},
		},
	}
}

// GetMenus 获取菜单配置
func GetMenus() []model.SysBaseMenu {
	return []model.SysBaseMenu{
		// 主菜单：陪玩服务
		{
			Path:      "playmate",
			Name:      "playmate",
			Hidden:    false,
			Component: "Layout",
			Sort:      1,
			ParentId:  0,
			Meta: model.Meta{
				Title:     "陪玩服务",
				Icon:      "gamepad",
				KeepAlive: true,
				CloseTab:  false,
			},
		},
		// 子菜单：陪玩管理
		{
			Path:      "playmateList",
			Name:      "playmateList",
			Hidden:    false,
			Component: "plugin/playmate/view/playmateList.vue",
			Sort:      1,
			ParentId:  1,
			Meta: model.Meta{
				Title:     "陪玩管理",
				Icon:      "menu",
				KeepAlive: true,
				CloseTab:  false,
			},
		},
		// 子菜单：用户管理
		{
			Path:      "userList",
			Name:      "userList",
			Hidden:    false,
			Component: "plugin/playmate/view/userList.vue",
			Sort:      2,
			ParentId:  1,
			Meta: model.Meta{
				Title:     "用户管理",
				Icon:      "user",
				KeepAlive: true,
				CloseTab:  false,
			},
		},
		// 子菜单：订单管理
		{
			Path:      "orderList",
			Name:      "orderList",
			Hidden:    false,
			Component: "plugin/playmate/view/orderList.vue",
			Sort:      3,
			ParentId:  1,
			Meta: model.Meta{
				Title:     "订单管理",
				Icon:      "shopping-cart",
				KeepAlive: true,
				CloseTab:  false,
			},
		},
		// 子菜单：评价管理
		{
			Path:      "reviewList",
			Name:      "reviewList",
			Hidden:    false,
			Component: "plugin/playmate/view/reviewList.vue",
			Sort:      4,
			ParentId:  1,
			Meta: model.Meta{
				Title:     "评价管理",
				Icon:      "star",
				KeepAlive: true,
				CloseTab:  false,
			},
		},
		// 子菜单：财务管理
		{
			Path:      "withdrawalList",
			Name:      "withdrawalList",
			Hidden:    false,
			Component: "plugin/playmate/view/withdrawalList.vue",
			Sort:      5,
			ParentId:  1,
			Meta: model.Meta{
				Title:     "财务管理",
				Icon:      "money",
				KeepAlive: true,
				CloseTab:  false,
			},
		},
		// 子菜单：社区管理
		{
			Path:      "communityList",
			Name:      "communityList",
			Hidden:    false,
			Component: "plugin/playmate/view/communityList.vue",
			Sort:      6,
			ParentId:  1,
			Meta: model.Meta{
				Title:     "社区管理",
				Icon:      "chat",
				KeepAlive: true,
				CloseTab:  false,
			},
		},
		// 子菜单：游戏管理
		{
			Path:      "gameList",
			Name:      "gameList",
			Hidden:    false,
			Component: "plugin/playmate/view/gameList.vue",
			Sort:      7,
			ParentId:  1,
			Meta: model.Meta{
				Title:     "游戏管理",
				Icon:      "gamepad",
				KeepAlive: true,
				CloseTab:  false,
			},
		},
		// 子菜单：活动管理
		{
			Path:      "activityList",
			Name:      "activityList",
			Hidden:    false,
			Component: "plugin/playmate/view/activityList.vue",
			Sort:      8,
			ParentId:  1,
			Meta: model.Meta{
				Title:     "活动管理",
				Icon:      "calendar",
				KeepAlive: true,
				CloseTab:  false,
			},
		},
		// 子菜单：游戏分类管理
		{
			Path:      "gameCategoryList",
			Name:      "gameCategoryList",
			Hidden:    false,
			Component: "plugin/playmate/view/gameCategoryList.vue",
			Sort:      9,
			ParentId:  1,
			Meta: model.Meta{
				Title:     "游戏分类管理",
				Icon:      "category",
				KeepAlive: true,
				CloseTab:  false,
			},
		},
		// 子菜单：专家认证管理
		{
			Path:      "expertVerificationList",
			Name:      "expertVerificationList",
			Hidden:    false,
			Component: "plugin/playmate/view/expertVerificationList.vue",
			Sort:      10,
			ParentId:  1,
			Meta: model.Meta{
				Title:     "专家认证管理",
				Icon:      "identity",
				KeepAlive: true,
				CloseTab:  false,
			},
		},
		// 子菜单：统计分析
		{
			Path:      "statsDashboard",
			Name:      "statsDashboard",
			Hidden:    false,
			Component: "plugin/playmate/view/statsDashboard.vue",
			Sort:      11,
			ParentId:  1,
			Meta: model.Meta{
				Title:     "统计分析",
				Icon:      "data-analysis",
				KeepAlive: true,
				CloseTab:  false,
			},
		},
	}
}
