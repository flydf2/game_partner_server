package playmate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/plugin/v2"
	"github.com/gin-gonic/gin"
)

// PlaymatePlugin 插件结构体
type PlaymatePlugin struct{}

// Register 注册插件
func (p *PlaymatePlugin) Register(engine *gin.Engine) {
	// 初始化数据库
	initialize.InitializeDB()

	// 初始化API
	initialize.InitializeAPI()

	// 初始化路由
	routerGroup := engine.Group("/")
	initialize.InitializeRouter(routerGroup)

	// 初始化菜单
	initialize.InitializeMenu()

	// 初始字典
	initialize.InitializeDictionaries()
}

// RouterPath 获取路由路径
func (p *PlaymatePlugin) RouterPath() string {
	return "/"
}

// Name 获取插件名称
func (p *PlaymatePlugin) Name() string {
	return "playmate"
}

// Init 初始化插件
func (p *PlaymatePlugin) Init() error {
	return nil
}

// Plugin 插件实例
var Plugin = new(PlaymatePlugin)

// init 注册插件
func init() {
	plugin.Register(Plugin)
}
