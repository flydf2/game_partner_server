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
		communityRouter.POST("/posts", api.ApiGroupApp.CommunityApi.CreatePost)
		communityRouter.GET("/posts/:postId", api.ApiGroupApp.CommunityApi.GetPostDetail)
		communityRouter.POST("/posts/:postId/like", api.ApiGroupApp.CommunityApi.LikePost)
		communityRouter.POST("/posts/:postId/comments", api.ApiGroupApp.CommunityApi.CommentPost)

		// 话题相关路由
		communityRouter.GET("/topics/:topicId", api.ApiGroupApp.CommunityApi.GetTopicDetail)
		communityRouter.GET("/topics/:topicId/posts", api.ApiGroupApp.CommunityApi.GetTopicPosts)
		communityRouter.POST("/topics/:topicId/follow", api.ApiGroupApp.CommunityApi.FollowTopic)
		communityRouter.DELETE("/topics/:topicId/follow", api.ApiGroupApp.CommunityApi.UnfollowTopic)
	}
}
