package gin

import (
	"blog/app/middleware"
	"blog/pkg/config"
	"blog/routes"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Start() {
	// 设置gin模式
	gin.SetMode(config.GetConf().Server.Mode)

	router := gin.Default()

	// 自定义异常处理
	router.Use(middleware.Recover)
	// 加载防跨域中间件
	router.Use(middleware.Cors())

	// 读取自定义路由
	setRoute(router)

	if err := router.Run(fmt.Sprintf("%s:%s", config.GetConf().Server.Host, config.GetConf().Server.Port)); err != nil {
		panic(err)
	}
}

// 注册自定义路由
func setRoute(r *gin.Engine) {
	routes.RegisterApiRoutes(r)
}
