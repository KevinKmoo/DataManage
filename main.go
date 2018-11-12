package main

import (
	"net/http"

	"github.com/kmookay/MyBittDataManage/model"

	"github.com/gin-contrib/cors"
	"github.com/kmookay/MyBittDataManage/controllers"
	"github.com/kmookay/MyBittDataManage/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	mybittDb := db.CreateMyBittDbConnection()

	userController := controllers.NewUserController()
	userRouter := r.Group("/user")
	{
		userRouter.POST("/login", userController.Login(mybittDb))
		userRouter.POST("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, model.ResultSuccess("success", gin.H{
				"name":   "kevin",
				"avatar": "http://img.wxcha.com/file/201803/12/aa106524ed.jpg",
				"roles":  []string{"admin"},
			}))
		})
	}

	projcetController := controllers.NewProjectController()
	projectRouter := r.Group("/project")
	{
		projectRouter.POST("/create", projcetController.CreateProject(mybittDb))
		projectRouter.POST("/list", projcetController.ListProject(mybittDb))
		projectRouter.POST("/update", projcetController.UpdateProject(mybittDb))
	}

	pageController := controllers.NewPageController()
	pageRouter := r.Group("/page")
	{
		pageRouter.POST("/create", pageController.Create(mybittDb))
	}

	r.Run(":8082")
}
