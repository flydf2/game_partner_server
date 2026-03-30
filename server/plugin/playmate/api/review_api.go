package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/gin-gonic/gin"
)

type ReviewApi struct{}

// SubmitReview 提交评价
// @Tags     Review
// @Summary  提交评价
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data body     request.SubmitReviewRequest true "评价信息"
// @Success  200  {object} response.Response{msg=string} "提交评价成功"
// @Router   /playmate/reviews [post]
func (a *ReviewApi) SubmitReview(c *gin.Context) {
	var req request.SubmitReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	// 从上下文获取用户ID
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.FailWithMessage("未获取到用户ID", c)
		return
	}
	if _, err := service.ServiceGroupApp.ReviewService.SubmitReview(userID, req); err != nil {
		response.FailWithMessage("提交评价失败", c)
		return
	}
	response.OkWithMessage("提交评价成功", c)
}

// GetExpertReviews 获取专家评价列表
// @Tags     Review
// @Summary  获取专家评价列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id path     string true "专家ID"
// @Param    page     query    int    false "页码"
// @Param    pageSize query    int    false "每页数量"
// @Success  200  {object} response.Response "获取专家评价列表成功"
// @Router   /playmate/experts/{id}/reviews [get]
func (a *ReviewApi) GetExpertReviews(c *gin.Context) {
	expertIdStr := c.Param("id")
	expertId, err := strconv.ParseUint(expertIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	reviews, total, err := service.ServiceGroupApp.PlaymateService.GetExpertReviews(uint(expertId), page, pageSize)
	if err != nil {
		response.FailWithMessage("获取专家评价列表失败", c)
		return
	}
	response.OkWithDetailed(gin.H{
		"reviews": reviews,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// GetReviews 获取评价列表
// @Tags     Review
// @Summary  获取评价列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    playmateId query    uint    false "陪玩ID"
// @Param    minRating  query    int     false "最低评分"
// @Param    maxRating  query    int     false "最高评分"
// @Param    game       query    string  false "游戏"
// @Param    keyword    query    string  false "关键词"
// @Param    startTime  query    string  false "开始时间"
// @Param    endTime    query    string  false "结束时间"
// @Param    page       query    int     false "页码"
// @Param    pageSize   query    int     false "每页数量"
// @Success  200        {object} response.Response{data=[]model.Review,pagination=map[string]int64} "获取评价列表成功"
// @Router   /playmate/reviews [get]
func (a *ReviewApi) GetReviews(c *gin.Context) {
	var search request.ReviewSearch
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

	reviews, total, err := service.ServiceGroupApp.ReviewService.GetReviews(search)
	if err != nil {
		response.FailWithMessage("获取评价列表失败", c)
		return
	}
	response.OkWithDetailed(gin.H{
		"data": reviews,
		"pagination": gin.H{
			"currentPage": search.Page,
			"totalPages":  (total + int64(search.PageSize) - 1) / int64(search.PageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}
