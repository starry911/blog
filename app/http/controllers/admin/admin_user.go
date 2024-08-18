package admin

import (
	"blog/app/http/requests"
	"blog/pkg/enum"
	"blog/pkg/response"
	"github.com/gin-gonic/gin"
	"regexp"
)

// LoginController 登录
func (c *BaseController) LoginController(cxt *gin.Context) {
	var req requests.BackLoginReq

	// 将请求参数绑定到结构体里面
	if err := cxt.ShouldBind(&req); err != nil {
		response.Fail(enum.InvalidArgument, "参数获取失败！").ToJson(cxt)
		return
	}

	// 验证参数
	if req.Account == "" {
		response.Fail(enum.InvalidArgument, "账号不能为空！").ToJson(cxt)
		return
	}
	if req.Password == "" {
		response.Fail(enum.InvalidArgument, "密码不能为空！").ToJson(cxt)
		return
	}

	// 判断是否账号格式是否正确
	re := regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9@$!%*?&]{4,19}$`)
	if !re.MatchString(req.Account) {
		response.Fail(enum.InvalidArgument, "账号格式错误！").ToJson(cxt)
		return
	}

	// 调用服务层
	resp := Svc.Admin.LoginService(cxt, &req)
	resp.ToJson(cxt)
	return
}

// LogoutController 退出
func (c *BaseController) LogoutController(cxt *gin.Context) {
	// 调用服务层
	resp := Svc.Admin.LogoutService(cxt)
	resp.ToJson(cxt)
	return
}

// UserInfoController 获取用户信息
func (c *BaseController) UserInfoController(cxt *gin.Context) {
	// 调用服务层
	resp := Svc.Admin.UserInfoService(cxt)
	resp.ToJson(cxt)
	return
}

// SetUserInfoController 修改用户信息
func (c *BaseController) SetUserInfoController(cxt *gin.Context) {
	var req requests.SetUserInfoReq

	// 将请求参数绑定到结构体里面
	if err := cxt.ShouldBind(&req); err != nil {
		response.Fail(enum.InvalidArgument, "参数获取失败！").ToJson(cxt)
		return
	}

	// 校验参数
	if req.Nickname != "" {
		// 校验长度
		if len(req.Nickname) >= 150 {
			response.Fail(enum.InvalidArgument, "昵称太长了！").ToJson(cxt)
			return
		}
	}

	// 调用服务层
	resp := Svc.Admin.SetUserInfoService(cxt, &req)
	resp.ToJson(cxt)
	return
}

// SetUserPasswordController 修改用户密码
func (c *BaseController) SetUserPasswordController(cxt *gin.Context) {
	var req requests.SetUserPasswordReq

	// 将请求参数绑定到结构体里面
	if err := cxt.ShouldBind(&req); err != nil {
		response.Fail(enum.InvalidArgument, "参数获取失败！").ToJson(cxt)
		return
	}

	// 校验参数
	if req.OldPassword == "" {
		response.Fail(enum.InvalidArgument, "原密码不能为空！").ToJson(cxt)
		return
	}

	if req.NewPassword == "" {
		response.Fail(enum.InvalidArgument, "新密码不能为空！").ToJson(cxt)
		return
	}

	if req.NewPassword == req.OldPassword {
		response.Fail(enum.InvalidArgument, "新密码不能和原密码相同！").ToJson(cxt)
		return
	}

	// 调用服务层
	resp := Svc.Admin.SetUserPasswordService(cxt, &req)
	resp.ToJson(cxt)
	return
}
