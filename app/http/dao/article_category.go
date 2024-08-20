package dao

import (
	"blog/app/http/requests"
	"blog/app/models"
	"blog/app/utils"
	"blog/pkg/database"
	"context"
	"errors"
	"gorm.io/gorm"
)

type IArticleCategory interface {
	// 分页查询文章分类
	FindArticleCategoryList(req *requests.ArticleCategoryListReq) ([]*models.ArticleCategory, int64, error)
	// 根据名称查询分类
	GetArticleCategoryByName(name string, id int64) (*models.ArticleCategory, error)
	// 根据别名查询分类
	GetArticleCategoryByAlias(name string, id int64) (*models.ArticleCategory, error)
	// 保存分类
	CreateArticleCategory(data *models.ArticleCategory) error
	// 根据id查询分类
	GetArticleCategoryById(id int64) (*models.ArticleCategory, error)
	// 根据id修改分类
	UpdateArticleCategory(id int64, data map[string]interface{}) error
	// 删除分类
	DeleteArticleCategory(id int64) error
	// 查询所有分类
	FindArticleCategorySelect() ([]*models.ArticleCategory, error)
}

// FindArticleCategoryList 分页查询文章分类
func (d *Dao) FindArticleCategoryList(req *requests.ArticleCategoryListReq) ([]*models.ArticleCategory, int64, error) {
	ctx := context.Background()
	sql := database.DB.MysqlConn.WithContext(ctx).Model(&models.ArticleCategory{})

	if req.Name != "" {
		sql = sql.Where("binary name like ?", "%"+req.Name+"%")
	}

	var count int64
	err := sql.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	var find []*models.ArticleCategory
	err = sql.Order("created_at DESC").Offset(utils.GetOffset(req.Page, req.Limit)).Limit(req.Limit).Scan(&find).Error
	if err != nil {
		return nil, 0, err
	}
	return find, count, err

}

// GetArticleCategoryByName 根据名称查询分类
func (d *Dao) GetArticleCategoryByName(name string, id int64) (*models.ArticleCategory, error) {
	ctx := context.Background()
	var category models.ArticleCategory
	sql := database.DB.MysqlConn.WithContext(ctx).Model(&models.ArticleCategory{}).Where("binary name = ?", name)
	if id > 0 {
		sql = sql.Where("id <> ?", id)
	}
	err := sql.First(&category).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// GetArticleCategoryByAlias 根据别名查询分类
func (d *Dao) GetArticleCategoryByAlias(name string, id int64) (*models.ArticleCategory, error) {
	ctx := context.Background()
	var category models.ArticleCategory
	sql := database.DB.MysqlConn.WithContext(ctx).Model(&models.ArticleCategory{}).Where("binary alias = ?", name)
	if id > 0 {
		sql = sql.Where("id <> ?", id)
	}
	err := sql.First(&category).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (d *Dao) CreateArticleCategory(data *models.ArticleCategory) error {
	return database.DB.MysqlConn.Model(&models.ArticleCategory{}).Create(data).Error
}

func (d *Dao) UpdateArticleCategory(id int64, data map[string]interface{}) error {
	return database.DB.MysqlConn.Model(&models.ArticleCategory{}).Where("id = ?", id).Updates(data).Error
}

func (d *Dao) DeleteArticleCategory(id int64) error {
	return database.DB.MysqlConn.Delete(&models.ArticleCategory{}, id).Error
}

func (d *Dao) GetArticleCategoryById(id int64) (*models.ArticleCategory, error) {
	var category models.ArticleCategory
	err := database.DB.MysqlConn.Model(&models.ArticleCategory{}).Where("id = ?", id).Find(&category).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (d *Dao) FindArticleCategorySelect() ([]*models.ArticleCategory, error) {
	var find []*models.ArticleCategory
	err := database.DB.MysqlConn.Model(&models.ArticleCategory{}).Order("created_at DESC").Find(&find).Error
	if err != nil {
		return nil, err
	}
	return find, nil
}
