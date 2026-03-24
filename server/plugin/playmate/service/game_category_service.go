package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
)

// GameCategoryService 游戏分类服务
type GameCategoryService struct{}

// GetCategories 获取游戏分类列表
func (s *GameCategoryService) GetCategories() ([]model.GameCategory, error) {
	var categories []model.GameCategory
	if err := global.GVA_DB.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

// GetGamesByCategory 根据分类获取游戏列表
func (s *GameCategoryService) GetGamesByCategory(category string) ([]model.Game, error) {
	var games []model.Game
	if err := global.GVA_DB.Where("category = ?", category).Find(&games).Error; err != nil {
		return nil, err
	}

	return games, nil
}