package service

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// UploadService 上传服务
type UploadService struct{}

// UploadFile 上传文件
func (s *UploadService) UploadFile(file *multipart.FileHeader, fileType string) (string, error) {
	// 确保上传目录存在
	uploadDir := "./uploads/" + fileType
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", err
	}

	// 生成唯一文件名
	extension := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), extension)
	filepath := filepath.Join(uploadDir, filename)

	// 保存文件
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// 这里应该实现文件拷贝逻辑

	// 返回文件URL
	fileURL := fmt.Sprintf("/uploads/%s/%s", fileType, filename)
	return fileURL, nil
}
