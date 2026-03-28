package api

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/gin-gonic/gin"
)

type CommunityApi struct{}

// GetPosts 获取社区帖子列表
// @Tags     Community
// @Summary  获取社区帖子列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    type       query    string  false "类型"
// @Param    status     query    string  false "状态"
// @Param    userId     query    uint    false "用户ID"
// @Param    game       query    string  false "游戏"
// @Param    keyword    query    string  false "关键词"
// @Param    startTime  query    string  false "开始时间"
// @Param    endTime    query    string  false "结束时间"
// @Param    page       query    int     false "页码"
// @Param    pageSize   query    int     false "每页数量"
// @Success  200        {object} response.Response{data=[]model.CommunityPost,pagination=map[string]int64} "获取社区帖子列表成功"
// @Router   /playmate/community/posts [get]
func (a *CommunityApi) GetPosts(c *gin.Context) {
	var search request.CommunitySearch
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

	posts, total, err := service.ServiceGroupApp.CommunityService.GetPosts(search)
	if err != nil {
		response.FailWithMessage("获取社区帖子列表失败", c)
		return
	}
	response.OkWithDetailed(gin.H{
		"data": posts,
		"pagination": gin.H{
			"currentPage": search.Page,
			"totalPages":  (total + int64(search.PageSize) - 1) / int64(search.PageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// GetPostDetail 获取帖子详情
// @Tags     Community
// @Summary  获取帖子详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    postId path     string true "帖子ID"
// @Success  200  {object} response.Response{data=model.CommunityPost} "获取帖子详情成功"
// @Router   /playmate/community/posts/{postId} [get]
func (a *CommunityApi) GetPostDetail(c *gin.Context) {
	postIdStr := c.Param("postId")
	postId, err := strconv.ParseUint(postIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	post, err := service.ServiceGroupApp.CommunityService.GetPostDetail(uint(postId))
	if err != nil {
		response.FailWithMessage("获取帖子详情失败", c)
		return
	}
	response.OkWithData(post, c)
}

// LikePost 点赞帖子
// @Tags     Community
// @Summary  点赞帖子
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    postId path     string true "帖子ID"
// @Success  200  {object} response.Response{msg=string} "点赞成功"
// @Router   /playmate/community/posts/{postId}/like [post]
func (a *CommunityApi) LikePost(c *gin.Context) {
	postIdStr := c.Param("postId")
	postId, err := strconv.ParseUint(postIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := service.ServiceGroupApp.CommunityService.LikePost(uint(postId)); err != nil {
		response.FailWithMessage("点赞帖子失败", c)
		return
	}
	response.OkWithMessage("点赞成功", c)
}

// CommentPost 评论帖子
// @Tags     Community
// @Summary  评论帖子
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    postId path     string true "帖子ID"
// @Param    data  body     map[string]string true "评论内容"
// @Success  200  {object} response.Response{data=model.Comment,msg=string} "评论成功"
// @Router   /playmate/community/posts/{postId}/comments [post]
func (a *CommunityApi) CommentPost(c *gin.Context) {
	postId := c.Param("postId")
	var req struct {
		Content string `json:"content" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	postID, err := strconv.ParseUint(postId, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	userID := uint(1)
	comment, err := service.ServiceGroupApp.CommunityService.CommentPost(userID, uint(postID), req.Content)
	if err != nil {
		response.FailWithMessage("评论帖子失败", c)
		return
	}
	response.OkWithData(comment, c)
}

// CreatePost 创建帖子
// @Tags     Community
// @Summary  创建帖子
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body     request.CreatePostRequest true "帖子内容"
// @Success  200  {object} response.Response{data=model.CommunityPost} "创建成功"
// @Router   /playmate/community/posts [post]
func (a *CommunityApi) CreatePost(c *gin.Context) {
	var req request.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	userID := uint(1)

	post, err := service.ServiceGroupApp.CommunityService.CreatePost(userID, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(post, "创建成功", c)
}

// GetTopicDetail 获取话题详情
// @Tags     Community
// @Summary  获取话题详情
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    topicId path     string true "话题ID"
// @Success  200  {object} response.Response{data=map[string]interface{}} "获取成功"
// @Router   /playmate/community/topics/{topicId} [get]
func (a *CommunityApi) GetTopicDetail(c *gin.Context) {
	topicIdStr := c.Param("topicId")
	topicId, err := strconv.ParseUint(topicIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	topic, err := service.ServiceGroupApp.CommunityService.GetTopicDetail(uint(topicId))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(topic, "获取成功", c)
}

// GetTopicPosts 获取话题帖子列表
// @Tags     Community
// @Summary  获取话题帖子列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    topicId path     string true "话题ID"
// @Param    page     query    int     false "页码"
// @Param    pageSize query    int     false "每页数量"
// @Success  200  {object} response.Response{data=[]model.CommunityPost,pagination=map[string]int64} "获取成功"
// @Router   /playmate/community/topics/{topicId}/posts [get]
func (a *CommunityApi) GetTopicPosts(c *gin.Context) {
	topicIdStr := c.Param("topicId")
	topicId, err := strconv.ParseUint(topicIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	posts, total, err := service.ServiceGroupApp.CommunityService.GetTopicPosts(uint(topicId), page, pageSize)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": posts,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// FollowTopic 关注话题
// @Tags     Community
// @Summary  关注话题
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    topicId path     string true "话题ID"
// @Success  200  {object} response.Response{msg=string} "关注成功"
// @Router   /playmate/community/topics/{topicId}/follow [post]
func (a *CommunityApi) FollowTopic(c *gin.Context) {
	topicIdStr := c.Param("topicId")
	topicId, err := strconv.ParseUint(topicIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	userID := uint(1)

	if err := service.ServiceGroupApp.CommunityService.FollowTopic(userID, uint(topicId)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("关注成功", c)
}

// UnfollowTopic 取消关注话题
// @Tags     Community
// @Summary  取消关注话题
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    topicId path     string true "话题ID"
// @Success  200  {object} response.Response{msg=string} "取消关注成功"
// @Router   /playmate/community/topics/{topicId}/follow [delete]
func (a *CommunityApi) UnfollowTopic(c *gin.Context) {
	topicIdStr := c.Param("topicId")
	topicId, err := strconv.ParseUint(topicIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	userID := uint(1)

	if err := service.ServiceGroupApp.CommunityService.UnfollowTopic(userID, uint(topicId)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("取消关注成功", c)
}

// GetBids 获取帖子投标列表
// @Tags     Community
// @Summary  获取帖子投标列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    postId path     string true "帖子ID"
// @Success  200  {object} response.Response{data=[]model.CommunityBid} "获取成功"
// @Router   /playmate/community/posts/{postId}/bids [get]
func (a *CommunityApi) GetBids(c *gin.Context) {
	postIdStr := c.Param("postId")
	postId, err := strconv.ParseUint(postIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	bids, err := service.ServiceGroupApp.CommunityService.GetBids(uint(postId))
	if err != nil {
		response.FailWithMessage("获取投标列表失败", c)
		return
	}

	response.OkWithData(bids, c)
}

// CreateBid 创建投标
// @Tags     Community
// @Summary  创建投标
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body     request.CreateBidRequest true "投标信息"
// @Success  200  {object} response.Response{data=model.CommunityBid} "创建成功"
// @Router   /playmate/community/bids [post]
func (a *CommunityApi) CreateBid(c *gin.Context) {
	var req request.CreateBidRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	userID := uint(1)

	bid, err := service.ServiceGroupApp.CommunityService.CreateBid(userID, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(bid, "创建成功", c)
}

// CancelBid 取消投标
// @Tags     Community
// @Summary  取消投标
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    bidId path     string true "投标ID"
// @Success  200  {object} response.Response{msg=string} "取消成功"
// @Router   /playmate/community/bids/{bidId}/cancel [post]
func (a *CommunityApi) CancelBid(c *gin.Context) {
	bidIdStr := c.Param("bidId")
	bidId, err := strconv.ParseUint(bidIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	userID := uint(1)

	if err := service.ServiceGroupApp.CommunityService.CancelBid(userID, uint(bidId)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("取消成功", c)
}

// AcceptBid 接受投标
// @Tags     Community
// @Summary  接受投标
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    bidId path     string true "投标ID"
// @Success  200  {object} response.Response{msg=string} "接受成功"
// @Router   /playmate/community/bids/{bidId}/accept [post]
func (a *CommunityApi) AcceptBid(c *gin.Context) {
	bidIdStr := c.Param("bidId")
	bidId, err := strconv.ParseUint(bidIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	userID := uint(1)

	if err := service.ServiceGroupApp.CommunityService.AcceptBid(userID, uint(bidId)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("接受成功", c)
}

// RejectBid 拒绝投标
// @Tags     Community
// @Summary  拒绝投标
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    bidId path     string true "投标ID"
// @Success  200  {object} response.Response{msg=string} "拒绝成功"
// @Router   /playmate/community/bids/{bidId}/reject [post]
func (a *CommunityApi) RejectBid(c *gin.Context) {
	bidIdStr := c.Param("bidId")
	bidId, err := strconv.ParseUint(bidIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	userID := uint(1)

	if err := service.ServiceGroupApp.CommunityService.RejectBid(userID, uint(bidId)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("拒绝成功", c)
}

// CompleteOrder 完成社区订单
// @Tags     Community
// @Summary  完成社区订单
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    orderId path     string true "订单ID"
// @Success  200  {object} response.Response{msg=string} "完成成功"
// @Router   /playmate/community/orders/{orderId}/complete [post]
func (a *CommunityApi) CompleteOrder(c *gin.Context) {
	orderIdStr := c.Param("orderId")
	orderId, err := strconv.ParseUint(orderIdStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	userID := uint(1)

	if err := service.ServiceGroupApp.CommunityService.CompleteOrder(userID, uint(orderId)); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("订单已完成，金钱已划拨", c)
}
