package service

import (
	"errors"

	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
)

// WithdrawalService 提现服务
type WithdrawalService struct{}

// SubmitWithdrawal 提交提现请求
func (s *WithdrawalService) SubmitWithdrawal(userID uint, req request.SubmitWithdrawalRequest) (model.Withdrawal, error) {
	// 检查用户钱包
	var wallet model.UserWallet
	if err := global.GVA_DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Withdrawal{}, errors.New("钱包不存在")
		}
		return model.Withdrawal{}, err
	}

	// 检查余额
	if wallet.Balance < req.Amount {
		return model.Withdrawal{}, errors.New("余额不足")
	}

	// 计算手续费
	fee := req.Amount * 0.01 // 1%手续费
	actualAmount := req.Amount - fee

	// 创建提现记录
	withdrawal := model.Withdrawal{
		UserID:       userID,
		Amount:       req.Amount,
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
	wallet.Balance -= req.Amount
	wallet.Frozen += req.Amount
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
func (s *WithdrawalService) GetWithdrawalRecords(userID uint) ([]model.Withdrawal, error) {
	var withdrawals []model.Withdrawal
	if err := global.GVA_DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&withdrawals).Error; err != nil {
		return nil, err
	}

	return withdrawals, nil
}