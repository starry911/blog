package admin

import (
	"blog/app/http/requests"
	"blog/app/utils"
	"blog/pkg/enum"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
	"time"
)

func (c *BaseController) ArticleListController(cxt *gin.Context) {
	var req requests.ArticleListReq

	// 将请求参数绑定到结构体里面
	if err := cxt.ShouldBind(&req); err != nil {
		response.Fail(enum.InvalidArgument, "参数获取失败！").ToJson(cxt)
		return
	}

	if req.CategoryId < 0 {
		response.Fail(enum.InvalidArgument, "分类id不能为空！")
		return
	}

	if req.Status != 0 && req.Status != 1 && req.Status != 2 {
		response.Fail(enum.InvalidArgument, "状态值错误！")
		return
	}

	// 验证时间格式是否正确
	if req.StartTime != "" {
		if ok := utils.VerifyTimeStr(time.DateTime, req.StartTime); !ok {
			response.Fail(enum.InvalidArgument, "开始时间格式不合法！")
			return
		}
	}
	if req.EndTime != "" {
		if ok := utils.VerifyTimeStr(time.DateTime, req.EndTime); !ok {
			response.Fail(enum.InvalidArgument, "结束时间格式不合法！")
			return
		}
	}

	// 调用服务层
	resp := Svc.Admin.ArticleListService(cxt, &req)
	resp.ToJson(cxt)
	return
}

func (c *BaseController) ArticleAddController(cxt *gin.Context) {
	var req requests.ArticleAddReq

	// 将请求参数绑定到结构体里面
	if err := cxt.ShouldBind(&req); err != nil {
		response.Fail(enum.InvalidArgument, "参数获取失败！").ToJson(cxt)
		return
	}

	if req.CategoryId <= 0 {
		response.Fail(enum.InvalidArgument, "分类id不能为空！")
		return
	}

	if req.Title == "" {
		response.Fail(enum.InvalidArgument, "标题不能为空！")
		return
	}

	if req.Describe == "" {
		response.Fail(enum.InvalidArgument, "描述不能为空！")
		return
	}

	if req.CoverImg == "" {
		response.Fail(enum.InvalidArgument, "请上传封面图！")
		return
	}

	if len(req.Tags) == 0 {
		response.Fail(enum.InvalidArgument, "请至少填写一个标签！")
		return
	}

	if req.PublishTime == "" {
		req.PublishTime = time.Now().Format(time.DateTime)
	} else {
		if ok := utils.VerifyTimeStr(time.DateTime, req.PublishTime); !ok {
			response.Fail(enum.InvalidArgument, "发布时间格式错误！")
			return
		}
	}

	if req.Status != 1 && req.Status != 2 {
		response.Fail(enum.InvalidArgument, "状态值错误！")
		return
	}

	if req.ContentHtml == "" {
		response.Fail(enum.InvalidArgument, "文章内容不能为空！")
		return
	}

	if req.ContentMd == "" {
		response.Fail(enum.InvalidArgument, "文章内容不能为空！")
		return
	}

	// 调用服务层
	resp := Svc.Admin.ArticleAddService(cxt, &req)
	resp.ToJson(cxt)
	return
}
