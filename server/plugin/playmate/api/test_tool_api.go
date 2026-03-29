package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/config"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/gin-gonic/gin"
)

// TestToolApi 测试工具API
type TestToolApi struct{}

// GetTestTokens 获取测试Token列表
// @Tags     TestTool
// @Summary  获取测试Token列表（仅用于自动化测试）
// @accept   application/json
// @Produce  application/json
// @Success  200  {object}  response.Response{data=map[string]interface{}} "获取成功"
// @Router   /playmate/test-tool/tokens [get]
func (a *TestToolApi) GetTestTokens(c *gin.Context) {
	cfg := config.DefaultTestAuthConfig()

	// 构建token列表（不包含敏感信息）
	tokens := make([]map[string]interface{}, 0)
	for token, userID := range cfg.UniversalTokens {
		tokens = append(tokens, map[string]interface{}{
			"token":    token,
			"userId":   userID,
			"description": getUserDescription(userID),
		})
	}

	response.OkWithDetailed(map[string]interface{}{
		"tokens":           tokens,
		"universalCaptcha": cfg.UniversalCaptcha,
		"enabled":          cfg.EnableUniversalToken,
	}, "获取测试Token列表成功", c)
}

// VerifyCaptcha 验证验证码（支持万能验证码）
// @Tags     TestTool
// @Summary  验证验证码（支持万能验证码123456）
// @accept   application/json
// @Produce  application/json
// @Param    data  body      VerifyCaptchaRequest  true "验证码信息"
// @Success  200   {object}  response.Response{data=map[string]interface{}} "验证成功"
// @Router   /playmate/test-tool/verify-captcha [post]
func (a *TestToolApi) VerifyCaptcha(c *gin.Context) {
	var req struct {
		CaptchaId   string `json:"captchaId" binding:"required"`
		CaptchaCode string `json:"captchaCode" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	valid := middleware.VerifyCaptcha(req.CaptchaId, req.CaptchaCode)

	response.OkWithDetailed(map[string]interface{}{
		"valid":   valid,
		"captchaId": req.CaptchaId,
	}, "验证码验证完成", c)
}

// GetTestCaptcha 获取万能验证码
// @Tags     TestTool
// @Summary  获取万能验证码（固定值123456）
// @accept   application/json
// @Produce  application/json
// @Success  200  {object}  response.Response{data=map[string]interface{}} "获取成功"
// @Router   /playmate/test-tool/captcha [get]
func (a *TestToolApi) GetTestCaptcha(c *gin.Context) {
	cfg := config.DefaultTestAuthConfig()

	response.OkWithDetailed(map[string]interface{}{
		"captchaCode": cfg.UniversalCaptcha,
		"tip":         "此验证码仅用于自动化测试，固定值为 123456",
	}, "获取万能验证码成功", c)
}

// VerifyCaptchaRequest 验证验证码请求
type VerifyCaptchaRequest struct {
	CaptchaId   string `json:"captchaId" binding:"required"`
	CaptchaCode string `json:"captchaCode" binding:"required"`
}

func getUserDescription(userID uint) string {
	switch userID {
	case 1:
		return "测试用户1 - 普通用户"
	case 2:
		return "测试用户2 - 陪玩用户"
	case 3:
		return "测试用户3 - 管理员"
	default:
		return "测试用户"
	}
}
