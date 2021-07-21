package controller

import (
	"go_web_demo/models"
	"go_web_demo/pkg/jwt"
	"go_web_demo/utils/gresponse"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context) {
	response := gresponse.Response{C: c}
	user, hasUser := c.Get("user")
	if !hasUser {
		response.Error("用户未登录")
		return
	}
	userInfo := user.(*jwt.Claims)
	res := make(map[string]interface{}, 2)
	res["role"] = models.GetUserRole(userInfo.Id)
	res["auth"] = models.GetUserAuth(userInfo.Id)
	response.SuccessData(res)
}
