package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
)

// StatsService 统计分析服务
type StatsService struct{}

// GetDashboardStats 获取仪表盘统计数据
func (s *StatsService) GetDashboardStats(startTime, endTime string) (map[string]interface{}, error) {
	// 订单统计
	var totalOrders int64
	var completedOrders int64
	var orderAmount float64

	orderQuery := global.GVA_DB.Model(&model.Order{})
	if startTime != "" {
		orderQuery = orderQuery.Where("created_at >= ?", startTime)
	}
	if endTime != "" {
		orderQuery = orderQuery.Where("created_at <= ?", endTime)
	}

	orderQuery.Count(&totalOrders)
	orderQuery.Where("status = ?", "completed").Count(&completedOrders)
	orderQuery.Select("COALESCE(SUM(amount), 0) as total_amount").Scan(&orderAmount)

	// 用户统计
	var totalUsers int64
	var newUsers int64

	userQuery := global.GVA_DB.Model(&model.User{})
	userQuery.Count(&totalUsers)
	if startTime != "" {
		userQuery.Where("created_at >= ?", startTime).Count(&newUsers)
	}

	// 专家统计
	var totalExperts int64
	var onlineExperts int64

	expertQuery := global.GVA_DB.Model(&model.Playmate{})
	expertQuery.Count(&totalExperts)
	expertQuery.Where("is_online = ?", true).Count(&onlineExperts)

	// 收入统计
	var totalRevenue float64
	var pendingRevenue float64

	revenueQuery := global.GVA_DB.Model(&model.Transaction{}).Where("type = ?", "income")
	if startTime != "" {
		revenueQuery = revenueQuery.Where("created_at >= ?", startTime)
	}
	if endTime != "" {
		revenueQuery = revenueQuery.Where("created_at <= ?", endTime)
	}

	revenueQuery.Select("COALESCE(SUM(amount), 0) as total_amount").Scan(&totalRevenue)

	pendingRevenueQuery := global.GVA_DB.Model(&model.Transaction{}).Where("type = ?", "income_pending")
	if startTime != "" {
		pendingRevenueQuery = pendingRevenueQuery.Where("created_at >= ?", startTime)
	}
	if endTime != "" {
		pendingRevenueQuery = pendingRevenueQuery.Where("created_at <= ?", endTime)
	}

	pendingRevenueQuery.Select("COALESCE(SUM(amount), 0) as total_amount").Scan(&pendingRevenue)

	// 构建仪表盘数据
	dashboardStats := map[string]interface{}{
		"orders": map[string]interface{}{
			"total":     totalOrders,
			"completed": completedOrders,
			"amount":    orderAmount,
		},
		"users": map[string]interface{}{
			"total": totalUsers,
			"new":   newUsers,
		},
		"experts": map[string]interface{}{
			"total":  totalExperts,
			"online": onlineExperts,
		},
		"revenue": map[string]interface{}{
			"total":   totalRevenue,
			"pending": pendingRevenue,
		},
	}

	return dashboardStats, nil
}

// GetOrderStats 获取订单统计数据
func (s *StatsService) GetOrderStats(startTime, endTime, game, status string) (map[string]interface{}, error) {
	var totalOrders int64
	var orderAmount float64

	query := global.GVA_DB.Model(&model.Order{})

	if startTime != "" {
		query = query.Where("created_at >= ?", startTime)
	}
	if endTime != "" {
		query = query.Where("created_at <= ?", endTime)
	}
	if game != "" {
		query = query.Where("game = ?", game)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&totalOrders)
	query.Select("COALESCE(SUM(amount), 0) as total_amount").Scan(&orderAmount)

	// 按状态统计
	var statusStats []map[string]interface{}
	query.Model(&model.Order{}).Select("status, COUNT(*) as count, COALESCE(SUM(amount), 0) as amount").Group("status").Scan(&statusStats)

	// 按游戏统计
	var gameStats []map[string]interface{}
	query.Model(&model.Order{}).Select("game, COUNT(*) as count, COALESCE(SUM(amount), 0) as amount").Group("game").Scan(&gameStats)

	orderStats := map[string]interface{}{
		"total":      totalOrders,
		"amount":     orderAmount,
		"statusStats": statusStats,
		"gameStats":   gameStats,
	}

	return orderStats, nil
}

