package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
	"gorm.io/gorm"
)

// RewardOrderService 悬赏订单服务
type RewardOrderService struct{}

// GetRewardOrders 获取悬赏订单列表
func (s *RewardOrderService) GetRewardOrders(search request.RewardOrderSearch) ([]model.RewardOrder, int64, error) {
	var orders []model.RewardOrder
	var total int64

	db := global.GVA_DB.Model(&model.RewardOrder{})

	// 过滤条件
	if search.Status != "" {
		db = db.Where("status = ?", search.Status)
	}
	if search.Game != "" {
		db = db.Where("game = ?", search.Game)
	}
	if search.Keyword != "" {
		db = db.Where("content LIKE ?", "%"+search.Keyword+"%")
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (search.Page - 1) * search.PageSize
	if err := db.Offset(offset).Limit(search.PageSize).Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

// GetMyRewardOrders 获取我的悬赏订单列表
func (s *RewardOrderService) GetMyRewardOrders(userID uint, page, pageSize int) ([]model.RewardOrder, int64, error) {
	var orders []model.RewardOrder
	var total int64

	db := global.GVA_DB.Model(&model.RewardOrder{}).Where("user_id = ?", userID)

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := db.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, total, nil
}

// GetRewardOrderDetail 获取悬赏订单详情
func (s *RewardOrderService) GetRewardOrderDetail(orderID uint) (model.RewardOrder, error) {
	var order model.RewardOrder
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, response.NewPlaymateError(response.ErrOrderNotFound)
		}
		return order, err
	}

	return order, nil
}

// CreateRewardOrder 创建悬赏订单
func (s *RewardOrderService) CreateRewardOrder(userID uint, req request.CreateRewardOrderRequest) (model.RewardOrder, error) {
	// 将标签和要求转换为逗号分隔的字符串
	tags, _ := json.Marshal(req.Tags)
	requirements, _ := json.Marshal(req.Requirements)

	order := model.RewardOrder{
		UserID:        userID,
		Game:          req.Game,
		Content:       req.Content,
		Reward:        req.Reward,
		PaymentMethod: req.PaymentMethod,
		Status:        "available",
		TimeLeft:      req.TimeLeft,
		GameRank:      req.GameRank,
		StartTime:     req.StartTime,
		Duration:      req.Duration,
		Location:      req.Location,
		Tags:          string(tags),
		Requirements:  string(requirements),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	if err := global.GVA_DB.Create(&order).Error; err != nil {
		return order, err
	}

	return order, nil
}

// UpdateRewardOrder 更新悬赏订单
func (s *RewardOrderService) UpdateRewardOrder(orderID uint, req request.UpdateRewardOrderRequest) (model.RewardOrder, error) {
	var order model.RewardOrder
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return order, response.NewPlaymateError(response.ErrOrderNotFound)
		}
		return order, err
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
	if len(req.Tags) > 0 {
		tags, _ := json.Marshal(req.Tags)
		order.Tags = string(tags)
	}
	if len(req.Requirements) > 0 {
		requirements, _ := json.Marshal(req.Requirements)
		order.Requirements = string(requirements)
	}

	order.UpdatedAt = time.Now()

	if err := global.GVA_DB.Save(&order).Error; err != nil {
		return order, err
	}

	return order, nil
}

// DeleteRewardOrder 删除悬赏订单
func (s *RewardOrderService) DeleteRewardOrder(orderID uint) error {
	if err := global.GVA_DB.Delete(&model.RewardOrder{}, orderID).Error; err != nil {
		return err
	}
	return nil
}

// GrabRewardOrder 抢单
func (s *RewardOrderService) GrabRewardOrder(orderID, userID uint, req request.GrabRewardOrderRequest) error {
	// 检查订单是否存在
	var order model.RewardOrder
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewPlaymateError(response.ErrOrderNotFound)
		}
		return err
	}

	// 检查订单状态
	if order.Status != "available" {
		return response.NewPlaymateError(response.ErrOrderNot抢able)
	}

	// 检查是否已经抢过单
	var existingApplicant model.RewardOrderApplicant
	result := global.GVA_DB.Where("order_id = ? AND user_id = ?", orderID, userID).First(&existingApplicant)
	if result.Error == nil {
		return response.NewPlaymateError(response.ErrAlready抢edOrder)
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	// 创建抢单申请
	applicant := model.RewardOrderApplicant{
		OrderID:         orderID,
		UserID:          userID,
		Recommendation:  req.Recommendation,
		VoiceUrl:        req.VoiceUrl,
		RecordUrl:       req.RecordUrl,
		Status:          "pending",
		AppliedAt:       time.Now(),
	}

	if err := global.GVA_DB.Create(&applicant).Error; err != nil {
		return err
	}

	return nil
}

// GetApplicants 获取抢单者列表
func (s *RewardOrderService) GetApplicants(orderID uint) ([]map[string]interface{}, error) {
	// 检查订单是否存在
	var order model.RewardOrder
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, response.NewPlaymateError(response.ErrOrderNotFound)
		}
		return nil, err
	}

	// 获取抢单申请
	var applicants []model.RewardOrderApplicant
	if err := global.GVA_DB.Where("order_id = ?", orderID).Find(&applicants).Error; err != nil {
		return nil, err
	}

	// 构建返回数据
	result := make([]map[string]interface{}, 0)
	for _, applicant := range applicants {
		// 这里应该查询用户信息，暂时模拟数据
		userInfo := map[string]interface{}{
			"id":         applicant.ID,
			"userId":     applicant.UserID,
			"name":       fmt.Sprintf("用户%d", applicant.UserID),
			"avatar":     "https://randomuser.me/api/portraits/men/32.jpg",
			"level":      24,
			"rating":     4.9,
			"specialty":  "擅长各种游戏",
			"orderCount": 100,
			"badges": []map[string]string{
				{"type": "verified", "text": "实名认证"},
				{"type": "deposit", "text": "保证金已缴"},
			},
			"recommendation": applicant.Recommendation,
			"appliedAt":      applicant.AppliedAt,
			"status":         applicant.Status,
		}
		result = append(result, userInfo)
	}

	return result, nil
}

