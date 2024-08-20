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

func (s *BaseService) ArticleTagsListService(cxt *gin.Context, req *requests.ArticleTagsListReq) *response.JsonResponse {
	// 查询标签列表
	find, count, err := dao.IDao.FindArticleTagsList(req)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleTagsListService---FindArticleTagsList---err:%v", err))
		return response.Fail(enum.HttpFail, "查询错误！")
	}

	list := make([]*requests.ArticleTagsListResp, 0)
	for _, tag := range find {
		list = append(list, &requests.ArticleTagsListResp{
			ID:        tag.ID,
			Name:      tag.Name,
			CreatedAt: tag.CreatedAt.Format(time.DateTime),
		})
	}

	return response.Pagination("获取成功！", count, list)
}

func (s *BaseService) ArticleTagsDelService(cxt *gin.Context, req *requests.ArticleTagsDelReq) *response.JsonResponse {
	find, err := dao.IDao.FindArticleTagsByIds(req.Ids)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleTagsDelService---FindArticleTagsList-find:%v", err))
		return response.Fail(enum.HttpFail, "查询失败！")
	}

	ids := make([]int64, len(find))
	for _, tag := range find {
		ids = append(ids, tag.ID)
	}

	if len(ids) == 0 {
		return response.Success("删除成功！")
	}

	// 删除文章
	err = dao.IDao.DeleteArticleTagsByIds(ids)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleTagsDelService---DeleteArticleTagsList-find:%v", err))
		return response.Fail(enum.HttpFail, "删除失败！")
	}

	return response.Success("删除成功！")
}

func (s *BaseService) ArticleTagsSelectService(cxt *gin.Context) *response.JsonResponse {
	find, err := dao.IDao.FindArticleTagsSelect()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleTagsSelectService---FindArticleTagsSelect-find:%v", err))
		return response.Fail(enum.HttpFail, "查询失败！")
	}

	list := make([]*requests.ArticleTagsSelectResp, 0)
	for _, tag := range find {
		list = append(list, &requests.ArticleTagsSelectResp{
			ID:   tag.ID,
			Name: tag.Name,
		})
	}

	return response.Success("获取成功！", list)
}
