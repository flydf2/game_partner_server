package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/gin-gonic/gin"
)

type TestToolRouter struct{}

// InitTestToolRouter 初始化测试工具路由
func (r *TestToolRouter) InitTestToolRouter(router *gin.RouterGroup) {
	testToolRouter := router.Group("/test-tool")
	{
		// 获取测试Token列表
		testToolRouter.GET("/tokens", api.ApiGroupApp.TestToolApi.GetTestTokens)
		// 获取万能验证码
		testToolRouter.GET("/captcha", api.ApiGroupApp.TestToolApi.GetTestCaptcha)
		// 验证验证码
		testToolRouter.POST("/verify-captcha", api.ApiGroupApp.TestToolApi.VerifyCaptcha)
	}
}
