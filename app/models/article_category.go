package models

import "gorm.io/gorm"

// ArticleCategory 文章分类表
type ArticleCategory struct {
	ID    int64  `json:"id" gorm:"id"`
	Name  string `json:"name" gorm:"name"`   // 分类名称
	Alias string `json:"alias" gorm:"alias"` // 别名，纯英文，用于跳转分类页面
	gorm.Model
}

// TableName 表名称
func (*ArticleCategory) TableName() string {
	return "article_category"
}
