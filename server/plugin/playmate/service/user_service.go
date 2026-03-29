package service

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
)

// UserService 用户服务
type UserService struct{}

// GetUserInfo 获取用户信息
func (s *UserService) GetUserInfo(userID uint) (model.User, error) {
	var user model.User
	if err := global.GVA_DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, response.NewPlaymateError(response.ErrUserNotFound)
		}
		return model.User{}, err
	}

	// 从钱包中获取最新的余额
	var wallet model.UserWallet
	if err := global.GVA_DB.Where("user_id = ?", userID).First(&wallet).Error; err == nil {
		// 更新用户余额为钱包余额
		user.Balance = wallet.Balance
		// 保存更新后的用户信息
		global.GVA_DB.Save(&user)
	}

	return user, nil
}

// Login 用户登录
func (s *UserService) Login(username, password string) (model.User, string, error) {
	var user model.User
	if err := global.GVA_DB.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, "", response.NewPlaymateError(response.ErrInvalidCredentials)
		}
		return model.User{}, "", err
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return model.User{}, "", response.NewPlaymateError(response.ErrInvalidCredentials)
	}

	// 从钱包中获取最新的余额
	var wallet model.UserWallet
	if err := global.GVA_DB.Where("user_id = ?", user.ID).First(&wallet).Error; err == nil {
		// 更新用户余额为钱包余额
		user.Balance = wallet.Balance
		// 保存更新后的用户信息
		global.GVA_DB.Save(&user)
	}

	// 生成token（这里简化处理，实际应该使用JWT）
	token := "mock_token_" + time.Now().String()

	return user, token, nil
}

// Register 用户注册
func (s *UserService) Register(req request.RegisterRequest) (model.User, string, error) {
	// 检查用户名是否已存在
	var existingUser model.User
	if err := global.GVA_DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		return model.User{}, "", response.NewPlaymateError(response.ErrUserAlreadyExists)
	}

	// 检查手机号是否已存在
	if err := global.GVA_DB.Where("phone = ?", req.Phone).First(&existingUser).Error; err == nil {
		return model.User{}, "", response.NewPlaymateError(response.ErrPhoneAlreadyRegistered)
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, "", err
	}

	// 创建用户
	user := model.User{
		Username:    req.Username,
		Password:    string(hashedPassword),
		Phone:       req.Phone,
		Avatar:      "https://randomuser.me/api/portraits/men/32.jpg",
		Nickname:    req.Username,
		VipLevel:    1,
		Balance:     0,
		CouponCount: 0,
	}

	if err := global.GVA_DB.Create(&user).Error; err != nil {
		return model.User{}, "", err
	}

	// 创建钱包
	wallet := model.UserWallet{
		UserID:       user.ID,
		Balance:      0,
		Frozen:       0,
		TotalIncome:  0,
		TotalExpense: 0,
	}

	if err := global.GVA_DB.Create(&wallet).Error; err != nil {
		return model.User{}, "", err
	}

	// 创建默认设置
	settings := model.UserSettings{
		UserID:        user.ID,
		Notifications: `{"order": true, "system": true, "promotion": false, "message": true}`,
		Privacy:       `{"showOnline": true, "allowMessages": true, "showOrders": false}`,
		Theme:         "light",
		Language:      "zh-CN",
	}

	if err := global.GVA_DB.Create(&settings).Error; err != nil {
		return model.User{}, "", err
	}

	// 生成token
	token := "mock_token_" + time.Now().String()

	return user, token, nil
}

// UpdateProfile 更新个人资料
func (s *UserService) UpdateProfile(userID uint, req request.UpdateProfileRequest) (model.User, error) {
	var user model.User
	if err := global.GVA_DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, response.NewPlaymateError(response.ErrUserNotFound)
		}
		return model.User{}, err
	}

	// 更新字段
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Phone != "" {
		// 检查手机号是否已被其他用户使用
		var existingUser model.User
		if err := global.GVA_DB.Where("phone = ? AND id != ?", req.Phone, userID).First(&existingUser).Error; err == nil {
			return model.User{}, response.NewPlaymateError(response.ErrPhoneAlreadyRegistered)
		}
		user.Phone = req.Phone
	}
	if req.Gender != "" {
		user.Gender = req.Gender
	}
	if req.Birthday != "" {
		user.Birthday = req.Birthday
	}
	if req.Bio != "" {
		user.Bio = req.Bio
	}
	if req.Location != "" {
		user.Location = req.Location
	}

	if err := global.GVA_DB.Save(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}

