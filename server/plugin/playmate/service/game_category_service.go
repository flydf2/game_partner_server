package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
)

// GameCategoryService 游戏分类服务
type GameCategoryService struct{}

// GetCategories 获取游戏分类列表
func (s *GameCategoryService) GetCategories(search request.GameCategorySearch) ([]model.GameCategory, int64, error) {
	var categories []model.GameCategory
	var total int64

	db := global.GVA_DB.Model(&model.GameCategory{})

	// 应用搜索条件
	if search.Status != "" {
		db = db.Where("status = ?", search.Status)
	}
	if search.ParentID > 0 {
		db = db.Where("parent_id = ?", search.ParentID)
	}
	if search.Keyword != "" {
		db = db.Where("name LIKE ? OR description LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%")
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页
	offset := (search.Page - 1) * search.PageSize

	// 执行查询
	if err := db.Offset(offset).Limit(search.PageSize).Order("sort_order ASC, id DESC").Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

// GetGamesByCategory 根据分类获取游戏列表
func (s *GameCategoryService) GetGamesByCategory(category string) ([]model.Game, error) {
	var games []model.Game
	if err := global.GVA_DB.Where("category = ?", category).Find(&games).Error; err != nil {
		return nil, err
	}

	return games, nil
}

// GetCategoryById 根据ID获取游戏分类
func (s *GameCategoryService) GetCategoryById(id uint) (model.GameCategory, error) {
	var category model.GameCategory
	if err := global.GVA_DB.First(&category, id).Error; err != nil {
		return category, err
	}
	return category, nil
}

// CreateCategory 创建游戏分类
func (s *GameCategoryService) CreateCategory(category model.GameCategory) (model.GameCategory, error) {
	if err := global.GVA_DB.Create(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

// UpdateCategory 更新游戏分类
func (s *GameCategoryService) UpdateCategory(category model.GameCategory) (model.GameCategory, error) {
	if err := global.GVA_DB.Save(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}

// DeleteCategory 删除游戏分类
func (s *GameCategoryService) DeleteCategory(id uint) error {
	if err := global.GVA_DB.Delete(&model.GameCategory{}, id).Error; err != nil {
		return err
	}
	return nil
}
