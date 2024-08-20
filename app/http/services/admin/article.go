package admin

import (
	"blog/app/http/dao"
	"blog/app/http/requests"
	"blog/pkg/enum"
	"blog/pkg/logger"
	"blog/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func (s *BaseService) ArticleListService(cxt *gin.Context, req *requests.ArticleListReq) *response.JsonResponse {
	// 查询文章列表
	find, count, err := dao.IDao.FindArticleList(req)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleListService---FindArticleList---err:%v", err))
		return response.Fail(enum.HttpFail, "查询错误！")
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
			CreatedAt:    v.CreatedAt.Format(time.DateTime),
		})
	}
	return response.Pagination("获取成功！", count, list)
}
