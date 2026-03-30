package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
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
// @Router   /playmate/user/info [get]
func (a *UserApi) GetUserInfo(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	user, err := service.ServiceGroupApp.UserService.GetUserInfo(userID)
	if err != nil {
		response.FailWithError(err, c)
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
// @Router   /playmate/auth/login [post]
func (a *UserApi) Login(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	user, token, err := service.ServiceGroupApp.UserService.Login(req.Username, req.Password)
	if err != nil {
		response.FailWithError(err, c)
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
// @Router   /playmate/auth/register [post]
func (a *UserApi) Register(c *gin.Context) {
	var req request.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	user, token, err := service.ServiceGroupApp.UserService.Register(req)
	if err != nil {
		response.FailWithError(err, c)
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
// @Router   /playmate/user/profile [put]
func (a *UserApi) UpdateProfile(c *gin.Context) {
	var req request.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	user, err := service.ServiceGroupApp.UserService.UpdateProfile(userID, req)
	if err != nil {
		response.FailWithError(err, c)
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
// @Router   /playmate/user/settings [get]
func (a *UserApi) GetSettings(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	settings, err := service.ServiceGroupApp.UserService.GetSettings(userID)
	if err != nil {
		response.FailWithError(err, c)
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
// @Router   /playmate/user/settings [put]
func (a *UserApi) UpdateSettings(c *gin.Context) {
	var req request.UpdateSettingsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	settings, err := service.ServiceGroupApp.UserService.UpdateSettings(userID, req)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(settings, "更新成功", c)
}

// Logout 用户登出
// @Tags     Auth
// @Summary  用户登出
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{message=string} "登出成功"
// @Router   /playmate/auth/logout [post]
func (a *UserApi) Logout(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	if userID != 0 {
		// 如果用户已登录，执行登出逻辑
		err := service.ServiceGroupApp.UserService.Logout(userID)
		if err != nil {
			response.FailWithError(err, c)
			return
		}
	}
	// 无论token是否有效，都返回成功（兼容token过期或无效的情况）
	response.OkWithMessage("登出成功", c)
}

// RefreshToken 刷新token
// @Tags     Auth
// @Summary  刷新token
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=map[string]string} "刷新成功"
// @Router   /playmate/auth/refresh [post]
func (a *UserApi) RefreshToken(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	token, err := service.ServiceGroupApp.UserService.RefreshToken(userID)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"token": token,
	}, "刷新成功", c)
}

// FollowUser 关注用户
// @Tags     User
// @Summary  关注用户
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    userId path uint true "用户ID"
// @Success  200  {object} response.Response{message=string} "关注成功"
// @Router   /playmate/user/following/{userId} [post]
func (a *UserApi) FollowUser(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	targetUserID, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	err = service.ServiceGroupApp.UserService.FollowUser(userID, uint(targetUserID))
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("关注成功", c)
}

// UnfollowUser 取消关注用户
// @Tags     User
// @Summary  取消关注用户
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    userId path uint true "用户ID"
// @Success  200  {object} response.Response{message=string} "取消关注成功"
// @Router   /playmate/user/following/{userId} [delete]
func (a *UserApi) UnfollowUser(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	targetUserID, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	err = service.ServiceGroupApp.UserService.UnfollowUser(userID, uint(targetUserID))
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("取消关注成功", c)
}

// RemoveFavorite 移除收藏
// @Tags     User
// @Summary  移除收藏
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    favoriteId path uint true "收藏ID"
// @Success  200  {object} response.Response{message=string} "移除成功"
// @Router   /playmate/user/favorites/{favoriteId} [delete]
func (a *UserApi) RemoveFavorite(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	favoriteID, err := strconv.ParseUint(c.Param("favoriteId"), 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	err = service.ServiceGroupApp.UserService.RemoveFavorite(userID, uint(favoriteID))
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("移除成功", c)
}

// ClearHistory 清空浏览历史
// @Tags     User
// @Summary  清空浏览历史
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{message=string} "清空成功"
// @Router   /playmate/user/history [delete]
func (a *UserApi) ClearHistory(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	err := service.ServiceGroupApp.UserService.ClearHistory(userID)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithMessage("清空成功", c)
}

// GetFollowing 获取关注列表
// @Tags     User
// @Summary  获取关注列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    page     query    int  false "页码"
// @Param    pageSize query    int  false "每页数量"
// @Success  200      {object} response.Response{data=[]model.Playmate,pagination=map[string]int64} "获取成功"
// @Router   /playmate/user/following [get]
func (a *UserApi) GetFollowing(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	following, total, err := service.ServiceGroupApp.UserService.GetFollowing(userID, page, pageSize)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": following,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// GetFavorites 获取收藏列表
// @Tags     User
// @Summary  获取收藏列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    page     query    int  false "页码"
// @Param    pageSize query    int  false "每页数量"
// @Success  200      {object} response.Response{data=[]model.Playmate,pagination=map[string]int64} "获取成功"
// @Router   /playmate/user/favorites [get]
func (a *UserApi) GetFavorites(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	favorites, total, err := service.ServiceGroupApp.UserService.GetFavorites(userID, page, pageSize)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": favorites,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// GetBrowseHistory 获取浏览历史
// @Tags     User
// @Summary  获取浏览历史
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    page     query    int  false "页码"
// @Param    pageSize query    int  false "每页数量"
// @Success  200      {object} response.Response{data=[]model.UserBrowseHistory,pagination=map[string]int64} "获取成功"
// @Router   /playmate/user/history [get]
func (a *UserApi) GetBrowseHistory(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	history, total, err := service.ServiceGroupApp.UserService.GetBrowseHistory(userID, page, pageSize)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": history,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// GetWallet 获取钱包信息
// @Tags     User
// @Summary  获取钱包信息
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=map[string]interface{}} "获取成功"
// @Router   /playmate/user/wallet [get]
func (a *UserApi) GetWallet(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	wallet, transactions, err := service.ServiceGroupApp.UserService.GetWallet(userID)
	if err != nil {
		response.FailWithError(err, c)
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

// GetUsers 获取用户列表
// @Tags     User
// @Summary  获取用户列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    page     query    int  false "页码"
// @Param    pageSize query    int  false "每页数量"
// @Success  200      {object} response.Response{data=[]model.User,pagination=map[string]int64} "获取成功"
// @Router   /playmate/users [get]
func (a *UserApi) GetUsers(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	users, total, err := service.ServiceGroupApp.UserService.GetUsers(page, pageSize)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": users,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// Recharge 充值
// @Tags     User
// @Summary  充值
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body      request.RechargeRequest  true "充值信息"
// @Success  200   {object} response.Response{data=map[string]interface{}} "充值成功"
// @Router   /playmate/user/recharge [post]
func (a *UserApi) Recharge(c *gin.Context) {
	var req request.RechargeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	result, err := service.ServiceGroupApp.UserService.Recharge(userID, req)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(result, "充值成功", c)
}

// GetTransactionList 获取交易记录列表
// @Tags     User
// @Summary  获取交易记录列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    page     query    int  false "页码"
// @Param    pageSize query    int  false "每页数量"
// @Success  200      {object} response.Response{data=[]model.Transaction,pagination=map[string]int64} "获取成功"
// @Router   /playmate/user/transactions [get]
func (a *UserApi) GetTransactionList(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	if userID == 0 {
		response.NoAuth("未登录或登录已过期", c)
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	transactions, total, err := service.ServiceGroupApp.UserService.GetTransactionList(userID, page, pageSize)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(gin.H{
		"data": transactions,
		"pagination": gin.H{
			"currentPage": page,
			"totalPages":  (total + int64(pageSize) - 1) / int64(pageSize),
			"totalCount":  total,
		},
	}, "获取成功", c)
}

// HandlePaymentCallback 处理支付回调
// @Tags     User
// @Summary  处理支付回调
// @accept   application/json
// @Produce  application/json
// @Param    orderId query    string true "订单号"
// @Param    status  query    string true "支付状态(success/failed)"
// @Success  200     {object} response.Response{data=map[string]interface{}} "处理成功"
// @Router   /playmate/user/payment/callback [get]
func (a *UserApi) HandlePaymentCallback(c *gin.Context) {
	orderID := c.Query("orderId")
	status := c.Query("status")

	if orderID == "" || status == "" {
		response.FailWithMessage("参数错误", c)
		return
	}

	result, err := service.ServiceGroupApp.UserService.HandlePaymentCallback(orderID, status)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	response.OkWithDetailed(result, "处理成功", c)
}
