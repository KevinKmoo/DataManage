package controllers

import (
	"database/sql"
	"log"
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

func (pc *PageController) Create(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var createPageDto request.CreatePageDto
		err := c.ShouldBindJSON(&createPageDto)
		if err != nil {
			log.Fatal(err)
		}
		page, err := pageModel.CreatePage(db, createPageDto.Name, createPageDto.Description, createPageDto.Key, createPageDto.VersionId)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, model.ResultSuccess("Success", page))
	}
}
