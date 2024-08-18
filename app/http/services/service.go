package services

import "blog/app/http/services/admin"

// Service 定义服务层结构体
type Service struct {
	Admin admin.BaseService
}