// GetUserStats 获取用户统计数据
func (s *StatsService) GetUserStats(startTime, endTime string) (map[string]interface{}, error) {
	var totalUsers int64
	var newUsers int64
	var activeUsers int64

	userQuery := global.GVA_DB.Model(&model.User{})
	userQuery.Count(&totalUsers)

	if startTime != "" {
		userQuery.Where("created_at >= ?", startTime).Count(&newUsers)
	}

	if startTime != "" {
		global.GVA_DB.Model(&model.Transaction{}).Where("created_at >= ?", startTime).Distinct("user_id").Count(&activeUsers)
	}

	// 按注册时间统计
	var registrationStats []map[string]interface{}
	if startTime != "" {
		global.GVA_DB.Model(&model.User{}).Select("DATE(created_at) as date, COUNT(*) as count").Where("created_at >= ?", startTime).Group("DATE(created_at)").Scan(&registrationStats)
	}

	userStats := map[string]interface{}{
		"total":             totalUsers,
		"new":               newUsers,
		"active":            activeUsers,
		"registrationStats": registrationStats,
	}

	return userStats, nil
}

// GetExpertStats 获取专家统计数据
func (s *StatsService) GetExpertStats(startTime, endTime, game string) (map[string]interface{}, error) {
	var totalExperts int64
	var onlineExperts int64

	query := global.GVA_DB.Model(&model.Playmate{})

	if game != "" {
		query = query.Where("game = ?", game)
	}

	query.Count(&totalExperts)
	query.Where("is_online = ?", true).Count(&onlineExperts)

	// 按游戏统计
	var gameStats []map[string]interface{}
	query.Model(&model.Playmate{}).Select("game, COUNT(*) as count").Group("game").Scan(&gameStats)

	// 按等级统计
	var levelStats []map[string]interface{}
	query.Model(&model.Playmate{}).Select("level, COUNT(*) as count").Group("level").Scan(&levelStats)

	expertStats := map[string]interface{}{
		"total":      totalExperts,
		"online":     onlineExperts,
		"gameStats":   gameStats,
		"levelStats":  levelStats,
	}

	return expertStats, nil
}

// GetRevenueStats 获取收入统计数据
func (s *StatsService) GetRevenueStats(startTime, endTime, game string) (map[string]interface{}, error) {
	var totalRevenue float64
	var pendingRevenue float64

	revenueQuery := global.GVA_DB.Model(&model.Transaction{}).Where("type = ?", "income")
	if startTime != "" {
		revenueQuery = revenueQuery.Where("created_at >= ?", startTime)
	}
	if endTime != "" {
		revenueQuery = revenueQuery.Where("created_at <= ?", endTime)
	}

	revenueQuery.Select("COALESCE(SUM(amount), 0) as total_amount").Scan(&totalRevenue)

	pendingRevenueQuery := global.GVA_DB.Model(&model.Transaction{}).Where("type = ?", "income_pending")
	if startTime != "" {
		pendingRevenueQuery = pendingRevenueQuery.Where("created_at >= ?", startTime)
	}
	if endTime != "" {
		pendingRevenueQuery = pendingRevenueQuery.Where("created_at <= ?", endTime)
	}

	pendingRevenueQuery.Select("COALESCE(SUM(amount), 0) as total_amount").Scan(&pendingRevenue)

	// 按日期统计
	var dailyStats []map[string]interface{}
	if startTime != "" {
		revenueQuery.Select("DATE(created_at) as date, COALESCE(SUM(amount), 0) as amount").Group("DATE(created_at)").Scan(&dailyStats)
	}

	// 按游戏统计
	var gameStats []map[string]interface{}
	if game != "" {
		// 这里需要通过订单关联游戏信息
		global.GVA_DB.Model(&model.Order{}).Select("game, COALESCE(SUM(amount), 0) as amount").Where("game = ?", game).Group("game").Scan(&gameStats)
	} else {
		global.GVA_DB.Model(&model.Order{}).Select("game, COALESCE(SUM(amount), 0) as amount").Group("game").Scan(&gameStats)
	}

	revenueStats := map[string]interface{}{
		"total":       totalRevenue,
		"pending":     pendingRevenue,
		"dailyStats":  dailyStats,
		"gameStats":   gameStats,
	}

	return revenueStats, nil
}

