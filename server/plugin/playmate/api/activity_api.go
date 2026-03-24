package api

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
)

type ActivityApi struct{}

// GetActivities 获取活动列表
// @Tags     Activity
// @Summary  获取活动列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]model.Activity} "获取活动列表成功"
// @Router   /activities [get]
func (a *ActivityApi) GetActivities(c *gin.Context) {
	activities, err := service.ServiceGroupApp.ActivityService.GetActivities()
	if err != nil {
		response.FailWithMessage("获取活动列表失败", c)
		return
	}
	response.OkWithData(activities, c)
}
