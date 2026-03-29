package api

import (
	"path/filepath"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/upload"
)

// UploadApi 文件上传API
type UploadApi struct{}

// UploadFile 上传文件
// @Tags     Upload
// @Summary  上传文件
// @Security ApiKeyAuth
// @accept   multipart/form-data
// @Produce  application/json
// @Param    file  formData  file    true  "文件"
// @Param    type  formData  string  false "类型 (avatar, image, voice, document)"
// @Success  200   {object}  response.Response{data=map[string]string} "上传成功"
// @Router   /playmate/upload [post]
func (a *UploadApi) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("请选择要上传的文件", c)
		return
	}

	uploadType := c.DefaultPostForm("type", "image")

	if file == nil {
		response.FailWithMessage("文件不能为空", c)
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext == "" {
		response.FailWithMessage("文件格式不正确", c)
		return
	}

	// 使用阿里云 OSS 上传
	oss := upload.NewOss()
	fileURL, _, err := oss.UploadFile(file)
	if err != nil {
		response.FailWithMessage("文件上传失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"url":      fileURL,
		"filename": file.Filename,
		"size":     file.Size,
		"type":     uploadType,
	}, "上传成功", c)
}

// UploadImage 上传图片
// @Tags     Upload
// @Summary  上传图片
// @Security ApiKeyAuth
// @accept   multipart/form-data
// @Produce  application/json
// @Param    file  formData  file  true  "图片文件"
// @Success  200   {object}  response.Response{data=map[string]string} "上传成功"
// @Router   /playmate/upload/image [post]
func (a *UploadApi) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("请选择要上传的图片", c)
		return
	}

	if file == nil {
		response.FailWithMessage("图片不能为空", c)
		return
	}

	ext := filepath.Ext(file.Filename)
	allowedExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true}
	if !allowedExts[ext] {
		response.FailWithMessage("只支持jpg、jpeg、png、gif、webp格式的图片", c)
		return
	}

	// 使用阿里云 OSS 上传
	oss := upload.NewOss()
	fileURL, _, err := oss.UploadFile(file)
	if err != nil {
		response.FailWithMessage("图片上传失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"url":      fileURL,
		"filename": file.Filename,
		"size":     file.Size,
	}, "上传成功", c)
}

// UploadVoice 上传语音
// @Tags     Upload
// @Summary  上传语音
// @Security ApiKeyAuth
// @accept   multipart/form-data
// @Produce  application/json
// @Param    file  formData  file  true  "语音文件"
// @Success  200   {object}  response.Response{data=map[string]string} "上传成功"
// @Router   /playmate/upload/voice [post]
func (a *UploadApi) UploadVoice(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.FailWithMessage("请选择要上传的语音", c)
		return
	}

	if file == nil {
		response.FailWithMessage("语音不能为空", c)
		return
	}

	ext := filepath.Ext(file.Filename)
	allowedExts := map[string]bool{".mp3": true, ".wav": true, ".m4a": true, ".aac": true}
	if !allowedExts[ext] {
		response.FailWithMessage("只支持mp3、wav、m4a、aac格式的语音", c)
		return
	}

	// 使用阿里云 OSS 上传
	oss := upload.NewOss()
	fileURL, _, err := oss.UploadFile(file)
	if err != nil {
		response.FailWithMessage("语音上传失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"url":      fileURL,
		"filename": file.Filename,
		"size":     file.Size,
	}, "上传成功", c)
}
