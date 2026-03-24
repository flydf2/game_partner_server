package service

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
)

// OrderService 订单服务
type OrderService struct{}

// GetOrders 获取订单列表
func (s *OrderService) GetOrders(userID uint, status string) ([]model.Order, error) {
	var orders []model.Order
	query := global.GVA_DB.Where("user_id = ?", userID)

	if status != "" && status != "all" {
		query = query.Where("status = ?", status)
	}

	if err := query.Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

// GetOrderDetail 获取订单详情
func (s *OrderService) GetOrderDetail(orderID uint) (model.Order, error) {
	var order model.Order
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Order{}, errors.New("订单不存在")
		}
		return model.Order{}, err
	}

	return order, nil
}

// CreateOrder 创建订单
func (s *OrderService) CreateOrder(userID uint, req request.CreateOrderRequest) (model.Order, error) {
	// 检查陪玩是否存在
	var playmate model.Playmate
	if err := global.GVA_DB.First(&playmate, req.PlaymateID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Order{}, errors.New("陪玩不存在")
		}
		return model.Order{}, err
	}

	// 检查用户余额
	var wallet model.UserWallet
	if err := global.GVA_DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Order{}, errors.New("钱包不存在")
		}
		return model.Order{}, err
	}

	if wallet.Balance < req.Amount {
		return model.Order{}, errors.New("余额不足")
	}

	// 生成订单号
	orderNumber := fmt.Sprintf("GP%s%d", time.Now().Format("20060102150405"), userID)

	// 创建订单
	order := model.Order{
		UserID:       userID,
		PlaymateID:   req.PlaymateID,
		Game:         req.Game,
		Skill:        req.Skill,
		Status:       "pending",
		ServiceTime:  req.ServiceTime,
		Amount:       req.Amount,
		OrderNumber:  orderNumber,
		PaymentMethod: "alipay",
	}

	// 开始事务
	tx := global.GVA_DB.Begin()

	// 创建订单
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return model.Order{}, err
	}

	// 扣减余额
	wallet.Balance -= req.Amount
	wallet.TotalExpense += req.Amount
	if err := tx.Save(&wallet).Error; err != nil {
		tx.Rollback()
		return model.Order{}, err
	}

	// 创建交易记录
	transaction := model.Transaction{
		UserID:      userID,
		Type:        "expense",
		Amount:      req.Amount,
		Description: "订单支付",
		Time:        time.Now(),
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return model.Order{}, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return model.Order{}, err
	}

	return order, nil
}

// GetOrderConfirmation 获取订单确认信息
func (s *OrderService) GetOrderConfirmation(orderID uint) (map[string]interface{}, error) {
	var order model.Order
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("订单不存在")
		}
		return nil, err
	}

	// 获取陪玩信息
	var playmate model.Playmate
	if err := global.GVA_DB.First(&playmate, order.PlaymateID).Error; err != nil {
		return nil, err
	}

	// 计算费用
	pricePerHour := playmate.Price
	duration := 1 // 默认1小时
	serviceFee := pricePerHour * 0.05 // 服务费5%
	totalAmount := pricePerHour*float64(duration) + serviceFee

	// 检查是否有可用优惠券
	coupon := map[string]interface{}{
		"available":   true,
		"discount":    10.00,
		"description": "新用户立减 ¥10.00",
	}

	// 应用优惠券
	if coupon["available"].(bool) {
		totalAmount -= coupon["discount"].(float64)
	}

	// 构建响应
	confirmation := map[string]interface{}{
		"orderId":      order.ID,
		"expert": map[string]interface{}{
			"id":         playmate.ID,
			"nickname":   playmate.Nickname,
			"avatar":     playmate.Avatar,
			"game":       playmate.Game,
			"rank":       playmate.Rank,
			"rating":     playmate.Rating,
			"reviews":    1200,
		},
		"pricePerHour": pricePerHour,
		"duration":     duration,
		"serviceFee":   serviceFee,
		"coupon":       coupon,
		"totalAmount":  totalAmount,
	}

	return confirmation, nil
}