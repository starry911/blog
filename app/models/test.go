package models

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	Name string `json:"name"`
}

// TableName 指定表名
func (Test) TableName() string {
	return "test"
}
