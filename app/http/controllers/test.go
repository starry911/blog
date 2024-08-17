package controllers

import (
	"blog/app/http/requests"
	"blog/pkg/enum"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
)

func (c *Controller) TestController(cxt *gin.Context) {
	var req requests.TestReq

	// 将请求参数绑定到结构体里面
	if err := cxt.ShouldBind(&req); err != nil {
		response.Fail(enum.InvalidArgument, "参数获取失败！").ToJson(cxt)
		return
	}

	// 验证参数
	if req.Name == "" {
		response.Fail(enum.InvalidArgument, "名称不能为空！").ToJson(cxt)
		return
	}

	// 调用服务层
	resp := Svc.TestService(cxt, &req)
	resp.ToJson(cxt)
	return
}
