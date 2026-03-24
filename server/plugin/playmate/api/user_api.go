package api

import (
	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
)

// UserApi 用户API
type UserApi struct{}

// GetUserInfo 获取用户信息
// @Tags     User
// @Summary  获取用户信息
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=model.User} "获取成功"
// @Router   /user/info [get]
func (a *UserApi) GetUserInfo(c *gin.Context) {
	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	user, err := service.ServiceGroupApp.UserService.GetUserInfo(userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(user, "获取成功", c)
}

// Login 用户登录
// @Tags     Auth
// @Summary  用户登录
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.LoginRequest  true "登录信息"
// @Success  200   {object} response.Response{data=map[string]interface{}} "登录成功"
// @Router   /auth/login [post]
func (a *UserApi) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	user, token, err := service.ServiceGroupApp.UserService.Login(req.Username, req.Password)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"token": token,
		"user":  user,
	}, "登录成功", c)
}

// Register 用户注册
// @Tags     Auth
// @Summary  用户注册
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.RegisterRequest  true "注册信息"
// @Success  200   {object} response.Response{data=map[string]interface{}} "注册成功"
// @Router   /auth/register [post]
func (a *UserApi) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	user, token, err := service.ServiceGroupApp.UserService.Register(req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"token": token,
		"user":  user,
	}, "注册成功", c)
}

// UpdateProfile 更新个人资料
// @Tags     User
// @Summary  更新个人资料
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.UpdateProfileRequest  true "个人资料"
// @Success  200   {object} response.Response{data=model.User} "更新成功"
// @Router   /user/profile [put]
func (a *UserApi) UpdateProfile(c *gin.Context) {
	var req request.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	user, err := service.ServiceGroupApp.UserService.UpdateProfile(userID, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(user, "更新成功", c)
}

// GetSettings 获取用户设置
// @Tags     User
// @Summary  获取用户设置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=model.UserSettings} "获取成功"
// @Router   /user/settings [get]
func (a *UserApi) GetSettings(c *gin.Context) {
	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	settings, err := service.ServiceGroupApp.UserService.GetSettings(userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(settings, "获取成功", c)
}

// UpdateSettings 更新用户设置
// @Tags     User
// @Summary  更新用户设置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.UpdateSettingsRequest  true "设置信息"
// @Success  200   {object} response.Response{data=model.UserSettings} "更新成功"
// @Router   /user/settings [put]
func (a *UserApi) UpdateSettings(c *gin.Context) {
	var req request.UpdateSettingsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	settings, err := service.ServiceGroupApp.UserService.UpdateSettings(userID, req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(settings, "更新成功", c)
}

// Logout 用户登出
// @Tags     Auth
// @Summary  用户登出
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{message=string} "登出成功"
// @Router   /auth/logout [post]
func (a *UserApi) Logout(c *gin.Context) {
	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	err := service.ServiceGroupApp.UserService.Logout(userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithMessage("登出成功", c)
}

// RefreshToken 刷新token
// @Tags     Auth
// @Summary  刷新token
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=map[string]string} "刷新成功"
// @Router   /auth/refresh [post]
func (a *UserApi) RefreshToken(c *gin.Context) {
	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	token, err := service.ServiceGroupApp.UserService.RefreshToken(userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"token": token,
	}, "刷新成功", c)
}

// GetFollowing 获取关注列表
// @Tags     User
// @Summary  获取关注列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]model.Playmate} "获取成功"
// @Router   /user/following [get]
func (a *UserApi) GetFollowing(c *gin.Context) {
	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	following, err := service.ServiceGroupApp.UserService.GetFollowing(userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(following, "获取成功", c)
}

// GetFavorites 获取收藏列表
// @Tags     User
// @Summary  获取收藏列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]model.Playmate} "获取成功"
// @Router   /user/favorites [get]
func (a *UserApi) GetFavorites(c *gin.Context) {
	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	favorites, err := service.ServiceGroupApp.UserService.GetFavorites(userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(favorites, "获取成功", c)
}

// GetBrowseHistory 获取浏览历史
// @Tags     User
// @Summary  获取浏览历史
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]model.UserBrowseHistory} "获取成功"
// @Router   /user/history [get]
func (a *UserApi) GetBrowseHistory(c *gin.Context) {
	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	history, err := service.ServiceGroupApp.UserService.GetBrowseHistory(userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(history, "获取成功", c)
}

// GetWallet 获取钱包信息
// @Tags     User
// @Summary  获取钱包信息
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=map[string]interface{}} "获取成功"
// @Router   /user/wallet [get]
func (a *UserApi) GetWallet(c *gin.Context) {
	// 这里应该从上下文获取用户ID
	userID := uint(1) // 临时值

	wallet, transactions, err := service.ServiceGroupApp.UserService.GetWallet(userID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(gin.H{
		"balance":      wallet.Balance,
		"frozen":       wallet.Frozen,
		"totalIncome":  wallet.TotalIncome,
		"totalExpense": wallet.TotalExpense,
		"transactions": transactions,
	}, "获取成功", c)
}