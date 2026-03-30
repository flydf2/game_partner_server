package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/model/response"
	svc "github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
)

// ExpertOrderSettingApi 专家订单设置API
type ExpertOrderSettingApi struct{}

// GetOrderSetting 获取专家订单设置
// @Tags     ExpertOrderSetting
// @Summary  获取专家订单设置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object}  response.Response{data=model.ExpertOrderSetting} "获取成功"
// @Router   /playmate/expert/order-settings [get]
func (a *ExpertOrderSettingApi) GetOrderSetting(c *gin.Context) {
	// 从上下文获取当前用户ID作为专家ID
	// 这里暂时使用默认值，实际应该从JWT token中获取
	expertID := uint(1)

	setting, err := svc.ServiceGroupApp.ExpertOrderSettingService.GetOrderSetting(expertID)
	if err != nil {
		response.FailWithMessage("获取订单设置失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(setting, "获取成功", c)
}

// UpdateOrderSetting 更新专家订单设置
// @Tags     ExpertOrderSetting
// @Summary  更新专家订单设置
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body      model.ExpertOrderSetting  true  "订单设置信息"
// @Success  200   {object}  response.Response{data=model.ExpertOrderSetting} "更新成功"
// @Router   /playmate/expert/order-settings [put]
func (a *ExpertOrderSettingApi) UpdateOrderSetting(c *gin.Context) {
	var setting model.ExpertOrderSetting
	if err := c.ShouldBindJSON(&setting); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 从上下文获取当前用户ID作为专家ID
	expertID := uint(1)
	setting.ExpertID = expertID

	updatedSetting, err := svc.ServiceGroupApp.ExpertOrderSettingService.UpdateOrderSetting(&setting)
	if err != nil {
		response.FailWithMessage("更新订单设置失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(updatedSetting, "更新成功", c)
}

// GetExpertServices 获取专家服务列表
// @Tags     ExpertOrderSetting
// @Summary  获取专家服务列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object}  response.Response{data=[]model.ExpertService} "获取成功"
// @Router   /playmate/expert/order-settings/services [get]
func (a *ExpertOrderSettingApi) GetExpertServices(c *gin.Context) {
	// 从上下文获取当前用户ID作为专家ID
	expertID := uint(1)

	services, err := svc.ServiceGroupApp.ExpertOrderSettingService.GetExpertServices(expertID)
	if err != nil {
		response.FailWithMessage("获取服务列表失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(services, "获取成功", c)
}

// CreateExpertService 创建专家服务
// @Tags     ExpertOrderSetting
// @Summary  创建专家服务
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body      model.ExpertService  true  "服务信息"
// @Success  200   {object}  response.Response{data=model.ExpertService} "创建成功"
// @Router   /playmate/expert/order-settings/services [post]
func (a *ExpertOrderSettingApi) CreateExpertService(c *gin.Context) {
	var s model.ExpertService
	if err := c.ShouldBindJSON(&s); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 从上下文获取当前用户ID作为专家ID
	expertID := uint(1)
	s.ExpertID = expertID

	createdService, err := svc.ServiceGroupApp.ExpertOrderSettingService.CreateExpertService(&s)
	if err != nil {
		response.FailWithMessage("创建服务失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(createdService, "创建成功", c)
}

// UpdateExpertService 更新专家服务
// @Tags     ExpertOrderSetting
// @Summary  更新专家服务
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id    path      uint                 true  "服务ID"
// @Param    data  body      model.ExpertService  true  "服务信息"
// @Success  200   {object}  response.Response{data=model.ExpertService} "更新成功"
// @Router   /playmate/expert/order-settings/services/{id} [put]
func (a *ExpertOrderSettingApi) UpdateExpertService(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var s model.ExpertService
	if err := c.ShouldBindJSON(&s); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	s.ID = uint(id)

	updatedService, err := svc.ServiceGroupApp.ExpertOrderSettingService.UpdateExpertService(&s)
	if err != nil {
		response.FailWithMessage("更新服务失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(updatedService, "更新成功", c)
}

// DeleteExpertService 删除专家服务
// @Tags     ExpertOrderSetting
// @Summary  删除专家服务
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id  path      uint  true  "服务ID"
// @Success  200 {object}  response.Response{msg=string} "删除成功"
// @Router   /playmate/expert/order-settings/services/{id} [delete]
func (a *ExpertOrderSettingApi) DeleteExpertService(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := svc.ServiceGroupApp.ExpertOrderSettingService.DeleteExpertService(uint(id)); err != nil {
		response.FailWithMessage("删除服务失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}

// GetTodayRecommendations 获取今日推荐列表
// @Tags     ExpertOrderSetting
// @Summary  获取今日推荐列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object}  response.Response{data=[]model.TodayRecommendation} "获取成功"
// @Router   /playmate/expert/order-settings/today-recommendations [get]
func (a *ExpertOrderSettingApi) GetTodayRecommendations(c *gin.Context) {
	// 从上下文获取当前用户ID作为专家ID
	expertID := uint(1)

	recommendations, err := svc.ServiceGroupApp.ExpertOrderSettingService.GetTodayRecommendations(expertID)
	if err != nil {
		response.FailWithMessage("获取今日推荐失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(recommendations, "获取成功", c)
}

// CreateTodayRecommendation 创建今日推荐
// @Tags     ExpertOrderSetting
// @Summary  创建今日推荐
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    data  body      model.TodayRecommendation  true  "推荐信息"
// @Success  200   {object}  response.Response{data=model.TodayRecommendation} "创建成功"
// @Router   /playmate/expert/order-settings/today-recommendations [post]
func (a *ExpertOrderSettingApi) CreateTodayRecommendation(c *gin.Context) {
	var recommendation model.TodayRecommendation
	if err := c.ShouldBindJSON(&recommendation); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	// 从上下文获取当前用户ID作为专家ID
	expertID := uint(1)
	recommendation.ExpertID = expertID

	createdRecommendation, err := svc.ServiceGroupApp.ExpertOrderSettingService.CreateTodayRecommendation(&recommendation)
	if err != nil {
		response.FailWithMessage("创建推荐失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(createdRecommendation, "创建成功", c)
}

// UpdateTodayRecommendation 更新今日推荐
// @Tags     ExpertOrderSetting
// @Summary  更新今日推荐
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id    path      uint                       true  "推荐ID"
// @Param    data  body      model.TodayRecommendation  true  "推荐信息"
// @Success  200   {object}  response.Response{data=model.TodayRecommendation} "更新成功"
// @Router   /playmate/expert/order-settings/today-recommendations/{id} [put]
func (a *ExpertOrderSettingApi) UpdateTodayRecommendation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	var recommendation model.TodayRecommendation
	if err := c.ShouldBindJSON(&recommendation); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	recommendation.ID = uint(id)

	updatedRecommendation, err := svc.ServiceGroupApp.ExpertOrderSettingService.UpdateTodayRecommendation(&recommendation)
	if err != nil {
		response.FailWithMessage("更新推荐失败: "+err.Error(), c)
		return
	}

	response.OkWithDetailed(updatedRecommendation, "更新成功", c)
}

// DeleteTodayRecommendation 删除今日推荐
// @Tags     ExpertOrderSetting
// @Summary  删除今日推荐
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Param    id  path      uint  true  "推荐ID"
// @Success  200 {object}  response.Response{msg=string} "删除成功"
// @Router   /playmate/expert/order-settings/today-recommendations/{id} [delete]
func (a *ExpertOrderSettingApi) DeleteTodayRecommendation(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	if err := svc.ServiceGroupApp.ExpertOrderSettingService.DeleteTodayRecommendation(uint(id)); err != nil {
		response.FailWithMessage("删除推荐失败: "+err.Error(), c)
		return
	}

	response.OkWithMessage("删除成功", c)
}
