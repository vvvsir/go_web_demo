package controller

import (
	"go_web_demo/middleware"
	"go_web_demo/models/request"
	"go_web_demo/pkg/logger"
	"go_web_demo/pkg/setting"
	"go_web_demo/service"
	"go_web_demo/utils/gresponse"

	"github.com/gin-gonic/gin"
)

// @Tags 登录操作
// @Summary 登录操作
// @Produce json
// @Param body body request.LoginUser true "用户"
// @Success 200
// @Router /login [post]
func Login(c *gin.Context) {
	response := gresponse.Response{C: c}
	user := request.LoginUser{}
	if err := c.ShouldBindJSON(&user); err != nil {
		response.SetValidateError(err)
		return
	}

	id, ok := c.Get("session_id")
	if !ok {
		response.ErrorMsg("GET SESSION ERROR")
		return
	}
	idStr := id.(string)
	token, err := service.Login(&user, idStr)
	if err != nil {
		response.SetOtherError(err)
		return
	}

	c.SetCookie(middleware.JwtName, token, setting.JwtConf.ExpTime*3600, "/", setting.AppConf.Doamin, false, true)
	response.SuccessData(token)
}

// @Tags 登录操作
// @Summary 验证码
// @Success 200
// @Router /captcha [get]
func Captcha(c *gin.Context) {
	id, ok := c.Get("session_id")
	if !ok {
		logger.Logger.Error("获取 session 异常")
		return
	}
	idStr := id.(string)
	image, err := service.NewCaptcha(idStr)
	if err != nil {
		logger.Logger.Error("创建验证码异常：", err)
		return
	}
	c.Header("content-type", "image/png")
	_, err = image.WriteTo(c.Writer)
	if err != nil {
		logger.Logger.Error("c.Writer.Write() 异常：", err)
	}
}
