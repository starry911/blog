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

		// 文章分类模块
		admin.GET("/article-category-list", controller.Admin.ArticleCategoryListController)
		admin.POST("/article-category-add", controller.Admin.ArticleCategoryAddController)
		admin.PUT("/article-category-edit", controller.Admin.ArticleCategoryEditController)
		admin.DELETE("/article-category-del", controller.Admin.ArticleCategoryDelController)
		admin.GET("/article-category-select", controller.Admin.ArticleCategorySelectController)

		// 文章标签模块
		admin.GET("/article-tags-list", controller.Admin.ArticleTagsListController)
		admin.POST("/article-tags-add", controller.Admin.ArticleTagsAddController)
		admin.DELETE("/article-tags-del", controller.Admin.ArticleTagsDelController)
		admin.GET("/article-tags-select", controller.Admin.ArticleTagsSelectController)

		// 文章模块
		admin.GET("/article-list", controller.Admin.ArticleListController)
		admin.POST("/article-add", controller.Admin.ArticleAddController)

	}
}
