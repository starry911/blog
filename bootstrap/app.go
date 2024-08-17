package bootstrap

import (
	"blog/app/crons"
	"blog/app/http/dao"
	cache2 "blog/pkg/cache"
	"blog/pkg/coroutine"
	"blog/pkg/database"
	"blog/pkg/gin"
	"blog/pkg/logger"
)

func Start() {
	// 初始化日志
	logger.InitLogger()

	// 启动协成池
	coroutine.ConnectPool()

	// 初始化数据库
	database.Init()
	// 初始化缓存
	cache2.Init()

	// 初始化Dao层
	dao.New()

	// 启动定时任务
	crons.New()

	// 启动http服务
	gin.Start()
}