// GetSettings 获取用户设置
func (s *UserService) GetSettings(userID uint) (model.UserSettings, error) {
	var settings model.UserSettings
	if err := global.GVA_DB.Where("user_id = ?", userID).First(&settings).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建默认设置
			settings = model.UserSettings{
				UserID:        userID,
				Notifications: `{"order": true, "system": true, "promotion": false, "message": true}`,
				Privacy:       `{"showOnline": true, "allowMessages": true, "showOrders": false}`,
				Theme:         "light",
				Language:      "zh-CN",
			}
			if err := global.GVA_DB.Create(&settings).Error; err != nil {
				return model.UserSettings{}, err
			}
			return settings, nil
		}
		return model.UserSettings{}, err
	}
	return settings, nil
}

// UpdateSettings 更新用户设置
func (s *UserService) UpdateSettings(userID uint, req request.UpdateSettingsRequest) (model.UserSettings, error) {
	var settings model.UserSettings
	if err := global.GVA_DB.Where("user_id = ?", userID).First(&settings).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建新设置
			settings = model.UserSettings{
				UserID: userID,
			}
		} else {
			return model.UserSettings{}, err
		}
	}

	// 更新设置
	if req.Notifications != nil {
		// 这里应该将map转换为JSON字符串
		settings.Notifications = `{"order": true, "system": true, "promotion": false, "message": true}`
	}
	if req.Privacy != nil {
		settings.Privacy = `{"showOnline": true, "allowMessages": true, "showOrders": false}`
	}
	if req.Theme != "" {
		settings.Theme = req.Theme
	}
	if req.Language != "" {
		settings.Language = req.Language
	}

	if settings.ID == 0 {
		if err := global.GVA_DB.Create(&settings).Error; err != nil {
			return model.UserSettings{}, err
		}
	} else {
		if err := global.GVA_DB.Save(&settings).Error; err != nil {
			return model.UserSettings{}, err
		}
	}

	return settings, nil
}

// Logout 用户登出
func (s *UserService) Logout(userID uint) error {
	// 这里可以实现token失效逻辑
	// 暂时返回成功
	return nil
}

