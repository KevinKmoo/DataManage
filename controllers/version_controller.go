package controllers

import (
	"database/sql"
	"log"
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
			c.JSON(http.StatusOK, model.ResultFail("未知错误", nil))
			return
		}
		c.JSON(http.StatusOK, model.ResultSuccess("成功", allVersions))
	}
}

/**
 * 创建版本
 */
func (controller *VersionController) Create(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var createVersionDto request.CreateVersionDto
		err := c.ShouldBind(&createVersionDto)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail("参数有误", nil))
			return
		}
		version, err := versionModel.CreateVersion(db, createVersionDto.Name, createVersionDto.Description, createVersionDto.PublishTime)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail("未知错误", nil))
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
