package service

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
)

// OrderService 订单服务
type OrderService struct{}

// GetOrders 获取订单列表
func (s *OrderService) GetOrders(userID uint, search request.OrderSearch) ([]model.Order, int64, error) {
	var orders []model.Order
	var total int64

	query := global.GVA_DB.Model(&model.Order{}).Where("user_id = ?", userID)

	// 应用搜索条件
	if search.Status != "" && search.Status != "all" {
		query = query.Where("status = ?", search.Status)
	}
	if search.Game != "" {
		query = query.Where("game = ?", search.Game)
	}
	if search.StartTime != "" {
		query = query.Where("created_at >= ?", search.StartTime)
	}
	if search.EndTime != "" {
		query = query.Where("created_at <= ?", search.EndTime)
	}
	if search.MinAmount > 0 {
		query = query.Where("amount >= ?", search.MinAmount)
	}
	if search.MaxAmount > 0 {
		query = query.Where("amount <= ?", search.MaxAmount)
	}
	if search.Quantity > 0 {
		query = query.Where("quantity = ?", search.Quantity)
	}
	if search.Keyword != "" {
		query = query.Where("game LIKE ? OR skill LIKE ? OR order_number LIKE ?", "%"+search.Keyword+"%", "%"+search.Keyword+"%", "%"+search.Keyword+"%")
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页
	offset := (search.Page - 1) * search.PageSize

	// 执行查询
	if err := query.Offset(offset).Limit(search.PageSize).Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

// GetOrderDetail 获取订单详情
func (s *OrderService) GetOrderDetail(orderID uint) (model.Order, error) {
	var order model.Order
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Order{}, response.NewPlaymateError(response.ErrOrderNotFound)
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
			return model.Order{}, response.NewPlaymateError(response.ErrPlaymateNotFound)
		}
		return model.Order{}, err
	}

	// 检查用户余额
	var wallet model.UserWallet
	if err := global.GVA_DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Order{}, response.NewPlaymateError(response.ErrWalletNotFound)
		}
		return model.Order{}, err
	}

	if wallet.Balance < req.Amount {
		return model.Order{}, response.NewPlaymateError(response.ErrInsufficientBalance)
	}

	// 生成订单号
	orderNumber := fmt.Sprintf("GP%s%d", time.Now().Format("20060102150405"), userID)

	// 设置技能ID，默认值1
	skillID := uint(1)
	if req.SkillID > 0 {
		skillID = req.SkillID
	}

	// 创建订单
	order := model.Order{
		UserID:        userID,
		PlaymateID:    req.PlaymateID,
		SkillID:       skillID,
		Game:          req.Game,
		Skill:         req.Skill,
		Status:        "pending",
		ServiceTime:   req.ServiceTime,
		Amount:        req.Amount,
		Quantity:      req.Quantity,
		OrderNumber:   orderNumber,
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
			return nil, response.NewPlaymateError(response.ErrOrderNotFound)
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
	duration := 1                     // 默认1小时
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
		"orderId": order.ID,
		"expert": map[string]interface{}{
			"id":       playmate.ID,
			"nickname": playmate.Nickname,
			"avatar":   playmate.Avatar,
			"game":     playmate.Game,
			"rank":     playmate.Rank,
			"rating":   playmate.Rating,
			"reviews":  1200,
		},
		"pricePerHour": pricePerHour,
		"duration":     duration,
		"quantity":     order.Quantity,
		"serviceFee":   serviceFee,
		"coupon":       coupon,
		"totalAmount":  totalAmount,
	}

	return confirmation, nil
}

