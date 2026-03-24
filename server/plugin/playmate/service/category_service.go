package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
)

// CategoryService 分类服务
type CategoryService struct{}

// GetCategories 获取分类列表
func (s *CategoryService) GetCategories() ([]model.Category, error) {
	var categories []model.Category
	if err := global.GVA_DB.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}