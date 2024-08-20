package admin

import (
	"blog/app/http/requests"
	"blog/pkg/enum"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

func (c *BaseController) ArticleTagsListController(cxt *gin.Context) {
	var req requests.ArticleTagsListReq

	// 将请求参数绑定到结构体里面
	if err := cxt.ShouldBind(&req); err != nil {
		response.Fail(enum.InvalidArgument, "参数获取失败！").ToJson(cxt)
		return
	}

	// 调用服务层
	resp := Svc.Admin.ArticleTagsListService(cxt, &req)
	resp.ToJson(cxt)
	return
}

func (c *BaseController) ArticleTagsDelController(cxt *gin.Context) {
	var req requests.ArticleTagsDelReq

	// 将请求参数绑定到结构体里面
	if err := cxt.ShouldBind(&req); err != nil {
		response.Fail(enum.InvalidArgument, "参数获取失败！").ToJson(cxt)
		return
	}

	if len(req.Ids) == 0 {
		response.Fail(enum.InvalidArgument, "至少选择一个标签！").ToJson(cxt)
	}

	if len(req.Ids) > 100 {
		response.Fail(enum.InvalidArgument, "一次最多只能删除100个标签！").ToJson(cxt)
	}

	// 调用服务层
	resp := Svc.Admin.ArticleTagsDelService(cxt, &req)
	resp.ToJson(cxt)
	return
}

func (c *BaseController) ArticleTagsSelectController(cxt *gin.Context) {
	// 调用服务层
	resp := Svc.Admin.ArticleTagsSelectService(cxt)
	resp.ToJson(cxt)
	return
}
