package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/config"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

// InitializeAPI 初始化API
func InitializeDictionaries() {
	// 从config获取字典配置
	dictionaries := config.GetDictionaries()
	utils.RegisterDictionaries(dictionaries...)
}
