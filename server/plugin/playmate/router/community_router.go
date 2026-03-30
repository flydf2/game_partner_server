package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/gin-gonic/gin"
)

type CommunityRouter struct{}

// InitCommunityRouter 初始化社区路由
func (r *CommunityRouter) InitCommunityRouter(router *gin.RouterGroup) {
	communityRouter := router.Group("/community")
	{
		// 不需要认证的路由
		communityRouter.GET("/posts", api.ApiGroupApp.CommunityApi.GetPosts)
		communityRouter.GET("/posts/:postId", api.ApiGroupApp.CommunityApi.GetPostDetail)
		communityRouter.GET("/posts/:postId/bids", api.ApiGroupApp.CommunityApi.GetBids)
		communityRouter.GET("/topics/:topicId", api.ApiGroupApp.CommunityApi.GetTopicDetail)
		communityRouter.GET("/topics/:topicId/posts", api.ApiGroupApp.CommunityApi.GetTopicPosts)

		// 需要认证的路由
		authRouter := communityRouter.Group("/")
		authRouter.Use(middleware.CombinedAuthMiddleware())
		{
			authRouter.GET("/my-posts", api.ApiGroupApp.CommunityApi.GetMyPosts)
			authRouter.POST("/posts", api.ApiGroupApp.CommunityApi.CreatePost)
			authRouter.DELETE("/posts/:postId", api.ApiGroupApp.CommunityApi.DeletePost)
			authRouter.POST("/posts/:postId/like", api.ApiGroupApp.CommunityApi.LikePost)
			authRouter.POST("/posts/:postId/comments", api.ApiGroupApp.CommunityApi.CommentPost)

			// 投标相关路由
			authRouter.POST("/bids", api.ApiGroupApp.CommunityApi.CreateBid)
			authRouter.POST("/bids/:bidId/cancel", api.ApiGroupApp.CommunityApi.CancelBid)
			authRouter.POST("/bids/:bidId/accept", api.ApiGroupApp.CommunityApi.AcceptBid)
			authRouter.POST("/bids/:bidId/reject", api.ApiGroupApp.CommunityApi.RejectBid)

			// 订单相关路由
			authRouter.POST("/orders/:orderId/complete", api.ApiGroupApp.CommunityApi.CompleteOrder)

			// 话题相关路由
			authRouter.POST("/topics/:topicId/follow", api.ApiGroupApp.CommunityApi.FollowTopic)
			authRouter.DELETE("/topics/:topicId/follow", api.ApiGroupApp.CommunityApi.UnfollowTopic)
		}
	}
}
