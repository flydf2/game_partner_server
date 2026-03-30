package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

// InitUserRouter 初始化用户路由
func (r *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	// 需要认证的用户路由
	userRouter := router.Group("/user")
	userRouter.Use(middleware.CombinedAuthMiddleware())
	{
		userRouter.GET("/info", api.ApiGroupApp.UserApi.GetUserInfo)
		userRouter.PUT("/profile", api.ApiGroupApp.UserApi.UpdateProfile)
		userRouter.GET("/settings", api.ApiGroupApp.UserApi.GetSettings)
		userRouter.PUT("/settings", api.ApiGroupApp.UserApi.UpdateSettings)
		userRouter.GET("/following", api.ApiGroupApp.UserApi.GetFollowing)
		userRouter.POST("/following/:userId", api.ApiGroupApp.UserApi.FollowUser)
		userRouter.DELETE("/following/:userId", api.ApiGroupApp.UserApi.UnfollowUser)
		userRouter.GET("/favorites", api.ApiGroupApp.UserApi.GetFavorites)
		userRouter.DELETE("/favorites/:favoriteId", api.ApiGroupApp.UserApi.RemoveFavorite)
		userRouter.GET("/history", api.ApiGroupApp.UserApi.GetBrowseHistory)
		userRouter.DELETE("/history", api.ApiGroupApp.UserApi.ClearHistory)
		userRouter.GET("/wallet", api.ApiGroupApp.UserApi.GetWallet)
		userRouter.POST("/recharge", api.ApiGroupApp.UserApi.Recharge)
		userRouter.GET("/transactions", api.ApiGroupApp.UserApi.GetTransactionList)
	}

	// 获取用户列表（需要认证）
	router.GET("/users", middleware.CombinedAuthMiddleware(), api.ApiGroupApp.UserApi.GetUsers)

	authRouter := router.Group("/auth")
	{
		// 不需要认证的路由
		authRouter.POST("/login", api.ApiGroupApp.UserApi.Login)
		authRouter.POST("/register", api.ApiGroupApp.UserApi.Register)
		authRouter.POST("/logout", api.ApiGroupApp.UserApi.Logout)
		// 需要认证的路由
		authRouter.POST("/refresh", middleware.CombinedAuthMiddleware(), api.ApiGroupApp.UserApi.RefreshToken)
	}
}
