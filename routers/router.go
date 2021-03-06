package routers

import (
	"go_web_demo/controller"
	_ "go_web_demo/docs"
	"go_web_demo/middleware"
	"go_web_demo/pkg/setting"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.AppConf.Env)
	r := gin.New()

	// swagger 文档输出
	if setting.AppConf.Env == "debug" {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// 加入通用中间件
	r.Use(
		gin.Recovery(),           // recovery 防止程序奔溃
		middleware.Logger(),      // 日志记录
		middleware.NoFound(),     // 404 处理
		middleware.MakeSession(), // session支持
		middleware.ErrorHandle(), // 错误处理
	)

	r.GET("/captcha", controller.Captcha)
	r.POST("/login", controller.Login)

	// 加入鉴权中间件
	r.Use(middleware.JWT())
	r.GET("/userInfo", controller.UserInfo)
	r.Use(middleware.Authentication())
	// 用户
	{
		r.GET("/auth/role", controller.RoleTree)
		r.GET("/adminUser", controller.UserIndex)
		r.GET("/adminUser/:id", controller.UserDetail)
		r.POST("/adminUser", controller.UserAdd)
		r.PUT("/adminUser/:id", controller.UserEdit)
		r.DELETE("/adminUser/:id", controller.UserDel)
	}

	// 角色
	{
		r.GET("/auth/tree", controller.AuthTree)
		r.GET("/role", controller.RoleIndex)
		r.GET("/role/:id", controller.RoleDetail)
		r.POST("/role", controller.RoleAdd)
		r.PUT("/role/:id", controller.RoleEdit)
		r.DELETE("/role/:id", controller.RoleDel)
	}

	// 权限
	{
		r.GET("/auth/treemenu", controller.TreeMenu)
	}
	return r
}
