package controllers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kmookay/MyBittDataManage/auth"
	"github.com/kmookay/MyBittDataManage/constant"
	"github.com/kmookay/MyBittDataManage/dto/request"
	"github.com/kmookay/MyBittDataManage/dto/response"
	"github.com/kmookay/MyBittDataManage/model"
)

type UserController struct{}

var userModel model.UserModel
var rolesModel model.RolesModel

func NewUserController() *UserController {
	userModel = model.UserModel{}
	rolesModel = model.RolesModel{}
	return &UserController{}
}

func (userController *UserController) Login(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginDto request.LoginDto
		err := c.ShouldBindJSON(&loginDto)
		if err != nil {
			log.Fatal(err)
			return
		}

		user, err := userModel.FindByUsernameAndPassword(db, loginDto.UserName, loginDto.Password)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultCommon(constant.STATUS_CODE_USER_OR_PWD_WRONG, "用户名或则密码错误", nil))
			return
		}

		tokenString, err := auth.GenerateToken(&user)
		if err != nil {
			log.Fatal(err)
			c.JSON(http.StatusOK, model.ResultFail("未知错误", nil))
			return
		}

		user, err = userModel.UpdateToken(db, user.UserName, user.Password, tokenString)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail("未知错误", nil))
			return
		}

		rolesId, err := strconv.Atoi(user.Roles)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail("未知错误", nil))
			return
		}

		roles, err := rolesModel.FindById(db, rolesId)
		if err != nil {
			c.JSON(http.StatusOK, model.ResultFail("未知错误", nil))
			return
		}
		c.JSON(http.StatusOK, model.ResultSuccess("登录成功", response.LoginVo{
			User:  user,
			Roles: roles.Name,
		}))
	}
}
