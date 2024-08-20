package dao

import (
	"blog/app/http/requests"
	"blog/app/models"
	"blog/app/utils"
	"blog/pkg/database"
	"context"
	"errors"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type IArticleTags interface {
	// 查询标签列表
	FindArticleTagsList(req *requests.ArticleTagsListReq) ([]*models.ArticleTags, int64, error)
	// 根据ids查询标签
	FindArticleTagsByIds(ids []int64) ([]*models.ArticleTags, error)
	// 删除标签
	DeleteArticleTagsByIds(ids []int64) error
	// 查询所有标签
	FindArticleTagsSelect() ([]*models.ArticleTags, error)
	// 根据名称查询标签
	GetArticleTagByName(name string) (*models.ArticleTags, error)
	// 根据ids查询标签
	FindArticleTagsByIdsString(ids string) ([]*models.ArticleTags, error)
}

func (d *Dao) FindArticleTagsList(req *requests.ArticleTagsListReq) ([]*models.ArticleTags, int64, error) {
	ctx := context.Background()
	sql := database.DB.MysqlConn.WithContext(ctx).Model(&models.ArticleTags{})

	if req.Name != "" {
		sql = sql.Where("binary name like ?", "%"+req.Name+"%")
	}

	var count int64
	err := sql.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	var find []*models.ArticleTags
	err = sql.Order("created_at desc").Offset(utils.GetOffset(req.Page, req.Limit)).Limit(req.Limit).Scan(&find).Error
	if err != nil {
		return nil, 0, err
	}
	return find, count, nil
}

func (d *Dao) FindArticleTagsByIds(ids []int64) ([]*models.ArticleTags, error) {
	var find []*models.ArticleTags
	err := database.DB.MysqlConn.Model(&models.ArticleTags{}).Where("id in (?)", ids).Scan(&find).Error
	if err != nil {
		return nil, err
	}
	return find, nil
}

func (d *Dao) DeleteArticleTagsByIds(ids []int64) error {
	return database.DB.MysqlConn.Delete(&models.ArticleTags{}, "id in (?)", ids).Error
}

func (d *Dao) FindArticleTagsSelect() ([]*models.ArticleTags, error) {
	var find []*models.ArticleTags
	err := database.DB.MysqlConn.Model(&models.ArticleTags{}).Order("created_at desc").Find(&find).Error
	if err != nil {
		return nil, err
	}
	return find, nil
}

func (d *Dao) GetArticleTagByName(name string) (*models.ArticleTags, error) {
	var tag models.ArticleTags
	err := database.DB.MysqlConn.Model(&models.ArticleTags{}).Where("binary name = ?", name).First(&tag).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (d *Dao) FindArticleTagsByIdsString(ids string) ([]*models.ArticleTags, error) {
	// 将id集合转为int
	idsSlice := strings.Split(ids, ",")
	idsInt64 := make([]int64, len(idsSlice))
	for i, v := range idsSlice {
		idsInt64[i], _ = strconv.ParseInt(v, 10, 64)
	}

	var find []*models.ArticleTags
	err := database.DB.MysqlConn.Model(&models.ArticleTags{}).Where("id in (?)", idsInt64).Scan(&find).Error
	if err != nil {
		return nil, err
	}
	return find, nil
}
