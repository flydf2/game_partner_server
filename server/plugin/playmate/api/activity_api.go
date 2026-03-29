package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/gin-gonic/gin"
)

type ActivityApi struct{}

// GetActivities 获取活动列表
// @Tags     Activity
// @Summary  获取活动列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    status     query    string  false "状态"
// @Param    type       query    string  false "类型"
// @Param    startTime  query    string  false "开始时间"
// @Param    endTime    query    string  false "结束时间"
// @Param    keyword    query    string  false "关键词"
// @Param    page       query    int     false "页码"
// @Param    pageSize   query    int     false "每页数量"
// @Success  200        {object} response.Response{data=[]model.Activity,pagination=map[string]int64} "获取活动列表成功"
// @Router   /playmate/activities [get]
func (a *ActivityApi) GetActivities(c *gin.Context) {
	var search request.ActivitySearch
	if err := c.ShouldBindQuery(&search); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 设置默认值
	if search.Page <= 0 {
		search.Page = 1
	}
	if search.PageSize <= 0 {
		search.PageSize = 10
	}

	activities, total, err := service.ServiceGroupApp.ActivityService.GetActivities(search)
	if err != nil {
		response.FailWithMessage("获取活动列表失败", c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": activities,
		"pagination": gin.H{
			"currentPage": search.Page,
			"totalPages":  (total + int64(search.PageSize) - 1) / int64(search.PageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}
