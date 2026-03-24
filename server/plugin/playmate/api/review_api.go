package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
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
// @Router   /reviews [post]
func (a *ReviewApi) SubmitReview(c *gin.Context) {
	var req request.SubmitReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值
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
// @Param    expertId path     string true "专家ID"
// @Param    page     query    int    false "页码"
// @Param    pageSize query    int    false "每页数量"
// @Success  200  {object} response.Response "获取专家评价列表成功"
// @Router   /experts/{expertId}/reviews [get]
func (a *ReviewApi) GetExpertReviews(c *gin.Context) {
	expertIdStr := c.Param("expertId")
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
// @Param    page     query    int    false "页码"
// @Param    pageSize query    int    false "每页数量"
// @Success  200  {object} response.Response "获取评价列表成功"
// @Router   /reviews [get]
func (a *ReviewApi) GetReviews(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	reviews, total, err := service.ServiceGroupApp.ReviewService.GetReviews(page, pageSize)
	if err != nil {
		response.FailWithMessage("获取评价列表失败", c)
		return
	}
	response.OkWithDetailed(gin.H{
		"data": reviews,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}