// GetTrendStats 获取趋势统计数据
func (s *StatsService) GetTrendStats(statsType, startTime, endTime, interval string) (map[string]interface{}, error) {
	var trendData []map[string]interface{}

	switch statsType {
	case "orders":
		query := global.GVA_DB.Model(&model.Order{})
		if startTime != "" {
			query = query.Where("created_at >= ?", startTime)
		}
		if endTime != "" {
			query = query.Where("created_at <= ?", endTime)
		}

		// 根据时间间隔选择分组方式
		var timeFormat string
		switch interval {
		case "day":
			timeFormat = "DATE(created_at)"
		case "week":
			timeFormat = "WEEK(created_at)"
		case "month":
			timeFormat = "DATE_FORMAT(created_at, '%Y-%m')"
		default:
			timeFormat = "DATE(created_at)"
		}

		query.Select(timeFormat+" as time, COUNT(*) as count, COALESCE(SUM(amount), 0) as amount").Group(timeFormat).Order("time").Scan(&trendData)

	case "users":
		query := global.GVA_DB.Model(&model.User{})
		if startTime != "" {
			query = query.Where("created_at >= ?", startTime)
		}
		if endTime != "" {
			query = query.Where("created_at <= ?", endTime)
		}

		var timeFormat string
		switch interval {
		case "day":
			timeFormat = "DATE(created_at)"
		case "week":
			timeFormat = "WEEK(created_at)"
		case "month":
			timeFormat = "DATE_FORMAT(created_at, '%Y-%m')"
		default:
			timeFormat = "DATE(created_at)"
		}

		query.Select(timeFormat+" as time, COUNT(*) as count").Group(timeFormat).Order("time").Scan(&trendData)

	case "revenue":
		query := global.GVA_DB.Model(&model.Transaction{}).Where("type = ?", "income")
		if startTime != "" {
			query = query.Where("created_at >= ?", startTime)
		}
		if endTime != "" {
			query = query.Where("created_at <= ?", endTime)
		}

		var timeFormat string
		switch interval {
		case "day":
			timeFormat = "DATE(created_at)"
		case "week":
			timeFormat = "WEEK(created_at)"
		case "month":
			timeFormat = "DATE_FORMAT(created_at, '%Y-%m')"
		default:
			timeFormat = "DATE(created_at)"
		}

		query.Select(timeFormat+" as time, COALESCE(SUM(amount), 0) as amount").Group(timeFormat).Order("time").Scan(&trendData)

	case "experts":
		query := global.GVA_DB.Model(&model.Playmate{})
		if startTime != "" {
			query = query.Where("created_at >= ?", startTime)
		}
		if endTime != "" {
			query = query.Where("created_at <= ?", endTime)
		}

		var timeFormat string
		switch interval {
		case "day":
			timeFormat = "DATE(created_at)"
		case "week":
			timeFormat = "WEEK(created_at)"
		case "month":
			timeFormat = "DATE_FORMAT(created_at, '%Y-%m')"
		default:
			timeFormat = "DATE(created_at)"
		}

		query.Select(timeFormat+" as time, COUNT(*) as count").Group(timeFormat).Order("time").Scan(&trendData)
	}

	trendStats := map[string]interface{}{
		"type": statsType,
		"data": trendData,
	}

	return trendStats, nil
}
