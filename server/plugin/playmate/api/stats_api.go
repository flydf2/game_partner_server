package api

import (
	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
)

// StatsApi 统计分析API
type StatsApi struct{}

// GetDashboardStats 获取仪表盘统计数据
// @Tags     Stats
// @Summary  获取仪表盘统计数据
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    startTime query    string  false "开始时间"
// @Param    endTime   query    string  false "结束时间"
// @Success  200       {object}  response.Response{data=map[string]interface{}} "获取成功"
// @Router   /playmate/stats/dashboard [get]
func (a *StatsApi) GetDashboardStats(c *gin.Context) {
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")

	stats, err := service.ServiceGroupApp.StatsService.GetDashboardStats(startTime, endTime)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(stats, "获取成功", c)
}

// GetOrderStats 获取订单统计数据
// @Tags     Stats
// @Summary  获取订单统计数据
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    startTime query    string  false "开始时间"
// @Param    endTime   query    string  false "结束时间"
// @Param    game      query    string  false "游戏"
// @Param    status    query    string  false "状态"
// @Success  200       {object}  response.Response{data=map[string]interface{}} "获取成功"
// @Router   /playmate/stats/orders [get]
func (a *StatsApi) GetOrderStats(c *gin.Context) {
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	game := c.Query("game")
	status := c.Query("status")

	stats, err := service.ServiceGroupApp.StatsService.GetOrderStats(startTime, endTime, game, status)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(stats, "获取成功", c)
}

// GetUserStats 获取用户统计数据
// @Tags     Stats
// @Summary  获取用户统计数据
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    startTime query    string  false "开始时间"
// @Param    endTime   query    string  false "结束时间"
// @Success  200       {object}  response.Response{data=map[string]interface{}} "获取成功"
// @Router   /playmate/stats/users [get]
func (a *StatsApi) GetUserStats(c *gin.Context) {
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")

	stats, err := service.ServiceGroupApp.StatsService.GetUserStats(startTime, endTime)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(stats, "获取成功", c)
}

// GetExpertStats 获取专家统计数据
// @Tags     Stats
// @Summary  获取专家统计数据
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    startTime query    string  false "开始时间"
// @Param    endTime   query    string  false "结束时间"
// @Param    game      query    string  false "游戏"
// @Success  200       {object}  response.Response{data=map[string]interface{}} "获取成功"
// @Router   /playmate/stats/experts [get]
func (a *StatsApi) GetExpertStats(c *gin.Context) {
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	game := c.Query("game")

	stats, err := service.ServiceGroupApp.StatsService.GetExpertStats(startTime, endTime, game)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(stats, "获取成功", c)
}

// GetRevenueStats 获取收入统计数据
// @Tags     Stats
// @Summary  获取收入统计数据
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    startTime query    string  false "开始时间"
// @Param    endTime   query    string  false "结束时间"
// @Param    game      query    string  false "游戏"
// @Success  200       {object}  response.Response{data=map[string]interface{}} "获取成功"
// @Router   /playmate/stats/revenue [get]
func (a *StatsApi) GetRevenueStats(c *gin.Context) {
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	game := c.Query("game")

	stats, err := service.ServiceGroupApp.StatsService.GetRevenueStats(startTime, endTime, game)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(stats, "获取成功", c)
}

// GetTrendStats 获取趋势统计数据
// @Tags     Stats
// @Summary  获取趋势统计数据
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    type      query    string  true  "统计类型: orders, users, revenue, experts"
// @Param    startTime query    string  false "开始时间"
// @Param    endTime   query    string  false "结束时间"
// @Param    interval  query    string  false "时间间隔: day, week, month"
// @Success  200       {object}  response.Response{data=map[string]interface{}} "获取成功"
// @Router   /playmate/stats/trend [get]
func (a *StatsApi) GetTrendStats(c *gin.Context) {
	statsType := c.Query("type")
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	interval := c.Query("interval")

	stats, err := service.ServiceGroupApp.StatsService.GetTrendStats(statsType, startTime, endTime, interval)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(stats, "获取成功", c)
}
