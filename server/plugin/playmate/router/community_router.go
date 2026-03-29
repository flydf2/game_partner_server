package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/gin-gonic/gin"
)

type CommunityRouter struct{}

// InitCommunityRouter 初始化社区路由
func (r *CommunityRouter) InitCommunityRouter(router *gin.RouterGroup) {
	communityRouter := router.Group("/community")
	{
		communityRouter.GET("/posts", api.ApiGroupApp.CommunityApi.GetPosts)
		communityRouter.GET("/my-posts", api.ApiGroupApp.CommunityApi.GetMyPosts)
		communityRouter.POST("/posts", api.ApiGroupApp.CommunityApi.CreatePost)
		communityRouter.GET("/posts/:postId", api.ApiGroupApp.CommunityApi.GetPostDetail)
		communityRouter.DELETE("/posts/:postId", api.ApiGroupApp.CommunityApi.DeletePost)
		communityRouter.POST("/posts/:postId/like", api.ApiGroupApp.CommunityApi.LikePost)
		communityRouter.POST("/posts/:postId/comments", api.ApiGroupApp.CommunityApi.CommentPost)
		communityRouter.GET("/posts/:postId/bids", api.ApiGroupApp.CommunityApi.GetBids)

		// 投标相关路由
		communityRouter.POST("/bids", api.ApiGroupApp.CommunityApi.CreateBid)
		communityRouter.POST("/bids/:bidId/cancel", api.ApiGroupApp.CommunityApi.CancelBid)
		communityRouter.POST("/bids/:bidId/accept", api.ApiGroupApp.CommunityApi.AcceptBid)
		communityRouter.POST("/bids/:bidId/reject", api.ApiGroupApp.CommunityApi.RejectBid)

		// 订单相关路由
		communityRouter.POST("/orders/:orderId/complete", api.ApiGroupApp.CommunityApi.CompleteOrder)

		// 话题相关路由
		communityRouter.GET("/topics/:topicId", api.ApiGroupApp.CommunityApi.GetTopicDetail)
		communityRouter.GET("/topics/:topicId/posts", api.ApiGroupApp.CommunityApi.GetTopicPosts)
		communityRouter.POST("/topics/:topicId/follow", api.ApiGroupApp.CommunityApi.FollowTopic)
		communityRouter.DELETE("/topics/:topicId/follow", api.ApiGroupApp.CommunityApi.UnfollowTopic)
	}
}
