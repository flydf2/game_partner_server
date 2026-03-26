package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"go.uber.org/zap"
)

// CategoryService 分类服务
type CategoryService struct{}

// GetCategories 获取分类列表
func (s *CategoryService) GetCategories(search request.CategorySearch) ([]model.Category, int64, error) {
	var categories []model.Category
	var total int64

	db := global.GVA_DB.Model(&model.Category{})

	// 应用搜索条件 - 只使用数据库中已有的字段
	if search.Keyword != "" {
		db = db.Where("name LIKE ?", "%"+search.Keyword+"%")
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		global.GVA_LOG.Error("获取分类总数失败", zap.Error(err))
		return nil, 0, err
	}

	// 分页
	offset := (search.Page - 1) * search.PageSize

	// 执行查询 - 只按 id 排序
	if err := db.Offset(offset).Limit(search.PageSize).Order("id DESC").Find(&categories).Error; err != nil {
		global.GVA_LOG.Error("获取分类列表失败", zap.Error(err))
		return nil, 0, err
	}

	return categories, total, nil
}