package models

import "time"

// ArticleContent 文章内容表
type ArticleContent struct {
	ID          int64     `json:"id" gorm:"id"`
	ArticleId   int64     `json:"article_id" gorm:"article_id"`     // 文章id
	ContentMd   string    `json:"content_md" gorm:"content_md"`     // 文章内容，md格式
	ContentHtml string    `json:"content_html" gorm:"content_html"` // 文章内容，html格式
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`     // 创建时间
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`     // 修改时间
}

// TableName 表名称
func (*ArticleContent) TableName() string {
	return "article_content"
}
