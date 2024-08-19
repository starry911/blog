package models

import "gorm.io/gorm"

// ArticleTags 文章标签表
type ArticleTags struct {
	ID   int64  `json:"id" gorm:"id"`
	Name string `json:"name" gorm:"name"` // 标签名称
	gorm.Model
}

// TableName 表名称
func (*ArticleTags) TableName() string {
	return "article_tags"
}
