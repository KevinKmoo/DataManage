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

	//用户相关的接口
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

	//项目相关的接口
	projcetController := controllers.NewProjectController()
	projectRouter := r.Group("/project")
	{
		projectRouter.POST("/create", projcetController.CreateProject(mybittDb))
		projectRouter.POST("/list", projcetController.ListProject(mybittDb))
		projectRouter.POST("/update", projcetController.UpdateProject(mybittDb))
	}

	//版本相关的接口
	versionController := controllers.NewVersionController()
	versionRouter := r.Group("/version")
	{
		versionRouter.POST("/list", versionController.GetAllVersion(mybittDb)) //版本列表
		versionRouter.POST("/create", versionController.Create(mybittDb))      //创建版本
		versionRouter.POST("/delete", versionController.Delete(mybittDb))      // 删除版本
	}

	//页面相关的接口
	pageController := controllers.NewPageController()
	pageRouter := r.Group("/page")
	{
		pageRouter.POST("/create", pageController.Create(mybittDb)) //创建页面
		pageRouter.POST("/update", pageController.Update(mybittDb)) //更新页面数据
	}

	//模块相关的接口
	moduleController := controllers.NewModuleController()
	moduleRouter := r.Group("/module")
	{
		moduleRouter.POST("/create", moduleController.Create(mybittDb))
		moduleRouter.POST("/update", moduleController.Update(mybittDb))
	}

	r.Run(":8082")
}
