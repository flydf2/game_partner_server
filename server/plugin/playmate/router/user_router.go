package router

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
)

type UserRouter struct{}

// InitUserRouter 初始化用户路由
func (r *UserRouter) InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	{
		userRouter.GET("/info", api.ApiGroupApp.UserApi.GetUserInfo)
		userRouter.PUT("/profile", api.ApiGroupApp.UserApi.UpdateProfile)
		userRouter.GET("/settings", api.ApiGroupApp.UserApi.GetSettings)
		userRouter.PUT("/settings", api.ApiGroupApp.UserApi.UpdateSettings)
		userRouter.GET("/following", api.ApiGroupApp.UserApi.GetFollowing)
		userRouter.GET("/favorites", api.ApiGroupApp.UserApi.GetFavorites)
		userRouter.GET("/history", api.ApiGroupApp.UserApi.GetBrowseHistory)
		userRouter.GET("/wallet", api.ApiGroupApp.UserApi.GetWallet)
	}

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", api.ApiGroupApp.UserApi.Login)
		authRouter.POST("/register", api.ApiGroupApp.UserApi.Register)
		authRouter.POST("/logout", api.ApiGroupApp.UserApi.Logout)
		authRouter.POST("/refresh", api.ApiGroupApp.UserApi.RefreshToken)
	}
}
