package mysql

import (
	"blog/pkg/config"
	logger2 "blog/pkg/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitMysql() *gorm.DB {
	var (
		host     = config.GetConf().Mysql.Host
		port     = config.GetConf().Mysql.Port
		database = config.GetConf().Mysql.Database
		username = config.GetConf().Mysql.Username
		password = config.GetConf().Mysql.Password
		charset  = config.GetConf().Mysql.Charset
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&collation=utf8mb4_unicode_ci&parseTime=true&loc=Local&timeout=30s", username, password, host, port, database, charset)

	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // 不记录SQL慢日志
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",   // 表前缀
			SingularTable: true, // 表名不加s
		},
	})

	if err != nil {
		logger2.Logger.Error("Mysql 连接异常: " + err.Error())
		panic(err)
	}

	//设置连接池
	db, _ := DB.DB()
	//空闲
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	db.SetMaxIdleConns(150)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(300)

	// 5s空闲时间
	db.SetConnMaxLifetime(5 * time.Second)

	err = db.Ping()
	if err != nil {
		logger2.Logger.Error("Mysql 无法Ping通: " + err.Error())
		panic(err)
	}

	return DB
}
