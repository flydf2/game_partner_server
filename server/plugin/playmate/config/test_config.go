package config

// TestAuthConfig 测试认证配置
type TestAuthConfig struct {
	// 是否启用万能Token（仅用于测试环境）
	EnableUniversalToken bool `mapstructure:"enable-universal-token" json:"enable-universal-token" yaml:"enable-universal-token"`
	// 万能Token列表，key为token值，value为用户ID
	UniversalTokens map[string]uint `mapstructure:"universal-tokens" json:"universal-tokens" yaml:"universal-tokens"`
	// 万能验证码（固定值）
	UniversalCaptcha string `mapstructure:"universal-captcha" json:"universal-captcha" yaml:"universal-captcha"`
}

// DefaultTestAuthConfig 默认测试认证配置
func DefaultTestAuthConfig() TestAuthConfig {
	return TestAuthConfig{
		EnableUniversalToken: true,
		UniversalTokens: map[string]uint{
			"test_auth_token_user_1": 1, // 测试用户1
			"test_auth_token_user_2": 2, // 测试用户2
			"test_auth_token_user_3": 3, // 测试用户3
		},
		UniversalCaptcha: "123456",
	}
}
