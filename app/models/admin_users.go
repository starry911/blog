package models

import (
	"gorm.io/gorm"
	"time"
)

// AdminUsers 后台用户表
type AdminUsers struct {
	ID       int64     `json:"id" gorm:"id"`
	Account  string    `json:"account" gorm:"account"`     // 登录账号
	Nickname string    `json:"nickname" gorm:"nickname"`   // 昵称
	CoverImg string    `json:"cover_img" gorm:"cover_img"` // 头像地址
	Password string    `json:"password" gorm:"password"`   // 密码
	Salt     string    `json:"salt" gorm:"salt"`           // 密码盐值
	LastIp   string    `json:"last_ip" gorm:"last_ip"`     // 最后登录IP
	LastTime time.Time `json:"last_time" gorm:"last_time"` // 最后登录时间
	gorm.Model
}

// TableName 表名称
func (*AdminUsers) TableName() string {
	return "admin_users"
}
