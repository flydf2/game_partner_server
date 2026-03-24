package api

import (
	"github.com/gin-gonic/gin"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/playmate/service"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
)

type CategoryApi struct{}

// GetCategories 获取分类列表
// @Tags     Category
// @Summary  获取分类列表
// @Security ApiKeyAuth
// @accept   application/json
// @Produce  application/json
// @Success  200  {object} response.Response{data=[]model.Category} "获取分类列表成功"
// @Router   /categories [get]
func (a *CategoryApi) GetCategories(c *gin.Context) {
	categories, err := service.ServiceGroupApp.CategoryService.GetCategories()
	if err != nil {
		response.FailWithMessage("获取分类列表失败", c)
		return
	}
	response.OkWithData(categories, c)
}
