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
	comment, err := service.ServiceGroupApp.CommunityService.CommentPost(uint(postID), req.Content)
	if err != nil {
		response.FailWithMessage("评论帖子失败", c)
		return
	}
	response.OkWithData(comment, c)
}
