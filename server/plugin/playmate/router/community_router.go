package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
)

type CommunityRouter struct{}

// InitCommunityRouter 初始化社区路由
func (r *CommunityRouter) InitCommunityRouter(router *gin.RouterGroup) {
	communityRouter := router.Group("/community")
	{
		communityRouter.GET("/posts", api.ApiGroupApp.CommunityApi.GetPosts)
		communityRouter.GET("/posts/:postId", api.ApiGroupApp.CommunityApi.GetPostDetail)
		communityRouter.POST("/posts/:postId/like", api.ApiGroupApp.CommunityApi.LikePost)
		communityRouter.POST("/posts/:postId/comments", api.ApiGroupApp.CommunityApi.CommentPost)
	}
}
