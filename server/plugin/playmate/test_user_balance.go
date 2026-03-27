package main

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// TestUserBalance 测试用户余额从钱包读取
func TestUserBalance() {
	// 初始化数据库连接（使用内存SQLite）
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库连接失败: %v\n", err)
		return
	}

	// 自动迁移表结构
	db.AutoMigrate(&model.User{}, &model.UserWallet{})

	// 保存原始的全局数据库连接
	originalDB := global.GVA_DB
	defer func() {
		global.GVA_DB = originalDB
	}()

	// 设置全局数据库连接为测试数据库
	global.GVA_DB = db

	// 创建测试用户
	user := model.User{
		Username:    "testuser",
		Password:    "password123",
		Phone:       "13800138000",
		Avatar:      "avatar.jpg",
		Nickname:    "Test User",
		VipLevel:    1,
		Balance:     0, // 初始余额为0
		CouponCount: 0,
	}

	if err := db.Create(&user).Error; err != nil {
		fmt.Printf("创建用户失败: %v\n", err)
		return
	}

	// 创建用户钱包，设置不同的余额
	wallet := model.UserWallet{
		UserID:       user.ID,
		Balance:      100.50, // 钱包余额为100.50
		Frozen:       0,
		TotalIncome:  100.50,
		TotalExpense: 0,
	}

	if err := db.Create(&wallet).Error; err != nil {
		fmt.Printf("创建钱包失败: %v\n", err)
		return
	}

	// 测试 GetUserInfo 方法
	fmt.Println("测试 GetUserInfo 方法:")
	userService := UserService{}
	userInfo, err := userService.GetUserInfo(user.ID)
	if err != nil {
		fmt.Printf("获取用户信息失败: %v\n", err)
		return
	}

	fmt.Printf("用户ID: %d\n", userInfo.ID)
	fmt.Printf("用户名: %s\n", userInfo.Username)
	fmt.Printf("用户余额: %.2f\n", userInfo.Balance)
	fmt.Printf("钱包余额: %.2f\n", wallet.Balance)

	if userInfo.Balance == wallet.Balance {
		fmt.Println("✓ 测试通过: 用户余额与钱包余额一致")
	} else {
		fmt.Println("✗ 测试失败: 用户余额与钱包余额不一致")
	}

	// 更新钱包余额
	wallet.Balance = 200.75
	if err := db.Save(&wallet).Error; err != nil {
		fmt.Printf("更新钱包失败: %v\n", err)
		return
	}

	// 再次测试 GetUserInfo 方法
	fmt.Println("\n测试更新钱包余额后:")
	userInfo, err = userService.GetUserInfo(user.ID)
	if err != nil {
		fmt.Printf("获取用户信息失败: %v\n", err)
		return
	}

	fmt.Printf("用户余额: %.2f\n", userInfo.Balance)
	fmt.Printf("钱包余额: %.2f\n", wallet.Balance)

	if userInfo.Balance == wallet.Balance {
		fmt.Println("✓ 测试通过: 用户余额与钱包余额一致")
	} else {
		fmt.Println("✗ 测试失败: 用户余额与钱包余额不一致")
	}

	fmt.Println("\n所有测试完成")
}

// UserService 复制 UserService 结构体用于测试
type UserService struct{}

// GetUserInfo 复制 GetUserInfo 方法用于测试
func (s *UserService) GetUserInfo(userID uint) (model.User, error) {
	var user model.User
	if err := global.GVA_DB.First(&user, userID).Error; err != nil {
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

func main() {
	TestUserBalance()
}
