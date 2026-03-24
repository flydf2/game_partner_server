package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
)

// RewardOrderService 奖励订单服务
type RewardOrderService struct{}

// GetRewardOrders 获取奖励订单列表
func (s *RewardOrderService) GetRewardOrders(params map[string]string) ([]model.RewardOrder, int64, error) {
	var orders []model.RewardOrder
	var total int64

	query := global.GVA_DB.Model(&model.RewardOrder{})

	// 应用过滤条件
	if game, ok := params["game"]; ok && game != "" {
		query = query.Where("game = ?", game)
	}

	if status, ok := params["status"]; ok && status != "" {
		query = query.Where("status = ?", status)
	}

	if paymentMethod, ok := params["paymentMethod"]; ok && paymentMethod != "" {
		query = query.Where("payment_method = ?", paymentMethod)
	}

	if keyword, ok := params["keyword"]; ok && keyword != "" {
		query = query.Where("content LIKE ? OR tags LIKE ? OR game LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 计算总数
	query.Count(&total)

	// 应用分页
	page := 1
	pageSize := 20

	if _, ok := params["page"]; ok {
		// 这里应该将字符串转换为整数
	}

	if _, ok := params["pageSize"]; ok {
		// 这里应该将字符串转换为整数
	}

	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	// 执行查询
	if err := query.Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

// GrabRewardOrder 抢奖励订单
func (s *RewardOrderService) GrabRewardOrder(orderID uint) error {
	var order model.RewardOrder
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("订单不存在")
		}
		return err
	}

	if order.Status != "available" {
		return errors.New("订单已被抢或已完成")
	}

	order.Status = "grabbed"
	if err := global.GVA_DB.Save(&order).Error; err != nil {
		return err
	}

	return nil
}