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
	api := router.Group("/api/v1")
	{
		api.GET("/test", controller.TestController)

		// 加载Jwt授权中间件
		api.Use(middleware.JwtAuth(enum.JwtKey))
	}
}
