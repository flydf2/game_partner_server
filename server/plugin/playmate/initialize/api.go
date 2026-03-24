package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/api"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/config"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
)

// InitializeAPI 初始化API
func InitializeAPI() {

	// 从config获取API配置
	entities := config.GetApis()
	utils.RegisterApis(entities...)
	// 从config获取字典配置
	dictionaries := config.GetDictionaries()
	utils.RegisterDictionaries(dictionaries...)

	// 注册API到全局ApiGroupApp
	api.ApiGroupApp = &api.ApiGroup{}

}
