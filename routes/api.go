package routes

import (
	"blog/app/http/controllers"
	"blog/app/middleware"
	"blog/pkg/enum"
	"github.com/gin-gonic/gin"
)

// RegisterApiRoutes 注册api路由
func RegisterApiRoutes(router *gin.Engine) {
	controller := controllers.Controller{}
	admin := router.Group("/admin/v1")
	{
		// 后台登录接口
		admin.POST("/login", controller.Admin.LoginController)

		// 加载Jwt授权中间件
		admin.Use(middleware.JwtAuth(enum.JwtKey))

		// 登录用户模块
		admin.GET("/logout", controller.Admin.LogoutController)
		admin.GET("/user-info", controller.Admin.UserInfoController)
		admin.PUT("/set-user-info", controller.Admin.SetUserInfoController)
		admin.PUT("/set-user-password", controller.Admin.SetUserPasswordController)

		// 文章模块
	}
}
