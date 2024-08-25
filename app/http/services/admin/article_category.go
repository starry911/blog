package admin

import (
	"blog/app/http/dao"
	"blog/app/http/requests"
	"blog/app/models"
	"blog/pkg/enum"
	"blog/pkg/logger"
	"blog/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func (s *BaseService) ArticleCategoryListService(cxt *gin.Context, req *requests.ArticleCategoryListReq) *response.JsonResponse {
	find, count, err := dao.IDao.FindArticleCategoryList(req)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleCategoryListService---FindArticleCategoryList---err:%v", err))
		return response.Fail(enum.HttpError, "查询失败！")
	}
	list := make([]*requests.ArticleCategoryListResp, 0)
	for _, category := range find {
		// 查询文章数量
		num, err := dao.IDao.GetArticleNumByCategory(category.ID)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("ArticleCategoryListService---GetArticleNumByCategory---err:%v", err))
		}
		list = append(list, &requests.ArticleCategoryListResp{
			ID:         category.ID,
			Name:       category.Name,
			Alias:      category.Alias,
			ArticleNum: num,
			CreatedAt:  category.CreatedAt.Format(time.DateTime),
		})
	}

	return response.Pagination("获取成功！", count, list)
}

func (s *BaseService) ArticleCategoryAddService(cxt *gin.Context, req *requests.ArticleCategoryAddReq) *response.JsonResponse {
	// 判断名称是否重复
	category1, err := dao.IDao.GetArticleCategoryByName(req.Name, 0)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleCategoryAddService---GetArticleCategoryByName---err:%v", err))
		return response.Fail(enum.HttpError, "查询错误！")
	}

	if category1 != nil {
		return response.Fail(enum.HttpFail, "该分类名称已经存在了！")
	}

	// 判断别名是否重复
	category2, err := dao.IDao.GetArticleCategoryByAlias(req.Alias, 0)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleCategoryAddService---GetArticleCategoryByAlias---err:%v", err))
		return response.Fail(enum.HttpError, "查询错误！")
	}
	if category2 != nil {
		return response.Fail(enum.HttpFail, "该分类别名已经存在了！")
	}

	err = dao.IDao.CreateArticleCategory(&models.ArticleCategory{
		Name:  req.Name,
		Alias: req.Alias,
	})
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleCategoryAddService---CreateArticleCategory---err:%v", err))
		return response.Fail(enum.HttpError, "新增失败！")
	}
	return response.Success("新增成功！")
}

func (s *BaseService) ArticleCategoryEditService(cxt *gin.Context, req *requests.ArticleCategoryEditReq) *response.JsonResponse {
	// 查询分类是否存在
	category, err := dao.IDao.GetArticleCategoryById(req.ID)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleCategoryEditService---GetArticleCategoryById---err:%v", err))
		return response.Fail(enum.HttpFail, "查询错误！")
	}
	if category == nil {
		return response.Fail(enum.HttpFail, "分类不存在！")
	}
	// 判断名称是否重复
	category1, err := dao.IDao.GetArticleCategoryByName(req.Name, category.ID)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleCategoryEditService---GetArticleCategoryByName---err:%v", err))
		return response.Fail(enum.HttpFail, "查询错误！")
	}

	if category1 != nil {
		return response.Fail(enum.HttpFail, "该分类名称已经存在了！")
	}

	// 判断别名是否重复
	category2, err := dao.IDao.GetArticleCategoryByAlias(req.Alias, category.ID)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleCategoryEditService---GetArticleCategoryByAlias---err:%v", err))
		return response.Fail(enum.HttpFail, "查询错误！")
	}
	if category2 != nil {
		return response.Fail(enum.HttpFail, "该分类别名已经存在了！")
	}

	err = dao.IDao.UpdateArticleCategory(category.ID, map[string]interface{}{
		"name":  req.Name,
		"alias": req.Alias,
	})
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleCategoryEditService---UpdateArticleCategory---err:%v", err))
		return response.Fail(enum.HttpFail, "修改失败！")
	}
	return response.Success("修改成功！")
}

func (s *BaseService) ArticleCategoryDelService(cxt *gin.Context, req *requests.ArticleCategoryDelReq) *response.JsonResponse {
	// 查询分类是否存在
	category, err := dao.IDao.GetArticleCategoryById(req.ID)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleCategoryDelService---GetArticleCategoryById---err:%v", err))
		return response.Fail(enum.HttpFail, "查询错误！")
	}
	if category == nil {
		return response.Fail(enum.HttpFail, "分类不存在！")
	}

	// 查询分类下面还有文章的，就不允许删除
	count, err := dao.IDao.GetArticleNumByCategory(category.ID)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleCategoryDelService---GetArticleNumByCategory---err:%v", err))
		return response.Fail(enum.HttpFail, "查询错误！")
	}
	if count > 0 {
		return response.Fail(enum.HttpFail, "该分类下还存在显示的文章，请先删除文章！")
	}

	err = dao.IDao.DeleteArticleCategory(category.ID)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleCategoryDelService---DeleteArticleCategory---err:%v", err))
		return response.Fail(enum.HttpFail, "删除失败！")
	}
	return response.Success("删除成功！")
}

func (s *BaseService) ArticleCategorySelectService(cxt *gin.Context) *response.JsonResponse {
	find, err := dao.IDao.FindArticleCategorySelect()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("ArticleCategorySelectService---FindArticleCategorySelect---err:%v", err))
		return response.Fail(enum.HttpError, "查询失败！")
	}
	list := make([]*requests.ArticleCategorySelectResp, 0)

	for _, v := range find {
		list = append(list, &requests.ArticleCategorySelectResp{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	return response.Success("获取成功！", list)
}
