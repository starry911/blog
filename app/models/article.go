package models

import (
	"gorm.io/gorm"
	"time"
)

// Article 文章表
type Article struct {
	ID          int64     `json:"id" gorm:"id"`
	Uuid        string    `json:"uuid" gorm:"uuid"`                 // 文章uuid
	Title       string    `json:"title" gorm:"title"`               // 文章标题
	Describe    string    `json:"describe" gorm:"describe"`         // 描述
	CategoryId  int64     `json:"category_id" gorm:"category_id"`   // 文章分类id
	CoverImg    string    `json:"cover_img" gorm:"cover_img"`       // 文章封面
	PublishTime time.Time `json:"publish_time" gorm:"publish_time"` // 文章发布时间
	TagIds      string    `json:"tag_ids" gorm:"tag_ids"`           // 文章标签集合，以英文逗号分割
	Status      int64     `json:"status" gorm:"status"`             // 状态，1：显示，2：隐藏
	ViewNum     int64     `json:"view_num" gorm:"view_num"`         // 文章浏览量
	LikeNum     int64     `json:"like_num" gorm:"like_num"`         // 文章点赞量
	gorm.Model
}

// TableName 表名称
func (*Article) TableName() string {
	return "article"
}
