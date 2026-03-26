package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
)

// GameService 游戏服务
type GameService struct{}

// GetGames 获取游戏列表
func (s *GameService) GetGames(search request.GameSearch) ([]model.Game, int64, error) {
	var games []model.Game
	var total int64

	db := global.GVA_DB.Model(&model.Game{})

	// 应用搜索条件
	if len(search.CategoryIDs) > 0 {
		db = db.Where("category_ids @> ?", search.CategoryIDs)
	}
	if search.Status != "" {
		db = db.Where("status = ?", search.Status)
	}
	if search.Platform != "" {
		db = db.Where("platform = ?", search.Platform)
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
	if err := db.Offset(offset).Limit(search.PageSize).Order("created_at DESC").Find(&games).Error; err != nil {
		return nil, 0, err
	}

	return games, total, nil
}

// GetGameByID 根据ID获取游戏
func (s *GameService) GetGameByID(id uint) (model.Game, error) {
	var game model.Game
	if err := global.GVA_DB.First(&game, id).Error; err != nil {
		return model.Game{}, err
	}

	return game, nil
}