// GetFollowing 获取关注列表
func (s *UserService) GetFollowing(userID uint, page, pageSize int) ([]model.Playmate, int64, error) {
	var follows []model.UserFollow
	var total int64

	// 获取总数
	if err := global.GVA_DB.Model(&model.UserFollow{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := global.GVA_DB.Where("user_id = ?", userID).Offset(offset).Limit(pageSize).Find(&follows).Error; err != nil {
		return nil, 0, err
	}

	var playmateIDs []uint
	for _, follow := range follows {
		playmateIDs = append(playmateIDs, follow.FollowID)
	}

	var playmates []model.Playmate
	if len(playmateIDs) > 0 {
		if err := global.GVA_DB.Where("id IN ?", playmateIDs).Find(&playmates).Error; err != nil {
			return nil, 0, err
		}
	}

	return playmates, total, nil
}

// GetFavorites 获取收藏列表
func (s *UserService) GetFavorites(userID uint, page, pageSize int) ([]model.Playmate, int64, error) {
	var favorites []model.UserFavorite
	var total int64

	// 获取总数
	if err := global.GVA_DB.Model(&model.UserFavorite{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	if err := global.GVA_DB.Where("user_id = ?", userID).Offset(offset).Limit(pageSize).Find(&favorites).Error; err != nil {
		return nil, 0, err
	}

	var playmateIDs []uint
	for _, favorite := range favorites {
		playmateIDs = append(playmateIDs, favorite.PlaymateID)
	}

	var playmates []model.Playmate
	if len(playmateIDs) > 0 {
		if err := global.GVA_DB.Where("id IN ?", playmateIDs).Find(&playmates).Error; err != nil {
			return nil, 0, err
		}
	}

	return playmates, total, nil
}

// GetBrowseHistory 获取浏览历史
func (s *UserService) GetBrowseHistory(userID uint, page, pageSize int) ([]model.UserBrowseHistory, int64, error) {
	var history []model.UserBrowseHistory
	var total int64

	// 获取总数
	if err := global.GVA_DB.Model(&model.UserBrowseHistory{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	if err := global.GVA_DB.Where("user_id = ?", userID).Offset(offset).Limit(pageSize).Order("viewed_at DESC").Find(&history).Error; err != nil {
		return nil, 0, err
	}

	return history, total, nil
}

// GetWallet 获取钱包信息
func (s *UserService) GetWallet(userID uint) (model.UserWallet, []model.Transaction, error) {
	var wallet model.UserWallet
	if err := global.GVA_DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建默认钱包
			wallet = model.UserWallet{
				UserID:       userID,
				Balance:      0,
				Frozen:       0,
				TotalIncome:  0,
				TotalExpense: 0,
			}
			if err := global.GVA_DB.Create(&wallet).Error; err != nil {
				return model.UserWallet{}, nil, err
			}
		} else {
			return model.UserWallet{}, nil, err
		}
	}

	// 获取交易记录
	var transactions []model.Transaction
	if err := global.GVA_DB.Where("user_id = ?", userID).Order("time DESC").Limit(20).Find(&transactions).Error; err != nil {
		return model.UserWallet{}, nil, err
	}

	return wallet, transactions, nil
}

// RefreshToken 刷新token
func (s *UserService) RefreshToken(userID uint) (string, error) {
	// 生成新token
	token := "mock_token_" + time.Now().String()
	return token, nil
}

// GetUsers 获取用户列表
func (s *UserService) GetUsers(page, pageSize int) ([]model.User, int64, error) {
	var users []model.User
	var total int64

	db := global.GVA_DB.Model(&model.User{})

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	if err := db.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// SendSmsCode 发送短信验证码
func (s *UserService) SendSmsCode(phone string) error {
	// 这里应该调用短信服务发送验证码
	// 暂时返回成功
	return nil
}

// FollowUser 关注用户
func (s *UserService) FollowUser(userID, targetUserID uint) error {
	// 检查是否已经关注
	var follow model.UserFollow
	result := global.GVA_DB.Where("user_id = ? AND follow_id = ?", userID, targetUserID).First(&follow)
	if result.Error == nil {
		return response.NewPlaymateError(response.ErrAlreadyFollowed)
	}

	// 创建关注记录
	follow = model.UserFollow{
		UserID:   userID,
		FollowID: targetUserID,
	}

	if err := global.GVA_DB.Create(&follow).Error; err != nil {
		return err
	}

	return nil
}

// UnfollowUser 取消关注用户
func (s *UserService) UnfollowUser(userID, targetUserID uint) error {
	result := global.GVA_DB.Where("user_id = ? AND follow_id = ?", userID, targetUserID).Delete(&model.UserFollow{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return response.NewPlaymateError(response.ErrNotFollowed)
	}

	return nil
}

// RemoveFavorite 移除收藏
func (s *UserService) RemoveFavorite(userID, favoriteID uint) error {
	result := global.GVA_DB.Where("id = ? AND user_id = ?", favoriteID, userID).Delete(&model.UserFavorite{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return response.NewPlaymateError(response.ErrFavoriteNotFound)
	}

	return nil
}

// ClearHistory 清空浏览历史
func (s *UserService) ClearHistory(userID uint) error {
	if err := global.GVA_DB.Where("user_id = ?", userID).Delete(&model.UserBrowseHistory{}).Error; err != nil {
		return err
	}

	return nil
}

// Recharge 充值
func (s *UserService) Recharge(userID uint, req request.RechargeRequest) (map[string]interface{}, error) {
	// 检查用户钱包
	var wallet model.UserWallet
	if err := global.GVA_DB.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 创建默认钱包
			wallet = model.UserWallet{
				UserID:       userID,
				Balance:      0,
				Frozen:       0,
				TotalIncome:  0,
				TotalExpense: 0,
			}
			if err := global.GVA_DB.Create(&wallet).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// 开始事务
	tx := global.GVA_DB.Begin()

	// 更新钱包余额
	wallet.Balance += req.Amount
	wallet.TotalIncome += req.Amount
	if err := tx.Save(&wallet).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 创建交易记录
	transaction := model.Transaction{
		UserID:      userID,
		Type:        "income",
		Amount:      req.Amount,
		Description: "充值",
		Time:        time.Now(),
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"transactionId": transaction.ID,
		"amountAdded":   req.Amount,
		"balance":       wallet.Balance,
		"method":        req.Method,
	}, nil
}
