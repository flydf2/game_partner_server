package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
)

// GameService 游戏服务
type GameService struct{}

// GetGames 获取游戏列表
func (s *GameService) GetGames() ([]model.Game, error) {
	var games []model.Game
	if err := global.GVA_DB.Find(&games).Error; err != nil {
		return nil, err
	}

	return games, nil
}

// GetGameByID 根据ID获取游戏
func (s *GameService) GetGameByID(id uint) (model.Game, error) {
	var game model.Game
	if err := global.GVA_DB.First(&game, id).Error; err != nil {
		return model.Game{}, err
	}

	return game, nil
}
