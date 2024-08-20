package admin

import (
	"blog/app/http/requests"
	"blog/pkg/enum"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
	"regexp"
)

func (c *BaseController) ArticleCategoryListController(cxt *gin.Context) {
	var req requests.ArticleCategoryListReq

	// 将请求参数绑定到结构体里面
	if err := cxt.ShouldBind(&req); err != nil {
		response.Fail(enum.InvalidArgument, "参数获取失败！").ToJson(cxt)
		return
	}

	// 调用服务层
	resp := Svc.Admin.ArticleCategoryListService(cxt, &req)
	resp.ToJson(cxt)
	return
}

func (c *BaseController) ArticleCategoryAddController(cxt *gin.Context) {
	var req requests.ArticleCategoryAddReq

	// 将请求参数绑定到结构体里面
	if err := cxt.ShouldBind(&req); err != nil {
		response.Fail(enum.InvalidArgument, "参数获取失败！").ToJson(cxt)
		return
	}

	//校验参数
	if req.Name == "" {
		response.Fail(enum.InvalidArgument, "名称不能为空！").ToJson(cxt)
	}

	if req.Alias == "" {
		response.Fail(enum.InvalidArgument, "别名不能为空！").ToJson(cxt)
	}

	if req.Alias != "" {
		// 判断是否为英文加数字
		re := regexp.MustCompile("^[a-zA-Z0-9]+$")
		if ok := re.MatchString(req.Alias); !ok {
			response.Fail(enum.InvalidArgument, "别名只能英文+数字组合！").ToJson(cxt)
		}
	}

	// 调用服务层
	resp := Svc.Admin.ArticleCategoryAddService(cxt, &req)
	resp.ToJson(cxt)
	return
}

func (c *BaseController) ArticleCategoryEditController(cxt *gin.Context) {
	var req requests.ArticleCategoryEditReq

	// 将请求参数绑定到结构体里面
	if err := cxt.ShouldBind(&req); err != nil {
		response.Fail(enum.InvalidArgument, "参数获取失败！").ToJson(cxt)
		return
	}

	//校验参数
	if req.ID <= 0 {
		response.Fail(enum.InvalidArgument, "分类ID不能为空！").ToJson(cxt)
	}
	if req.Name == "" {
		response.Fail(enum.InvalidArgument, "名称不能为空！").ToJson(cxt)
	}

	if req.Alias == "" {
		response.Fail(enum.InvalidArgument, "别名不能为空！").ToJson(cxt)
	}

	if req.Alias != "" {
		// 判断是否为英文加数字
		re := regexp.MustCompile("^[a-zA-Z0-9]+$")
		if ok := re.MatchString(req.Alias); !ok {
			response.Fail(enum.InvalidArgument, "别名只能英文+数字组合！").ToJson(cxt)
		}
	}

	// 调用服务层
	resp := Svc.Admin.ArticleCategoryEditService(cxt, &req)
	resp.ToJson(cxt)
	return
}

func (c *BaseController) ArticleCategoryDelController(cxt *gin.Context) {
	var req requests.ArticleCategoryDelReq

	// 将请求参数绑定到结构体里面
	if err := cxt.ShouldBind(&req); err != nil {
		response.Fail(enum.InvalidArgument, "参数获取失败！").ToJson(cxt)
		return
	}

	//校验参数
	if req.ID <= 0 {
		response.Fail(enum.InvalidArgument, "分类ID不能为空！").ToJson(cxt)
	}

	// 调用服务层
	resp := Svc.Admin.ArticleCategoryDelService(cxt, &req)
	resp.ToJson(cxt)
	return
}

func (c *BaseController) ArticleCategorySelectController(cxt *gin.Context) {
	// 调用服务层
	resp := Svc.Admin.ArticleCategorySelectService(cxt)
	resp.ToJson(cxt)
	return
}
