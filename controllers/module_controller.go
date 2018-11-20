package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kmookay/MyBittDataManage/dto/request"
	"github.com/kmookay/MyBittDataManage/model"
)

type ModuleController struct {
}

var moduleModle model.ModuleModel

func NewModuleController() ModuleController {
	moduleModle = model.ModuleModel{}
	return ModuleController{}
}

/**
 * 创建模块
 */
func (controller *ModuleController) Create(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var createModuleDto request.CreateModuleDto
		err := c.ShouldBindJSON(&createModuleDto)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}
		result, err := moduleModle.CrateModule(db, createModuleDto.Name, createModuleDto.Description, createModuleDto.ProjectId, createModuleDto.VersionId)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}
		c.JSON(http.StatusOK, model.ResultSuccess("操作成功", model.ResultSuccess("操作成功", result)))
	}
}

/**
 * 更新模块
 */
func (controller *ModuleController) Update(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateModuleDto request.UpdateModuleDto
		err := c.ShouldBindJSON(&updateModuleDto)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}
		module, err := moduleModle.UpdateModule(db, updateModuleDto.Id, updateModuleDto.Name, updateModuleDto.Description, updateModuleDto.ProjectId, updateModuleDto.VersionId)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}
		c.JSON(http.StatusOK, model.ResultSuccess("success", module))
	}
}

/**
 * 获取一个项目下的模块列表
 */
func (controller *ModuleController) List(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var getProjectModuleListDto request.GetModuleListDto
		err := c.BindJSON(&getProjectModuleListDto)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}

		moduleList, err := moduleModle.GetModuleListByProjectId(db, getProjectModuleListDto.ProjectId)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}

		c.JSON(http.StatusOK, model.ResultSuccess("Success", moduleList))
	}
}
