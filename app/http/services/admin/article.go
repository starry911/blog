package admin

import (
	"blog/app/http/dao"
	"blog/app/http/requests"
	"blog/app/models"
	"blog/app/utils"
	"blog/pkg/enum"
	"blog/pkg/logger"
	"blog/pkg/response"
	"blog/pkg/uuidStr"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)

func (s *BaseService) ArticleListService(cxt *gin.Context, req *requests.ArticleListReq) *response.JsonResponse {
	// 查询文章列表
	find, count, err := dao.IDao.FindArticleList(req)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleListService---FindArticleList---err:%v", err))
		return response.Fail(enum.HttpError, "查询错误！")
	}

	list := make([]*requests.ArticleListResp, 0)

	for _, v := range find {
		// 分类
		category, err := dao.IDao.GetArticleCategoryById(v.CategoryId)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("ArticleListService---GetArticleCategoryById---err:%v", err))
		}
		var CategoryName string
		if category != nil {
			CategoryName = category.Name
		}

		// 标签
		var TagsName []string
		tags, err := dao.IDao.FindArticleTagsByIdsString(v.TagIds)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("ArticleListService---FindArticleTagsByIdsString---err:%v", err))
		}

		for _, tag := range tags {
			TagsName = append(TagsName, tag.Name)
		}

		list = append(list, &requests.ArticleListResp{
			ID:           v.ID,
			UUID:         v.Uuid,
			Title:        v.Title,
			CategoryName: CategoryName,
			CoverImg:     v.CoverImg,
			PublishTime:  v.PublishTime.Format(time.DateTime),
			Status:       v.Status,
			Tags:         TagsName,
			ViewNum:      v.ViewNum,
			LikeNum:      v.LikeNum,
			CreatedAt:    v.CreatedAt.Format(time.DateTime),
		})
	}
	return response.Pagination("获取成功！", count, list)
}

func (s *BaseService) ArticleAddService(cxt *gin.Context, req *requests.ArticleAddReq) *response.JsonResponse {
	// 查询分类
	category, err := dao.IDao.GetArticleCategoryById(req.CategoryId)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleAddService---GetArticleCategoryById---err:%v", err))
		return response.Fail(enum.HttpError, "查询错误！")
	}
	if category == nil {
		return response.Fail(enum.HttpFail, "分类不存在！")
	}

	// 查询文章标题是否已经重复
	article, err := dao.IDao.GetArticleByTitle(req.Title)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleAddService---GetArticleByTitle---err:%v", err))
		return response.Fail(enum.HttpError, "查询错误！")
	}

	if article != nil {
		return response.Fail(enum.HttpFail, "已经存在相同标题的文章了！")
	}

	// 查询标签
	tags, err := dao.IDao.FindArticleTagsByIds(req.Tags)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleAddService---FindArticleTagsByIds---err:%v", err))
		return response.Fail(enum.HttpError, "查询失败！")
	}

	var tagIds []string
	for _, tag := range tags {
		tagIds = append(tagIds, strconv.FormatInt(tag.ID, 10))
	}
	tagIdsStr := strings.Join(tagIds, ",")

	PublishTime, _ := time.ParseInLocation(time.DateTime, req.PublishTime, time.Local)

	articleModel := &models.Article{
		Uuid:        utils.UUIDGetHex(uuidStr.UUID4()),
		Title:       req.Title,
		Describe:    req.Describe,
		CategoryId:  req.CategoryId,
		CoverImg:    req.CoverImg,
		PublishTime: PublishTime,
		TagIds:      tagIdsStr,
		Status:      req.Status,
		Model:       gorm.Model{},
	}
	content := &models.ArticleContent{
		ArticleId:   0,
		ContentMd:   req.ContentMd,
		ContentHtml: req.ContentHtml,
	}

	// 保存文章
	err = dao.IDao.CreateArticle(articleModel, content)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleAddService---CreateArticle---err:%v", err))
		return response.Fail(enum.HttpError, "添加失败！")
	}
	return response.Success("添加成功！")
}
