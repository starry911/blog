package controllers

import (
	"blog/app/http/controllers/admin"
)

// Controller 定义控制器结构体
type Controller struct {
	Admin admin.BaseController
}
