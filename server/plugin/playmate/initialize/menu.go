package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/config"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
	"go.uber.org/zap"
)

// InitializeMenu 初始化菜单
func InitializeMenu() {
	// 从config获取菜单配置
	menus := config.GetMenus()
	// 从config获取菜单配置
	utils.RegisterMenus(menus...)

	var parentId uint
	parentId = 0
	// 首先创建父级菜单（ParentId为0的菜单）
	for _, menu := range menus {
		var existingMenu system.SysBaseMenu
		if menu.ParentId != 0 {
			menu.ParentId = parentId
		}
		if err := global.GVA_DB.Where("name = ?", menu.Name).First(&existingMenu).Error; err != nil {
			// 菜单不存在，创建新菜单
			if err := global.GVA_DB.Create(&menu).Error; err != nil {
				global.GVA_LOG.Error("创建父级菜单失败", zap.Error(err))
				continue
			}
			// 获取创建的菜单ID
			if err := global.GVA_DB.Where("name = ?", menu.Name).First(&existingMenu).Error; err != nil {
				global.GVA_LOG.Error("获取父级菜单ID失败", zap.Error(err))
				continue
			}
		}
		if menu.ParentId == 0 {
			parentId = existingMenu.ID
		}
	}

	global.GVA_LOG.Info("菜单初始化完成")

}