// SelectApplicant 选择抢单者
func (s *RewardOrderService) SelectApplicant(orderID, applicantID uint) error {
	// 检查订单是否存在
	var order model.RewardOrder
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewPlaymateError(response.ErrOrderNotFound)
		}
		return err
	}

	// 检查抢单申请是否存在
	var applicant model.RewardOrderApplicant
	if err := global.GVA_DB.First(&applicant, applicantID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewPlaymateError(response.Err抢单ApplicationNotFound)
		}
		return err
	}

	// 检查抢单申请是否属于该订单
	if applicant.OrderID != orderID {
		return response.NewPlaymateError(response.Err抢单ApplicationNotMatch)
	}

	// 开始事务
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新订单状态为进行中
	order.Status = "ongoing"
	order.UpdatedAt = time.Now()
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 更新抢单申请状态为已通过
	applicant.Status = "approved"
	if err := tx.Save(&applicant).Error; err != nil {
		tx.Rollback()
		return err
	}

	// 其他抢单申请状态更新为已拒绝
	if err := tx.Model(&model.RewardOrderApplicant{}).Where("order_id = ? AND id != ?", orderID, applicantID).Update("status", "rejected").Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// PayRewardOrder 支付订单
func (s *RewardOrderService) PayRewardOrder(orderID uint, req request.PayRewardOrderRequest) (string, error) {
	// 检查订单是否存在
	var order model.RewardOrder
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", response.NewPlaymateError(response.ErrOrderNotFound)
		}
		return "", err
	}

	// 检查订单状态
	if order.Status != "available" && order.Status != "ongoing" {
		return "", response.NewPlaymateError(response.ErrOrderStatusNotAllowPay)
	}

	// 检查金额是否匹配
	if req.Amount != order.Reward {
		return "", response.NewPlaymateError(response.ErrPayAmountMismatch)
	}

	// 创建支付记录
	payment := model.RewardOrderPayment{
		OrderID:         orderID,
		Amount:          req.Amount,
		PaymentMethod:   req.PaymentMethod,
		TransactionID:   req.TransactionID,
		PaymentStatus:   "success",
		PaidAt:          time.Now(),
	}

	if err := global.GVA_DB.Create(&payment).Error; err != nil {
		return "", err
	}

	return req.TransactionID, nil
}

// ConfirmService 确认服务
func (s *RewardOrderService) ConfirmService(orderID uint, req request.ConfirmServiceRequest) (float64, error) {
	// 检查订单是否存在
	var order model.RewardOrder
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, response.NewPlaymateError(response.ErrOrderNotFound)
		}
		return 0, err
	}

	// 检查订单状态
	if order.Status != "ongoing" {
		return 0, response.NewPlaymateError(response.ErrOrderStatusNotAllowConfirm)
	}

	// 开始事务
	tx := global.GVA_DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新订单状态为已完成
	order.Status = "completed"
	order.UpdatedAt = time.Now()
	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	// 转换图片数组为字符串
	images, _ := json.Marshal(req.Images)

	// 创建评价记录
	review := model.RewardOrderReview{
		OrderID:    orderID,
		Rating:     req.Rating,
		Review:     req.Review,
		Images:     string(images),
		ReviewedAt: time.Now(),
	}

	if err := tx.Create(&review).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	// 计算结算金额（这里简化处理，实际应该有更复杂的逻辑）
	settlementAmount := order.Reward * 0.9 // 假设平台抽成10%

	return settlementAmount, nil
}

// PublishRewardOrder 发布订单
func (s *RewardOrderService) PublishRewardOrder(orderID uint) error {
	// 检查订单是否存在
	var order model.RewardOrder
	if err := global.GVA_DB.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.NewPlaymateError(response.ErrOrderNotFound)
		}
		return err
	}

	// 更新订单状态为可抢单
	order.Status = "available"
	order.UpdatedAt = time.Now()

	if err := global.GVA_DB.Save(&order).Error; err != nil {
		return err
	}

	return nil
}

// ShareRewardOrder 分享悬赏订单
func (s *RewardOrderService) ShareRewardOrder(orderID, userID uint, platform string) (map[string]interface{}, error) {
	// 查找订单
	var order model.RewardOrder
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
	shareCode := fmt.Sprintf("GR%s%d", time.Now().Format("20060102150405"), orderID)

	// 构建分享URL
	shareURL := fmt.Sprintf("http://127.0.0.1:8080/share/reward?code=%s", shareCode)

	// 创建分享记录
	share := model.Share{
		UserID:        userID,
		OrderID:       nil,
		RewardOrderID: &orderID,
		ShareType:     "reward",
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
		"reward":        order.Reward,
		"sharePlatform": platform,
		"createdAt":     share.CreatedAt,
	}

	return shareData, nil
}

