package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/kmookay/MyBittDataManage/constant"

	"github.com/gin-gonic/gin"
	"github.com/kmookay/MyBittDataManage/dto/request"
	"github.com/kmookay/MyBittDataManage/model"
)

type ProjectController struct {
}

func NewProjectController() ProjectController {
	return ProjectController{}
}

func (p ProjectController) CreateProject(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var createProjectDto request.CreateProjectDto
		if err := c.BindJSON(&createProjectDto); err != nil {
			log.Fatal(err)
			return
		}
		createdProject, _ := model.CreateProject(db, createProjectDto.Name)
		c.JSON(http.StatusOK, model.ResultSuccess("Success", createdProject))
	}
}

func (p ProjectController) ListProject(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagingDto request.PagingDto
		if err := c.ShouldBindJSON(&pagingDto); err != nil {
			pagingDto = request.PagingDto{
				PageIndex: constant.PAGE_DEFAULT_INDEX,
				PageSize:  constant.PAGE_DEFAULT_SIZE,
			}
		}
		projectList, _ := model.ProjectList(db, pagingDto.PageIndex, pagingDto.PageSize)
		c.JSON(http.StatusOK, model.ResultSuccess("Success", projectList))
	}
}

func (p ProjectController) UpdateProject(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateProjectDto request.UpdateProjectDto
		if err := c.ShouldBindJSON(&updateProjectDto); err != nil {
			log.Fatal(err)
			return
		}
		updatedProject, _ := model.UpdateProject(db, updateProjectDto.Id, updateProjectDto.Name)
		c.JSON(http.StatusOK, model.ResultSuccess("Success", updatedProject))
	}
}
