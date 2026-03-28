package middleware

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/config"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
)

// 全局测试配置（可以从配置文件加载）
var testAuthConfig = config.DefaultTestAuthConfig()

// TestAuthMiddleware 测试认证中间件
// 支持通过 x-test-auth-token 头部传入万能Token，跳过正常JWT验证
func TestAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查是否启用了万能Token
		if !testAuthConfig.EnableUniversalToken {
			c.Next()
			return
		}

		// 获取万能Token
		testToken := c.GetHeader("x-test-auth-token")
		if testToken == "" {
			c.Next()
			return
		}

		// 验证万能Token
		if userID, ok := testAuthConfig.UniversalTokens[testToken]; ok {
			// 设置测试用户上下文
			c.Set("test_user_id", userID)
			c.Set("is_test_auth", true)
			c.Set("test_token", testToken)
			c.Next()
			return
		}

		// Token无效
		response.NoAuth("无效的测试认证Token", c)
		c.Abort()
	}
}

// GetTestUserID 从上下文获取测试用户ID
func GetTestUserID(c *gin.Context) (uint, bool) {
	if userID, exists := c.Get("test_user_id"); exists {
		if id, ok := userID.(uint); ok {
			return id, true
		}
	}
	return 0, false
}

// IsTestAuth 检查是否是通过测试认证
func IsTestAuth(c *gin.Context) bool {
	if isTest, exists := c.Get("is_test_auth"); exists {
		if test, ok := isTest.(bool); ok {
			return test
		}
	}
	return false
}

// CombinedAuthMiddleware 组合认证中间件
// 先尝试JWT认证，如果失败再尝试测试Token认证
func CombinedAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 先检查是否有测试Token
		testToken := c.GetHeader("x-test-auth-token")
		if testToken != "" && testAuthConfig.EnableUniversalToken {
			if userID, ok := testAuthConfig.UniversalTokens[testToken]; ok {
				c.Set("test_user_id", userID)
				c.Set("is_test_auth", true)
				c.Set("test_token", testToken)
				c.Next()
				return
			}
		}

		// 没有测试Token或无效，尝试JWT认证
		token := utils.GetToken(c)
		if token == "" {
			response.NoAuth("未登录或非法访问，请登录", c)
			c.Abort()
			return
		}

		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			response.NoAuth("登录已过期或无效，请重新登录", c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}

// GetCurrentUserID 获取当前用户ID（支持JWT和测试Token）
func GetCurrentUserID(c *gin.Context) uint {
	// 先检查测试认证
	if userID, ok := GetTestUserID(c); ok {
		return userID
	}

	// 再检查JWT认证
	if claims, exists := c.Get("claims"); exists {
		if jwtClaims, ok := claims.(*request.CustomClaims); ok {
			return jwtClaims.BaseClaims.ID
		}
	}

	return 0
}

// VerifyCaptcha 验证验证码（支持万能验证码）
func VerifyCaptcha(captchaId, captchaCode string) bool {
	// 检查是否是万能验证码
	if captchaCode == testAuthConfig.UniversalCaptcha {
		return true
	}

	// 这里可以添加正常的验证码验证逻辑
	// 例如从Redis或内存存储中验证
	return false
}
