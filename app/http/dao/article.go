package dao

import (
	"blog/app/http/requests"
	"blog/app/models"
	"blog/app/utils"
	"blog/pkg/database"
	"context"
)

type IArticle interface {
	// 根据分类id，查询文章数量
	GetArticleNumByCategory(categoryId int64) (int64, error)
	// 查询文章列表
	FindArticleList(req *requests.ArticleListReq) ([]*models.Article, int64, error)
}

func (d *Dao) GetArticleNumByCategory(categoryId int64) (int64, error) {
	var count int64
	err := database.DB.MysqlConn.Model(&models.Article{}).Where("category_id = ?", categoryId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (d *Dao) FindArticleList(req *requests.ArticleListReq) ([]*models.Article, int64, error) {
	ctx := context.Background()
	sql := database.DB.MysqlConn.WithContext(ctx).Model(&models.Article{})

	if req.Title != "" {
		sql = sql.Where("title like ?", "%"+req.Title+"%")
	}

	if req.CategoryId != 0 {
		sql = sql.Where("category_id = ?", req.CategoryId)
	}

	if req.Status != 0 {
		sql = sql.Where("status = ?", req.Status)
	}

	if req.StartTime != "" {
		sql = sql.Where("created_at >= ?", req.StartTime)
	}

	if req.EndTime != "" {
		sql = sql.Where("created_at <= ?", req.EndTime)
	}

	var count int64
	err := sql.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	var list []*models.Article
	err = sql.Order("created_at desc").Offset(utils.GetOffset(req.Page, req.Limit)).Limit(req.Limit).Scan(&list).Error
	if err != nil {
		return nil, 0, err
	}
	return list, count, nil
}
