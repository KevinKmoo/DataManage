package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kmookay/MyBittDataManage/dto/request"
	"github.com/kmookay/MyBittDataManage/model"
)

type VersionController struct{}

var versionModel model.VersionModel

func NewVersionController() VersionController {
	return VersionController{}
}

/**
 * 获取所有的版本
 */
func (controller *VersionController) GetAllVersion(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		allVersions, err := versionModel.FindAllVersion(db)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}
		c.JSON(http.StatusOK, model.ResultSuccess("成功", allVersions))
	}
}

/**
 * 查找项目下的版本列表
 */
func (controller *VersionController) GetVersionListByProjectId(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var getVersionByProjectIdDto request.GetVersionListByProjectIdDto
		err := c.BindJSON(&getVersionByProjectIdDto)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}
		versionList, err := versionModel.FindVerionListByProjectId(db, getVersionByProjectIdDto.ProjectId)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}
		c.JSON(http.StatusOK, model.ResultSuccess("Success", versionList))
	}
}

/**
 * 创建版本
 */
func (controller *VersionController) Create(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var createVersionDto request.CreateVersionDto
		err := c.BindJSON(&createVersionDto)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail("参数有误", nil))
			return
		}
		version, err := versionModel.CreateVersion(db, createVersionDto.Name, createVersionDto.Description, createVersionDto.ProjectId, createVersionDto.PublishTime)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}
		c.JSON(http.StatusOK, model.ResultSuccess("成功", version))
	}
}

/**
 * 更新版本
 */
func (controller *VersionController) Update(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateVersionDto request.UpdateVersionDto
		err := c.BindJSON(&updateVersionDto)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail("参数有误", nil))
			return
		}
		version, err := versionModel.UpdateVersion(db, updateVersionDto.Id, updateVersionDto.Name, updateVersionDto.Description)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}
		c.JSON(http.StatusOK, model.ResultSuccess("Success", version))
	}
}

/**
 * 删除版本
 */
func (controller *VersionController) Delete(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var deleteVersionDto request.DeleteVersionDto
		err := c.ShouldBind(&deleteVersionDto)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail("参数有误", nil))
			return
		}
		c.JSON(http.StatusOK, model.ResultSuccess("成功", nil))
	}
}
