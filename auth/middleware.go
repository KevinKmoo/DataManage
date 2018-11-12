package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
	"github.com/kmookay/MyBittDataManage/constant"
	"github.com/kmookay/MyBittDataManage/model"
)

func GenerateToken(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.UserName,
		"userId":   user.Id,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte("secret"))
}

func middlewareHandler(next http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusOK, model.ResultCommon(constant.STATUS_CODE_NEED_LOGIN, "need login", nil))
		} else {
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					c.JSON(http.StatusOK, model.ResultCommon(constant.STATUS_CODE_NEED_LOGIN, "need login", nil))
				}
				return []byte("secret"), nil
			})
			if !token.Valid {
				c.JSON(http.StatusOK, model.ResultCommon(constant.STATUS_CODE_NEED_LOGIN, "need login", nil))
			} else {
				c.Next()
			}
		}
	}
}
