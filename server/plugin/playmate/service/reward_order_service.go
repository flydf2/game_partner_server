package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
)

// RewardOrderService 奖励订单服务
type RewardOrderService struct{}

// GetRewardOrders 获取奖励订单列表
func (s *RewardOrderService) GetRewardOrders(search request.RewardOrderSearch) ([]model.RewardOrder, int64, error) {
	var orders []model.RewardOrder
	var total int64

	query := global.GVA_DB.Model(&model.RewardOrder{})

	// 应用过滤条件
	if search.Game != "" {
		query = query.Where("game = ?", search.Game)
	}

	if search.Status != "" {
		query = query.Where("status = ?", search.Status)
	}

	if search.PaymentMethod != "" {
		query = query.Where("payment_method = ?", search.PaymentMethod)
	}

	if search.Keyword != "" {
		query = query.Where("content LIKE ? OR tags LIKE ? OR game LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%", "%"+search.Keyword+"%")
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 应用分页
	offset := (search.Page - 1) * search.PageSize
	query = query.Offset(offset).Limit(search.PageSize)

	// 执行查询，关联用户信息
	if err := query.Preload("User").Preload("GrabUser").Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

// GetRewardOrderDetail 获取奖励订单详情
func (s *RewardOrderService) GetRewardOrderDetail(orderID uint) (*model.RewardOrder, error) {
	var order model.RewardOrder
	if err := global.GVA_DB.Preload("User").Preload("GrabUser").First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("订单不存在")
		}
		return nil, err
	}
	return &order, nil
}

// CreateRewardOrder 创建奖励订单
func (s *RewardOrderService) CreateRewardOrder(userID uint, req request.CreateRewardOrderRequest) (*model.RewardOrder, error) {
	order := model.RewardOrder{
		UserID:        userID,
		Game:          req.Game,
		Content:       req.Content,
		Reward:        req.Reward,
		PaymentMethod: req.PaymentMethod,
		Status:        "available",
		Tags:          req.Tags,
	}

	if err := global.GVA_DB.Create(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

// UpdateRewardOrder 更新奖励订单
func (s *RewardOrderService) UpdateRewardOrder(orderID uint, req request.UpdateRewardOrderRequest) (*model.RewardOrder, error) {
	var order model.RewardOrder
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("订单不存在")
		}
		return nil, err
	}

	// 更新字段
	if req.Game != "" {
		order.Game = req.Game
	}
	if req.Content != "" {
		order.Content = req.Content
	}
	if req.Reward > 0 {
		order.Reward = req.Reward
	}
	if req.PaymentMethod != "" {
		order.PaymentMethod = req.PaymentMethod
	}
	if req.Status != "" {
		order.Status = req.Status
	}
	if req.Tags != "" {
		order.Tags = req.Tags
	}

	if err := global.GVA_DB.Save(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}

// DeleteRewardOrder 删除奖励订单
func (s *RewardOrderService) DeleteRewardOrder(orderID uint) error {
	var order model.RewardOrder
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("订单不存在")
		}
		return err
	}

	if err := global.GVA_DB.Delete(&order).Error; err != nil {
		return err
	}

	return nil
}

// GrabRewardOrder 抢奖励订单
func (s *RewardOrderService) GrabRewardOrder(orderID uint, userID uint) error {
	// 使用事务和悲观锁防止并发抢单
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var order model.RewardOrder
		// 使用FOR UPDATE锁定记录
		if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&order, orderID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("订单不存在")
			}
			return err
		}

		if order.Status != "available" {
			return errors.New("订单已被抢或已完成")
		}

		// 更新订单状态
		order.Status = "grabbed"
		order.GrabUserID = &userID
		if err := tx.Save(&order).Error; err != nil {
			return err
		}

		return nil
	})
}
