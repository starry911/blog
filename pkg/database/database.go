package database

import (
	"blog/pkg/database/mysql"
	"gorm.io/gorm"
)

var DB *Database

type Database struct {
	MysqlConn *gorm.DB
}

func Init() *Database {
	DB = &Database{
		MysqlConn: mysql.InitMysql(),
	}
	return DB
}
