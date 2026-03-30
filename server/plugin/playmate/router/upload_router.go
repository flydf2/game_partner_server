package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/gin-gonic/gin"
)

type UploadRouter struct{}

// InitUploadRouter 初始化文件上传路由
func (r *UploadRouter) InitUploadRouter(router *gin.RouterGroup) {
	uploadRouter := router.Group("/upload")
	{
		uploadRouter.POST("", api.ApiGroupApp.UploadApi.UploadFile)
		uploadRouter.POST("/image", api.ApiGroupApp.UploadApi.UploadImage)
		uploadRouter.POST("/voice", api.ApiGroupApp.UploadApi.UploadVoice)
	}
}
