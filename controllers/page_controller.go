package controllers

import (
	"database/sql"
	"net/http"

	"github.com/kmookay/MyBittDataManage/model"

	"github.com/gin-gonic/gin"
	"github.com/kmookay/MyBittDataManage/dto/request"
)

type PageController struct {
}

func NewPageController() *PageController {
	pageModel = model.PageModel{}
	return &PageController{}
}

var pageModel model.PageModel

/**
 * 创建页面
 */
func (pc *PageController) Create(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var createPageDto request.CreatePageDto
		err := c.ShouldBindJSON(&createPageDto)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}
		page, err := pageModel.CreatePage(db, createPageDto.Name, createPageDto.Description, createPageDto.Key, createPageDto.VersionId, createPageDto.ModuleId)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}
		c.JSON(http.StatusOK, model.ResultSuccess("Success", page))
	}
}

/**
 * 更新页面
 */
func (controller *PageController) Update(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updatePageDto request.UpdatePageDto
		err := c.ShouldBindJSON(&updatePageDto)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}
		page, err := pageModel.UpdatePage(db, updatePageDto.Id, updatePageDto.Name, updatePageDto.Description)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}
		c.JSON(http.StatusOK, model.ResultSuccess("Success", page))
	}
}

/**
 * 获取页面的列表
 */
func (controller *PageController) List(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pageListDto request.PageListDto
		err := c.ShouldBindJSON(&pageListDto)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}
		result, err := pageModel.PageList(db, pageListDto.ProjectId, pageListDto.ModuelId, pageListDto.VersionId, pageListDto.Key)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail(err.Error(), nil))
			return
		}
		c.JSON(http.StatusOK, model.ResultSuccess("Success", result))
	}
}