// CancelOrder 取消订单
func (s *OrderService) CancelOrder(orderID, userID uint) error {
	// 查找订单
	var order model.Order
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewPlaymateError(response.ErrOrderNotFound)
		}
		return err
	}

	// 检查订单是否属于当前用户
	if order.UserID != userID {
		return response.NewPlaymateError(response.ErrUnauthorizedOperation)
	}

	// 检查订单状态是否可以取消
	if order.Status != "pending" && order.Status != "confirmed" {
		return response.NewPlaymateError(response.ErrOrderStatusNotAllowCancel)
	}

	// 开始事务
	tx := global.GVA_DB.Begin()

	// 更新订单状态
	order.Status = "cancelled"
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 退款逻辑：将金额退回用户钱包
	var wallet model.UserWallet
	if err := tx.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 增加余额
	wallet.Balance += order.Amount
	if err := tx.Save(&wallet).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 创建退款交易记录
	transaction := model.Transaction{
		UserID:      userID,
		Type:        "refund",
		Amount:      order.Amount,
		Description: "订单取消退款",
		Time:        time.Now(),
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// ConfirmOrder 确认订单
func (s *OrderService) ConfirmOrder(orderID, userID uint) error {
	// 查找订单
	var order model.Order
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewPlaymateError(response.ErrOrderNotFound)
		}
		return err
	}

	// 检查订单是否属于当前用户
	if order.UserID != userID {
		return response.NewPlaymateError(response.ErrUnauthorizedOperation)
	}

	// 检查订单状态是否可以确认
	if order.Status != "completed" {
		return response.NewPlaymateError(response.ErrOrderStatusNotAllowConfirm)
	}

	// 更新订单状态
	order.Status = "confirmed"
	if err := global.GVA_DB.Save(&order).Error; err != nil {
		return err
	}

	return nil
}

// AcceptOrder 接受订单
func (s *OrderService) AcceptOrder(orderID, userID uint) error {
	// 查找订单
	var order model.Order
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewPlaymateError(response.ErrOrderNotFound)
		}
		return err
	}

	// 检查订单是否属于当前陪玩
	if order.PlaymateID != userID {
		return response.NewPlaymateError(response.ErrUnauthorizedOperation)
	}

	// 检查订单状态是否可以接受
	if order.Status != "pending" {
		return response.NewPlaymateError(response.ErrOrderStatusNotAllowAccept)
	}

	// 更新订单状态
	order.Status = "accepted"
	if err := global.GVA_DB.Save(&order).Error; err != nil {
		return err
	}

	return nil
}

// RejectOrder 拒绝订单
func (s *OrderService) RejectOrder(orderID, userID uint) error {
	// 查找订单
	var order model.Order
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewPlaymateError(response.ErrOrderNotFound)
		}
		return err
	}

	// 检查订单是否属于当前陪玩
	if order.PlaymateID != userID {
		return response.NewPlaymateError(response.ErrUnauthorizedOperation)
	}

	// 检查订单状态是否可以拒绝
	if order.Status != "pending" {
		return response.NewPlaymateError(response.ErrOrderStatusNotAllowReject)
	}

	// 开始事务
	tx := global.GVA_DB.Begin()

	// 更新订单状态
	order.Status = "rejected"
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 退款逻辑：将金额退回用户钱包
	var wallet model.UserWallet
	if err := tx.Where("user_id = ?", order.UserID).First(&wallet).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 增加余额
	wallet.Balance += order.Amount
	if err := tx.Save(&wallet).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 创建退款交易记录
	transaction := model.Transaction{
		UserID:      order.UserID,
		Type:        "refund",
		Amount:      order.Amount,
		Description: "订单拒绝退款",
		Time:        time.Now(),
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// ShareOrder 分享订单
func (s *OrderService) ShareOrder(orderID, userID uint, platform string) (map[string]interface{}, error) {
	// 查找订单
	var order model.Order
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewPlaymateError(response.ErrOrderNotFound)
		}
		return nil, err
	}

	// 检查订单是否属于当前用户
	if order.UserID != userID {
		return nil, response.NewPlaymateError(response.ErrUnauthorizedOperation)
	}

	// 生成分享码
	shareCode := fmt.Sprintf("GP%s%d", time.Now().Format("20060102150405"), orderID)

	// 构建分享URL
	shareURL := fmt.Sprintf("http://127.0.0.1:8080/share/order?code=%s", shareCode)

	// 创建分享记录
	share := model.Share{
		UserID:        userID,
		OrderID:       &orderID,
		RewardOrderID: nil,
		ShareType:     "order",
		SharePlatform: platform,
		ShareURL:      shareURL,
		ShareCode:     shareCode,
		ClickCount:    0,
		Status:        "active",
	}

	if err := global.GVA_DB.Create(&share).Error; err != nil {
		return nil, err
	}

	// 构建返回数据
	shareData := map[string]interface{}{
		"shareCode":     shareCode,
		"shareURL":      shareURL,
		"orderID":       orderID,
		"orderNumber":   order.OrderNumber,
		"sharePlatform": platform,
		"createdAt":     share.CreatedAt,
	}

	return shareData, nil
}
