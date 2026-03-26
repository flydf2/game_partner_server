package service

import (
	"errors"
	"strconv"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
)

// WithdrawalService 提现服务
type WithdrawalService struct{}

// SubmitWithdrawal 提交提现请求
func (s *WithdrawalService) SubmitWithdrawal(userID uint, req request.SubmitWithdrawalRequest) (model.Withdrawal, error) {
	// 检查用户钱包，选择余额最大的钱包
	var wallet model.UserWallet
	if err := global.GVA_DB.Where("user_id = ?", userID).Order("balance DESC").First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Withdrawal{}, errors.New("钱包不存在")
		}
		return model.Withdrawal{}, err
	}

	// 转换金额字符串为float64
	amount, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		return model.Withdrawal{}, errors.New("金额格式错误")
	}

	// 检查余额
	if wallet.Balance < amount {
		return model.Withdrawal{}, errors.New("余额不足")
	}

	// 计算手续费
	fee := amount * 0.01 // 1%手续费
	actualAmount := amount - fee

	// 创建提现记录
	withdrawal := model.Withdrawal{
		UserID:       userID,
		Amount:       amount,
		Fee:          fee,
		ActualAmount: actualAmount,
		Method:       req.Method,
		Status:       "pending",
	}

	// 开始事务
	tx := global.GVA_DB.Begin()

	// 创建提现记录
	if err := tx.Create(&withdrawal).Error; err != nil {
		tx.Rollback()
		return model.Withdrawal{}, err
	}

	// 冻结金额
	wallet.Balance -= amount
	wallet.Frozen += amount
	if err := tx.Save(&wallet).Error; err != nil {
		tx.Rollback()
		return model.Withdrawal{}, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return model.Withdrawal{}, err
	}

	return withdrawal, nil
}

// GetWithdrawalRecords 获取提现记录
func (s *WithdrawalService) GetWithdrawalRecords(userID uint, search request.WithdrawalSearch) ([]model.Withdrawal, int64, error) {
	var withdrawals []model.Withdrawal
	var total int64

	// 构建查询
	query := global.GVA_DB.Model(&model.Withdrawal{}).Where("user_id = ?", userID)

	// 应用搜索条件
	if search.Status != "" {
		query = query.Where("status = ?", search.Status)
	}
	if search.Method != "" {
		query = query.Where("method = ?", search.Method)
	}
	if search.MinAmount > 0 {
		query = query.Where("amount >= ?", search.MinAmount)
	}
	if search.MaxAmount > 0 {
		query = query.Where("amount <= ?", search.MaxAmount)
	}
	if search.StartTime != "" {
		query = query.Where("created_at >= ?", search.StartTime)
	}
	if search.EndTime != "" {
		query = query.Where("created_at <= ?", search.EndTime)
	}

	// 计算总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页
	offset := (search.Page - 1) * search.PageSize

	// 执行查询
	if err := query.Offset(offset).Limit(search.PageSize).Order("created_at DESC").Find(&withdrawals).Error; err != nil {
		return nil, 0, err
	}

	return withdrawals, total, nil
}