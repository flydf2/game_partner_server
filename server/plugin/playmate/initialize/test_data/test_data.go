package initialize

import (
	"fmt"
	"log"

	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// TestInitializeData 测试数据初始化
func TestInitializeData() {
	// 初始化系统
	global.GVA_VP = core.Viper() // 初始化Viper
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	global.GVA_DB = core.Gorm() // gorm连接数据库

	if global.GVA_DB == nil {
		log.Fatal("Failed to connect to database")
	}

	// 自动迁移表结构
	InitializeDB()

	// 初始化测试数据
	InitializeData()

	fmt.Println("Test data initialized successfully!")
}
